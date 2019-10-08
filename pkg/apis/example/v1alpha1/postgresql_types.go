package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PostgresqlSpec defines the desired state of Postgresql
// +k8s:openapi-gen=true
type PostgresqlSpec struct {
	// Instances specify the number of instances that this Postgres cluster will have
	Instances int32 `json:"instances"`
}

// PostgresqlStatus defines the observed state of Postgresql
// +k8s:openapi-gen=true
type PostgresqlStatus struct {
	// ConnectionUrl allows the database consumer to connect to Postgres cluster
	ConnectionUrl string `json:"connectionUrl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Postgresql is the Schema for the postgresqls API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Postgresql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgresqlSpec   `json:"spec,omitempty"`
	Status PostgresqlStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresqlList contains a list of Postgresql
type PostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Postgresql `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Postgresql{}, &PostgresqlList{})
}
