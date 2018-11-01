// +build !ignore_autogenerated

/*
Copyright 2018 The CDI Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolume":             schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolume(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeList":         schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeList(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSource":       schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSource(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceHTTP":   schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourceHTTP(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourcePVC":    schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourcePVC(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceS3":     schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourceS3(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceUpload": schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourceUpload(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSpec":         schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSpec(ref),
		"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeStatus":       schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeStatus(ref),
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolume(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolume provides a representation of our data volume",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
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
							Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeStatus"),
						},
					},
				},
				Required: []string{"spec", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSpec", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeStatus"},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeList provides the needed parameters to do request a list of Data Volumes from the system",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "Items provides a list of DataVolumes",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolume"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolume"},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeSource represents the source for our Data Volume, this can be HTTP, S3 or an existing PVC",
				Properties: map[string]spec.Schema{
					"http": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceHTTP"),
						},
					},
					"s3": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceS3"),
						},
					},
					"pvc": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourcePVC"),
						},
					},
					"upload": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceUpload"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceHTTP", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourcePVC", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceS3", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSourceUpload"},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourceHTTP(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeSourceHTTP provides the parameters to create a Data Volume from an HTTP source",
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is the URL of the http source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretRef provides the secret reference needed to access the HTTP source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourcePVC(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC",
				Properties: map[string]spec.Schema{
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourceS3(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source",
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is the url of the S3 source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretRef provides the secret reference needed to access the S3 source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSourceUpload(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeSourceUpload provides the parameters to create a Data Volume by uploading the source",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeSpec defines our specification for a DataVolume type",
				Properties: map[string]spec.Schema{
					"source": {
						SchemaProps: spec.SchemaProps{
							Description: "Source is the src of the data for the requested DataVolume",
							Ref:         ref("kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSource"),
						},
					},
					"pvc": {
						SchemaProps: spec.SchemaProps{
							Description: "PVC is a pointer to the PVC Spec we want to use",
							Ref:         ref("k8s.io/api/core/v1.PersistentVolumeClaimSpec"),
						},
					},
				},
				Required: []string{"source", "pvc"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.PersistentVolumeClaimSpec", "kubevirt.io/containerized-data-importer/pkg/apis/datavolumecontroller/v1alpha1.DataVolumeSource"},
	}
}

func schema_pkg_apis_datavolumecontroller_v1alpha1_DataVolumeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataVolumeStatus provides the parameters to store the phase of the Data Volume",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is the current phase of the data volume",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
