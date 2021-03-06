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

package v1

import (
	"time"

	v1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	scheme "github.com/F5Networks/k8s-bigip-ctlr/config/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TransportServersGetter has a method to return a TransportServerInterface.
// A group's client should implement this interface.
type TransportServersGetter interface {
	TransportServers(namespace string) TransportServerInterface
}

// TransportServerInterface has methods to work with TransportServer resources.
type TransportServerInterface interface {
	Create(*v1.TransportServer) (*v1.TransportServer, error)
	Update(*v1.TransportServer) (*v1.TransportServer, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.TransportServer, error)
	List(opts metav1.ListOptions) (*v1.TransportServerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TransportServer, err error)
	TransportServerExpansion
}

// transportServers implements TransportServerInterface
type transportServers struct {
	client rest.Interface
	ns     string
}

// newTransportServers returns a TransportServers
func newTransportServers(c *K8sV1Client, namespace string) *transportServers {
	return &transportServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the transportServer, and returns the corresponding transportServer object, and an error if there is any.
func (c *transportServers) Get(name string, options metav1.GetOptions) (result *v1.TransportServer, err error) {
	result = &v1.TransportServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transportservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TransportServers that match those selectors.
func (c *transportServers) List(opts metav1.ListOptions) (result *v1.TransportServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TransportServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transportservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested transportServers.
func (c *transportServers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("transportservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a transportServer and creates it.  Returns the server's representation of the transportServer, and an error, if there is any.
func (c *transportServers) Create(transportServer *v1.TransportServer) (result *v1.TransportServer, err error) {
	result = &v1.TransportServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("transportservers").
		Body(transportServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a transportServer and updates it. Returns the server's representation of the transportServer, and an error, if there is any.
func (c *transportServers) Update(transportServer *v1.TransportServer) (result *v1.TransportServer, err error) {
	result = &v1.TransportServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transportservers").
		Name(transportServer.Name).
		Body(transportServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the transportServer and deletes it. Returns an error if one occurs.
func (c *transportServers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transportservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *transportServers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transportservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched transportServer.
func (c *transportServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TransportServer, err error) {
	result = &v1.TransportServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("transportservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
