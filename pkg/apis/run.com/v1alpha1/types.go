package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PipelineRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineRunSpec   `json:"spec"`
	Status PipelineRunStatus `json:"status"`
}

type PipelineRunSpec struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type PipelineRunStatus struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PipelineRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PipelineRun `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TaskRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TaskRunSpec   `json:"spec"`
	Status TaskRunStatus `json:"status"`
}

type TaskRunSpec struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type TaskRunStatus struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TaskRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TaskRun `json:"items"`
}

func (el *PipelineRun) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("PipelineRun")
}
