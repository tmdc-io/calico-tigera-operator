// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	projectcalicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// GlobalNetworkSetLister helps list GlobalNetworkSets.
// All objects returned here must be treated as read-only.
type GlobalNetworkSetLister interface {
	// List lists all GlobalNetworkSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*projectcalicov3.GlobalNetworkSet, err error)
	// Get retrieves the GlobalNetworkSet from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*projectcalicov3.GlobalNetworkSet, error)
	GlobalNetworkSetListerExpansion
}

// globalNetworkSetLister implements the GlobalNetworkSetLister interface.
type globalNetworkSetLister struct {
	listers.ResourceIndexer[*projectcalicov3.GlobalNetworkSet]
}

// NewGlobalNetworkSetLister returns a new GlobalNetworkSetLister.
func NewGlobalNetworkSetLister(indexer cache.Indexer) GlobalNetworkSetLister {
	return &globalNetworkSetLister{listers.New[*projectcalicov3.GlobalNetworkSet](indexer, projectcalicov3.Resource("globalnetworkset"))}
}
