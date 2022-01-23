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

type StaticSiteObservation struct {
	APIKey *string `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	DefaultHostName *string `json:"defaultHostName,omitempty" tf:"default_host_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StaticSiteParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SkuSize *string `json:"skuSize,omitempty" tf:"sku_size,omitempty"`

	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StaticSiteSpec defines the desired state of StaticSite
type StaticSiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StaticSiteParameters `json:"forProvider"`
}

// StaticSiteStatus defines the observed state of StaticSite.
type StaticSiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StaticSiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StaticSite is the Schema for the StaticSites API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type StaticSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StaticSiteSpec   `json:"spec"`
	Status            StaticSiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticSiteList contains a list of StaticSites
type StaticSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticSite `json:"items"`
}

// Repository type metadata.
var (
	StaticSite_Kind             = "StaticSite"
	StaticSite_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StaticSite_Kind}.String()
	StaticSite_KindAPIVersion   = StaticSite_Kind + "." + CRDGroupVersion.String()
	StaticSite_GroupVersionKind = CRDGroupVersion.WithKind(StaticSite_Kind)
)

func init() {
	SchemeBuilder.Register(&StaticSite{}, &StaticSiteList{})
}