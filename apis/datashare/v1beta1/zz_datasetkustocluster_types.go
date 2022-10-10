/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DataSetKustoClusterObservation struct {

	// The name of the Data Share Dataset.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The resource ID of the Data Share Kusto Cluster Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the Kusto Cluster.
	KustoClusterLocation *string `json:"kustoClusterLocation,omitempty" tf:"kusto_cluster_location,omitempty"`
}

type DataSetKustoClusterParameters struct {

	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KustoClusterID *string `json:"kustoClusterId,omitempty" tf:"kusto_cluster_id,omitempty"`

	// Reference to a Cluster in kusto to populate kustoClusterId.
	// +kubebuilder:validation:Optional
	KustoClusterIDRef *v1.Reference `json:"kustoClusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate kustoClusterId.
	// +kubebuilder:validation:Optional
	KustoClusterIDSelector *v1.Selector `json:"kustoClusterIdSelector,omitempty" tf:"-"`

	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	// +crossplane:generate:reference:type=DataShare
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// Reference to a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDRef *v1.Reference `json:"shareIdRef,omitempty" tf:"-"`

	// Selector for a DataShare to populate shareId.
	// +kubebuilder:validation:Optional
	ShareIDSelector *v1.Selector `json:"shareIdSelector,omitempty" tf:"-"`
}

// DataSetKustoClusterSpec defines the desired state of DataSetKustoCluster
type DataSetKustoClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetKustoClusterParameters `json:"forProvider"`
}

// DataSetKustoClusterStatus defines the observed state of DataSetKustoCluster.
type DataSetKustoClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetKustoClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetKustoCluster is the Schema for the DataSetKustoClusters API. Manages a Data Share Kusto Cluster Dataset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetKustoCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetKustoClusterSpec   `json:"spec"`
	Status            DataSetKustoClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetKustoClusterList contains a list of DataSetKustoClusters
type DataSetKustoClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetKustoCluster `json:"items"`
}

// Repository type metadata.
var (
	DataSetKustoCluster_Kind             = "DataSetKustoCluster"
	DataSetKustoCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetKustoCluster_Kind}.String()
	DataSetKustoCluster_KindAPIVersion   = DataSetKustoCluster_Kind + "." + CRDGroupVersion.String()
	DataSetKustoCluster_GroupVersionKind = CRDGroupVersion.WithKind(DataSetKustoCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetKustoCluster{}, &DataSetKustoClusterList{})
}