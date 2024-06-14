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

type AzureActiveDirectoryAuthenticationInitParameters struct {

	// The Audience which should be used for authentication.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Issuer which should be used for authentication.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The Tenant which should be used for authentication.
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`
}

type AzureActiveDirectoryAuthenticationObservation struct {

	// The Audience which should be used for authentication.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Issuer which should be used for authentication.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The Tenant which should be used for authentication.
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`
}

type AzureActiveDirectoryAuthenticationParameters struct {

	// The Audience which should be used for authentication.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience" tf:"audience,omitempty"`

	// The Issuer which should be used for authentication.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// The Tenant which should be used for authentication.
	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant" tf:"tenant,omitempty"`
}

type ClientRevokedCertificateInitParameters struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Thumbprint of the Certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ClientRevokedCertificateObservation struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Thumbprint of the Certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ClientRevokedCertificateParameters struct {

	// A name used to uniquely identify this certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`
}

type ClientRootCertificateInitParameters struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Public Key Data associated with the Certificate.
	PublicCertData *string `json:"publicCertData,omitempty" tf:"public_cert_data,omitempty"`
}

type ClientRootCertificateObservation struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Public Key Data associated with the Certificate.
	PublicCertData *string `json:"publicCertData,omitempty" tf:"public_cert_data,omitempty"`
}

type ClientRootCertificateParameters struct {

	// A name used to uniquely identify this certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Public Key Data associated with the Certificate.
	// +kubebuilder:validation:Optional
	PublicCertData *string `json:"publicCertData" tf:"public_cert_data,omitempty"`
}

type RadiusClientRootCertificateInitParameters struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Thumbprint of the Certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type RadiusClientRootCertificateObservation struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Thumbprint of the Certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type RadiusClientRootCertificateParameters struct {

	// A name used to uniquely identify this certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`
}

type RadiusInitParameters struct {

	// One or more client_root_certificate blocks as defined below.
	ClientRootCertificate []RadiusClientRootCertificateInitParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`

	// One or more server blocks as defined below.
	Server []ServerInitParameters `json:"server,omitempty" tf:"server,omitempty"`

	// One or more server_root_certificate blocks as defined below.
	ServerRootCertificate []ServerRootCertificateInitParameters `json:"serverRootCertificate,omitempty" tf:"server_root_certificate,omitempty"`
}

type RadiusObservation struct {

	// One or more client_root_certificate blocks as defined below.
	ClientRootCertificate []RadiusClientRootCertificateObservation `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`

	// One or more server blocks as defined below.
	Server []ServerObservation `json:"server,omitempty" tf:"server,omitempty"`

	// One or more server_root_certificate blocks as defined below.
	ServerRootCertificate []ServerRootCertificateObservation `json:"serverRootCertificate,omitempty" tf:"server_root_certificate,omitempty"`
}

type RadiusParameters struct {

	// One or more client_root_certificate blocks as defined below.
	// +kubebuilder:validation:Optional
	ClientRootCertificate []RadiusClientRootCertificateParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`

	// One or more server blocks as defined below.
	// +kubebuilder:validation:Optional
	Server []ServerParameters `json:"server,omitempty" tf:"server,omitempty"`

	// One or more server_root_certificate blocks as defined below.
	// +kubebuilder:validation:Optional
	ServerRootCertificate []ServerRootCertificateParameters `json:"serverRootCertificate,omitempty" tf:"server_root_certificate,omitempty"`
}

type ServerInitParameters struct {

	// The Address of the Radius Server.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The Score of the Radius Server determines the priority of the server. Ranges from 1 to 30.
	Score *float64 `json:"score,omitempty" tf:"score,omitempty"`

	// The Secret used to communicate with the Radius Server.
	SecretSecretRef v1.SecretKeySelector `json:"secretSecretRef" tf:"-"`
}

type ServerObservation struct {

	// The Address of the Radius Server.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The Score of the Radius Server determines the priority of the server. Ranges from 1 to 30.
	Score *float64 `json:"score,omitempty" tf:"score,omitempty"`
}

type ServerParameters struct {

	// The Address of the Radius Server.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// The Score of the Radius Server determines the priority of the server. Ranges from 1 to 30.
	// +kubebuilder:validation:Optional
	Score *float64 `json:"score" tf:"score,omitempty"`

	// The Secret used to communicate with the Radius Server.
	// +kubebuilder:validation:Optional
	SecretSecretRef v1.SecretKeySelector `json:"secretSecretRef" tf:"-"`
}

