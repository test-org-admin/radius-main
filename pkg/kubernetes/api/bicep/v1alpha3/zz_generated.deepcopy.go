//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTemplate) DeepCopyInto(out *DeploymentTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTemplate.
func (in *DeploymentTemplate) DeepCopy() *DeploymentTemplate {
	if in == nil {
		return nil
	}
	out := new(DeploymentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTemplateList) DeepCopyInto(out *DeploymentTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTemplateList.
func (in *DeploymentTemplateList) DeepCopy() *DeploymentTemplateList {
	if in == nil {
		return nil
	}
	out := new(DeploymentTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTemplateOperation) DeepCopyInto(out *DeploymentTemplateOperation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTemplateOperation.
func (in *DeploymentTemplateOperation) DeepCopy() *DeploymentTemplateOperation {
	if in == nil {
		return nil
	}
	out := new(DeploymentTemplateOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTemplateSpec) DeepCopyInto(out *DeploymentTemplateSpec) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTemplateSpec.
func (in *DeploymentTemplateSpec) DeepCopy() *DeploymentTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTemplateStatus) DeepCopyInto(out *DeploymentTemplateStatus) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]DeploymentTemplateOperation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTemplateStatus.
func (in *DeploymentTemplateStatus) DeepCopy() *DeploymentTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
