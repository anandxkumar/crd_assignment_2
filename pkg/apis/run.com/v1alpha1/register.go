package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{
	Group:   "run.com",
	Version: "v1alpha1",
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
// func Kind(kind string) schema.GroupKind {
// 	return SchemeGroupVersion.WithKind(kind).GroupKind()
// }

// // Resource takes an unqualified resource and returns a Group qualified GroupResource
// func Resource(resource string) schema.GroupResource {
// 	return SchemeGroupVersion.WithResource(resource).GroupResource()
// }

var (
	// SchemeBuilder initializes a scheme builder
	SchemeBuilder runtime.SchemeBuilder
	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

// Adding resource function for listers
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&PipelineRun{},
		&PipelineRunList{},
		&TaskRun{},
		&TaskRunList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
