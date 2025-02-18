package controllers

import (
	"context"

	phoenixv1beta1 "github.com/setimozac/phoenix/api/v1beta1"
	// batchv1 "k8s.io/api/batch/v1"
	// corev1 "k8s.io/api/core/v1"
	// apimeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	appsv1 "k8s.io/api/apps/v1"
	// ref "k8s.io/client-go/tools/reference"
	ctl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type EnvManagerReconciler struct{
	client.Client
	Scheme *runtime.Scheme
}

func (r *EnvManagerReconciler) Reconcile(ctx context.Context, req ctl.Request) (ctl.Result, error) {
	log := log.FromContext(ctx)
	log.WithName("Reconciler")

	var envManager phoenixv1beta1.EnvManager
	if err := r.Get(ctx, req.NamespacedName, &envManager); err != nil {
		log.Error(err, "unable to get the envmanager")

		return ctl.Result{}, client.IgnoreNotFound(err)
	}
	log.V(1).Info("testing the controller get function", "envManager", envManager)

	var envManagerList phoenixv1beta1.EnvManagerList
	if err := r.List(ctx, &envManagerList, client.MatchingFields{"spec.name": envManager.Spec.Name}); err != nil {
		log.Error(err, "unable to list the envmanagers")

		return ctl.Result{}, err
	}
	log.V(1).Info("testing the controller get function", "envManagerList", envManagerList)

	var deployment appsv1.Deployment
	deploymentName := types.NamespacedName{
		Name: envManager.Spec.Name,
		Namespace: req.Namespace,
	}
	if err := r.Get(ctx, deploymentName, &deployment); err != nil {
		log.Error(err, "unable to get the deployment")
		return ctl.Result{}, client.IgnoreNotFound(err)
	}
	log.V(1).Info("testing the deployment get function", "deployment", deployment)

	log.V(1).Info("<------ Updating the deployment ----->")
	var desiredReplicas int32
	if !*envManager.Spec.UIEnabled {
		desiredReplicas = int32(*envManager.Spec.MinReplica)
	} else {
		desiredReplicas = 0
	}
	
	if deployment.Spec.Replicas != nil && deployment.Spec.Replicas == &desiredReplicas {
		log.V(1).Info("The deployment replicas matches the CR", "deployment.repicas:", deployment.Spec.Replicas, "CR.MinReplica", desiredReplicas)
		return ctl.Result{}, nil
	}

	// update the deployment
	deployment.Spec.Replicas = &desiredReplicas
	err := r.Client.Update(ctx, &deployment)
	if err != nil {
		log.Error(err, "failed to update the deployment")
		return ctl.Result{}, nil
	}

	// requeue to ensure the update was applied
	return ctl.Result{Requeue: true}, nil
}

func (r *EnvManagerReconciler) SetupWithManager(mgr ctl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &phoenixv1beta1.EnvManager{}, "spec.name", func(obj client.Object) []string{
		envM := obj.(*phoenixv1beta1.EnvManager)
		return []string{envM.Spec.Name}
	}); err != nil {
		log.Log.V(1).Error(err, err.Error())
		return err
	}
	// var deployment appsv1.Deployment

	return ctl.NewControllerManagedBy(mgr).
	For(&phoenixv1beta1.EnvManager{}).
	Watches(
		&appsv1.Deployment{},
		handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
			d, ok := obj.(*appsv1.Deployment)
			if !ok {
				log.Log.V(1).Info("getting obj in the Wacher in SetupWithManager()")
				return nil
			}

			var envManagerList phoenixv1beta1.EnvManagerList
			if err := r.List(ctx, &envManagerList, client.MatchingFields{"spec.name": d.Name}); err != nil {
				log.Log.Error(err, "unable to list the envmanagers")

				return nil
			}

			var requests []reconcile.Request
			if len(envManagerList.Items) > 0 {
				for _, i := range envManagerList.Items {
					requests = append(requests, reconcile.Request{
						NamespacedName: client.ObjectKey{
							Namespace: i.Namespace,
							Name: i.Name,
						},
					})
				}
			}
			return requests
		}),
	).
	Named("envManager").
	Complete(r)
	
}