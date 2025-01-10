package main

import (
	"crypto/tls"
	"flag"
	"os"

	phoenixv1beta1 "github.com/setimozac/phoenix/api/v1beta1"
	"github.com/setimozac/phoenix/internal/controllers"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var (
	scheme = runtime.NewScheme()
	setupLog = ctl.Log.WithName("manager-setup")
)


func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(phoenixv1beta1.SchemeBuilder.AddToScheme(scheme))
}

func main() {
	var tlsOpts []func(*tls.Config)
	var healthProbe string
	var leaderElection bool
	
	flag.StringVar(&healthProbe, "health-probe-addr", ":8081", "The address for probe endpoint")
	flag.BoolVar(&leaderElection,"leader-election", true, "Enable/Disable the leader election." + 
			"Enable by default and it ensures there is only one manager.")
	logOptions := zap.Options{
		Development: true,
	}
	logOptions.BindFlags(flag.CommandLine)
	flag.Parse()
	// setting up logger for the controller runtime
	ctl.SetLogger(zap.New(zap.UseFlagOptions(&logOptions)))


	http1 := func(c *tls.Config) {
		c.NextProtos = []string{"http/1.1"}
	}
	tlsOpts = append(tlsOpts, http1)

	// namespace restrictions for manager

	namespaces := []string{"default"} // list of namespaces
	defaultNameSpaces := make(map[string]cache.Config)
	for _, namespace := range namespaces {
		defaultNameSpaces[namespace] = cache.Config{}
	}

	// initializing a new webhook
	webhookServer := webhook.NewServer(webhook.Options{TLSOpts: tlsOpts})

	mgr, err := ctl.NewManager(ctl.GetConfigOrDie(), ctl.Options{
		Scheme: scheme,
		Cache: cache.Options{DefaultNamespaces: defaultNameSpaces},
		WebhookServer: webhookServer,
		PprofBindAddress: healthProbe,
		LeaderElection: leaderElection,
		LeaderElectionID: "envmanager.setimozac.phoenix",
	})
	if err != nil {
		setupLog.Error(err, "unable to start the manager")
	}

	if err := (&controllers.EnvManagerReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "EnvManager")
		os.Exit(1)
	}

	if err := mgr.AddHealthzCheck("health_check", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to setup the healt_check")
		os.Exit(1)
	}

	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to setup the readyz")
		os.Exit(1)
	}

	setupLog.Info("starting the manager...")
	if err := mgr.Start(ctl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "unable to start the manager")
		os.Exit(1)
	}

}
