/*
Copyright 2017 The etcd-operator Authors

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
	v1alpha1 "github.com/coreos/etcd-operator/pkg/apis/galera/v1alpha1"
	scheme "github.com/coreos/etcd-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GaleraBackupsGetter has a method to return a GaleraBackupInterface.
// A group's client should implement this interface.
type GaleraBackupsGetter interface {
	GaleraBackups(namespace string) GaleraBackupInterface
}

// GaleraBackupInterface has methods to work with GaleraBackup resources.
type GaleraBackupInterface interface {
	Create(*v1alpha1.GaleraBackup) (*v1alpha1.GaleraBackup, error)
	Update(*v1alpha1.GaleraBackup) (*v1alpha1.GaleraBackup, error)
	UpdateStatus(*v1alpha1.GaleraBackup) (*v1alpha1.GaleraBackup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GaleraBackup, error)
	List(opts v1.ListOptions) (*v1alpha1.GaleraBackupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GaleraBackup, err error)
	GaleraBackupExpansion
}

// galeraBackups implements GaleraBackupInterface
type galeraBackups struct {
	client rest.Interface
	ns     string
}

// newGaleraBackups returns a GaleraBackups
func newGaleraBackups(c *GaleraV1alpha1Client, namespace string) *galeraBackups {
	return &galeraBackups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the galeraBackup, and returns the corresponding galeraBackup object, and an error if there is any.
func (c *galeraBackups) Get(name string, options v1.GetOptions) (result *v1alpha1.GaleraBackup, err error) {
	result = &v1alpha1.GaleraBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galerabackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GaleraBackups that match those selectors.
func (c *galeraBackups) List(opts v1.ListOptions) (result *v1alpha1.GaleraBackupList, err error) {
	result = &v1alpha1.GaleraBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galerabackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested galeraBackups.
func (c *galeraBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("galerabackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a galeraBackup and creates it.  Returns the server's representation of the galeraBackup, and an error, if there is any.
func (c *galeraBackups) Create(galeraBackup *v1alpha1.GaleraBackup) (result *v1alpha1.GaleraBackup, err error) {
	result = &v1alpha1.GaleraBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("galerabackups").
		Body(galeraBackup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a galeraBackup and updates it. Returns the server's representation of the galeraBackup, and an error, if there is any.
func (c *galeraBackups) Update(galeraBackup *v1alpha1.GaleraBackup) (result *v1alpha1.GaleraBackup, err error) {
	result = &v1alpha1.GaleraBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galerabackups").
		Name(galeraBackup.Name).
		Body(galeraBackup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *galeraBackups) UpdateStatus(galeraBackup *v1alpha1.GaleraBackup) (result *v1alpha1.GaleraBackup, err error) {
	result = &v1alpha1.GaleraBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galerabackups").
		Name(galeraBackup.Name).
		SubResource("status").
		Body(galeraBackup).
		Do().
		Into(result)
	return
}

// Delete takes name of the galeraBackup and deletes it. Returns an error if one occurs.
func (c *galeraBackups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galerabackups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *galeraBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galerabackups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched galeraBackup.
func (c *galeraBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GaleraBackup, err error) {
	result = &v1alpha1.GaleraBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("galerabackups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
