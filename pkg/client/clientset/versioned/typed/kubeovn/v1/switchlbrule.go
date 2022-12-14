/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	scheme "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SwitchLBRulesGetter has a method to return a SwitchLBRuleInterface.
// A group's client should implement this interface.
type SwitchLBRulesGetter interface {
	SwitchLBRules() SwitchLBRuleInterface
}

// SwitchLBRuleInterface has methods to work with SwitchLBRule resources.
type SwitchLBRuleInterface interface {
	Create(ctx context.Context, switchLBRule *v1.SwitchLBRule, opts metav1.CreateOptions) (*v1.SwitchLBRule, error)
	Update(ctx context.Context, switchLBRule *v1.SwitchLBRule, opts metav1.UpdateOptions) (*v1.SwitchLBRule, error)
	UpdateStatus(ctx context.Context, switchLBRule *v1.SwitchLBRule, opts metav1.UpdateOptions) (*v1.SwitchLBRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SwitchLBRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SwitchLBRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SwitchLBRule, err error)
	SwitchLBRuleExpansion
}

// switchLBRules implements SwitchLBRuleInterface
type switchLBRules struct {
	client rest.Interface
}

// newSwitchLBRules returns a SwitchLBRules
func newSwitchLBRules(c *KubeovnV1Client) *switchLBRules {
	return &switchLBRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the switchLBRule, and returns the corresponding switchLBRule object, and an error if there is any.
func (c *switchLBRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SwitchLBRule, err error) {
	result = &v1.SwitchLBRule{}
	err = c.client.Get().
		Resource("switch-lb-rules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SwitchLBRules that match those selectors.
func (c *switchLBRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SwitchLBRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SwitchLBRuleList{}
	err = c.client.Get().
		Resource("switch-lb-rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested switchLBRules.
func (c *switchLBRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("switch-lb-rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a switchLBRule and creates it.  Returns the server's representation of the switchLBRule, and an error, if there is any.
func (c *switchLBRules) Create(ctx context.Context, switchLBRule *v1.SwitchLBRule, opts metav1.CreateOptions) (result *v1.SwitchLBRule, err error) {
	result = &v1.SwitchLBRule{}
	err = c.client.Post().
		Resource("switch-lb-rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(switchLBRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a switchLBRule and updates it. Returns the server's representation of the switchLBRule, and an error, if there is any.
func (c *switchLBRules) Update(ctx context.Context, switchLBRule *v1.SwitchLBRule, opts metav1.UpdateOptions) (result *v1.SwitchLBRule, err error) {
	result = &v1.SwitchLBRule{}
	err = c.client.Put().
		Resource("switch-lb-rules").
		Name(switchLBRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(switchLBRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *switchLBRules) UpdateStatus(ctx context.Context, switchLBRule *v1.SwitchLBRule, opts metav1.UpdateOptions) (result *v1.SwitchLBRule, err error) {
	result = &v1.SwitchLBRule{}
	err = c.client.Put().
		Resource("switch-lb-rules").
		Name(switchLBRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(switchLBRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the switchLBRule and deletes it. Returns an error if one occurs.
func (c *switchLBRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("switch-lb-rules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *switchLBRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("switch-lb-rules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched switchLBRule.
func (c *switchLBRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SwitchLBRule, err error) {
	result = &v1.SwitchLBRule{}
	err = c.client.Patch(pt).
		Resource("switch-lb-rules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
