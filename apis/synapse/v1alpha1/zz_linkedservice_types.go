/*
Copyright 2022 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IntegrationRuntimeObservation struct {
}

type IntegrationRuntimeParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkedServiceParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntime []IntegrationRuntimeParameters `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	SynapseWorkspaceID *string `json:"synapseWorkspaceId" tf:"synapse_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	TypePropertiesJSON *string `json:"typePropertiesJson" tf:"type_properties_json,omitempty"`
}

// LinkedServiceSpec defines the desired state of LinkedService
type LinkedServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceParameters `json:"forProvider"`
}

// LinkedServiceStatus defines the observed state of LinkedService.
type LinkedServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedService is the Schema for the LinkedServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LinkedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceSpec   `json:"spec"`
	Status            LinkedServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceList contains a list of LinkedServices
type LinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedService `json:"items"`
}

// Repository type metadata.
var (
	LinkedService_Kind             = "LinkedService"
	LinkedService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedService_Kind}.String()
	LinkedService_KindAPIVersion   = LinkedService_Kind + "." + CRDGroupVersion.String()
	LinkedService_GroupVersionKind = CRDGroupVersion.WithKind(LinkedService_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedService{}, &LinkedServiceList{})
}