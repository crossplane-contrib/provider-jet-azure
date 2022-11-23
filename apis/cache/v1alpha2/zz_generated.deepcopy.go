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

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleObservation) DeepCopyInto(out *ModuleObservation) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleObservation.
func (in *ModuleObservation) DeepCopy() *ModuleObservation {
	if in == nil {
		return nil
	}
	out := new(ModuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleParameters) DeepCopyInto(out *ModuleParameters) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleParameters.
func (in *ModuleParameters) DeepCopy() *ModuleParameters {
	if in == nil {
		return nil
	}
	out := new(ModuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchScheduleObservation) DeepCopyInto(out *PatchScheduleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchScheduleObservation.
func (in *PatchScheduleObservation) DeepCopy() *PatchScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(PatchScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchScheduleParameters) DeepCopyInto(out *PatchScheduleParameters) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchScheduleParameters.
func (in *PatchScheduleParameters) DeepCopy() *PatchScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(PatchScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCache) DeepCopyInto(out *RedisCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCache.
func (in *RedisCache) DeepCopy() *RedisCache {
	if in == nil {
		return nil
	}
	out := new(RedisCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheList) DeepCopyInto(out *RedisCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheList.
func (in *RedisCacheList) DeepCopy() *RedisCacheList {
	if in == nil {
		return nil
	}
	out := new(RedisCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisCacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheObservation) DeepCopyInto(out *RedisCacheObservation) {
	*out = *in
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = make([]RedisConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSLPort != nil {
		in, out := &in.SSLPort, &out.SSLPort
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheObservation.
func (in *RedisCacheObservation) DeepCopy() *RedisCacheObservation {
	if in == nil {
		return nil
	}
	out := new(RedisCacheObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheParameters) DeepCopyInto(out *RedisCacheParameters) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.EnableNonSSLPort != nil {
		in, out := &in.EnableNonSSLPort, &out.EnableNonSSLPort
		*out = new(bool)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
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
	if in.PatchSchedule != nil {
		in, out := &in.PatchSchedule, &out.PatchSchedule
		*out = make([]PatchScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateStaticIPAddress != nil {
		in, out := &in.PrivateStaticIPAddress, &out.PrivateStaticIPAddress
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = make([]RedisConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(float64)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(float64)
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
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(float64)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
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
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
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
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
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
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheParameters.
func (in *RedisCacheParameters) DeepCopy() *RedisCacheParameters {
	if in == nil {
		return nil
	}
	out := new(RedisCacheParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheSpec) DeepCopyInto(out *RedisCacheSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheSpec.
func (in *RedisCacheSpec) DeepCopy() *RedisCacheSpec {
	if in == nil {
		return nil
	}
	out := new(RedisCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheStatus) DeepCopyInto(out *RedisCacheStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheStatus.
func (in *RedisCacheStatus) DeepCopy() *RedisCacheStatus {
	if in == nil {
		return nil
	}
	out := new(RedisCacheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfigurationObservation) DeepCopyInto(out *RedisConfigurationObservation) {
	*out = *in
	if in.Maxclients != nil {
		in, out := &in.Maxclients, &out.Maxclients
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfigurationObservation.
func (in *RedisConfigurationObservation) DeepCopy() *RedisConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(RedisConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfigurationParameters) DeepCopyInto(out *RedisConfigurationParameters) {
	*out = *in
	if in.AofBackupEnabled != nil {
		in, out := &in.AofBackupEnabled, &out.AofBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofStorageConnectionString0SecretRef != nil {
		in, out := &in.AofStorageConnectionString0SecretRef, &out.AofStorageConnectionString0SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AofStorageConnectionString1SecretRef != nil {
		in, out := &in.AofStorageConnectionString1SecretRef, &out.AofStorageConnectionString1SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.EnableAuthentication != nil {
		in, out := &in.EnableAuthentication, &out.EnableAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.NotifyKeySpaceEvents != nil {
		in, out := &in.NotifyKeySpaceEvents, &out.NotifyKeySpaceEvents
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(float64)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(float64)
		**out = **in
	}
	if in.RdbStorageConnectionStringSecretRef != nil {
		in, out := &in.RdbStorageConnectionStringSecretRef, &out.RdbStorageConnectionStringSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfigurationParameters.
func (in *RedisConfigurationParameters) DeepCopy() *RedisConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(RedisConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseCluster) DeepCopyInto(out *RedisEnterpriseCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseCluster.
func (in *RedisEnterpriseCluster) DeepCopy() *RedisEnterpriseCluster {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseClusterList) DeepCopyInto(out *RedisEnterpriseClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterpriseCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseClusterList.
func (in *RedisEnterpriseClusterList) DeepCopy() *RedisEnterpriseClusterList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseClusterObservation) DeepCopyInto(out *RedisEnterpriseClusterObservation) {
	*out = *in
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseClusterObservation.
func (in *RedisEnterpriseClusterObservation) DeepCopy() *RedisEnterpriseClusterObservation {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseClusterParameters) DeepCopyInto(out *RedisEnterpriseClusterParameters) {
	*out = *in
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
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
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
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseClusterParameters.
func (in *RedisEnterpriseClusterParameters) DeepCopy() *RedisEnterpriseClusterParameters {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseClusterSpec) DeepCopyInto(out *RedisEnterpriseClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseClusterSpec.
func (in *RedisEnterpriseClusterSpec) DeepCopy() *RedisEnterpriseClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseClusterStatus) DeepCopyInto(out *RedisEnterpriseClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseClusterStatus.
func (in *RedisEnterpriseClusterStatus) DeepCopy() *RedisEnterpriseClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabase) DeepCopyInto(out *RedisEnterpriseDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabase.
func (in *RedisEnterpriseDatabase) DeepCopy() *RedisEnterpriseDatabase {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseList) DeepCopyInto(out *RedisEnterpriseDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterpriseDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseList.
func (in *RedisEnterpriseDatabaseList) DeepCopy() *RedisEnterpriseDatabaseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseObservation) DeepCopyInto(out *RedisEnterpriseDatabaseObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = make([]ModuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseObservation.
func (in *RedisEnterpriseDatabaseObservation) DeepCopy() *RedisEnterpriseDatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseParameters) DeepCopyInto(out *RedisEnterpriseDatabaseParameters) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(string)
		**out = **in
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(string)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(string)
		**out = **in
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = make([]ModuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
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
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseParameters.
func (in *RedisEnterpriseDatabaseParameters) DeepCopy() *RedisEnterpriseDatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseSpec) DeepCopyInto(out *RedisEnterpriseDatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseSpec.
func (in *RedisEnterpriseDatabaseSpec) DeepCopy() *RedisEnterpriseDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseStatus) DeepCopyInto(out *RedisEnterpriseDatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseStatus.
func (in *RedisEnterpriseDatabaseStatus) DeepCopy() *RedisEnterpriseDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRule) DeepCopyInto(out *RedisFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRule.
func (in *RedisFirewallRule) DeepCopy() *RedisFirewallRule {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleList) DeepCopyInto(out *RedisFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleList.
func (in *RedisFirewallRuleList) DeepCopy() *RedisFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleObservation) DeepCopyInto(out *RedisFirewallRuleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleObservation.
func (in *RedisFirewallRuleObservation) DeepCopy() *RedisFirewallRuleObservation {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleParameters) DeepCopyInto(out *RedisFirewallRuleParameters) {
	*out = *in
	if in.EndIP != nil {
		in, out := &in.EndIP, &out.EndIP
		*out = new(string)
		**out = **in
	}
	if in.RedisCacheName != nil {
		in, out := &in.RedisCacheName, &out.RedisCacheName
		*out = new(string)
		**out = **in
	}
	if in.RedisCacheNameRef != nil {
		in, out := &in.RedisCacheNameRef, &out.RedisCacheNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisCacheNameSelector != nil {
		in, out := &in.RedisCacheNameSelector, &out.RedisCacheNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StartIP != nil {
		in, out := &in.StartIP, &out.StartIP
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleParameters.
func (in *RedisFirewallRuleParameters) DeepCopy() *RedisFirewallRuleParameters {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleSpec) DeepCopyInto(out *RedisFirewallRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleSpec.
func (in *RedisFirewallRuleSpec) DeepCopy() *RedisFirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleStatus) DeepCopyInto(out *RedisFirewallRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleStatus.
func (in *RedisFirewallRuleStatus) DeepCopy() *RedisFirewallRuleStatus {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServer) DeepCopyInto(out *RedisLinkedServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServer.
func (in *RedisLinkedServer) DeepCopy() *RedisLinkedServer {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisLinkedServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerList) DeepCopyInto(out *RedisLinkedServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisLinkedServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerList.
func (in *RedisLinkedServerList) DeepCopy() *RedisLinkedServerList {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisLinkedServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerObservation) DeepCopyInto(out *RedisLinkedServerObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerObservation.
func (in *RedisLinkedServerObservation) DeepCopy() *RedisLinkedServerObservation {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerParameters) DeepCopyInto(out *RedisLinkedServerParameters) {
	*out = *in
	if in.LinkedRedisCacheID != nil {
		in, out := &in.LinkedRedisCacheID, &out.LinkedRedisCacheID
		*out = new(string)
		**out = **in
	}
	if in.LinkedRedisCacheIDRef != nil {
		in, out := &in.LinkedRedisCacheIDRef, &out.LinkedRedisCacheIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LinkedRedisCacheIDSelector != nil {
		in, out := &in.LinkedRedisCacheIDSelector, &out.LinkedRedisCacheIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.LinkedRedisCacheLocation != nil {
		in, out := &in.LinkedRedisCacheLocation, &out.LinkedRedisCacheLocation
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
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerRole != nil {
		in, out := &in.ServerRole, &out.ServerRole
		*out = new(string)
		**out = **in
	}
	if in.TargetRedisCacheName != nil {
		in, out := &in.TargetRedisCacheName, &out.TargetRedisCacheName
		*out = new(string)
		**out = **in
	}
	if in.TargetRedisCacheNameRef != nil {
		in, out := &in.TargetRedisCacheNameRef, &out.TargetRedisCacheNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetRedisCacheNameSelector != nil {
		in, out := &in.TargetRedisCacheNameSelector, &out.TargetRedisCacheNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerParameters.
func (in *RedisLinkedServerParameters) DeepCopy() *RedisLinkedServerParameters {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerSpec) DeepCopyInto(out *RedisLinkedServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerSpec.
func (in *RedisLinkedServerSpec) DeepCopy() *RedisLinkedServerSpec {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerStatus) DeepCopyInto(out *RedisLinkedServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerStatus.
func (in *RedisLinkedServerStatus) DeepCopy() *RedisLinkedServerStatus {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerStatus)
	in.DeepCopyInto(out)
	return out
}
