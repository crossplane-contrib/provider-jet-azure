//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainService) DeepCopyInto(out *ActiveDirectoryDomainService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainService.
func (in *ActiveDirectoryDomainService) DeepCopy() *ActiveDirectoryDomainService {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceList) DeepCopyInto(out *ActiveDirectoryDomainServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveDirectoryDomainService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceList.
func (in *ActiveDirectoryDomainServiceList) DeepCopy() *ActiveDirectoryDomainServiceList {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceObservation) DeepCopyInto(out *ActiveDirectoryDomainServiceObservation) {
	*out = *in
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SyncOwner != nil {
		in, out := &in.SyncOwner, &out.SyncOwner
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceObservation.
func (in *ActiveDirectoryDomainServiceObservation) DeepCopy() *ActiveDirectoryDomainServiceObservation {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceParameters) DeepCopyInto(out *ActiveDirectoryDomainServiceParameters) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.FilteredSyncEnabled != nil {
		in, out := &in.FilteredSyncEnabled, &out.FilteredSyncEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InitialReplicaSet != nil {
		in, out := &in.InitialReplicaSet, &out.InitialReplicaSet
		*out = make([]InitialReplicaSetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]NotificationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecureLdap != nil {
		in, out := &in.SecureLdap, &out.SecureLdap
		*out = make([]SecureLdapParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make([]SecurityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceParameters.
func (in *ActiveDirectoryDomainServiceParameters) DeepCopy() *ActiveDirectoryDomainServiceParameters {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSet) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSet.
func (in *ActiveDirectoryDomainServiceReplicaSet) DeepCopy() *ActiveDirectoryDomainServiceReplicaSet {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainServiceReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetList) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveDirectoryDomainServiceReplicaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetList.
func (in *ActiveDirectoryDomainServiceReplicaSetList) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainServiceReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetObservation) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetObservation) {
	*out = *in
	if in.DomainControllerIPAddresses != nil {
		in, out := &in.DomainControllerIPAddresses, &out.DomainControllerIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExternalAccessIPAddress != nil {
		in, out := &in.ExternalAccessIPAddress, &out.ExternalAccessIPAddress
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ServiceStatus != nil {
		in, out := &in.ServiceStatus, &out.ServiceStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetObservation.
func (in *ActiveDirectoryDomainServiceReplicaSetObservation) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetObservation {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetParameters) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetParameters) {
	*out = *in
	if in.DomainServiceID != nil {
		in, out := &in.DomainServiceID, &out.DomainServiceID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetParameters.
func (in *ActiveDirectoryDomainServiceReplicaSetParameters) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetParameters {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetSpec) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetSpec.
func (in *ActiveDirectoryDomainServiceReplicaSetSpec) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetStatus) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetStatus.
func (in *ActiveDirectoryDomainServiceReplicaSetStatus) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceSpec) DeepCopyInto(out *ActiveDirectoryDomainServiceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceSpec.
func (in *ActiveDirectoryDomainServiceSpec) DeepCopy() *ActiveDirectoryDomainServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceStatus) DeepCopyInto(out *ActiveDirectoryDomainServiceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceStatus.
func (in *ActiveDirectoryDomainServiceStatus) DeepCopy() *ActiveDirectoryDomainServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicaSetObservation) DeepCopyInto(out *InitialReplicaSetObservation) {
	*out = *in
	if in.DomainControllerIPAddresses != nil {
		in, out := &in.DomainControllerIPAddresses, &out.DomainControllerIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExternalAccessIPAddress != nil {
		in, out := &in.ExternalAccessIPAddress, &out.ExternalAccessIPAddress
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ServiceStatus != nil {
		in, out := &in.ServiceStatus, &out.ServiceStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicaSetObservation.
func (in *InitialReplicaSetObservation) DeepCopy() *InitialReplicaSetObservation {
	if in == nil {
		return nil
	}
	out := new(InitialReplicaSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicaSetParameters) DeepCopyInto(out *InitialReplicaSetParameters) {
	*out = *in
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicaSetParameters.
func (in *InitialReplicaSetParameters) DeepCopy() *InitialReplicaSetParameters {
	if in == nil {
		return nil
	}
	out := new(InitialReplicaSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsObservation) DeepCopyInto(out *NotificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsObservation.
func (in *NotificationsObservation) DeepCopy() *NotificationsObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsParameters) DeepCopyInto(out *NotificationsParameters) {
	*out = *in
	if in.AdditionalRecipients != nil {
		in, out := &in.AdditionalRecipients, &out.AdditionalRecipients
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotifyDcAdmins != nil {
		in, out := &in.NotifyDcAdmins, &out.NotifyDcAdmins
		*out = new(bool)
		**out = **in
	}
	if in.NotifyGlobalAdmins != nil {
		in, out := &in.NotifyGlobalAdmins, &out.NotifyGlobalAdmins
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsParameters.
func (in *NotificationsParameters) DeepCopy() *NotificationsParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureLdapObservation) DeepCopyInto(out *SecureLdapObservation) {
	*out = *in
	if in.CertificateExpiry != nil {
		in, out := &in.CertificateExpiry, &out.CertificateExpiry
		*out = new(string)
		**out = **in
	}
	if in.CertificateThumbprint != nil {
		in, out := &in.CertificateThumbprint, &out.CertificateThumbprint
		*out = new(string)
		**out = **in
	}
	if in.PublicCertificate != nil {
		in, out := &in.PublicCertificate, &out.PublicCertificate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureLdapObservation.
func (in *SecureLdapObservation) DeepCopy() *SecureLdapObservation {
	if in == nil {
		return nil
	}
	out := new(SecureLdapObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureLdapParameters) DeepCopyInto(out *SecureLdapParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExternalAccessEnabled != nil {
		in, out := &in.ExternalAccessEnabled, &out.ExternalAccessEnabled
		*out = new(bool)
		**out = **in
	}
	out.PfxCertificatePasswordSecretRef = in.PfxCertificatePasswordSecretRef
	out.PfxCertificateSecretRef = in.PfxCertificateSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureLdapParameters.
func (in *SecureLdapParameters) DeepCopy() *SecureLdapParameters {
	if in == nil {
		return nil
	}
	out := new(SecureLdapParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityObservation) DeepCopyInto(out *SecurityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityObservation.
func (in *SecurityObservation) DeepCopy() *SecurityObservation {
	if in == nil {
		return nil
	}
	out := new(SecurityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityParameters) DeepCopyInto(out *SecurityParameters) {
	*out = *in
	if in.NtlmV1Enabled != nil {
		in, out := &in.NtlmV1Enabled, &out.NtlmV1Enabled
		*out = new(bool)
		**out = **in
	}
	if in.SyncKerberosPasswords != nil {
		in, out := &in.SyncKerberosPasswords, &out.SyncKerberosPasswords
		*out = new(bool)
		**out = **in
	}
	if in.SyncNtlmPasswords != nil {
		in, out := &in.SyncNtlmPasswords, &out.SyncNtlmPasswords
		*out = new(bool)
		**out = **in
	}
	if in.SyncOnPremPasswords != nil {
		in, out := &in.SyncOnPremPasswords, &out.SyncOnPremPasswords
		*out = new(bool)
		**out = **in
	}
	if in.TLSV1Enabled != nil {
		in, out := &in.TLSV1Enabled, &out.TLSV1Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityParameters.
func (in *SecurityParameters) DeepCopy() *SecurityParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityParameters)
	in.DeepCopyInto(out)
	return out
}
