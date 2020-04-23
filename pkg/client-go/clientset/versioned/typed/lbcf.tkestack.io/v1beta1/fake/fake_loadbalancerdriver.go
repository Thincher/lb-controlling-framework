/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "tkestack.io/lb-controlling-framework/pkg/apis/lbcf.tkestack.io/v1beta1"
)

// FakeLoadBalancerDrivers implements LoadBalancerDriverInterface
type FakeLoadBalancerDrivers struct {
	Fake *FakeLbcfV1beta1
	ns   string
}

var loadbalancerdriversResource = schema.GroupVersionResource{Group: "lbcf.tkestack.io", Version: "v1beta1", Resource: "loadbalancerdrivers"}

var loadbalancerdriversKind = schema.GroupVersionKind{Group: "lbcf.tkestack.io", Version: "v1beta1", Kind: "LoadBalancerDriver"}

// Get takes name of the loadBalancerDriver, and returns the corresponding loadBalancerDriver object, and an error if there is any.
func (c *FakeLoadBalancerDrivers) Get(name string, options v1.GetOptions) (result *v1beta1.LoadBalancerDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loadbalancerdriversResource, c.ns, name), &v1beta1.LoadBalancerDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoadBalancerDriver), err
}

// List takes label and field selectors, and returns the list of LoadBalancerDrivers that match those selectors.
func (c *FakeLoadBalancerDrivers) List(opts v1.ListOptions) (result *v1beta1.LoadBalancerDriverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loadbalancerdriversResource, loadbalancerdriversKind, c.ns, opts), &v1beta1.LoadBalancerDriverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.LoadBalancerDriverList{ListMeta: obj.(*v1beta1.LoadBalancerDriverList).ListMeta}
	for _, item := range obj.(*v1beta1.LoadBalancerDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loadBalancerDrivers.
func (c *FakeLoadBalancerDrivers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loadbalancerdriversResource, c.ns, opts))

}

// Create takes the representation of a loadBalancerDriver and creates it.  Returns the server's representation of the loadBalancerDriver, and an error, if there is any.
func (c *FakeLoadBalancerDrivers) Create(loadBalancerDriver *v1beta1.LoadBalancerDriver) (result *v1beta1.LoadBalancerDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loadbalancerdriversResource, c.ns, loadBalancerDriver), &v1beta1.LoadBalancerDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoadBalancerDriver), err
}

// Update takes the representation of a loadBalancerDriver and updates it. Returns the server's representation of the loadBalancerDriver, and an error, if there is any.
func (c *FakeLoadBalancerDrivers) Update(loadBalancerDriver *v1beta1.LoadBalancerDriver) (result *v1beta1.LoadBalancerDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loadbalancerdriversResource, c.ns, loadBalancerDriver), &v1beta1.LoadBalancerDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoadBalancerDriver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoadBalancerDrivers) UpdateStatus(loadBalancerDriver *v1beta1.LoadBalancerDriver) (*v1beta1.LoadBalancerDriver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loadbalancerdriversResource, "status", c.ns, loadBalancerDriver), &v1beta1.LoadBalancerDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoadBalancerDriver), err
}

// Delete takes name of the loadBalancerDriver and deletes it. Returns an error if one occurs.
func (c *FakeLoadBalancerDrivers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loadbalancerdriversResource, c.ns, name), &v1beta1.LoadBalancerDriver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoadBalancerDrivers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loadbalancerdriversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.LoadBalancerDriverList{})
	return err
}

// Patch applies the patch and returns the patched loadBalancerDriver.
func (c *FakeLoadBalancerDrivers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.LoadBalancerDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loadbalancerdriversResource, c.ns, name, pt, data, subresources...), &v1beta1.LoadBalancerDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LoadBalancerDriver), err
}