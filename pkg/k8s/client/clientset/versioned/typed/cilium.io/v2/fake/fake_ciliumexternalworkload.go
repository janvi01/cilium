// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumExternalWorkloads implements CiliumExternalWorkloadInterface
type FakeCiliumExternalWorkloads struct {
	Fake *FakeCiliumV2
}

var ciliumexternalworkloadsResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumexternalworkloads"}

var ciliumexternalworkloadsKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2", Kind: "CiliumExternalWorkload"}

// Get takes name of the ciliumExternalWorkload, and returns the corresponding ciliumExternalWorkload object, and an error if there is any.
func (c *FakeCiliumExternalWorkloads) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumExternalWorkload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumexternalworkloadsResource, name), &v2.CiliumExternalWorkload{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumExternalWorkload), err
}

// List takes label and field selectors, and returns the list of CiliumExternalWorkloads that match those selectors.
func (c *FakeCiliumExternalWorkloads) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumExternalWorkloadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumexternalworkloadsResource, ciliumexternalworkloadsKind, opts), &v2.CiliumExternalWorkloadList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CiliumExternalWorkloadList{ListMeta: obj.(*v2.CiliumExternalWorkloadList).ListMeta}
	for _, item := range obj.(*v2.CiliumExternalWorkloadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumExternalWorkloads.
func (c *FakeCiliumExternalWorkloads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumexternalworkloadsResource, opts))
}

// Create takes the representation of a ciliumExternalWorkload and creates it.  Returns the server's representation of the ciliumExternalWorkload, and an error, if there is any.
func (c *FakeCiliumExternalWorkloads) Create(ctx context.Context, ciliumExternalWorkload *v2.CiliumExternalWorkload, opts v1.CreateOptions) (result *v2.CiliumExternalWorkload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumexternalworkloadsResource, ciliumExternalWorkload), &v2.CiliumExternalWorkload{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumExternalWorkload), err
}

// Update takes the representation of a ciliumExternalWorkload and updates it. Returns the server's representation of the ciliumExternalWorkload, and an error, if there is any.
func (c *FakeCiliumExternalWorkloads) Update(ctx context.Context, ciliumExternalWorkload *v2.CiliumExternalWorkload, opts v1.UpdateOptions) (result *v2.CiliumExternalWorkload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumexternalworkloadsResource, ciliumExternalWorkload), &v2.CiliumExternalWorkload{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumExternalWorkload), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCiliumExternalWorkloads) UpdateStatus(ctx context.Context, ciliumExternalWorkload *v2.CiliumExternalWorkload, opts v1.UpdateOptions) (*v2.CiliumExternalWorkload, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ciliumexternalworkloadsResource, "status", ciliumExternalWorkload), &v2.CiliumExternalWorkload{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumExternalWorkload), err
}

// Delete takes name of the ciliumExternalWorkload and deletes it. Returns an error if one occurs.
func (c *FakeCiliumExternalWorkloads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumexternalworkloadsResource, name, opts), &v2.CiliumExternalWorkload{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumExternalWorkloads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumexternalworkloadsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CiliumExternalWorkloadList{})
	return err
}

// Patch applies the patch and returns the patched ciliumExternalWorkload.
func (c *FakeCiliumExternalWorkloads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumExternalWorkload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumexternalworkloadsResource, name, pt, data, subresources...), &v2.CiliumExternalWorkload{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumExternalWorkload), err
}
