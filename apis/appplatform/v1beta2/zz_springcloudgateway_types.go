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

type APIMetadataInitParameters struct {

	// Detailed description of the APIs available on the Gateway instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Location of additional documentation for the APIs available on the Gateway instance.
	DocumentationURL *string `json:"documentationUrl,omitempty" tf:"documentation_url,omitempty"`

	// Base URL that API consumers will use to access APIs on the Gateway instance.
	ServerURL *string `json:"serverUrl,omitempty" tf:"server_url,omitempty"`

	// Specifies the title describing the context of the APIs available on the Gateway instance.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// Specifies the version of APIs available on this Gateway instance.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type APIMetadataObservation struct {

	// Detailed description of the APIs available on the Gateway instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Location of additional documentation for the APIs available on the Gateway instance.
	DocumentationURL *string `json:"documentationUrl,omitempty" tf:"documentation_url,omitempty"`

	// Base URL that API consumers will use to access APIs on the Gateway instance.
	ServerURL *string `json:"serverUrl,omitempty" tf:"server_url,omitempty"`

	// Specifies the title describing the context of the APIs available on the Gateway instance.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// Specifies the version of APIs available on this Gateway instance.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type APIMetadataParameters struct {

	// Detailed description of the APIs available on the Gateway instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Location of additional documentation for the APIs available on the Gateway instance.
	// +kubebuilder:validation:Optional
	DocumentationURL *string `json:"documentationUrl,omitempty" tf:"documentation_url,omitempty"`

	// Base URL that API consumers will use to access APIs on the Gateway instance.
	// +kubebuilder:validation:Optional
	ServerURL *string `json:"serverUrl,omitempty" tf:"server_url,omitempty"`

	// Specifies the title describing the context of the APIs available on the Gateway instance.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// Specifies the version of APIs available on this Gateway instance.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClientAuthorizationInitParameters struct {

	// Specifies the Spring Cloud Certificate IDs of the Spring Cloud Gateway.
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// Specifies whether the client certificate verification is enabled.
	VerificationEnabled *bool `json:"verificationEnabled,omitempty" tf:"verification_enabled,omitempty"`
}

type ClientAuthorizationObservation struct {

	// Specifies the Spring Cloud Certificate IDs of the Spring Cloud Gateway.
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// Specifies whether the client certificate verification is enabled.
	VerificationEnabled *bool `json:"verificationEnabled,omitempty" tf:"verification_enabled,omitempty"`
}

type ClientAuthorizationParameters struct {

	// Specifies the Spring Cloud Certificate IDs of the Spring Cloud Gateway.
	// +kubebuilder:validation:Optional
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// Specifies whether the client certificate verification is enabled.
	// +kubebuilder:validation:Optional
	VerificationEnabled *bool `json:"verificationEnabled,omitempty" tf:"verification_enabled,omitempty"`
}

