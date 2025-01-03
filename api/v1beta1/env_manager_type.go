package v1beta1

import (
	k8sbatch "k8s.io/api/batch/v1"
	k8score "k8s.io/api/core/v1"
	apimeta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvManagerSpec struct {
	Enable bool             `json:"enable"`
	MinReplica int          `json:"minReplica,omitempty"`
}

type EnvManagerStatus struct {
	ControlledBy string     `json:"controlledBy,omitempty"`
}

type EnvManager struct {
	apimeta.TypeMeta        `json:",inline"`
	apimeta.ObjectMeta      `json:"metadata,omitempty"`

	Spec EnvManagerSpec     `json:"spec,omitempty"`
	Status EnvManagerStatus `json:"status,omitempty"`
}

type EnvManagerList struct {
	apimeta.TypeMeta        `json:",inline"`
	apimeta.ObjectMeta      `json:"metadata,omitempty"`
	Items []EnvManager      `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnvManager{}, &EnvManagerList{})
}