// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package readinesschecks

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomReadinessCheckConfiguration) DeepCopyInto(out *CustomReadinessCheckConfiguration) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1alpha1.Duration)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make([]v1alpha1.TypedObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(LabelSelectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Requirements != nil {
		in, out := &in.Requirements, &out.Requirements
		*out = make([]RequirementSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomReadinessCheckConfiguration.
func (in *CustomReadinessCheckConfiguration) DeepCopy() *CustomReadinessCheckConfiguration {
	if in == nil {
		return nil
	}
	out := new(CustomReadinessCheckConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSelectorSpec) DeepCopyInto(out *LabelSelectorSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSelectorSpec.
func (in *LabelSelectorSpec) DeepCopy() *LabelSelectorSpec {
	if in == nil {
		return nil
	}
	out := new(LabelSelectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheckConfiguration) DeepCopyInto(out *ReadinessCheckConfiguration) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1alpha1.Duration)
		**out = **in
	}
	if in.CustomReadinessChecks != nil {
		in, out := &in.CustomReadinessChecks, &out.CustomReadinessChecks
		*out = make([]CustomReadinessCheckConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheckConfiguration.
func (in *ReadinessCheckConfiguration) DeepCopy() *ReadinessCheckConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheckConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequirementSpec) DeepCopyInto(out *RequirementSpec) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequirementSpec.
func (in *RequirementSpec) DeepCopy() *RequirementSpec {
	if in == nil {
		return nil
	}
	out := new(RequirementSpec)
	in.DeepCopyInto(out)
	return out
}