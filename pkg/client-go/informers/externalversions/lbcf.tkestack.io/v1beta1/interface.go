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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "tkestack.io/lb-controlling-framework/pkg/client-go/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackendGroups returns a BackendGroupInformer.
	BackendGroups() BackendGroupInformer
	// BackendRecords returns a BackendRecordInformer.
	BackendRecords() BackendRecordInformer
	// LoadBalancers returns a LoadBalancerInformer.
	LoadBalancers() LoadBalancerInformer
	// LoadBalancerDrivers returns a LoadBalancerDriverInformer.
	LoadBalancerDrivers() LoadBalancerDriverInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BackendGroups returns a BackendGroupInformer.
func (v *version) BackendGroups() BackendGroupInformer {
	return &backendGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackendRecords returns a BackendRecordInformer.
func (v *version) BackendRecords() BackendRecordInformer {
	return &backendRecordInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LoadBalancers returns a LoadBalancerInformer.
func (v *version) LoadBalancers() LoadBalancerInformer {
	return &loadBalancerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LoadBalancerDrivers returns a LoadBalancerDriverInformer.
func (v *version) LoadBalancerDrivers() LoadBalancerDriverInformer {
	return &loadBalancerDriverInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
