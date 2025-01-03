package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	// k8score "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	GroupVersion = schema.GroupVersion{
		Group: "phoenix.setimozak",
		Version: "v1beta1",
	}
	
	SchemeBuilder = &scheme.Builder{
		GroupVersion: GroupVersion,
	}
	
	AddToScheme = SchemeBuilder.AddToScheme
)

func (in *EnvManager) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	newObj := new(EnvManager)
	// primitive types don't require deep copy
	newObj.Spec.Enable = in.Spec.Enable
	newObj.Spec.MinReplica = in.Spec.MinReplica

	return newObj
}

func (in *EnvManagerList) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	newObj := new(EnvManagerList)
	copy(newObj.Items, in.Items)
	
	return newObj
}





