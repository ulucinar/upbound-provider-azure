// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LaunchInitParameters struct {

	// Specifies a map of non-sensitive properties for launchProperties.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Specifies a map of sensitive properties for launchProperties.
	// +mapType=granular
	Secrets map[string]*string `json:"secrets,omitempty" tf:"secrets,omitempty"`
}

type LaunchObservation struct {

	// Specifies a map of non-sensitive properties for launchProperties.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Specifies a map of sensitive properties for launchProperties.
	// +mapType=granular
	Secrets map[string]*string `json:"secrets,omitempty" tf:"secrets,omitempty"`
}

type LaunchParameters struct {

	// Specifies a map of non-sensitive properties for launchProperties.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Specifies a map of sensitive properties for launchProperties.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Secrets map[string]*string `json:"secrets,omitempty" tf:"secrets,omitempty"`
}

type SpringCloudBuildPackBindingInitParameters struct {

	// Specifies the Build Pack Binding Type. Allowed values are ApacheSkyWalking, AppDynamics, ApplicationInsights, Dynatrace, ElasticAPM and NewRelic.
	BindingType *string `json:"bindingType,omitempty" tf:"binding_type,omitempty"`

	// A launch block as defined below.
	Launch *LaunchInitParameters `json:"launch,omitempty" tf:"launch,omitempty"`
}

type SpringCloudBuildPackBindingObservation struct {

	// Specifies the Build Pack Binding Type. Allowed values are ApacheSkyWalking, AppDynamics, ApplicationInsights, Dynatrace, ElasticAPM and NewRelic.
	BindingType *string `json:"bindingType,omitempty" tf:"binding_type,omitempty"`

	// The ID of the Spring Cloud Build Pack Binding.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A launch block as defined below.
	Launch *LaunchObservation `json:"launch,omitempty" tf:"launch,omitempty"`

	// The ID of the Spring Cloud Builder. Changing this forces a new Spring Cloud Build Pack Binding to be created.
	SpringCloudBuilderID *string `json:"springCloudBuilderId,omitempty" tf:"spring_cloud_builder_id,omitempty"`
}

type SpringCloudBuildPackBindingParameters struct {

	// Specifies the Build Pack Binding Type. Allowed values are ApacheSkyWalking, AppDynamics, ApplicationInsights, Dynatrace, ElasticAPM and NewRelic.
	// +kubebuilder:validation:Optional
	BindingType *string `json:"bindingType,omitempty" tf:"binding_type,omitempty"`

	// A launch block as defined below.
	// +kubebuilder:validation:Optional
	Launch *LaunchParameters `json:"launch,omitempty" tf:"launch,omitempty"`

	// The ID of the Spring Cloud Builder. Changing this forces a new Spring Cloud Build Pack Binding to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta2.SpringCloudBuilder
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudBuilderID *string `json:"springCloudBuilderId,omitempty" tf:"spring_cloud_builder_id,omitempty"`

	// Reference to a SpringCloudBuilder in appplatform to populate springCloudBuilderId.
	// +kubebuilder:validation:Optional
	SpringCloudBuilderIDRef *v1.Reference `json:"springCloudBuilderIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudBuilder in appplatform to populate springCloudBuilderId.
	// +kubebuilder:validation:Optional
	SpringCloudBuilderIDSelector *v1.Selector `json:"springCloudBuilderIdSelector,omitempty" tf:"-"`
}

// SpringCloudBuildPackBindingSpec defines the desired state of SpringCloudBuildPackBinding
type SpringCloudBuildPackBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudBuildPackBindingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SpringCloudBuildPackBindingInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudBuildPackBindingStatus defines the observed state of SpringCloudBuildPackBinding.
type SpringCloudBuildPackBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudBuildPackBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SpringCloudBuildPackBinding is the Schema for the SpringCloudBuildPackBindings API. Manages a Spring Cloud Build Pack Binding.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudBuildPackBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudBuildPackBindingSpec   `json:"spec"`
	Status            SpringCloudBuildPackBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudBuildPackBindingList contains a list of SpringCloudBuildPackBindings
type SpringCloudBuildPackBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudBuildPackBinding `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudBuildPackBinding_Kind             = "SpringCloudBuildPackBinding"
	SpringCloudBuildPackBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudBuildPackBinding_Kind}.String()
	SpringCloudBuildPackBinding_KindAPIVersion   = SpringCloudBuildPackBinding_Kind + "." + CRDGroupVersion.String()
	SpringCloudBuildPackBinding_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudBuildPackBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudBuildPackBinding{}, &SpringCloudBuildPackBindingList{})
}