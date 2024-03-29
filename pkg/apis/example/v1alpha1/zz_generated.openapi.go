// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.Postgresql":       schema_pkg_apis_example_v1alpha1_Postgresql(ref),
		"github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.PostgresqlSpec":   schema_pkg_apis_example_v1alpha1_PostgresqlSpec(ref),
		"github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.PostgresqlStatus": schema_pkg_apis_example_v1alpha1_PostgresqlStatus(ref),
	}
}

func schema_pkg_apis_example_v1alpha1_Postgresql(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Postgresql is the Schema for the postgresqls API",
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
							Ref: ref("github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.PostgresqlSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.PostgresqlStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.PostgresqlSpec", "github.com/aliok/postgres-operator/pkg/apis/example/v1alpha1.PostgresqlStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_example_v1alpha1_PostgresqlSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PostgresqlSpec defines the desired state of Postgresql",
				Properties: map[string]spec.Schema{
					"instances": {
						SchemaProps: spec.SchemaProps{
							Description: "Instances specify the number of instances that this Postgres cluster will have",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"instances"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_example_v1alpha1_PostgresqlStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PostgresqlStatus defines the observed state of Postgresql",
				Properties: map[string]spec.Schema{
					"connectionUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "ConnectionUrl allows the database consumer to connect to Postgres cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"connectionUrl"},
			},
		},
		Dependencies: []string{},
	}
}
