/*
Copyright 2021.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterGroupUpgradeSpecApplyConfiguration represents an declarative configuration of the ClusterGroupUpgradeSpec type for use
// with apply.
type ClusterGroupUpgradeSpecApplyConfiguration struct {
	Backup                *bool                                      `json:"backup,omitempty"`
	PreCaching            *bool                                      `json:"preCaching,omitempty"`
	PreCachingConfigRef   *PreCachingConfigCRApplyConfiguration      `json:"preCachingConfigRef,omitempty"`
	Enable                *bool                                      `json:"enable,omitempty"`
	Clusters              []string                                   `json:"clusters,omitempty"`
	ClusterSelector       []string                                   `json:"clusterSelector,omitempty"`
	ClusterLabelSelectors []v1.LabelSelector                         `json:"clusterLabelSelectors,omitempty"`
	RemediationStrategy   *RemediationStrategySpecApplyConfiguration `json:"remediationStrategy,omitempty"`
	ManagedPolicies       []string                                   `json:"managedPolicies,omitempty"`
	ManifestWorkTemplates []string                                   `json:"manifestWorkTemplates,omitempty"`
	BlockingCRs           []BlockingCRApplyConfiguration             `json:"blockingCRs,omitempty"`
	Actions               *ActionsApplyConfiguration                 `json:"actions,omitempty"`
	BatchTimeoutAction    *string                                    `json:"batchTimeoutAction,omitempty"`
}

// ClusterGroupUpgradeSpecApplyConfiguration constructs an declarative configuration of the ClusterGroupUpgradeSpec type for use with
// apply.
func ClusterGroupUpgradeSpec() *ClusterGroupUpgradeSpecApplyConfiguration {
	return &ClusterGroupUpgradeSpecApplyConfiguration{}
}

// WithBackup sets the Backup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Backup field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithBackup(value bool) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.Backup = &value
	return b
}

// WithPreCaching sets the PreCaching field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreCaching field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithPreCaching(value bool) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.PreCaching = &value
	return b
}

// WithPreCachingConfigRef sets the PreCachingConfigRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreCachingConfigRef field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithPreCachingConfigRef(value *PreCachingConfigCRApplyConfiguration) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.PreCachingConfigRef = value
	return b
}

// WithEnable sets the Enable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enable field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithEnable(value bool) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.Enable = &value
	return b
}

// WithClusters adds the given value to the Clusters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Clusters field.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithClusters(values ...string) *ClusterGroupUpgradeSpecApplyConfiguration {
	for i := range values {
		b.Clusters = append(b.Clusters, values[i])
	}
	return b
}

// WithClusterSelector adds the given value to the ClusterSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ClusterSelector field.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithClusterSelector(values ...string) *ClusterGroupUpgradeSpecApplyConfiguration {
	for i := range values {
		b.ClusterSelector = append(b.ClusterSelector, values[i])
	}
	return b
}

// WithClusterLabelSelectors adds the given value to the ClusterLabelSelectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ClusterLabelSelectors field.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithClusterLabelSelectors(values ...v1.LabelSelector) *ClusterGroupUpgradeSpecApplyConfiguration {
	for i := range values {
		b.ClusterLabelSelectors = append(b.ClusterLabelSelectors, values[i])
	}
	return b
}

// WithRemediationStrategy sets the RemediationStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemediationStrategy field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithRemediationStrategy(value *RemediationStrategySpecApplyConfiguration) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.RemediationStrategy = value
	return b
}

// WithManagedPolicies adds the given value to the ManagedPolicies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ManagedPolicies field.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithManagedPolicies(values ...string) *ClusterGroupUpgradeSpecApplyConfiguration {
	for i := range values {
		b.ManagedPolicies = append(b.ManagedPolicies, values[i])
	}
	return b
}

// WithManifestWorkTemplates adds the given value to the ManifestWorkTemplates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ManifestWorkTemplates field.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithManifestWorkTemplates(values ...string) *ClusterGroupUpgradeSpecApplyConfiguration {
	for i := range values {
		b.ManifestWorkTemplates = append(b.ManifestWorkTemplates, values[i])
	}
	return b
}

// WithBlockingCRs adds the given value to the BlockingCRs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the BlockingCRs field.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithBlockingCRs(values ...*BlockingCRApplyConfiguration) *ClusterGroupUpgradeSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithBlockingCRs")
		}
		b.BlockingCRs = append(b.BlockingCRs, *values[i])
	}
	return b
}

// WithActions sets the Actions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Actions field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithActions(value *ActionsApplyConfiguration) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.Actions = value
	return b
}

// WithBatchTimeoutAction sets the BatchTimeoutAction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BatchTimeoutAction field is set to the value of the last call.
func (b *ClusterGroupUpgradeSpecApplyConfiguration) WithBatchTimeoutAction(value string) *ClusterGroupUpgradeSpecApplyConfiguration {
	b.BatchTimeoutAction = &value
	return b
}