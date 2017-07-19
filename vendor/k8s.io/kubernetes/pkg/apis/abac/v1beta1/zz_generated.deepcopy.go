// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Policy).DeepCopyInto(out.(*Policy))
			return nil
		}, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PolicySpec).DeepCopyInto(out.(*PolicySpec))
			return nil
		}, InType: reflect.TypeOf(&PolicySpec{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Spec = in.Spec
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (x *Policy) DeepCopy() *Policy {
	if x == nil {
		return nil
	}
	out := new(Policy)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Policy) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (x *PolicySpec) DeepCopy() *PolicySpec {
	if x == nil {
		return nil
	}
	out := new(PolicySpec)
	x.DeepCopyInto(out)
	return out
}
