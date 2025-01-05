package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	GroupVersion = schema.GroupVersion{
		Group:   "phoenix.setimozak",
		Version: "v1beta1",
	}

	SchemeBuilder = &scheme.Builder{
		GroupVersion: GroupVersion,
	}

	AddToScheme = SchemeBuilder.AddToScheme
)
