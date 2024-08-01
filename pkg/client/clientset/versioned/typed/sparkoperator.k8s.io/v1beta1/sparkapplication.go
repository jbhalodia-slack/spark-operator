// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/kubeflow/spark-operator/api/v1beta1"
	scheme "github.com/kubeflow/spark-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SparkApplicationsGetter has a method to return a SparkApplicationInterface.
// A group's client should implement this interface.
type SparkApplicationsGetter interface {
	SparkApplications(namespace string) SparkApplicationInterface
}

// SparkApplicationInterface has methods to work with SparkApplication resources.
type SparkApplicationInterface interface {
	Create(ctx context.Context, sparkApplication *v1beta1.SparkApplication, opts v1.CreateOptions) (*v1beta1.SparkApplication, error)
	Update(ctx context.Context, sparkApplication *v1beta1.SparkApplication, opts v1.UpdateOptions) (*v1beta1.SparkApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.SparkApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SparkApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SparkApplication, err error)
	SparkApplicationExpansion
}

// sparkApplications implements SparkApplicationInterface
type sparkApplications struct {
	client rest.Interface
	ns     string
}

// newSparkApplications returns a SparkApplications
func newSparkApplications(c *SparkoperatorV1beta1Client, namespace string) *sparkApplications {
	return &sparkApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sparkApplication, and returns the corresponding sparkApplication object, and an error if there is any.
func (c *sparkApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SparkApplication, err error) {
	result = &v1beta1.SparkApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SparkApplications that match those selectors.
func (c *sparkApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SparkApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SparkApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sparkApplications.
func (c *sparkApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sparkApplication and creates it.  Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *sparkApplications) Create(ctx context.Context, sparkApplication *v1beta1.SparkApplication, opts v1.CreateOptions) (result *v1beta1.SparkApplication, err error) {
	result = &v1beta1.SparkApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sparkApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sparkApplication and updates it. Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *sparkApplications) Update(ctx context.Context, sparkApplication *v1beta1.SparkApplication, opts v1.UpdateOptions) (result *v1beta1.SparkApplication, err error) {
	result = &v1beta1.SparkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(sparkApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sparkApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sparkApplication and deletes it. Returns an error if one occurs.
func (c *sparkApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sparkApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sparkapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sparkApplication.
func (c *sparkApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SparkApplication, err error) {
	result = &v1beta1.SparkApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sparkapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
