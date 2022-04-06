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

type HPCCacheBlobNFSTargetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HPCCacheBlobNFSTargetParameters struct {

	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// +kubebuilder:validation:Required
	CacheName *string `json:"cacheName" tf:"cache_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespacePath *string `json:"namespacePath" tf:"namespace_path,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	StorageContainerID *string `json:"storageContainerId" tf:"storage_container_id,omitempty"`

	// +kubebuilder:validation:Required
	UsageModel *string `json:"usageModel" tf:"usage_model,omitempty"`
}

// HPCCacheBlobNFSTargetSpec defines the desired state of HPCCacheBlobNFSTarget
type HPCCacheBlobNFSTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheBlobNFSTargetParameters `json:"forProvider"`
}

// HPCCacheBlobNFSTargetStatus defines the observed state of HPCCacheBlobNFSTarget.
type HPCCacheBlobNFSTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheBlobNFSTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobNFSTarget is the Schema for the HPCCacheBlobNFSTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type HPCCacheBlobNFSTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HPCCacheBlobNFSTargetSpec   `json:"spec"`
	Status            HPCCacheBlobNFSTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobNFSTargetList contains a list of HPCCacheBlobNFSTargets
type HPCCacheBlobNFSTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCacheBlobNFSTarget `json:"items"`
}

// Repository type metadata.
var (
	HPCCacheBlobNFSTarget_Kind             = "HPCCacheBlobNFSTarget"
	HPCCacheBlobNFSTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCacheBlobNFSTarget_Kind}.String()
	HPCCacheBlobNFSTarget_KindAPIVersion   = HPCCacheBlobNFSTarget_Kind + "." + CRDGroupVersion.String()
	HPCCacheBlobNFSTarget_GroupVersionKind = CRDGroupVersion.WithKind(HPCCacheBlobNFSTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCacheBlobNFSTarget{}, &HPCCacheBlobNFSTargetList{})
}
