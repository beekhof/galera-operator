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
package fake

import (
	v1alpha1 "github.com/coreos/etcd-operator/pkg/apis/galera/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGaleraBackups implements GaleraBackupInterface
type FakeGaleraBackups struct {
	Fake *FakeGaleraV1alpha1
	ns   string
}

var galerabackupsResource = schema.GroupVersionResource{Group: "galera.database.beekhof.net", Version: "v1alpha1", Resource: "galerabackups"}

var galerabackupsKind = schema.GroupVersionKind{Group: "galera.database.beekhof.net", Version: "v1alpha1", Kind: "GaleraBackup"}

// Get takes name of the galeraBackup, and returns the corresponding galeraBackup object, and an error if there is any.
func (c *FakeGaleraBackups) Get(name string, options v1.GetOptions) (result *v1alpha1.GaleraBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(galerabackupsResource, c.ns, name), &v1alpha1.GaleraBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GaleraBackup), err
}

// List takes label and field selectors, and returns the list of GaleraBackups that match those selectors.
func (c *FakeGaleraBackups) List(opts v1.ListOptions) (result *v1alpha1.GaleraBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(galerabackupsResource, galerabackupsKind, c.ns, opts), &v1alpha1.GaleraBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GaleraBackupList{}
	for _, item := range obj.(*v1alpha1.GaleraBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested galeraBackups.
func (c *FakeGaleraBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(galerabackupsResource, c.ns, opts))

}

// Create takes the representation of a galeraBackup and creates it.  Returns the server's representation of the galeraBackup, and an error, if there is any.
func (c *FakeGaleraBackups) Create(galeraBackup *v1alpha1.GaleraBackup) (result *v1alpha1.GaleraBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(galerabackupsResource, c.ns, galeraBackup), &v1alpha1.GaleraBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GaleraBackup), err
}

// Update takes the representation of a galeraBackup and updates it. Returns the server's representation of the galeraBackup, and an error, if there is any.
func (c *FakeGaleraBackups) Update(galeraBackup *v1alpha1.GaleraBackup) (result *v1alpha1.GaleraBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(galerabackupsResource, c.ns, galeraBackup), &v1alpha1.GaleraBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GaleraBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGaleraBackups) UpdateStatus(galeraBackup *v1alpha1.GaleraBackup) (*v1alpha1.GaleraBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(galerabackupsResource, "status", c.ns, galeraBackup), &v1alpha1.GaleraBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GaleraBackup), err
}

// Delete takes name of the galeraBackup and deletes it. Returns an error if one occurs.
func (c *FakeGaleraBackups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(galerabackupsResource, c.ns, name), &v1alpha1.GaleraBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGaleraBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(galerabackupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GaleraBackupList{})
	return err
}

// Patch applies the patch and returns the patched galeraBackup.
func (c *FakeGaleraBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GaleraBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(galerabackupsResource, c.ns, name, data, subresources...), &v1alpha1.GaleraBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GaleraBackup), err
}
