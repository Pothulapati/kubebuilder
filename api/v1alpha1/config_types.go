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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type InstallerStatusType string

func (i InstallerStatusType) String() string {
	return string(i)
}

const (
	InstallerStatusTypePending  InstallerStatusType = "PENDING"
	InstallerStatusTypeRunning  InstallerStatusType = "RUNNING"
	InstallerStatusTypeCleaning InstallerStatusType = "CLEANING"
)

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Domain is the domain name that the Gitpod instance is run on
	Domain string `json:"domain"`

	// The full path to the Installer image - typically "eu.gcr.io/gitpod-core-dev/build/installer:<tag>"
	InstallerImage string `json:"installerImage"`

	// This flag tells the controller render experimental config.  Defaults to false.
	// +optional
	UseExperimentConfig *bool `json:"useExperimentalConfig,omitempty"`

	// @todo(sje): remove
	ClientId            string `json:"clientId,omitempty"`
	ContainerImage      string `json:"containerImage,omitempty"`
	ContainerTag        string `json:"containerTag,omitempty"`
	ContainerEntrypoint string `json:"containerEntrypoint,omitempty"`
}

// ConfigStatus defines the observed state of Config
type ConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of pointers to currently running jobs.
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`

	// RFC 3339 date and time at which the object was acknowledged by the Kubelet.
	// This is before the Kubelet pulled the container image(s) for the pod.
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`

	// Status of the config
	// +optional
	InstallerStatus InstallerStatusType `json:"status,omitempty"`

	// Last pod name
	// +optional
	LastPodName string `json:"lastPodName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Config is the Schema for the configs API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