type ServerRootCertificateInitParameters struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Public Key Data associated with the Certificate.
	PublicCertData *string `json:"publicCertData,omitempty" tf:"public_cert_data,omitempty"`
}

type ServerRootCertificateObservation struct {

	// A name used to uniquely identify this certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Public Key Data associated with the Certificate.
	PublicCertData *string `json:"publicCertData,omitempty" tf:"public_cert_data,omitempty"`
}

type ServerRootCertificateParameters struct {

	// A name used to uniquely identify this certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Public Key Data associated with the Certificate.
	// +kubebuilder:validation:Optional
	PublicCertData *string `json:"publicCertData" tf:"public_cert_data,omitempty"`
}

type VPNServerConfigurationInitParameters struct {

	// A azure_active_directory_authentication block as defined below.
	AzureActiveDirectoryAuthentication []AzureActiveDirectoryAuthenticationInitParameters `json:"azureActiveDirectoryAuthentication,omitempty" tf:"azure_active_directory_authentication,omitempty"`

	// One or more client_revoked_certificate blocks as defined below.
	ClientRevokedCertificate []ClientRevokedCertificateInitParameters `json:"clientRevokedCertificate,omitempty" tf:"client_revoked_certificate,omitempty"`

	// One or more client_root_certificate blocks as defined below.
	ClientRootCertificate []ClientRootCertificateInitParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`

	// A ipsec_policy block as defined below.
	IpsecPolicy *VPNServerConfigurationIpsecPolicyInitParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A radius block as defined below.
	Radius *RadiusInitParameters `json:"radius,omitempty" tf:"radius,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are AAD (Azure Active Directory), Certificate and Radius.
	VPNAuthenticationTypes []*string `json:"vpnAuthenticationTypes,omitempty" tf:"vpn_authentication_types,omitempty"`

	// A list of VPN Protocols to use for this Server Configuration. Possible values are IkeV2 and OpenVPN.
	// +listType=set
	VPNProtocols []*string `json:"vpnProtocols,omitempty" tf:"vpn_protocols,omitempty"`
}

type VPNServerConfigurationIpsecPolicyInitParameters struct {

	// The DH Group, used in IKE Phase 1. Possible values include DHGroup1, DHGroup2, DHGroup14, DHGroup24, DHGroup2048, ECP256, ECP384 and None.
	DhGroup *string `json:"dhGroup,omitempty" tf:"dh_group,omitempty"`

	// The IKE encryption algorithm, used for IKE Phase 2. Possible values include AES128, AES192, AES256, DES, DES3, GCMAES128 and GCMAES256.
	IkeEncryption *string `json:"ikeEncryption,omitempty" tf:"ike_encryption,omitempty"`

	// The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include GCMAES128, GCMAES256, MD5, SHA1, SHA256 and SHA384.
	IkeIntegrity *string `json:"ikeIntegrity,omitempty" tf:"ike_integrity,omitempty"`

	// The IPSec encryption algorithm, used for IKE phase 1. Possible values include AES128, AES192, AES256, DES, DES3, GCMAES128, GCMAES192, GCMAES256 and None.
	IpsecEncryption *string `json:"ipsecEncryption,omitempty" tf:"ipsec_encryption,omitempty"`

	// The IPSec integrity algorithm, used for IKE phase 1. Possible values include GCMAES128, GCMAES192, GCMAES256, MD5, SHA1 and SHA256.
	IpsecIntegrity *string `json:"ipsecIntegrity,omitempty" tf:"ipsec_integrity,omitempty"`

	// The Pfs Group, used in IKE Phase 2. Possible values include ECP256, ECP384, PFS1, PFS2, PFS14, PFS24, PFS2048, PFSMM and None.
	PfsGroup *string `json:"pfsGroup,omitempty" tf:"pfs_group,omitempty"`

	// The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
	SaDataSizeKilobytes *float64 `json:"saDataSizeKilobytes,omitempty" tf:"sa_data_size_kilobytes,omitempty"`

	// The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.
	SaLifetimeSeconds *float64 `json:"saLifetimeSeconds,omitempty" tf:"sa_lifetime_seconds,omitempty"`
}

