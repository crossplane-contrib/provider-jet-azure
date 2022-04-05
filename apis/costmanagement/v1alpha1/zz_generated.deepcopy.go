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
func (in *DeliveryInfoObservation) DeepCopyInto(out *DeliveryInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliveryInfoObservation.
func (in *DeliveryInfoObservation) DeepCopy() *DeliveryInfoObservation {
	if in == nil {
		return nil
	}
	out := new(DeliveryInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliveryInfoParameters) DeepCopyInto(out *DeliveryInfoParameters) {
	*out = *in
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.RootFolderPath != nil {
		in, out := &in.RootFolderPath, &out.RootFolderPath
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliveryInfoParameters.
func (in *DeliveryInfoParameters) DeepCopy() *DeliveryInfoParameters {
	if in == nil {
		return nil
	}
	out := new(DeliveryInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportResourceGroup) DeepCopyInto(out *ExportResourceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportResourceGroup.
func (in *ExportResourceGroup) DeepCopy() *ExportResourceGroup {
	if in == nil {
		return nil
	}
	out := new(ExportResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExportResourceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportResourceGroupList) DeepCopyInto(out *ExportResourceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExportResourceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportResourceGroupList.
func (in *ExportResourceGroupList) DeepCopy() *ExportResourceGroupList {
	if in == nil {
		return nil
	}
	out := new(ExportResourceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExportResourceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportResourceGroupObservation) DeepCopyInto(out *ExportResourceGroupObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportResourceGroupObservation.
func (in *ExportResourceGroupObservation) DeepCopy() *ExportResourceGroupObservation {
	if in == nil {
		return nil
	}
	out := new(ExportResourceGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportResourceGroupParameters) DeepCopyInto(out *ExportResourceGroupParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.DeliveryInfo != nil {
		in, out := &in.DeliveryInfo, &out.DeliveryInfo
		*out = make([]DeliveryInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = make([]QueryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecurrencePeriodEnd != nil {
		in, out := &in.RecurrencePeriodEnd, &out.RecurrencePeriodEnd
		*out = new(string)
		**out = **in
	}
	if in.RecurrencePeriodStart != nil {
		in, out := &in.RecurrencePeriodStart, &out.RecurrencePeriodStart
		*out = new(string)
		**out = **in
	}
	if in.RecurrenceType != nil {
		in, out := &in.RecurrenceType, &out.RecurrenceType
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupIDRef != nil {
		in, out := &in.ResourceGroupIDRef, &out.ResourceGroupIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupIDSelector != nil {
		in, out := &in.ResourceGroupIDSelector, &out.ResourceGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportResourceGroupParameters.
func (in *ExportResourceGroupParameters) DeepCopy() *ExportResourceGroupParameters {
	if in == nil {
		return nil
	}
	out := new(ExportResourceGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportResourceGroupSpec) DeepCopyInto(out *ExportResourceGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportResourceGroupSpec.
func (in *ExportResourceGroupSpec) DeepCopy() *ExportResourceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ExportResourceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportResourceGroupStatus) DeepCopyInto(out *ExportResourceGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportResourceGroupStatus.
func (in *ExportResourceGroupStatus) DeepCopy() *ExportResourceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ExportResourceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryObservation) DeepCopyInto(out *QueryObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryObservation.
func (in *QueryObservation) DeepCopy() *QueryObservation {
	if in == nil {
		return nil
	}
	out := new(QueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryParameters) DeepCopyInto(out *QueryParameters) {
	*out = *in
	if in.TimeFrame != nil {
		in, out := &in.TimeFrame, &out.TimeFrame
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameters.
func (in *QueryParameters) DeepCopy() *QueryParameters {
	if in == nil {
		return nil
	}
	out := new(QueryParameters)
	in.DeepCopyInto(out)
	return out
}