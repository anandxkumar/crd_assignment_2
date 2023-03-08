/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/anandxkumar/crd_assignment_2/pkg/apis/run.com/v1alpha1"
	scheme "github.com/anandxkumar/crd_assignment_2/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PipelineRunsGetter has a method to return a PipelineRunInterface.
// A group's client should implement this interface.
type PipelineRunsGetter interface {
	PipelineRuns(namespace string) PipelineRunInterface
}

// PipelineRunInterface has methods to work with PipelineRun resources.
type PipelineRunInterface interface {
	Create(ctx context.Context, pipelineRun *v1alpha1.PipelineRun, opts v1.CreateOptions) (*v1alpha1.PipelineRun, error)
	Update(ctx context.Context, pipelineRun *v1alpha1.PipelineRun, opts v1.UpdateOptions) (*v1alpha1.PipelineRun, error)
	UpdateStatus(ctx context.Context, pipelineRun *v1alpha1.PipelineRun, opts v1.UpdateOptions) (*v1alpha1.PipelineRun, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PipelineRun, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PipelineRunList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PipelineRun, err error)
	PipelineRunExpansion
}

// pipelineRuns implements PipelineRunInterface
type pipelineRuns struct {
	client rest.Interface
	ns     string
}

// newPipelineRuns returns a PipelineRuns
func newPipelineRuns(c *RunV1alpha1Client, namespace string) *pipelineRuns {
	return &pipelineRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pipelineRun, and returns the corresponding pipelineRun object, and an error if there is any.
func (c *pipelineRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PipelineRun, err error) {
	result = &v1alpha1.PipelineRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pipelineruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PipelineRuns that match those selectors.
func (c *pipelineRuns) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PipelineRunList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PipelineRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pipelineruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pipelineRuns.
func (c *pipelineRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pipelineruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pipelineRun and creates it.  Returns the server's representation of the pipelineRun, and an error, if there is any.
func (c *pipelineRuns) Create(ctx context.Context, pipelineRun *v1alpha1.PipelineRun, opts v1.CreateOptions) (result *v1alpha1.PipelineRun, err error) {
	result = &v1alpha1.PipelineRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pipelineruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pipelineRun).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pipelineRun and updates it. Returns the server's representation of the pipelineRun, and an error, if there is any.
func (c *pipelineRuns) Update(ctx context.Context, pipelineRun *v1alpha1.PipelineRun, opts v1.UpdateOptions) (result *v1alpha1.PipelineRun, err error) {
	result = &v1alpha1.PipelineRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pipelineruns").
		Name(pipelineRun.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pipelineRun).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *pipelineRuns) UpdateStatus(ctx context.Context, pipelineRun *v1alpha1.PipelineRun, opts v1.UpdateOptions) (result *v1alpha1.PipelineRun, err error) {
	result = &v1alpha1.PipelineRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pipelineruns").
		Name(pipelineRun.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pipelineRun).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pipelineRun and deletes it. Returns an error if one occurs.
func (c *pipelineRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pipelineruns").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pipelineRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pipelineruns").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pipelineRun.
func (c *pipelineRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PipelineRun, err error) {
	result = &v1alpha1.PipelineRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pipelineruns").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
