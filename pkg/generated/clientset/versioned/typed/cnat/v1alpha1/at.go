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

package v1alpha1

import (
	context "context"

	cnatv1alpha1 "github.com/soaib024/cnat-controller/pkg/apis/cnat/v1alpha1"
	scheme "github.com/soaib024/cnat-controller/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AtsGetter has a method to return a AtInterface.
// A group's client should implement this interface.
type AtsGetter interface {
	Ats(namespace string) AtInterface
}

// AtInterface has methods to work with At resources.
type AtInterface interface {
	Create(ctx context.Context, at *cnatv1alpha1.At, opts v1.CreateOptions) (*cnatv1alpha1.At, error)
	Update(ctx context.Context, at *cnatv1alpha1.At, opts v1.UpdateOptions) (*cnatv1alpha1.At, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, at *cnatv1alpha1.At, opts v1.UpdateOptions) (*cnatv1alpha1.At, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*cnatv1alpha1.At, error)
	List(ctx context.Context, opts v1.ListOptions) (*cnatv1alpha1.AtList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cnatv1alpha1.At, err error)
	AtExpansion
}

// ats implements AtInterface
type ats struct {
	*gentype.ClientWithList[*cnatv1alpha1.At, *cnatv1alpha1.AtList]
}

// newAts returns a Ats
func newAts(c *CnatV1alpha1Client, namespace string) *ats {
	return &ats{
		gentype.NewClientWithList[*cnatv1alpha1.At, *cnatv1alpha1.AtList](
			"ats",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *cnatv1alpha1.At { return &cnatv1alpha1.At{} },
			func() *cnatv1alpha1.AtList { return &cnatv1alpha1.AtList{} },
		),
	}
}