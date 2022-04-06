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

type ManagedDiskObservation struct {
}

type ManagedDiskParameters struct {

	// +kubebuilder:validation:Optional
	DiskID *string `json:"diskId,omitempty" tf:"disk_id"`

	// +kubebuilder:validation:Optional
	StagingStorageAccountID *string `json:"stagingStorageAccountId,omitempty" tf:"staging_storage_account_id"`

	// +kubebuilder:validation:Optional
	TargetDiskEncryptionSetID *string `json:"targetDiskEncryptionSetId,omitempty" tf:"target_disk_encryption_set_id"`

	// +kubebuilder:validation:Optional
	TargetDiskType *string `json:"targetDiskType,omitempty" tf:"target_disk_type"`

	// +kubebuilder:validation:Optional
	TargetReplicaDiskType *string `json:"targetReplicaDiskType,omitempty" tf:"target_replica_disk_type"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceGroupID *string `json:"targetResourceGroupId,omitempty" tf:"target_resource_group_id"`

	// +kubebuilder:validation:Optional
	TargetResourceGroupIDRef *v1.Reference `json:"targetResourceGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetResourceGroupIDSelector *v1.Selector `json:"targetResourceGroupIdSelector,omitempty" tf:"-"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	RecoveryPublicIPAddressID *string `json:"recoveryPublicIpAddressId,omitempty" tf:"recovery_public_ip_address_id"`

	// +kubebuilder:validation:Optional
	SourceNetworkInterfaceID *string `json:"sourceNetworkInterfaceId,omitempty" tf:"source_network_interface_id"`

	// +kubebuilder:validation:Optional
	TargetStaticIP *string `json:"targetStaticIp,omitempty" tf:"target_static_ip"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +kubebuilder:validation:Optional
	TargetSubnetName *string `json:"targetSubnetName,omitempty" tf:"target_subnet_name"`

	// +kubebuilder:validation:Optional
	TargetSubnetNameRef *v1.Reference `json:"targetSubnetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetSubnetNameSelector *v1.Selector `json:"targetSubnetNameSelector,omitempty" tf:"-"`
}

type SiteRecoveryReplicatedVMObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SiteRecoveryReplicatedVMParameters struct {

	// +kubebuilder:validation:Optional
	ManagedDisk []ManagedDiskParameters `json:"managedDisk,omitempty" tf:"managed_disk,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryReplicationPolicyID *string `json:"recoveryReplicationPolicyId" tf:"recovery_replication_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SourceRecoveryFabricName *string `json:"sourceRecoveryFabricName" tf:"source_recovery_fabric_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceRecoveryProtectionContainerName *string `json:"sourceRecoveryProtectionContainerName" tf:"source_recovery_protection_container_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceVMID *string `json:"sourceVmId" tf:"source_vm_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetAvailabilitySetID *string `json:"targetAvailabilitySetId,omitempty" tf:"target_availability_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetNetworkID *string `json:"targetNetworkId,omitempty" tf:"target_network_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetRecoveryFabricID *string `json:"targetRecoveryFabricId" tf:"target_recovery_fabric_id,omitempty"`

	// +kubebuilder:validation:Required
	TargetRecoveryProtectionContainerID *string `json:"targetRecoveryProtectionContainerId" tf:"target_recovery_protection_container_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceGroupID *string `json:"targetResourceGroupId,omitempty" tf:"target_resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceGroupIDRef *v1.Reference `json:"targetResourceGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetResourceGroupIDSelector *v1.Selector `json:"targetResourceGroupIdSelector,omitempty" tf:"-"`
}

// SiteRecoveryReplicatedVMSpec defines the desired state of SiteRecoveryReplicatedVM
type SiteRecoveryReplicatedVMSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryReplicatedVMParameters `json:"forProvider"`
}

// SiteRecoveryReplicatedVMStatus defines the observed state of SiteRecoveryReplicatedVM.
type SiteRecoveryReplicatedVMStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryReplicatedVMObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryReplicatedVM is the Schema for the SiteRecoveryReplicatedVMs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SiteRecoveryReplicatedVM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryReplicatedVMSpec   `json:"spec"`
	Status            SiteRecoveryReplicatedVMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryReplicatedVMList contains a list of SiteRecoveryReplicatedVMs
type SiteRecoveryReplicatedVMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryReplicatedVM `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryReplicatedVM_Kind             = "SiteRecoveryReplicatedVM"
	SiteRecoveryReplicatedVM_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryReplicatedVM_Kind}.String()
	SiteRecoveryReplicatedVM_KindAPIVersion   = SiteRecoveryReplicatedVM_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryReplicatedVM_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryReplicatedVM_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryReplicatedVM{}, &SiteRecoveryReplicatedVMList{})
}
