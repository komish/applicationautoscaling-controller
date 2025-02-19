// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScheduledActionSpec defines the desired state of ScheduledAction
type ScheduledActionSpec struct {
	EndTime *metav1.Time `json:"endTime,omitempty"`
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceID"`
	// +kubebuilder:validation:Required
	ScalableDimension    *string               `json:"scalableDimension"`
	ScalableTargetAction *ScalableTargetAction `json:"scalableTargetAction,omitempty"`
	Schedule             *string               `json:"schedule,omitempty"`
	// +kubebuilder:validation:Required
	ScheduledActionName *string `json:"scheduledActionName"`
	// +kubebuilder:validation:Required
	ServiceNamespace *string      `json:"serviceNamespace"`
	StartTime        *metav1.Time `json:"startTime,omitempty"`
	Timezone         *string      `json:"timezone,omitempty"`
}

// ScheduledActionStatus defines the observed state of ScheduledAction
type ScheduledActionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// ScheduledAction is the Schema for the ScheduledActions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ScheduledAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduledActionSpec   `json:"spec,omitempty"`
	Status            ScheduledActionStatus `json:"status,omitempty"`
}

// ScheduledActionList contains a list of ScheduledAction
// +kubebuilder:object:root=true
type ScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledAction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScheduledAction{}, &ScheduledActionList{})
}
