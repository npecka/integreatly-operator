// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcast":       schema_pkg_apis_apps_v1alpha1_APIcast(ref),
		"github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastSpec":   schema_pkg_apis_apps_v1alpha1_APIcastSpec(ref),
		"github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastStatus": schema_pkg_apis_apps_v1alpha1_APIcastStatus(ref),
	}
}

func schema_pkg_apis_apps_v1alpha1_APIcast(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIcast is the Schema for the apicasts API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastSpec", "github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_apps_v1alpha1_APIcastSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIcastSpec defines the desired state of APIcast",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"adminPortalCredentialsRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
					"embeddedConfigurationSecretRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"exposedHost": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastExposedHost"),
						},
					},
					"deploymentEnvironment": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"dnsResolverAddress": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"enabledServices": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"configurationLoadMode": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"logLevel": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"pathRoutingEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"responseCodesIncluded": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"cacheConfigurationSeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"managementAPIScope": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"openSSLPeerVerificationEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastExposedHost", "k8s.io/api/core/v1.LocalObjectReference"},
	}
}

func schema_pkg_apis_apps_v1alpha1_APIcastStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIcastStatus defines the observed state of APIcast",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Represents the latest available observations of a replica set's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastCondition"),
									},
								},
							},
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "The image being used in the APIcast deployment",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/3scale/apicast-operator/pkg/apis/apps/v1alpha1.APIcastCondition"},
	}
}