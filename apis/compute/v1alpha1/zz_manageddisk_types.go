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

type DiskEncryptionKeyObservation struct {
}

type DiskEncryptionKeyParameters struct {

	// +kubebuilder:validation:Required
	SecretURL *string `json:"secretUrl" tf:"secret_url,omitempty"`

	// +kubebuilder:validation:Required
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsObservation struct {
}

type EncryptionSettingsParameters struct {

	// +kubebuilder:validation:Optional
	DiskEncryptionKey []DiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyEncryptionKey []KeyEncryptionKeyParameters `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type KeyEncryptionKeyObservation struct {
}

type KeyEncryptionKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyURL *string `json:"keyUrl" tf:"key_url,omitempty"`

	// +kubebuilder:validation:Required
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type ManagedDiskObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagedDiskParameters struct {

	// +kubebuilder:validation:Required
	CreateOption *string `json:"createOption" tf:"create_option,omitempty"`

	// +kubebuilder:validation:Optional
	DiskAccessID *string `json:"diskAccessId,omitempty" tf:"disk_access_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskIopsReadWrite *float64 `json:"diskIopsReadWrite,omitempty" tf:"disk_iops_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	DiskMbpsReadWrite *float64 `json:"diskMbpsReadWrite,omitempty" tf:"disk_mbps_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionSettings []EncryptionSettingsParameters `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`

	// +kubebuilder:validation:Optional
	ImageReferenceID *string `json:"imageReferenceId,omitempty" tf:"image_reference_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkAccessPolicy *string `json:"networkAccessPolicy,omitempty" tf:"network_access_policy,omitempty"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// ManagedDiskSpec defines the desired state of ManagedDisk
type ManagedDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedDiskParameters `json:"forProvider"`
}

// ManagedDiskStatus defines the observed state of ManagedDisk.
type ManagedDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDisk is the Schema for the ManagedDisks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagedDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedDiskSpec   `json:"spec"`
	Status            ManagedDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDiskList contains a list of ManagedDisks
type ManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedDisk `json:"items"`
}

// Repository type metadata.
var (
	ManagedDisk_Kind             = "ManagedDisk"
	ManagedDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedDisk_Kind}.String()
	ManagedDisk_KindAPIVersion   = ManagedDisk_Kind + "." + CRDGroupVersion.String()
	ManagedDisk_GroupVersionKind = CRDGroupVersion.WithKind(ManagedDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedDisk{}, &ManagedDiskList{})
}