/*
Copyright 2022 The Katalyst Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/tide/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TideNodePoolLister helps list TideNodePools.
// All objects returned here must be treated as read-only.
type TideNodePoolLister interface {
	// List lists all TideNodePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TideNodePool, err error)
	// Get retrieves the TideNodePool from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TideNodePool, error)
	TideNodePoolListerExpansion
}

// tideNodePoolLister implements the TideNodePoolLister interface.
type tideNodePoolLister struct {
	indexer cache.Indexer
}

// NewTideNodePoolLister returns a new TideNodePoolLister.
func NewTideNodePoolLister(indexer cache.Indexer) TideNodePoolLister {
	return &tideNodePoolLister{indexer: indexer}
}

// List lists all TideNodePools in the indexer.
func (s *tideNodePoolLister) List(selector labels.Selector) (ret []*v1alpha1.TideNodePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TideNodePool))
	})
	return ret, err
}

// Get retrieves the TideNodePool from the index for a given name.
func (s *tideNodePoolLister) Get(name string) (*v1alpha1.TideNodePool, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tidenodepool"), name)
	}
	return obj.(*v1alpha1.TideNodePool), nil
}
