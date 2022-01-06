// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureadAdministratorObservation) DeepCopyInto(out *AzureadAdministratorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureadAdministratorObservation.
func (in *AzureadAdministratorObservation) DeepCopy() *AzureadAdministratorObservation {
	if in == nil {
		return nil
	}
	out := new(AzureadAdministratorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureadAdministratorParameters) DeepCopyInto(out *AzureadAdministratorParameters) {
	*out = *in
	if in.LoginUsername != nil {
		in, out := &in.LoginUsername, &out.LoginUsername
		*out = new(string)
		**out = **in
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureadAdministratorParameters.
func (in *AzureadAdministratorParameters) DeepCopy() *AzureadAdministratorParameters {
	if in == nil {
		return nil
	}
	out := new(AzureadAdministratorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedAuditingPolicyObservation) DeepCopyInto(out *ExtendedAuditingPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedAuditingPolicyObservation.
func (in *ExtendedAuditingPolicyObservation) DeepCopy() *ExtendedAuditingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ExtendedAuditingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedAuditingPolicyParameters) DeepCopyInto(out *ExtendedAuditingPolicyParameters) {
	*out = *in
	if in.LogMonitoringEnabled != nil {
		in, out := &in.LogMonitoringEnabled, &out.LogMonitoringEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int64)
		**out = **in
	}
	if in.StorageAccountAccessKeyIsSecondary != nil {
		in, out := &in.StorageAccountAccessKeyIsSecondary, &out.StorageAccountAccessKeyIsSecondary
		*out = new(bool)
		**out = **in
	}
	if in.StorageAccountAccessKeySecretRef != nil {
		in, out := &in.StorageAccountAccessKeySecretRef, &out.StorageAccountAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.StorageEndpoint != nil {
		in, out := &in.StorageEndpoint, &out.StorageEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedAuditingPolicyParameters.
func (in *ExtendedAuditingPolicyParameters) DeepCopy() *ExtendedAuditingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ExtendedAuditingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServer) DeepCopyInto(out *MSSQLServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServer.
func (in *MSSQLServer) DeepCopy() *MSSQLServer {
	if in == nil {
		return nil
	}
	out := new(MSSQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MSSQLServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerList) DeepCopyInto(out *MSSQLServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MSSQLServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerList.
func (in *MSSQLServerList) DeepCopy() *MSSQLServerList {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MSSQLServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerObservation) DeepCopyInto(out *MSSQLServerObservation) {
	*out = *in
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RestorableDroppedDatabaseIds != nil {
		in, out := &in.RestorableDroppedDatabaseIds, &out.RestorableDroppedDatabaseIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerObservation.
func (in *MSSQLServerObservation) DeepCopy() *MSSQLServerObservation {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerParameters) DeepCopyInto(out *MSSQLServerParameters) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	out.AdministratorLoginPasswordSecretRef = in.AdministratorLoginPasswordSecretRef
	if in.AzureadAdministrator != nil {
		in, out := &in.AzureadAdministrator, &out.AzureadAdministrator
		*out = make([]AzureadAdministratorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectionPolicy != nil {
		in, out := &in.ConnectionPolicy, &out.ConnectionPolicy
		*out = new(string)
		**out = **in
	}
	if in.ExtendedAuditingPolicy != nil {
		in, out := &in.ExtendedAuditingPolicy, &out.ExtendedAuditingPolicy
		*out = make([]ExtendedAuditingPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerParameters.
func (in *MSSQLServerParameters) DeepCopy() *MSSQLServerParameters {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerSpec) DeepCopyInto(out *MSSQLServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerSpec.
func (in *MSSQLServerSpec) DeepCopy() *MSSQLServerSpec {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerStatus) DeepCopyInto(out *MSSQLServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerStatus.
func (in *MSSQLServerStatus) DeepCopy() *MSSQLServerStatus {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerExtendedAuditingPolicyObservation) DeepCopyInto(out *ServerExtendedAuditingPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerExtendedAuditingPolicyObservation.
func (in *ServerExtendedAuditingPolicyObservation) DeepCopy() *ServerExtendedAuditingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ServerExtendedAuditingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerExtendedAuditingPolicyParameters) DeepCopyInto(out *ServerExtendedAuditingPolicyParameters) {
	*out = *in
	if in.LogMonitoringEnabled != nil {
		in, out := &in.LogMonitoringEnabled, &out.LogMonitoringEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int64)
		**out = **in
	}
	if in.StorageAccountAccessKeyIsSecondary != nil {
		in, out := &in.StorageAccountAccessKeyIsSecondary, &out.StorageAccountAccessKeyIsSecondary
		*out = new(bool)
		**out = **in
	}
	if in.StorageAccountAccessKeySecretRef != nil {
		in, out := &in.StorageAccountAccessKeySecretRef, &out.StorageAccountAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.StorageEndpoint != nil {
		in, out := &in.StorageEndpoint, &out.StorageEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerExtendedAuditingPolicyParameters.
func (in *ServerExtendedAuditingPolicyParameters) DeepCopy() *ServerExtendedAuditingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ServerExtendedAuditingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerIdentityObservation) DeepCopyInto(out *ServerIdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerIdentityObservation.
func (in *ServerIdentityObservation) DeepCopy() *ServerIdentityObservation {
	if in == nil {
		return nil
	}
	out := new(ServerIdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerIdentityParameters) DeepCopyInto(out *ServerIdentityParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerIdentityParameters.
func (in *ServerIdentityParameters) DeepCopy() *ServerIdentityParameters {
	if in == nil {
		return nil
	}
	out := new(ServerIdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerObservation) DeepCopyInto(out *ServerObservation) {
	*out = *in
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerObservation.
func (in *ServerObservation) DeepCopy() *ServerObservation {
	if in == nil {
		return nil
	}
	out := new(ServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerParameters) DeepCopyInto(out *ServerParameters) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	out.AdministratorLoginPasswordSecretRef = in.AdministratorLoginPasswordSecretRef
	if in.ConnectionPolicy != nil {
		in, out := &in.ConnectionPolicy, &out.ConnectionPolicy
		*out = new(string)
		**out = **in
	}
	if in.ExtendedAuditingPolicy != nil {
		in, out := &in.ExtendedAuditingPolicy, &out.ExtendedAuditingPolicy
		*out = make([]ServerExtendedAuditingPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]ServerIdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
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
	if in.ThreatDetectionPolicy != nil {
		in, out := &in.ThreatDetectionPolicy, &out.ThreatDetectionPolicy
		*out = make([]ThreatDetectionPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerParameters.
func (in *ServerParameters) DeepCopy() *ServerParameters {
	if in == nil {
		return nil
	}
	out := new(ServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThreatDetectionPolicyObservation) DeepCopyInto(out *ThreatDetectionPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThreatDetectionPolicyObservation.
func (in *ThreatDetectionPolicyObservation) DeepCopy() *ThreatDetectionPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ThreatDetectionPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThreatDetectionPolicyParameters) DeepCopyInto(out *ThreatDetectionPolicyParameters) {
	*out = *in
	if in.DisabledAlerts != nil {
		in, out := &in.DisabledAlerts, &out.DisabledAlerts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EmailAccountAdmins != nil {
		in, out := &in.EmailAccountAdmins, &out.EmailAccountAdmins
		*out = new(bool)
		**out = **in
	}
	if in.EmailAddresses != nil {
		in, out := &in.EmailAddresses, &out.EmailAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RetentionDays != nil {
		in, out := &in.RetentionDays, &out.RetentionDays
		*out = new(int64)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountAccessKeySecretRef != nil {
		in, out := &in.StorageAccountAccessKeySecretRef, &out.StorageAccountAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.StorageEndpoint != nil {
		in, out := &in.StorageEndpoint, &out.StorageEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThreatDetectionPolicyParameters.
func (in *ThreatDetectionPolicyParameters) DeepCopy() *ThreatDetectionPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ThreatDetectionPolicyParameters)
	in.DeepCopyInto(out)
	return out
}
