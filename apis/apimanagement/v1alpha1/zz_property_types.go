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

type PropertyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PropertyParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

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
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// PropertySpec defines the desired state of Property
type PropertySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PropertyParameters `json:"forProvider"`
}

// PropertyStatus defines the observed state of Property.
type PropertyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PropertyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Property is the Schema for the Propertys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Property struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PropertySpec   `json:"spec"`
	Status            PropertyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PropertyList contains a list of Propertys
type PropertyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Property `json:"items"`
}

// Repository type metadata.
var (
	Property_Kind             = "Property"
	Property_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Property_Kind}.String()
	Property_KindAPIVersion   = Property_Kind + "." + CRDGroupVersion.String()
	Property_GroupVersionKind = CRDGroupVersion.WithKind(Property_Kind)
)

func init() {
	SchemeBuilder.Register(&Property{}, &PropertyList{})
}