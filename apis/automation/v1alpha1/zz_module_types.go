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

type HashObservation struct {
}

type HashParameters struct {

	// +kubebuilder:validation:Required
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ModuleLinkObservation struct {
}

type ModuleLinkParameters struct {

	// +kubebuilder:validation:Optional
	Hash []HashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ModuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ModuleParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Required
	ModuleLink []ModuleLinkParameters `json:"moduleLink" tf:"module_link,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// ModuleSpec defines the desired state of Module
type ModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModuleParameters `json:"forProvider"`
}

// ModuleStatus defines the observed state of Module.
type ModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Module is the Schema for the Modules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Module struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModuleSpec   `json:"spec"`
	Status            ModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModuleList contains a list of Modules
type ModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Module `json:"items"`
}

// Repository type metadata.
var (
	Module_Kind             = "Module"
	Module_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Module_Kind}.String()
	Module_KindAPIVersion   = Module_Kind + "." + CRDGroupVersion.String()
	Module_GroupVersionKind = CRDGroupVersion.WithKind(Module_Kind)
)

func init() {
	SchemeBuilder.Register(&Module{}, &ModuleList{})
}
