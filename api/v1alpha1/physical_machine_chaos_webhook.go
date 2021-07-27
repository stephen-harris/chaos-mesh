// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var physicalmachinechaoslog = logf.Log.WithName("physicalmachinechaos-resource")

// +kubebuilder:webhook:path=/mutate-chaos-mesh-org-v1alpha1-physicalmachinechaos,mutating=true,failurePolicy=fail,groups=chaos-mesh.org,resources=physicalmachinechaos,verbs=create;update,versions=v1alpha1,name=mphysicalmachinechaos.kb.io

var _ webhook.Defaulter = &PhysicalMachineChaos{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (in *PhysicalMachineChaos) Default() {
	physicalmachinechaoslog.Info("default", "name", in.Name)
	in.Spec.Default()
}

func (in *PhysicalMachineChaosSpec) Default() {
	if len(in.UID) == 0 {
		in.UID = uuid.New().String()
	}

	// add http prefix for address
	addressArray := strings.Split(in.Address, ",")
	for i := range addressArray {
		if !strings.HasPrefix(addressArray[i], "http") {
			addressArray[i] = fmt.Sprintf("http://%s", addressArray[i])
		}
	}
	in.Address = strings.Join(addressArray, ",")
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-chaos-mesh-org-v1alpha1-physicalmachinechaos,mutating=false,failurePolicy=fail,groups=chaos-mesh.org,resources=physicalmachinechaos,versions=v1alpha1,name=vphysicalmachinechaos.kb.io

var _ webhook.Validator = &PhysicalMachineChaos{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (in *PhysicalMachineChaos) ValidateCreate() error {
	physicalmachinechaoslog.Info("validate create", "name", in.Name)
	return in.Validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (in *PhysicalMachineChaos) ValidateUpdate(old runtime.Object) error {
	physicalmachinechaoslog.Info("validate update", "name", in.Name)
	if !reflect.DeepEqual(in.Spec, old.(*PhysicalMachineChaos).Spec) {
		return ErrCanNotUpdateChaos
	}
	return in.Validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (in *PhysicalMachineChaos) ValidateDelete() error {
	physicalmachinechaoslog.Info("validate delete", "name", in.Name)

	// Nothing to do?
	return nil
}

// Validate validates chaos object
func (in *PhysicalMachineChaos) Validate() error {
	allErrs := in.Spec.Validate()
	if len(allErrs) > 0 {
		return fmt.Errorf(allErrs.ToAggregate().Error())
	}

	return nil
}

func (in *PhysicalMachineChaosSpec) Validate() field.ErrorList {
	allErrs := field.ErrorList{}
	return allErrs
}
