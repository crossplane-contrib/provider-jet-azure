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

type NetworkPeeringObservation struct {
}

type NetworkPeeringParameters struct {

	// +kubebuilder:validation:Optional
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty" tf:"allow_forwarded_traffic,omitempty"`

	// +kubebuilder:validation:Optional
	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty" tf:"allow_gateway_transit,omitempty"`

	// +kubebuilder:validation:Optional
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty" tf:"allow_virtual_network_access,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	RemoteVirtualNetworkID *string `json:"remoteVirtualNetworkId,omitempty" tf:"remote_virtual_network_id,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteVirtualNetworkIDRef *v1.Reference `json:"remoteVirtualNetworkIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RemoteVirtualNetworkIDSelector *v1.Selector `json:"remoteVirtualNetworkIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UseRemoteGateways *bool `json:"useRemoteGateways,omitempty" tf:"use_remote_gateways,omitempty"`

	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	VirtualNetworkName *string `json:"virtualNetworkName,omitempty" tf:"virtual_network_name,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkNameRef *v1.Reference `json:"virtualNetworkNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualNetworkNameSelector *v1.Selector `json:"virtualNetworkNameSelector,omitempty" tf:"-"`
}

// NetworkPeeringSpec defines the desired state of NetworkPeering
type NetworkPeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkPeeringParameters `json:"forProvider"`
}

// NetworkPeeringStatus defines the observed state of NetworkPeering.
type NetworkPeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkPeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPeering is the Schema for the NetworkPeerings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type NetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPeeringSpec   `json:"spec"`
	Status            NetworkPeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPeeringList contains a list of NetworkPeerings
type NetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPeering `json:"items"`
}

// Repository type metadata.
var (
	NetworkPeeringKind             = "NetworkPeering"
	NetworkPeeringGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkPeeringKind}.String()
	NetworkPeeringKindAPIVersion   = NetworkPeeringKind + "." + GroupVersion.String()
	NetworkPeeringGroupVersionKind = GroupVersion.WithKind(NetworkPeeringKind)
)

func init() {
	SchemeBuilder.Register(&NetworkPeering{}, &NetworkPeeringList{})
}
