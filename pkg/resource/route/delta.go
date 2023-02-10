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

package route

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.APIID, b.ko.Spec.APIID) {
		delta.Add("Spec.APIID", a.ko.Spec.APIID, b.ko.Spec.APIID)
	} else if a.ko.Spec.APIID != nil && b.ko.Spec.APIID != nil {
		if *a.ko.Spec.APIID != *b.ko.Spec.APIID {
			delta.Add("Spec.APIID", a.ko.Spec.APIID, b.ko.Spec.APIID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.APIKeyRequired, b.ko.Spec.APIKeyRequired) {
		delta.Add("Spec.APIKeyRequired", a.ko.Spec.APIKeyRequired, b.ko.Spec.APIKeyRequired)
	} else if a.ko.Spec.APIKeyRequired != nil && b.ko.Spec.APIKeyRequired != nil {
		if *a.ko.Spec.APIKeyRequired != *b.ko.Spec.APIKeyRequired {
			delta.Add("Spec.APIKeyRequired", a.ko.Spec.APIKeyRequired, b.ko.Spec.APIKeyRequired)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.APIRef, b.ko.Spec.APIRef) {
		delta.Add("Spec.APIRef", a.ko.Spec.APIRef, b.ko.Spec.APIRef)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.AuthorizationScopes, b.ko.Spec.AuthorizationScopes) {
		delta.Add("Spec.AuthorizationScopes", a.ko.Spec.AuthorizationScopes, b.ko.Spec.AuthorizationScopes)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizationType, b.ko.Spec.AuthorizationType) {
		delta.Add("Spec.AuthorizationType", a.ko.Spec.AuthorizationType, b.ko.Spec.AuthorizationType)
	} else if a.ko.Spec.AuthorizationType != nil && b.ko.Spec.AuthorizationType != nil {
		if *a.ko.Spec.AuthorizationType != *b.ko.Spec.AuthorizationType {
			delta.Add("Spec.AuthorizationType", a.ko.Spec.AuthorizationType, b.ko.Spec.AuthorizationType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizerID, b.ko.Spec.AuthorizerID) {
		delta.Add("Spec.AuthorizerID", a.ko.Spec.AuthorizerID, b.ko.Spec.AuthorizerID)
	} else if a.ko.Spec.AuthorizerID != nil && b.ko.Spec.AuthorizerID != nil {
		if *a.ko.Spec.AuthorizerID != *b.ko.Spec.AuthorizerID {
			delta.Add("Spec.AuthorizerID", a.ko.Spec.AuthorizerID, b.ko.Spec.AuthorizerID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.AuthorizerRef, b.ko.Spec.AuthorizerRef) {
		delta.Add("Spec.AuthorizerRef", a.ko.Spec.AuthorizerRef, b.ko.Spec.AuthorizerRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelSelectionExpression, b.ko.Spec.ModelSelectionExpression) {
		delta.Add("Spec.ModelSelectionExpression", a.ko.Spec.ModelSelectionExpression, b.ko.Spec.ModelSelectionExpression)
	} else if a.ko.Spec.ModelSelectionExpression != nil && b.ko.Spec.ModelSelectionExpression != nil {
		if *a.ko.Spec.ModelSelectionExpression != *b.ko.Spec.ModelSelectionExpression {
			delta.Add("Spec.ModelSelectionExpression", a.ko.Spec.ModelSelectionExpression, b.ko.Spec.ModelSelectionExpression)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OperationName, b.ko.Spec.OperationName) {
		delta.Add("Spec.OperationName", a.ko.Spec.OperationName, b.ko.Spec.OperationName)
	} else if a.ko.Spec.OperationName != nil && b.ko.Spec.OperationName != nil {
		if *a.ko.Spec.OperationName != *b.ko.Spec.OperationName {
			delta.Add("Spec.OperationName", a.ko.Spec.OperationName, b.ko.Spec.OperationName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RequestModels, b.ko.Spec.RequestModels) {
		delta.Add("Spec.RequestModels", a.ko.Spec.RequestModels, b.ko.Spec.RequestModels)
	} else if a.ko.Spec.RequestModels != nil && b.ko.Spec.RequestModels != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.RequestModels, b.ko.Spec.RequestModels) {
			delta.Add("Spec.RequestModels", a.ko.Spec.RequestModels, b.ko.Spec.RequestModels)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters) {
		delta.Add("Spec.RequestParameters", a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters)
	} else if a.ko.Spec.RequestParameters != nil && b.ko.Spec.RequestParameters != nil {
		if !reflect.DeepEqual(a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters) {
			delta.Add("Spec.RequestParameters", a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RouteKey, b.ko.Spec.RouteKey) {
		delta.Add("Spec.RouteKey", a.ko.Spec.RouteKey, b.ko.Spec.RouteKey)
	} else if a.ko.Spec.RouteKey != nil && b.ko.Spec.RouteKey != nil {
		if *a.ko.Spec.RouteKey != *b.ko.Spec.RouteKey {
			delta.Add("Spec.RouteKey", a.ko.Spec.RouteKey, b.ko.Spec.RouteKey)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RouteResponseSelectionExpression, b.ko.Spec.RouteResponseSelectionExpression) {
		delta.Add("Spec.RouteResponseSelectionExpression", a.ko.Spec.RouteResponseSelectionExpression, b.ko.Spec.RouteResponseSelectionExpression)
	} else if a.ko.Spec.RouteResponseSelectionExpression != nil && b.ko.Spec.RouteResponseSelectionExpression != nil {
		if *a.ko.Spec.RouteResponseSelectionExpression != *b.ko.Spec.RouteResponseSelectionExpression {
			delta.Add("Spec.RouteResponseSelectionExpression", a.ko.Spec.RouteResponseSelectionExpression, b.ko.Spec.RouteResponseSelectionExpression)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Target, b.ko.Spec.Target) {
		delta.Add("Spec.Target", a.ko.Spec.Target, b.ko.Spec.Target)
	} else if a.ko.Spec.Target != nil && b.ko.Spec.Target != nil {
		if *a.ko.Spec.Target != *b.ko.Spec.Target {
			delta.Add("Spec.Target", a.ko.Spec.Target, b.ko.Spec.Target)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.TargetRef, b.ko.Spec.TargetRef) {
		delta.Add("Spec.TargetRef", a.ko.Spec.TargetRef, b.ko.Spec.TargetRef)
	}

	return delta
}
