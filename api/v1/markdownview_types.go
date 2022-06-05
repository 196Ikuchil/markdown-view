/*
Copyright 2022.

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

// MarkdownViewSpec defines the desired state of MarkdownView
type MarkdownViewSpec struct {
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:MinProperties=1
	Markdowns map[string]string `json:"markdowns,omitempty"`

	// Replicas is the number of viewrs.
	//+kubebuilder:default=1
	//+optional
	Replicas int32 `json:"replicas,omitempty"`

	//+optional
	ViewerImage string `json:"viewerImage,omitempty"`
}

// MarkdownViewStatus defines the observed state of MarkdownView
//+kubebuilder:validation:Enum=NotReady;Available;Healthy
type MarkdownViewStatus string

const (
	MarkdownViewNotReady  = MarkdownViewStatus("NotReady")
	MarkdownViewAvailable = MarkdownViewStatus("Available")
	MarkdownViewHealthy   = MarkdownViewStatus("Healthy")
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="REPLOCAS",type="integer",JSONPath=".spec.replicas"
//+kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status"

// MarkdownView is the Schema for the markdownviews API
type MarkdownView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MarkdownViewSpec   `json:"spec,omitempty"`
	Status MarkdownViewStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MarkdownViewList contains a list of MarkdownView
type MarkdownViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MarkdownView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MarkdownView{}, &MarkdownViewList{})
}
