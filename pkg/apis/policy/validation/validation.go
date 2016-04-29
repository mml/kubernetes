/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
package validation

import (
	"k8s.io/kubernetes/pkg/apis/disruption"
	extensionsvalidation "k8s.io/kubernetes/pkg/apis/extensions/validation"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

func ValidatePodDisruptionBudgetSpec(spec disruption.PodDisruptionBudgetSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, extensionsvalidation.ValidatePositiveIntOrPercent(spec.MinAvailable, fldPath.Child("minAvailable"))...)
	allErrs = append(allErrs, extensionsvalidation.IsNotMoreThan100Percent(spec.MinAvailable, fldPath.Child("minAvailable"))...)

	return allErrs
}
