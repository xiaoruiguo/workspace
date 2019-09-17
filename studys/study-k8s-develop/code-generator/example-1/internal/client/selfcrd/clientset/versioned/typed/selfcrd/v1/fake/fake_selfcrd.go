/*
Copyright 2019 The Lijiao Authors.

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

package fake

import (
	selfcrdv1 "github.com/lijiaocn/workspace/studys/study-k8s-develop/code-generator/example-1/internal/apis/selfcrd/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSelfCRDs implements SelfCRDInterface
type FakeSelfCRDs struct {
	Fake *FakeCrdV1
	ns   string
}

var selfcrdsResource = schema.GroupVersionResource{Group: "crd.lijiaocn.com", Version: "v1", Resource: "selfcrds"}

var selfcrdsKind = schema.GroupVersionKind{Group: "crd.lijiaocn.com", Version: "v1", Kind: "SelfCRD"}

// Get takes name of the selfCRD, and returns the corresponding selfCRD object, and an error if there is any.
func (c *FakeSelfCRDs) Get(name string, options v1.GetOptions) (result *selfcrdv1.SelfCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(selfcrdsResource, c.ns, name), &selfcrdv1.SelfCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*selfcrdv1.SelfCRD), err
}

// List takes label and field selectors, and returns the list of SelfCRDs that match those selectors.
func (c *FakeSelfCRDs) List(opts v1.ListOptions) (result *selfcrdv1.SelfCRDList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(selfcrdsResource, selfcrdsKind, c.ns, opts), &selfcrdv1.SelfCRDList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &selfcrdv1.SelfCRDList{ListMeta: obj.(*selfcrdv1.SelfCRDList).ListMeta}
	for _, item := range obj.(*selfcrdv1.SelfCRDList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested selfCRDs.
func (c *FakeSelfCRDs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(selfcrdsResource, c.ns, opts))

}

// Create takes the representation of a selfCRD and creates it.  Returns the server's representation of the selfCRD, and an error, if there is any.
func (c *FakeSelfCRDs) Create(selfCRD *selfcrdv1.SelfCRD) (result *selfcrdv1.SelfCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(selfcrdsResource, c.ns, selfCRD), &selfcrdv1.SelfCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*selfcrdv1.SelfCRD), err
}

// Update takes the representation of a selfCRD and updates it. Returns the server's representation of the selfCRD, and an error, if there is any.
func (c *FakeSelfCRDs) Update(selfCRD *selfcrdv1.SelfCRD) (result *selfcrdv1.SelfCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(selfcrdsResource, c.ns, selfCRD), &selfcrdv1.SelfCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*selfcrdv1.SelfCRD), err
}

// Delete takes name of the selfCRD and deletes it. Returns an error if one occurs.
func (c *FakeSelfCRDs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(selfcrdsResource, c.ns, name), &selfcrdv1.SelfCRD{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSelfCRDs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(selfcrdsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &selfcrdv1.SelfCRDList{})
	return err
}

// Patch applies the patch and returns the patched selfCRD.
func (c *FakeSelfCRDs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *selfcrdv1.SelfCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(selfcrdsResource, c.ns, name, pt, data, subresources...), &selfcrdv1.SelfCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*selfcrdv1.SelfCRD), err
}
