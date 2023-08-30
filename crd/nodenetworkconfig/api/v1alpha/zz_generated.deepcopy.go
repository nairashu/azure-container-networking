//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignment) DeepCopyInto(out *IPAssignment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignment.
func (in *IPAssignment) DeepCopy() *IPAssignment {
	if in == nil {
		return nil
	}
	out := new(IPAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainer) DeepCopyInto(out *NetworkContainer) {
	*out = *in
	if in.IPAssignments != nil {
		in, out := &in.IPAssignments, &out.IPAssignments
		*out = make([]IPAssignment, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainer.
func (in *NetworkContainer) DeepCopy() *NetworkContainer {
	if in == nil {
		return nil
	}
	out := new(NetworkContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContainerStatus) DeepCopyInto(out *NetworkContainerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContainerStatus.
func (in *NetworkContainerStatus) DeepCopy() *NetworkContainerStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkContainerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetworkConfig) DeepCopyInto(out *NodeNetworkConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetworkConfig.
func (in *NodeNetworkConfig) DeepCopy() *NodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NodeNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeNetworkConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetworkConfigList) DeepCopyInto(out *NodeNetworkConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeNetworkConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetworkConfigList.
func (in *NodeNetworkConfigList) DeepCopy() *NodeNetworkConfigList {
	if in == nil {
		return nil
	}
	out := new(NodeNetworkConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeNetworkConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetworkConfigSpec) DeepCopyInto(out *NodeNetworkConfigSpec) {
	*out = *in
	if in.IPsNotInUse != nil {
		in, out := &in.IPsNotInUse, &out.IPsNotInUse
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetworkConfigSpec.
func (in *NodeNetworkConfigSpec) DeepCopy() *NodeNetworkConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeNetworkConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetworkConfigStatus) DeepCopyInto(out *NodeNetworkConfigStatus) {
	*out = *in
	out.Scaler = in.Scaler
	if in.NetworkContainers != nil {
		in, out := &in.NetworkContainers, &out.NetworkContainers
		*out = make([]NetworkContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkContainerStatuses != nil {
		in, out := &in.NetworkContainerStatuses, &out.NetworkContainerStatuses
		*out = make([]NetworkContainerStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetworkConfigStatus.
func (in *NodeNetworkConfigStatus) DeepCopy() *NodeNetworkConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NodeNetworkConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scaler) DeepCopyInto(out *Scaler) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scaler.
func (in *Scaler) DeepCopy() *Scaler {
	if in == nil {
		return nil
	}
	out := new(Scaler)
	in.DeepCopyInto(out)
	return out
}
