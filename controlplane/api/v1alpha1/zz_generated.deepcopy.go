//go:build !ignore_autogenerated

/*
Copyright 2024 SUSE LLC.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/rancher-sandbox/cluster-api-provider-rke2/bootstrap/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisableComponents) DeepCopyInto(out *DisableComponents) {
	*out = *in
	if in.KubernetesComponents != nil {
		in, out := &in.KubernetesComponents, &out.KubernetesComponents
		*out = make([]DisabledKubernetesComponent, len(*in))
		copy(*out, *in)
	}
	if in.PluginComponents != nil {
		in, out := &in.PluginComponents, &out.PluginComponents
		*out = make([]DisabledPluginComponent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisableComponents.
func (in *DisableComponents) DeepCopy() *DisableComponents {
	if in == nil {
		return nil
	}
	out := new(DisableComponents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdBackupConfig) DeepCopyInto(out *EtcdBackupConfig) {
	*out = *in
	if in.DisableAutomaticSnapshots != nil {
		in, out := &in.DisableAutomaticSnapshots, &out.DisableAutomaticSnapshots
		*out = new(bool)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(EtcdS3)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdBackupConfig.
func (in *EtcdBackupConfig) DeepCopy() *EtcdBackupConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConfig) DeepCopyInto(out *EtcdConfig) {
	*out = *in
	in.BackupConfig.DeepCopyInto(&out.BackupConfig)
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = new(apiv1alpha1.ComponentConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConfig.
func (in *EtcdConfig) DeepCopy() *EtcdConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdS3) DeepCopyInto(out *EtcdS3) {
	*out = *in
	if in.EndpointCASecret != nil {
		in, out := &in.EndpointCASecret, &out.EndpointCASecret
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	out.S3CredentialSecret = in.S3CredentialSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdS3.
func (in *EtcdS3) DeepCopy() *EtcdS3 {
	if in == nil {
		return nil
	}
	out := new(EtcdS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlane) DeepCopyInto(out *RKE2ControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlane.
func (in *RKE2ControlPlane) DeepCopy() *RKE2ControlPlane {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKE2ControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneList) DeepCopyInto(out *RKE2ControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKE2ControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneList.
func (in *RKE2ControlPlaneList) DeepCopy() *RKE2ControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKE2ControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneSpec) DeepCopyInto(out *RKE2ControlPlaneSpec) {
	*out = *in
	in.RKE2ConfigSpec.DeepCopyInto(&out.RKE2ConfigSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.ServerConfig.DeepCopyInto(&out.ServerConfig)
	out.ManifestsConfigMapReference = in.ManifestsConfigMapReference
	out.InfrastructureRef = in.InfrastructureRef
	if in.NodeDrainTimeout != nil {
		in, out := &in.NodeDrainTimeout, &out.NodeDrainTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RolloutStrategy != nil {
		in, out := &in.RolloutStrategy, &out.RolloutStrategy
		*out = new(RolloutStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneSpec.
func (in *RKE2ControlPlaneSpec) DeepCopy() *RKE2ControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneStatus) DeepCopyInto(out *RKE2ControlPlaneStatus) {
	*out = *in
	if in.DataSecretName != nil {
		in, out := &in.DataSecretName, &out.DataSecretName
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AvailableServerIPs != nil {
		in, out := &in.AvailableServerIPs, &out.AvailableServerIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneStatus.
func (in *RKE2ControlPlaneStatus) DeepCopy() *RKE2ControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneTemplate) DeepCopyInto(out *RKE2ControlPlaneTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneTemplate.
func (in *RKE2ControlPlaneTemplate) DeepCopy() *RKE2ControlPlaneTemplate {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKE2ControlPlaneTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneTemplateList) DeepCopyInto(out *RKE2ControlPlaneTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKE2ControlPlaneTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneTemplateList.
func (in *RKE2ControlPlaneTemplateList) DeepCopy() *RKE2ControlPlaneTemplateList {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKE2ControlPlaneTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneTemplateSpec) DeepCopyInto(out *RKE2ControlPlaneTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneTemplateSpec.
func (in *RKE2ControlPlaneTemplateSpec) DeepCopy() *RKE2ControlPlaneTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ControlPlaneTemplateStatus) DeepCopyInto(out *RKE2ControlPlaneTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ControlPlaneTemplateStatus.
func (in *RKE2ControlPlaneTemplateStatus) DeepCopy() *RKE2ControlPlaneTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(RKE2ControlPlaneTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2ServerConfig) DeepCopyInto(out *RKE2ServerConfig) {
	*out = *in
	if in.AuditPolicySecret != nil {
		in, out := &in.AuditPolicySecret, &out.AuditPolicySecret
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.TLSSan != nil {
		in, out := &in.TLSSan, &out.TLSSan
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.DisableComponents.DeepCopyInto(&out.DisableComponents)
	in.Etcd.DeepCopyInto(&out.Etcd)
	if in.KubeAPIServer != nil {
		in, out := &in.KubeAPIServer, &out.KubeAPIServer
		*out = new(apiv1alpha1.ComponentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeControllerManager != nil {
		in, out := &in.KubeControllerManager, &out.KubeControllerManager
		*out = new(apiv1alpha1.ComponentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeScheduler != nil {
		in, out := &in.KubeScheduler, &out.KubeScheduler
		*out = new(apiv1alpha1.ComponentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudControllerManager != nil {
		in, out := &in.CloudControllerManager, &out.CloudControllerManager
		*out = new(apiv1alpha1.ComponentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudProviderConfigMap != nil {
		in, out := &in.CloudProviderConfigMap, &out.CloudProviderConfigMap
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2ServerConfig.
func (in *RKE2ServerConfig) DeepCopy() *RKE2ServerConfig {
	if in == nil {
		return nil
	}
	out := new(RKE2ServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdate) DeepCopyInto(out *RollingUpdate) {
	*out = *in
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdate.
func (in *RollingUpdate) DeepCopy() *RollingUpdate {
	if in == nil {
		return nil
	}
	out := new(RollingUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}
