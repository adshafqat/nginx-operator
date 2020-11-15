/*
Copyright 2020 Adeel Shafqat.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ZaynNginxSpec defines the desired state of ZaynNginx
type ZaynNginxSpec struct {
	Replicas int32 `json:"replicas"`
}

// ZaynNginxStatus defines the observed state of ZaynNginx
type ZaynNginxStatus struct {
	PodNames []string `json:"podNames"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ZaynNginx is the Schema for the zaynnginxes API
type ZaynNginx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZaynNginxSpec   `json:"spec,omitempty"`
	Status ZaynNginxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZaynNginxList contains a list of ZaynNginx
type ZaynNginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZaynNginx `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZaynNginx{}, &ZaynNginxList{})
}
