/*
Copyright The Kubectl Bind contributors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
	scheme "github.com/kube-bind/kube-bind/pkg/client/clientset/versioned/scheme"
)

// ServiceNamespacesGetter has a method to return a ServiceNamespaceInterface.
// A group's client should implement this interface.
type ServiceNamespacesGetter interface {
	ServiceNamespaces(namespace string) ServiceNamespaceInterface
}

// ServiceNamespaceInterface has methods to work with ServiceNamespace resources.
type ServiceNamespaceInterface interface {
	Create(ctx context.Context, serviceNamespace *v1alpha1.ServiceNamespace, opts v1.CreateOptions) (*v1alpha1.ServiceNamespace, error)
	Update(ctx context.Context, serviceNamespace *v1alpha1.ServiceNamespace, opts v1.UpdateOptions) (*v1alpha1.ServiceNamespace, error)
	UpdateStatus(ctx context.Context, serviceNamespace *v1alpha1.ServiceNamespace, opts v1.UpdateOptions) (*v1alpha1.ServiceNamespace, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceNamespace, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceNamespaceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceNamespace, err error)
	ServiceNamespaceExpansion
}

// serviceNamespaces implements ServiceNamespaceInterface
type serviceNamespaces struct {
	client rest.Interface
	ns     string
}

// newServiceNamespaces returns a ServiceNamespaces
func newServiceNamespaces(c *KubeBindV1alpha1Client, namespace string) *serviceNamespaces {
	return &serviceNamespaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceNamespace, and returns the corresponding serviceNamespace object, and an error if there is any.
func (c *serviceNamespaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceNamespace, err error) {
	result = &v1alpha1.ServiceNamespace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceNamespaces that match those selectors.
func (c *serviceNamespaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceNamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceNamespaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceNamespaces.
func (c *serviceNamespaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicenamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceNamespace and creates it.  Returns the server's representation of the serviceNamespace, and an error, if there is any.
func (c *serviceNamespaces) Create(ctx context.Context, serviceNamespace *v1alpha1.ServiceNamespace, opts v1.CreateOptions) (result *v1alpha1.ServiceNamespace, err error) {
	result = &v1alpha1.ServiceNamespace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicenamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNamespace).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceNamespace and updates it. Returns the server's representation of the serviceNamespace, and an error, if there is any.
func (c *serviceNamespaces) Update(ctx context.Context, serviceNamespace *v1alpha1.ServiceNamespace, opts v1.UpdateOptions) (result *v1alpha1.ServiceNamespace, err error) {
	result = &v1alpha1.ServiceNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenamespaces").
		Name(serviceNamespace.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNamespace).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceNamespaces) UpdateStatus(ctx context.Context, serviceNamespace *v1alpha1.ServiceNamespace, opts v1.UpdateOptions) (result *v1alpha1.ServiceNamespace, err error) {
	result = &v1alpha1.ServiceNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenamespaces").
		Name(serviceNamespace.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNamespace).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceNamespace and deletes it. Returns an error if one occurs.
func (c *serviceNamespaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenamespaces").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceNamespaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenamespaces").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceNamespace.
func (c *serviceNamespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceNamespace, err error) {
	result = &v1alpha1.ServiceNamespace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicenamespaces").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}