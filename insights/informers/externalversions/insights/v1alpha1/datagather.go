// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	insightsv1alpha1 "github.com/openshift/api/insights/v1alpha1"
	versioned "github.com/openshift/client-go/insights/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/insights/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/client-go/insights/listers/insights/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DataGatherInformer provides access to a shared informer and lister for
// DataGathers.
type DataGatherInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DataGatherLister
}

type dataGatherInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDataGatherInformer constructs a new informer for DataGather type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDataGatherInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDataGatherInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDataGatherInformer constructs a new informer for DataGather type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDataGatherInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InsightsV1alpha1().DataGathers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InsightsV1alpha1().DataGathers(namespace).Watch(context.TODO(), options)
			},
		},
		&insightsv1alpha1.DataGather{},
		resyncPeriod,
		indexers,
	)
}

func (f *dataGatherInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDataGatherInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dataGatherInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&insightsv1alpha1.DataGather{}, f.defaultInformer)
}

func (f *dataGatherInformer) Lister() v1alpha1.DataGatherLister {
	return v1alpha1.NewDataGatherLister(f.Informer().GetIndexer())
}
