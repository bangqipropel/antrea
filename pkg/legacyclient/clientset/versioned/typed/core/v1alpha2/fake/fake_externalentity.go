// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "antrea.io/antrea/pkg/legacyapis/core/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExternalEntities implements ExternalEntityInterface
type FakeExternalEntities struct {
	Fake *FakeCoreV1alpha2
	ns   string
}

var externalentitiesResource = schema.GroupVersionResource{Group: "core.antrea.tanzu.vmware.com", Version: "v1alpha2", Resource: "externalentities"}

var externalentitiesKind = schema.GroupVersionKind{Group: "core.antrea.tanzu.vmware.com", Version: "v1alpha2", Kind: "ExternalEntity"}

// Get takes name of the externalEntity, and returns the corresponding externalEntity object, and an error if there is any.
func (c *FakeExternalEntities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ExternalEntity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalentitiesResource, c.ns, name), &v1alpha2.ExternalEntity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ExternalEntity), err
}

// List takes label and field selectors, and returns the list of ExternalEntities that match those selectors.
func (c *FakeExternalEntities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ExternalEntityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalentitiesResource, externalentitiesKind, c.ns, opts), &v1alpha2.ExternalEntityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ExternalEntityList{ListMeta: obj.(*v1alpha2.ExternalEntityList).ListMeta}
	for _, item := range obj.(*v1alpha2.ExternalEntityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalEntities.
func (c *FakeExternalEntities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalentitiesResource, c.ns, opts))

}

// Create takes the representation of a externalEntity and creates it.  Returns the server's representation of the externalEntity, and an error, if there is any.
func (c *FakeExternalEntities) Create(ctx context.Context, externalEntity *v1alpha2.ExternalEntity, opts v1.CreateOptions) (result *v1alpha2.ExternalEntity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalentitiesResource, c.ns, externalEntity), &v1alpha2.ExternalEntity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ExternalEntity), err
}

// Update takes the representation of a externalEntity and updates it. Returns the server's representation of the externalEntity, and an error, if there is any.
func (c *FakeExternalEntities) Update(ctx context.Context, externalEntity *v1alpha2.ExternalEntity, opts v1.UpdateOptions) (result *v1alpha2.ExternalEntity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalentitiesResource, c.ns, externalEntity), &v1alpha2.ExternalEntity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ExternalEntity), err
}

// Delete takes name of the externalEntity and deletes it. Returns an error if one occurs.
func (c *FakeExternalEntities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(externalentitiesResource, c.ns, name), &v1alpha2.ExternalEntity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalEntities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalentitiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ExternalEntityList{})
	return err
}

// Patch applies the patch and returns the patched externalEntity.
func (c *FakeExternalEntities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ExternalEntity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalentitiesResource, c.ns, name, pt, data, subresources...), &v1alpha2.ExternalEntity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ExternalEntity), err
}
