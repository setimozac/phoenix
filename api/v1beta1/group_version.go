package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is used to register the objects/resources.

	// EXAMPLE: in CR manifest yaml file:
	// {
		// "Kind": "EnvManager",
		// "apiVersion": "phoenix.setimozak/v1beta1"
	// }
	GroupVersion = schema.GroupVersion{
		Group:   "phoenix.setimozak",
		Version: "v1beta1",
	}

	// SchemeBuilder is used to register Kind's Go types the GVK scheme
	SchemeBuilder = &scheme.Builder{
		GroupVersion: GroupVersion,
	}

	// AddToScheme = SchemeBuilder.AddToScheme
)