type VPNServerConfigurationIpsecPolicyObservation struct {

	// The DH Group, used in IKE Phase 1. Possible values include DHGroup1, DHGroup2, DHGroup14, DHGroup24, DHGroup2048, ECP256, ECP384 and None.
	DhGroup *string `json:"dhGroup,omitempty" tf:"dh_group,omitempty"`

	// The IKE encryption algorithm, used for IKE Phase 2. Possible values include AES128, AES192, AES256, DES, DES3, GCMAES128 and GCMAES256.
	IkeEncryption *string `json:"ikeEncryption,omitempty" tf:"ike_encryption,omitempty"`

	// The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include GCMAES128, GCMAES256, MD5, SHA1, SHA256 and SHA384.
	IkeIntegrity *string `json:"ikeIntegrity,omitempty" tf:"ike_integrity,omitempty"`

	// The IPSec encryption algorithm, used for IKE phase 1. Possible values include AES128, AES192, AES256, DES, DES3, GCMAES128, GCMAES192, GCMAES256 and None.
	IpsecEncryption *string `json:"ipsecEncryption,omitempty" tf:"ipsec_encryption,omitempty"`

	// The IPSec integrity algorithm, used for IKE phase 1. Possible values include GCMAES128, GCMAES192, GCMAES256, MD5, SHA1 and SHA256.
	IpsecIntegrity *string `json:"ipsecIntegrity,omitempty" tf:"ipsec_integrity,omitempty"`

	// The Pfs Group, used in IKE Phase 2. Possible values include ECP256, ECP384, PFS1, PFS2, PFS14, PFS24, PFS2048, PFSMM and None.
	PfsGroup *string `json:"pfsGroup,omitempty" tf:"pfs_group,omitempty"`

	// The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
	SaDataSizeKilobytes *float64 `json:"saDataSizeKilobytes,omitempty" tf:"sa_data_size_kilobytes,omitempty"`

	// The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.
	SaLifetimeSeconds *float64 `json:"saLifetimeSeconds,omitempty" tf:"sa_lifetime_seconds,omitempty"`
}

type VPNServerConfigurationIpsecPolicyParameters struct {

	// The DH Group, used in IKE Phase 1. Possible values include DHGroup1, DHGroup2, DHGroup14, DHGroup24, DHGroup2048, ECP256, ECP384 and None.
	// +kubebuilder:validation:Optional
	DhGroup *string `json:"dhGroup" tf:"dh_group,omitempty"`

	// The IKE encryption algorithm, used for IKE Phase 2. Possible values include AES128, AES192, AES256, DES, DES3, GCMAES128 and GCMAES256.
	// +kubebuilder:validation:Optional
	IkeEncryption *string `json:"ikeEncryption" tf:"ike_encryption,omitempty"`

	// The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include GCMAES128, GCMAES256, MD5, SHA1, SHA256 and SHA384.
	// +kubebuilder:validation:Optional
	IkeIntegrity *string `json:"ikeIntegrity" tf:"ike_integrity,omitempty"`

	// The IPSec encryption algorithm, used for IKE phase 1. Possible values include AES128, AES192, AES256, DES, DES3, GCMAES128, GCMAES192, GCMAES256 and None.
	// +kubebuilder:validation:Optional
	IpsecEncryption *string `json:"ipsecEncryption" tf:"ipsec_encryption,omitempty"`

	// The IPSec integrity algorithm, used for IKE phase 1. Possible values include GCMAES128, GCMAES192, GCMAES256, MD5, SHA1 and SHA256.
	// +kubebuilder:validation:Optional
	IpsecIntegrity *string `json:"ipsecIntegrity" tf:"ipsec_integrity,omitempty"`

	// The Pfs Group, used in IKE Phase 2. Possible values include ECP256, ECP384, PFS1, PFS2, PFS14, PFS24, PFS2048, PFSMM and None.
	// +kubebuilder:validation:Optional
	PfsGroup *string `json:"pfsGroup" tf:"pfs_group,omitempty"`

	// The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
	// +kubebuilder:validation:Optional
	SaDataSizeKilobytes *float64 `json:"saDataSizeKilobytes" tf:"sa_data_size_kilobytes,omitempty"`

	// The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.
	// +kubebuilder:validation:Optional
	SaLifetimeSeconds *float64 `json:"saLifetimeSeconds" tf:"sa_lifetime_seconds,omitempty"`
}

