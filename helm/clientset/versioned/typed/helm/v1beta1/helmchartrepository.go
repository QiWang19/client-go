// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/openshift/api/helm/v1beta1"
	scheme "github.com/openshift/client-go/helm/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HelmChartRepositoriesGetter has a method to return a HelmChartRepositoryInterface.
// A group's client should implement this interface.
type HelmChartRepositoriesGetter interface {
	HelmChartRepositories() HelmChartRepositoryInterface
}

// HelmChartRepositoryInterface has methods to work with HelmChartRepository resources.
type HelmChartRepositoryInterface interface {
	Create(ctx context.Context, helmChartRepository *v1beta1.HelmChartRepository, opts v1.CreateOptions) (*v1beta1.HelmChartRepository, error)
	Update(ctx context.Context, helmChartRepository *v1beta1.HelmChartRepository, opts v1.UpdateOptions) (*v1beta1.HelmChartRepository, error)
	UpdateStatus(ctx context.Context, helmChartRepository *v1beta1.HelmChartRepository, opts v1.UpdateOptions) (*v1beta1.HelmChartRepository, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.HelmChartRepository, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.HelmChartRepositoryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.HelmChartRepository, err error)
	HelmChartRepositoryExpansion
}

// helmChartRepositories implements HelmChartRepositoryInterface
type helmChartRepositories struct {
	client rest.Interface
}

// newHelmChartRepositories returns a HelmChartRepositories
func newHelmChartRepositories(c *HelmV1beta1Client) *helmChartRepositories {
	return &helmChartRepositories{
		client: c.RESTClient(),
	}
}

// Get takes name of the helmChartRepository, and returns the corresponding helmChartRepository object, and an error if there is any.
func (c *helmChartRepositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.HelmChartRepository, err error) {
	result = &v1beta1.HelmChartRepository{}
	err = c.client.Get().
		Resource("helmchartrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelmChartRepositories that match those selectors.
func (c *helmChartRepositories) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.HelmChartRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.HelmChartRepositoryList{}
	err = c.client.Get().
		Resource("helmchartrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helmChartRepositories.
func (c *helmChartRepositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("helmchartrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a helmChartRepository and creates it.  Returns the server's representation of the helmChartRepository, and an error, if there is any.
func (c *helmChartRepositories) Create(ctx context.Context, helmChartRepository *v1beta1.HelmChartRepository, opts v1.CreateOptions) (result *v1beta1.HelmChartRepository, err error) {
	result = &v1beta1.HelmChartRepository{}
	err = c.client.Post().
		Resource("helmchartrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmChartRepository).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a helmChartRepository and updates it. Returns the server's representation of the helmChartRepository, and an error, if there is any.
func (c *helmChartRepositories) Update(ctx context.Context, helmChartRepository *v1beta1.HelmChartRepository, opts v1.UpdateOptions) (result *v1beta1.HelmChartRepository, err error) {
	result = &v1beta1.HelmChartRepository{}
	err = c.client.Put().
		Resource("helmchartrepositories").
		Name(helmChartRepository.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmChartRepository).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *helmChartRepositories) UpdateStatus(ctx context.Context, helmChartRepository *v1beta1.HelmChartRepository, opts v1.UpdateOptions) (result *v1beta1.HelmChartRepository, err error) {
	result = &v1beta1.HelmChartRepository{}
	err = c.client.Put().
		Resource("helmchartrepositories").
		Name(helmChartRepository.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmChartRepository).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the helmChartRepository and deletes it. Returns an error if one occurs.
func (c *helmChartRepositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("helmchartrepositories").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helmChartRepositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("helmchartrepositories").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched helmChartRepository.
func (c *helmChartRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.HelmChartRepository, err error) {
	result = &v1beta1.HelmChartRepository{}
	err = c.client.Patch(pt).
		Resource("helmchartrepositories").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}