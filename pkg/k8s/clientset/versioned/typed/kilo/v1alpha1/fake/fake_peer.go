// Copyright 2021 the Kilo authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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

	v1alpha1 "github.com/squat/kilo/pkg/k8s/apis/kilo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePeers implements PeerInterface
type FakePeers struct {
	Fake *FakeKiloV1alpha1
}

var peersResource = schema.GroupVersionResource{Group: "kilo", Version: "v1alpha1", Resource: "peers"}

var peersKind = schema.GroupVersionKind{Group: "kilo", Version: "v1alpha1", Kind: "Peer"}

// Get takes name of the peer, and returns the corresponding peer object, and an error if there is any.
func (c *FakePeers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Peer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(peersResource, name), &v1alpha1.Peer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Peer), err
}

// List takes label and field selectors, and returns the list of Peers that match those selectors.
func (c *FakePeers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PeerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(peersResource, peersKind, opts), &v1alpha1.PeerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PeerList{ListMeta: obj.(*v1alpha1.PeerList).ListMeta}
	for _, item := range obj.(*v1alpha1.PeerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested peers.
func (c *FakePeers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(peersResource, opts))
}

// Create takes the representation of a peer and creates it.  Returns the server's representation of the peer, and an error, if there is any.
func (c *FakePeers) Create(ctx context.Context, peer *v1alpha1.Peer, opts v1.CreateOptions) (result *v1alpha1.Peer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(peersResource, peer), &v1alpha1.Peer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Peer), err
}

// Update takes the representation of a peer and updates it. Returns the server's representation of the peer, and an error, if there is any.
func (c *FakePeers) Update(ctx context.Context, peer *v1alpha1.Peer, opts v1.UpdateOptions) (result *v1alpha1.Peer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(peersResource, peer), &v1alpha1.Peer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Peer), err
}

// Delete takes name of the peer and deletes it. Returns an error if one occurs.
func (c *FakePeers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(peersResource, name), &v1alpha1.Peer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePeers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(peersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PeerList{})
	return err
}

// Patch applies the patch and returns the patched peer.
func (c *FakePeers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Peer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(peersResource, name, pt, data, subresources...), &v1alpha1.Peer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Peer), err
}