type VPNServerConfigurationObservation struct {

	// A azure_active_directory_authentication block as defined below.
	AzureActiveDirectoryAuthentication []AzureActiveDirectoryAuthenticationObservation `json:"azureActiveDirectoryAuthentication,omitempty" tf:"azure_active_directory_authentication,omitempty"`

	// One or more client_revoked_certificate blocks as defined below.
	ClientRevokedCertificate []ClientRevokedCertificateObservation `json:"clientRevokedCertificate,omitempty" tf:"client_revoked_certificate,omitempty"`

	// One or more client_root_certificate blocks as defined below.
	ClientRootCertificate []ClientRootCertificateObservation `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`

	// The ID of the VPN Server Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A ipsec_policy block as defined below.
	IpsecPolicy *VPNServerConfigurationIpsecPolicyObservation `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A radius block as defined below.
	Radius *RadiusObservation `json:"radius,omitempty" tf:"radius,omitempty"`

	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are AAD (Azure Active Directory), Certificate and Radius.
	VPNAuthenticationTypes []*string `json:"vpnAuthenticationTypes,omitempty" tf:"vpn_authentication_types,omitempty"`

	// A list of VPN Protocols to use for this Server Configuration. Possible values are IkeV2 and OpenVPN.
	// +listType=set
	VPNProtocols []*string `json:"vpnProtocols,omitempty" tf:"vpn_protocols,omitempty"`
}

type VPNServerConfigurationParameters struct {

	// A azure_active_directory_authentication block as defined below.
	// +kubebuilder:validation:Optional
	AzureActiveDirectoryAuthentication []AzureActiveDirectoryAuthenticationParameters `json:"azureActiveDirectoryAuthentication,omitempty" tf:"azure_active_directory_authentication,omitempty"`

	// One or more client_revoked_certificate blocks as defined below.
	// +kubebuilder:validation:Optional
	ClientRevokedCertificate []ClientRevokedCertificateParameters `json:"clientRevokedCertificate,omitempty" tf:"client_revoked_certificate,omitempty"`

	// One or more client_root_certificate blocks as defined below.
	// +kubebuilder:validation:Optional
	ClientRootCertificate []ClientRootCertificateParameters `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`

	// A ipsec_policy block as defined below.
	// +kubebuilder:validation:Optional
	IpsecPolicy *VPNServerConfigurationIpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A radius block as defined below.
	// +kubebuilder:validation:Optional
	Radius *RadiusParameters `json:"radius,omitempty" tf:"radius,omitempty"`

	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are AAD (Azure Active Directory), Certificate and Radius.
	// +kubebuilder:validation:Optional
	VPNAuthenticationTypes []*string `json:"vpnAuthenticationTypes,omitempty" tf:"vpn_authentication_types,omitempty"`

	// A list of VPN Protocols to use for this Server Configuration. Possible values are IkeV2 and OpenVPN.
	// +kubebuilder:validation:Optional
	// +listType=set
	VPNProtocols []*string `json:"vpnProtocols,omitempty" tf:"vpn_protocols,omitempty"`
}

// VPNServerConfigurationSpec defines the desired state of VPNServerConfiguration
type VPNServerConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNServerConfigurationParameters `json:"forProvider"`
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
	InitProvider VPNServerConfigurationInitParameters `json:"initProvider,omitempty"`
}

// VPNServerConfigurationStatus defines the observed state of VPNServerConfiguration.
type VPNServerConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNServerConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VPNServerConfiguration is the Schema for the VPNServerConfigurations API. Manages a VPN Server Configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VPNServerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpnAuthenticationTypes) || (has(self.initProvider) && has(self.initProvider.vpnAuthenticationTypes))",message="spec.forProvider.vpnAuthenticationTypes is a required parameter"
	Spec   VPNServerConfigurationSpec   `json:"spec"`
	Status VPNServerConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNServerConfigurationList contains a list of VPNServerConfigurations
type VPNServerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNServerConfiguration `json:"items"`
}

// Repository type metadata.
var (
	VPNServerConfiguration_Kind             = "VPNServerConfiguration"
	VPNServerConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNServerConfiguration_Kind}.String()
	VPNServerConfiguration_KindAPIVersion   = VPNServerConfiguration_Kind + "." + CRDGroupVersion.String()
	VPNServerConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(VPNServerConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNServerConfiguration{}, &VPNServerConfigurationList{})
}