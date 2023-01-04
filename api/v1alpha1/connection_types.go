/*
Copyright 2022.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Protocol is the protocol used to connect to the host.
type Protocol string

const (
	// ProtocolSSH is the TCP protocol.
	ProtocolSSH Protocol = "SSH"
)

// OSInfo describes the operating system of a system.
type OSInfo struct {
	// Name is the name of the operating system.
	Name string `json:"name"`
	// Version is the version of the operating system.
	Version string `json:"version"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Protocol is the protocol used to connect to the host.
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Enum=SSH
	Protocol Protocol `json:"protocol"`

	// SecretRef is the reference to a secret containing sensitive connection credentials.
	//+kubebuilder:validation:Required
	SecretRef corev1.SecretReference `json:"secretRef"`
}

// ConnectionStatus defines the observed state of Connection
type ConnectionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// OS contains information about the discovered operating system of the system.
	OS OSInfo `json:"os,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=conns,path=connections,categories=kubestack
//+kubebuilder:printcolumn:name="Protocol",type=string,JSONPath=`.spec.protocol`
//+kubebuilder:printcolumn:name="OS-Name",type=string,JSONPath=`.status.os.name`
//+kubebuilder:printcolumn:name="OS-Version",type=string,JSONPath=`.status.os.version`

// Connection is the Schema for the connections API
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectionSpec   `json:"spec,omitempty"`
	Status ConnectionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConnectionList contains a list of Connection
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
