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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Alarm struct {
	AlarmARN  *string `json:"alarmARN,omitempty"`
	AlarmName *string `json:"alarmName,omitempty"`
}

type CustomizedMetricSpecification struct {
	Dimensions []*MetricDimension `json:"dimensions,omitempty"`
	MetricName *string            `json:"metricName,omitempty"`
	Namespace  *string            `json:"namespace,omitempty"`
	Statistic  *string            `json:"statistic,omitempty"`
	Unit       *string            `json:"unit,omitempty"`
}

type MetricDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type PredefinedMetricSpecification struct {
	PredefinedMetricType *string `json:"predefinedMetricType,omitempty"`
	ResourceLabel        *string `json:"resourceLabel,omitempty"`
}

type ScalableTargetAction struct {
	MaxCapacity *int64 `json:"maxCapacity,omitempty"`
	MinCapacity *int64 `json:"minCapacity,omitempty"`
}

type ScalableTarget_SDK struct {
	CreationTime      *metav1.Time    `json:"creationTime,omitempty"`
	MaxCapacity       *int64          `json:"maxCapacity,omitempty"`
	MinCapacity       *int64          `json:"minCapacity,omitempty"`
	ResourceID        *string         `json:"resourceID,omitempty"`
	RoleARN           *string         `json:"roleARN,omitempty"`
	ScalableDimension *string         `json:"scalableDimension,omitempty"`
	ServiceNamespace  *string         `json:"serviceNamespace,omitempty"`
	SuspendedState    *SuspendedState `json:"suspendedState,omitempty"`
}

type ScalingActivity struct {
	ActivityID        *string      `json:"activityID,omitempty"`
	Cause             *string      `json:"cause,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Details           *string      `json:"details,omitempty"`
	EndTime           *metav1.Time `json:"endTime,omitempty"`
	ResourceID        *string      `json:"resourceID,omitempty"`
	ScalableDimension *string      `json:"scalableDimension,omitempty"`
	ServiceNamespace  *string      `json:"serviceNamespace,omitempty"`
	StartTime         *metav1.Time `json:"startTime,omitempty"`
	StatusMessage     *string      `json:"statusMessage,omitempty"`
}

type ScalingPolicy_SDK struct {
	Alarms                                   []*Alarm                                  `json:"alarms,omitempty"`
	CreationTime                             *metav1.Time                              `json:"creationTime,omitempty"`
	PolicyARN                                *string                                   `json:"policyARN,omitempty"`
	PolicyName                               *string                                   `json:"policyName,omitempty"`
	PolicyType                               *string                                   `json:"policyType,omitempty"`
	ResourceID                               *string                                   `json:"resourceID,omitempty"`
	ScalableDimension                        *string                                   `json:"scalableDimension,omitempty"`
	ServiceNamespace                         *string                                   `json:"serviceNamespace,omitempty"`
	StepScalingPolicyConfiguration           *StepScalingPolicyConfiguration           `json:"stepScalingPolicyConfiguration,omitempty"`
	TargetTrackingScalingPolicyConfiguration *TargetTrackingScalingPolicyConfiguration `json:"targetTrackingScalingPolicyConfiguration,omitempty"`
}

type ScheduledAction_SDK struct {
	CreationTime         *metav1.Time          `json:"creationTime,omitempty"`
	EndTime              *metav1.Time          `json:"endTime,omitempty"`
	ResourceID           *string               `json:"resourceID,omitempty"`
	ScalableDimension    *string               `json:"scalableDimension,omitempty"`
	ScalableTargetAction *ScalableTargetAction `json:"scalableTargetAction,omitempty"`
	Schedule             *string               `json:"schedule,omitempty"`
	ScheduledActionARN   *string               `json:"scheduledActionARN,omitempty"`
	ScheduledActionName  *string               `json:"scheduledActionName,omitempty"`
	ServiceNamespace     *string               `json:"serviceNamespace,omitempty"`
	StartTime            *metav1.Time          `json:"startTime,omitempty"`
	Timezone             *string               `json:"timezone,omitempty"`
}

type StepAdjustment struct {
	MetricIntervalLowerBound *float64 `json:"metricIntervalLowerBound,omitempty"`
	MetricIntervalUpperBound *float64 `json:"metricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int64   `json:"scalingAdjustment,omitempty"`
}

type StepScalingPolicyConfiguration struct {
	AdjustmentType         *string           `json:"adjustmentType,omitempty"`
	Cooldown               *int64            `json:"cooldown,omitempty"`
	MetricAggregationType  *string           `json:"metricAggregationType,omitempty"`
	MinAdjustmentMagnitude *int64            `json:"minAdjustmentMagnitude,omitempty"`
	StepAdjustments        []*StepAdjustment `json:"stepAdjustments,omitempty"`
}

type SuspendedState struct {
	DynamicScalingInSuspended  *bool `json:"dynamicScalingInSuspended,omitempty"`
	DynamicScalingOutSuspended *bool `json:"dynamicScalingOutSuspended,omitempty"`
	ScheduledScalingSuspended  *bool `json:"scheduledScalingSuspended,omitempty"`
}

type TargetTrackingScalingPolicyConfiguration struct {
	CustomizedMetricSpecification *CustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty"`
	DisableScaleIn                *bool                          `json:"disableScaleIn,omitempty"`
	PredefinedMetricSpecification *PredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty"`
	ScaleInCooldown               *int64                         `json:"scaleInCooldown,omitempty"`
	ScaleOutCooldown              *int64                         `json:"scaleOutCooldown,omitempty"`
	TargetValue                   *float64                       `json:"targetValue,omitempty"`
}
