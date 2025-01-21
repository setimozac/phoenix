package v1beta1

import (
	apimeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type EnvManagerSpec struct {
	Enabled     *bool  `json:"enabled"`
	UIEnabled   *bool  `json:"uiEnabled,omitempty"`
	MinReplica *int32 `json:"minReplica,omitempty"`
	Name       string `json:"name"`
	LastUpdate *int64 `json:"lastUpdate,omitempty"`
}

type EnvManagerStatus struct {
	ControlledBy string `json:"controlledBy,omitempty"`
}

type EnvManager struct {
	apimeta.TypeMeta   `json:",inline"`
	apimeta.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvManagerSpec   `json:"spec,omitempty"`
	Status EnvManagerStatus `json:"status,omitempty"`
}

type EnvManagerList struct {
	apimeta.TypeMeta   `json:",inline"`
	apimeta.ListMeta `json:"metadata,omitempty"`
	Items              []EnvManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnvManager{}, &EnvManagerList{})
}

// <---------- Starting EnvManagerList's DeepCopyObject and it's deepcopy functions ---------->
func (in *EnvManager) DeepCopyObject() runtime.Object {
	
	return in.DeepCopy()
}

func (in *EnvManager) DeepCopy() *EnvManager {
	if in == nil {
		return nil
	}
	out := new(EnvManager)
	*out = *in

	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)

	return out
}

func (in *EnvManagerSpec) DeepCopyInto(out *EnvManagerSpec) {
	*out = *in

	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		*out = *in
	}

	if in.MinReplica != nil {
		in, out := &in.MinReplica, &out.MinReplica
		*out = new(int32)
		*out = *in
	}

	if in.UIEnabled != nil {
		in, out := &in.UIEnabled, &out.UIEnabled
		*out = new(bool)
		*out = *in
	}

	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(int64)
		*out = *in
	}

}

func (in *EnvManagerStatus) DeepCopyInto(out *EnvManagerStatus) {
	*out = *in
}

// <---------- Starting EnvManagerList's DeepCopyObject and it's deepcopy functions ---------->
func (in *EnvManagerList) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}

	out := new(EnvManagerList)
	*out = *in

	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)

	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvManager, len(*in))
		for i := range *in {
			(*out)[i] = *((*in)[i].DeepCopy())
		}

	}

	return out
}

func (in *EnvManagerList) GetContinue() string{
	return in.ListMeta.Continue
}

func (in *EnvManagerList) SetContinue(c string) {
	in.Continue = c
}

func (in *EnvManagerList) GetRemainingItemCount() *int64{
	return in.ListMeta.RemainingItemCount
}

func (in *EnvManagerList) SetRemainingItemCount(c *int64) {
	in.RemainingItemCount = c
}

// func DeepCopyTypeMeta(in *apimeta.TypeMeta, out *apimeta.TypeMeta) {
// 	*out = *in

// }

// func DeepCopyObjectMeta(in *apimeta.ObjectMeta, out *apimeta.ObjectMeta) {
// 	*out = *in

// 	// Deepcopy CreationTimestamp

// 	in.CreationTimestamp.DeepCopyInto(out.DeletionTimestamp )
// 	fmt.Println("CreationTimestamp", &out.CreationTimestamp, &in.CreationTimestamp)

// 	// Deepcopy DeletionTimestamp
// 	if in.DeletionTimestamp != nil {
// 		out.DeletionTimestamp = in.DeletionTimestamp.DeepCopy()
// 		fmt.Println("DeletionTimestamp ", &out.DeletionTimestamp, &in.DeletionTimestamp)
// 	}

// 	// Deepcopy Labels
// 	if in.Labels != nil {
// 		in, out := &in.Labels, &out.Labels
// 		*out = make(map[string]string, len(*in))
// 		for key, val := range *in {
// 			(*out)[key] = val
// 		}
// 		fmt.Println("Labels ", &in, &out)
// 	}

// 	// Deepcopy Annotations
// 	if in.Annotations != nil {
// 		in, out := &in.Annotations, &out.Annotations
// 		*out = make(map[string]string, len(*in))
// 		for key, val := range *in {
// 			(*out)[key] = val
// 		}

// 	}

// 	// Deepcopy OwnerReferences
// 	if in.OwnerReferences != nil {
// 		in, out := &in.OwnerReferences, &out.OwnerReferences
// 		*out = make([]apimeta.OwnerReference, len(*in))

// 		for i := range *in {
// 			(*out)[i] = func(or *apimeta.OwnerReference) (apimeta.OwnerReference){
// 				orOut := new(apimeta.OwnerReference)
// 				*orOut = *or

// 				if or.Controller != nil {
// 					or, orOut := &or.Controller, &orOut.Controller
// 					*orOut = new(bool)
// 					**orOut = **or
// 				}

// 				if or.BlockOwnerDeletion != nil {
// 					or, orOut := &or.BlockOwnerDeletion, &orOut.BlockOwnerDeletion
// 					*orOut = new(bool)
// 					**orOut = **or
// 				}

// 				return *orOut
// 			}(&(*in)[i])
// 		}

// 	}

// 	// Deepcopy Finalizers
// 	if in.Finalizers != nil {
// 		in, out := &in.Finalizers, &out.Finalizers
// 		*out = make([]string, len(*in))
// 		copy(*out, *in)

// 	}

// 	// Deepcopy ManagedFields
// 	if in.ManagedFields != nil {
// 		in, out := &in.ManagedFields, &out.ManagedFields
// 		*out = make([]apimeta.ManagedFieldsEntry, len(*in))
// 		for i := range *in {
// 			(*out)[i] = func(mfi *apimeta.ManagedFieldsEntry) (apimeta.ManagedFieldsEntry){
// 				mfiOut := new(apimeta.ManagedFieldsEntry)
// 				*mfiOut = *mfi

// 				if mfi.Time != nil {
// 					*mfiOut.Time = *mfi.Time.DeepCopy()
// 				}

// 				if mfi.FieldsV1 != nil {
// 					mfi, mfiOut := &mfi.FieldsV1, &mfiOut.FieldsV1
// 					*mfiOut = new(apimeta.FieldsV1)
// 					**mfiOut = **mfi
// 					if (*mfi).Raw != nil {
// 						mfi, mfiOut := &(*mfi).Raw, &(*mfiOut).Raw
// 						*mfiOut = make([]byte, len((*mfi)))
// 						copy(*mfiOut, *mfi)
// 					}
// 				}

// 				return *mfiOut
// 			}(&(*in)[i])
// 		}

// 	}

// 	// Deepcopy DeletionGracePeriodSeconds
// 	if in.DeletionGracePeriodSeconds != nil {
// 		in, out := &in.DeletionGracePeriodSeconds, &out.DeletionGracePeriodSeconds
// 		*out = new(int64)
// 		**out = **in
// 	}

// }
