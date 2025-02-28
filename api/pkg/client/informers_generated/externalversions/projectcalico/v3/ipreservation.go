// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	context "context"
	time "time"

	apisprojectcalicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	clientset "github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/projectcalico/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	projectcalicov3 "github.com/projectcalico/api/pkg/client/listers_generated/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IPReservationInformer provides access to a shared informer and lister for
// IPReservations.
type IPReservationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() projectcalicov3.IPReservationLister
}

type iPReservationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIPReservationInformer constructs a new informer for IPReservation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIPReservationInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIPReservationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIPReservationInformer constructs a new informer for IPReservation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIPReservationInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().IPReservations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().IPReservations().Watch(context.TODO(), options)
			},
		},
		&apisprojectcalicov3.IPReservation{},
		resyncPeriod,
		indexers,
	)
}

func (f *iPReservationInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIPReservationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iPReservationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisprojectcalicov3.IPReservation{}, f.defaultInformer)
}

func (f *iPReservationInformer) Lister() projectcalicov3.IPReservationLister {
	return projectcalicov3.NewIPReservationLister(f.Informer().GetIndexer())
}
