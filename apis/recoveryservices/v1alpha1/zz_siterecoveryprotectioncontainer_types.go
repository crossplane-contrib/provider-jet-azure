/*
Copyright 2020 The Crossplane Authors.

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

type SiteRecoveryProtectionContainerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SiteRecoveryProtectionContainerParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryFabricName *string `json:"recoveryFabricName" tf:"recovery_fabric_name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// SiteRecoveryProtectionContainerSpec defines the desired state of SiteRecoveryProtectionContainer
type SiteRecoveryProtectionContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryProtectionContainerParameters `json:"forProvider"`
}

// SiteRecoveryProtectionContainerStatus defines the observed state of SiteRecoveryProtectionContainer.
type SiteRecoveryProtectionContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryProtectionContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainer is the Schema for the SiteRecoveryProtectionContainers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SiteRecoveryProtectionContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryProtectionContainerSpec   `json:"spec"`
	Status            SiteRecoveryProtectionContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainerList contains a list of SiteRecoveryProtectionContainers
type SiteRecoveryProtectionContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryProtectionContainer `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryProtectionContainer_Kind             = "SiteRecoveryProtectionContainer"
	SiteRecoveryProtectionContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryProtectionContainer_Kind}.String()
	SiteRecoveryProtectionContainer_KindAPIVersion   = SiteRecoveryProtectionContainer_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryProtectionContainer_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryProtectionContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryProtectionContainer{}, &SiteRecoveryProtectionContainerList{})
}