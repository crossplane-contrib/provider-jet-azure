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

type AppActionHTTPObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppActionHTTPParameters struct {

	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Required
	LogicAppID *string `json:"logicAppId" tf:"logic_app_id,omitempty"`

	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RunAfter []RunAfterParameters `json:"runAfter,omitempty" tf:"run_after,omitempty"`

	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type RunAfterObservation struct {
}

type RunAfterParameters struct {

	// +kubebuilder:validation:Required
	ActionName *string `json:"actionName" tf:"action_name,omitempty"`

	// +kubebuilder:validation:Required
	ActionResult *string `json:"actionResult" tf:"action_result,omitempty"`
}

// AppActionHTTPSpec defines the desired state of AppActionHTTP
type AppActionHTTPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppActionHTTPParameters `json:"forProvider"`
}

// AppActionHTTPStatus defines the observed state of AppActionHTTP.
type AppActionHTTPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppActionHTTPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppActionHTTP is the Schema for the AppActionHTTPs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppActionHTTP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppActionHTTPSpec   `json:"spec"`
	Status            AppActionHTTPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppActionHTTPList contains a list of AppActionHTTPs
type AppActionHTTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppActionHTTP `json:"items"`
}

// Repository type metadata.
var (
	AppActionHTTP_Kind             = "AppActionHTTP"
	AppActionHTTP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppActionHTTP_Kind}.String()
	AppActionHTTP_KindAPIVersion   = AppActionHTTP_Kind + "." + CRDGroupVersion.String()
	AppActionHTTP_GroupVersionKind = CRDGroupVersion.WithKind(AppActionHTTP_Kind)
)

func init() {
	SchemeBuilder.Register(&AppActionHTTP{}, &AppActionHTTPList{})
}