type CorsInitParameters struct {

	// Allowed headers in cross-site requests. The special value * allows actual requests to send any header.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Allowed HTTP methods on cross-site requests. The special value * allows all methods. If not set, GET and HEAD are allowed by default. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS and PUT.
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Allowed origin patterns to make cross-site requests.
	// +listType=set
	AllowedOriginPatterns []*string `json:"allowedOriginPatterns,omitempty" tf:"allowed_origin_patterns,omitempty"`

	// Allowed origins to make cross-site requests. The special value * allows all domains.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// is user credentials are supported on cross-site requests?
	CredentialsAllowed *bool `json:"credentialsAllowed,omitempty" tf:"credentials_allowed,omitempty"`

	// HTTP response headers to expose for cross-site requests.
	// +listType=set
	ExposedHeaders []*string `json:"exposedHeaders,omitempty" tf:"exposed_headers,omitempty"`

	// How long, in seconds, the response from a pre-flight request can be cached by clients.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsObservation struct {

	// Allowed headers in cross-site requests. The special value * allows actual requests to send any header.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Allowed HTTP methods on cross-site requests. The special value * allows all methods. If not set, GET and HEAD are allowed by default. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS and PUT.
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Allowed origin patterns to make cross-site requests.
	// +listType=set
	AllowedOriginPatterns []*string `json:"allowedOriginPatterns,omitempty" tf:"allowed_origin_patterns,omitempty"`

	// Allowed origins to make cross-site requests. The special value * allows all domains.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// is user credentials are supported on cross-site requests?
	CredentialsAllowed *bool `json:"credentialsAllowed,omitempty" tf:"credentials_allowed,omitempty"`

	// HTTP response headers to expose for cross-site requests.
	// +listType=set
	ExposedHeaders []*string `json:"exposedHeaders,omitempty" tf:"exposed_headers,omitempty"`

	// How long, in seconds, the response from a pre-flight request can be cached by clients.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsParameters struct {

	// Allowed headers in cross-site requests. The special value * allows actual requests to send any header.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Allowed HTTP methods on cross-site requests. The special value * allows all methods. If not set, GET and HEAD are allowed by default. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS and PUT.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Allowed origin patterns to make cross-site requests.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedOriginPatterns []*string `json:"allowedOriginPatterns,omitempty" tf:"allowed_origin_patterns,omitempty"`

	// Allowed origins to make cross-site requests. The special value * allows all domains.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// is user credentials are supported on cross-site requests?
	// +kubebuilder:validation:Optional
	CredentialsAllowed *bool `json:"credentialsAllowed,omitempty" tf:"credentials_allowed,omitempty"`

	// HTTP response headers to expose for cross-site requests.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExposedHeaders []*string `json:"exposedHeaders,omitempty" tf:"exposed_headers,omitempty"`

	// How long, in seconds, the response from a pre-flight request can be cached by clients.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type LocalResponseCachePerInstanceInitParameters struct {

	// Specifies the maximum size of cache (10MB, 900KB, 1GB...) to determine if the cache needs to evict some entries.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the time before a cached entry is expired (300s, 5m, 1h...).
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type LocalResponseCachePerInstanceObservation struct {

	// Specifies the maximum size of cache (10MB, 900KB, 1GB...) to determine if the cache needs to evict some entries.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the time before a cached entry is expired (300s, 5m, 1h...).
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type LocalResponseCachePerInstanceParameters struct {

	// Specifies the maximum size of cache (10MB, 900KB, 1GB...) to determine if the cache needs to evict some entries.
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the time before a cached entry is expired (300s, 5m, 1h...).
	// +kubebuilder:validation:Optional
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type LocalResponseCachePerRouteInitParameters struct {

	// Specifies the maximum size of cache (10MB, 900KB, 1GB...) to determine if the cache needs to evict some entries.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the time before a cached entry is expired (300s, 5m, 1h...).
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type LocalResponseCachePerRouteObservation struct {

	// Specifies the maximum size of cache (10MB, 900KB, 1GB...) to determine if the cache needs to evict some entries.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the time before a cached entry is expired (300s, 5m, 1h...).
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type LocalResponseCachePerRouteParameters struct {

	// Specifies the maximum size of cache (10MB, 900KB, 1GB...) to determine if the cache needs to evict some entries.
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the time before a cached entry is expired (300s, 5m, 1h...).
	// +kubebuilder:validation:Optional
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type SpringCloudGatewayInitParameters struct {

	// A api_metadata block as defined below.
	APIMetadata *APIMetadataInitParameters `json:"apiMetadata,omitempty" tf:"api_metadata,omitempty"`

	// Specifies a list of Spring Cloud Application Performance Monitoring IDs.
	ApplicationPerformanceMonitoringIds []*string `json:"applicationPerformanceMonitoringIds,omitempty" tf:"application_performance_monitoring_ids,omitempty"`

	// Specifies a list of application performance monitoring types used in the Spring Cloud Gateway. The allowed values are AppDynamics, ApplicationInsights, Dynatrace, ElasticAPM and NewRelic.
	ApplicationPerformanceMonitoringTypes []*string `json:"applicationPerformanceMonitoringTypes,omitempty" tf:"application_performance_monitoring_types,omitempty"`

	// A client_authorization block as defined below.
	ClientAuthorization *ClientAuthorizationInitParameters `json:"clientAuthorization,omitempty" tf:"client_authorization,omitempty"`

	// A cors block as defined below.
	Cors *CorsInitParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies the environment variables of the Spring Cloud Gateway as a map of key-value pairs. Changing this forces a new resource to be created.
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// is only https is allowed?
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// Specifies the required instance count of the Spring Cloud Gateway. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// A local_response_cache_per_instance block as defined below. Only one of local_response_cache_per_instance or local_response_cache_per_route can be specified.
	LocalResponseCachePerInstance *LocalResponseCachePerInstanceInitParameters `json:"localResponseCachePerInstance,omitempty" tf:"local_response_cache_per_instance,omitempty"`

	// A local_response_cache_per_route block as defined below. Only one of local_response_cache_per_instance or local_response_cache_per_route can be specified.
	LocalResponseCachePerRoute *LocalResponseCachePerRouteInitParameters `json:"localResponseCachePerRoute,omitempty" tf:"local_response_cache_per_route,omitempty"`

	// Indicates whether the Spring Cloud Gateway exposes endpoint.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A quota block as defined below.
	Quota *SpringCloudGatewayQuotaInitParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	SensitiveEnvironmentVariables map[string]*string `json:"sensitiveEnvironmentVariablesSecretRef,omitempty" tf:"-"`

	// A sso block as defined below.
	Sso *SpringCloudGatewaySsoInitParameters `json:"sso,omitempty" tf:"sso,omitempty"`
}

type SpringCloudGatewayObservation struct {

	// A api_metadata block as defined below.
	APIMetadata *APIMetadataObservation `json:"apiMetadata,omitempty" tf:"api_metadata,omitempty"`

	// Specifies a list of Spring Cloud Application Performance Monitoring IDs.
	ApplicationPerformanceMonitoringIds []*string `json:"applicationPerformanceMonitoringIds,omitempty" tf:"application_performance_monitoring_ids,omitempty"`

	// Specifies a list of application performance monitoring types used in the Spring Cloud Gateway. The allowed values are AppDynamics, ApplicationInsights, Dynatrace, ElasticAPM and NewRelic.
	ApplicationPerformanceMonitoringTypes []*string `json:"applicationPerformanceMonitoringTypes,omitempty" tf:"application_performance_monitoring_types,omitempty"`

	// A client_authorization block as defined below.
	ClientAuthorization *ClientAuthorizationObservation `json:"clientAuthorization,omitempty" tf:"client_authorization,omitempty"`

	// A cors block as defined below.
	Cors *CorsObservation `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies the environment variables of the Spring Cloud Gateway as a map of key-value pairs. Changing this forces a new resource to be created.
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// is only https is allowed?
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// The ID of the Spring Cloud Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the required instance count of the Spring Cloud Gateway. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// A local_response_cache_per_instance block as defined below. Only one of local_response_cache_per_instance or local_response_cache_per_route can be specified.
	LocalResponseCachePerInstance *LocalResponseCachePerInstanceObservation `json:"localResponseCachePerInstance,omitempty" tf:"local_response_cache_per_instance,omitempty"`

	// A local_response_cache_per_route block as defined below. Only one of local_response_cache_per_instance or local_response_cache_per_route can be specified.
	LocalResponseCachePerRoute *LocalResponseCachePerRouteObservation `json:"localResponseCachePerRoute,omitempty" tf:"local_response_cache_per_route,omitempty"`

	// Indicates whether the Spring Cloud Gateway exposes endpoint.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A quota block as defined below.
	Quota *SpringCloudGatewayQuotaObservation `json:"quota,omitempty" tf:"quota,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Gateway to be created.
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// A sso block as defined below.
	Sso *SpringCloudGatewaySsoObservation `json:"sso,omitempty" tf:"sso,omitempty"`

	// URL of the Spring Cloud Gateway, exposed when 'public_network_access_enabled' is true.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SpringCloudGatewayParameters struct {

	// A api_metadata block as defined below.
	// +kubebuilder:validation:Optional
	APIMetadata *APIMetadataParameters `json:"apiMetadata,omitempty" tf:"api_metadata,omitempty"`

	// Specifies a list of Spring Cloud Application Performance Monitoring IDs.
	// +kubebuilder:validation:Optional
	ApplicationPerformanceMonitoringIds []*string `json:"applicationPerformanceMonitoringIds,omitempty" tf:"application_performance_monitoring_ids,omitempty"`

	// Specifies a list of application performance monitoring types used in the Spring Cloud Gateway. The allowed values are AppDynamics, ApplicationInsights, Dynatrace, ElasticAPM and NewRelic.
	// +kubebuilder:validation:Optional
	ApplicationPerformanceMonitoringTypes []*string `json:"applicationPerformanceMonitoringTypes,omitempty" tf:"application_performance_monitoring_types,omitempty"`

	// A client_authorization block as defined below.
	// +kubebuilder:validation:Optional
	ClientAuthorization *ClientAuthorizationParameters `json:"clientAuthorization,omitempty" tf:"client_authorization,omitempty"`

	// A cors block as defined below.
	// +kubebuilder:validation:Optional
	Cors *CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies the environment variables of the Spring Cloud Gateway as a map of key-value pairs. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// is only https is allowed?
	// +kubebuilder:validation:Optional
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// Specifies the required instance count of the Spring Cloud Gateway. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// A local_response_cache_per_instance block as defined below. Only one of local_response_cache_per_instance or local_response_cache_per_route can be specified.
	// +kubebuilder:validation:Optional
	LocalResponseCachePerInstance *LocalResponseCachePerInstanceParameters `json:"localResponseCachePerInstance,omitempty" tf:"local_response_cache_per_instance,omitempty"`

	// A local_response_cache_per_route block as defined below. Only one of local_response_cache_per_instance or local_response_cache_per_route can be specified.
	// +kubebuilder:validation:Optional
	LocalResponseCachePerRoute *LocalResponseCachePerRouteParameters `json:"localResponseCachePerRoute,omitempty" tf:"local_response_cache_per_route,omitempty"`

	// Indicates whether the Spring Cloud Gateway exposes endpoint.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A quota block as defined below.
	// +kubebuilder:validation:Optional
	Quota *SpringCloudGatewayQuotaParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	// Specifies the sensitive environment variables of the Spring Cloud Gateway as a map of key-value pairs. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SensitiveEnvironmentVariablesSecretRef *v1.SecretReference `json:"sensitiveEnvironmentVariablesSecretRef,omitempty" tf:"-"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Gateway to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta2.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`

	// A sso block as defined below.
	// +kubebuilder:validation:Optional
	Sso *SpringCloudGatewaySsoParameters `json:"sso,omitempty" tf:"sso,omitempty"`
}

type SpringCloudGatewayQuotaInitParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 2Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudGatewayQuotaObservation struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 2Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudGatewayQuotaParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 2Gi if not specified.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudGatewaySsoInitParameters struct {

	// The public identifier for the application.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The secret known only to the application and the authorization server.
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// The URI of Issuer Identifier.
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// It defines the specific actions applications can be allowed to do on a user's behalf.
	// +listType=set
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type SpringCloudGatewaySsoObservation struct {

	// The public identifier for the application.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The secret known only to the application and the authorization server.
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// The URI of Issuer Identifier.
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// It defines the specific actions applications can be allowed to do on a user's behalf.
	// +listType=set
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type SpringCloudGatewaySsoParameters struct {

	// The public identifier for the application.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The secret known only to the application and the authorization server.
	// +kubebuilder:validation:Optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// The URI of Issuer Identifier.
	// +kubebuilder:validation:Optional
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// It defines the specific actions applications can be allowed to do on a user's behalf.
	// +kubebuilder:validation:Optional
	// +listType=set
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// SpringCloudGatewaySpec defines the desired state of SpringCloudGateway
type SpringCloudGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudGatewayParameters `json:"forProvider"`
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
	InitProvider SpringCloudGatewayInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudGatewayStatus defines the observed state of SpringCloudGateway.
type SpringCloudGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SpringCloudGateway is the Schema for the SpringCloudGateways API. Manages a Spring Cloud Gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudGatewaySpec   `json:"spec"`
	Status            SpringCloudGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudGatewayList contains a list of SpringCloudGateways
type SpringCloudGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudGateway `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudGateway_Kind             = "SpringCloudGateway"
	SpringCloudGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudGateway_Kind}.String()
	SpringCloudGateway_KindAPIVersion   = SpringCloudGateway_Kind + "." + CRDGroupVersion.String()
	SpringCloudGateway_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudGateway{}, &SpringCloudGatewayList{})
}