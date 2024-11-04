//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	apiv1alpha1 "github.com/rancher/cluster-api-provider-rke2/bootstrap/api/v1alpha1"
	apiv1beta1 "github.com/rancher/cluster-api-provider-rke2/bootstrap/api/v1beta1"
	v1beta1 "github.com/rancher/cluster-api-provider-rke2/controlplane/api/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	clusterapiapiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DisableComponents)(nil), (*v1beta1.DisableComponents)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DisableComponents_To_v1beta1_DisableComponents(a.(*DisableComponents), b.(*v1beta1.DisableComponents), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DisableComponents)(nil), (*DisableComponents)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DisableComponents_To_v1alpha1_DisableComponents(a.(*v1beta1.DisableComponents), b.(*DisableComponents), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdBackupConfig)(nil), (*v1beta1.EtcdBackupConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EtcdBackupConfig_To_v1beta1_EtcdBackupConfig(a.(*EtcdBackupConfig), b.(*v1beta1.EtcdBackupConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdBackupConfig)(nil), (*EtcdBackupConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdBackupConfig_To_v1alpha1_EtcdBackupConfig(a.(*v1beta1.EtcdBackupConfig), b.(*EtcdBackupConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdConfig)(nil), (*v1beta1.EtcdConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EtcdConfig_To_v1beta1_EtcdConfig(a.(*EtcdConfig), b.(*v1beta1.EtcdConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdConfig)(nil), (*EtcdConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdConfig_To_v1alpha1_EtcdConfig(a.(*v1beta1.EtcdConfig), b.(*EtcdConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EtcdS3)(nil), (*v1beta1.EtcdS3)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EtcdS3_To_v1beta1_EtcdS3(a.(*EtcdS3), b.(*v1beta1.EtcdS3), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EtcdS3)(nil), (*EtcdS3)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EtcdS3_To_v1alpha1_EtcdS3(a.(*v1beta1.EtcdS3), b.(*EtcdS3), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RKE2ControlPlane)(nil), (*v1beta1.RKE2ControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane(a.(*RKE2ControlPlane), b.(*v1beta1.RKE2ControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RKE2ControlPlane)(nil), (*RKE2ControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane(a.(*v1beta1.RKE2ControlPlane), b.(*RKE2ControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RKE2ControlPlaneList)(nil), (*v1beta1.RKE2ControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneList_To_v1beta1_RKE2ControlPlaneList(a.(*RKE2ControlPlaneList), b.(*v1beta1.RKE2ControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RKE2ControlPlaneList)(nil), (*RKE2ControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneList_To_v1alpha1_RKE2ControlPlaneList(a.(*v1beta1.RKE2ControlPlaneList), b.(*RKE2ControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RKE2ControlPlaneSpec)(nil), (*v1beta1.RKE2ControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneSpec_To_v1beta1_RKE2ControlPlaneSpec(a.(*RKE2ControlPlaneSpec), b.(*v1beta1.RKE2ControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RKE2ControlPlaneTemplate)(nil), (*v1beta1.RKE2ControlPlaneTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate(a.(*RKE2ControlPlaneTemplate), b.(*v1beta1.RKE2ControlPlaneTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RKE2ControlPlaneTemplate)(nil), (*RKE2ControlPlaneTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate(a.(*v1beta1.RKE2ControlPlaneTemplate), b.(*RKE2ControlPlaneTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RKE2ControlPlaneTemplateList)(nil), (*v1beta1.RKE2ControlPlaneTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneTemplateList_To_v1beta1_RKE2ControlPlaneTemplateList(a.(*RKE2ControlPlaneTemplateList), b.(*v1beta1.RKE2ControlPlaneTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RKE2ControlPlaneTemplateList)(nil), (*RKE2ControlPlaneTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneTemplateList_To_v1alpha1_RKE2ControlPlaneTemplateList(a.(*v1beta1.RKE2ControlPlaneTemplateList), b.(*RKE2ControlPlaneTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RKE2ServerConfig)(nil), (*v1beta1.RKE2ServerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ServerConfig_To_v1beta1_RKE2ServerConfig(a.(*RKE2ServerConfig), b.(*v1beta1.RKE2ServerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RKE2ServerConfig)(nil), (*RKE2ServerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(a.(*v1beta1.RKE2ServerConfig), b.(*RKE2ServerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RollingUpdate)(nil), (*v1beta1.RollingUpdate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RollingUpdate_To_v1beta1_RollingUpdate(a.(*RollingUpdate), b.(*v1beta1.RollingUpdate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RollingUpdate)(nil), (*RollingUpdate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RollingUpdate_To_v1alpha1_RollingUpdate(a.(*v1beta1.RollingUpdate), b.(*RollingUpdate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RolloutStrategy)(nil), (*v1beta1.RolloutStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RolloutStrategy_To_v1beta1_RolloutStrategy(a.(*RolloutStrategy), b.(*v1beta1.RolloutStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RolloutStrategy)(nil), (*RolloutStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RolloutStrategy_To_v1alpha1_RolloutStrategy(a.(*v1beta1.RolloutStrategy), b.(*RolloutStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1alpha1.RKE2ConfigSpec)(nil), (*apiv1beta1.RKE2ConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ConfigSpec_To_v1beta1_RKE2ConfigSpec(a.(*apiv1alpha1.RKE2ConfigSpec), b.(*apiv1beta1.RKE2ConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*RKE2ControlPlaneStatus)(nil), (*v1beta1.RKE2ControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneStatus_To_v1beta1_RKE2ControlPlaneStatus(a.(*RKE2ControlPlaneStatus), b.(*v1beta1.RKE2ControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*RKE2ControlPlaneTemplateSpec)(nil), (*v1beta1.RKE2ControlPlaneTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneTemplateSpec_To_v1beta1_RKE2ControlPlaneTemplateSpec(a.(*RKE2ControlPlaneTemplateSpec), b.(*v1beta1.RKE2ControlPlaneTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*RKE2ControlPlaneTemplateStatus)(nil), (*v1beta1.RKE2ControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RKE2ControlPlaneTemplateStatus_To_v1beta1_RKE2ControlPlaneStatus(a.(*RKE2ControlPlaneTemplateStatus), b.(*v1beta1.RKE2ControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiv1beta1.RKE2ConfigSpec)(nil), (*apiv1alpha1.RKE2ConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ConfigSpec_To_v1alpha1_RKE2ConfigSpec(a.(*apiv1beta1.RKE2ConfigSpec), b.(*apiv1alpha1.RKE2ConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.RKE2ControlPlaneSpec)(nil), (*RKE2ControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneSpec_To_v1alpha1_RKE2ControlPlaneSpec(a.(*v1beta1.RKE2ControlPlaneSpec), b.(*RKE2ControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.RKE2ControlPlaneStatus)(nil), (*RKE2ControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneStatus(a.(*v1beta1.RKE2ControlPlaneStatus), b.(*RKE2ControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.RKE2ControlPlaneStatus)(nil), (*RKE2ControlPlaneTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneTemplateStatus(a.(*v1beta1.RKE2ControlPlaneStatus), b.(*RKE2ControlPlaneTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.RKE2ControlPlaneTemplateSpec)(nil), (*RKE2ControlPlaneTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RKE2ControlPlaneTemplateSpec_To_v1alpha1_RKE2ControlPlaneTemplateSpec(a.(*v1beta1.RKE2ControlPlaneTemplateSpec), b.(*RKE2ControlPlaneTemplateSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_DisableComponents_To_v1beta1_DisableComponents(in *DisableComponents, out *v1beta1.DisableComponents, s conversion.Scope) error {
	out.KubernetesComponents = *(*[]v1beta1.DisabledKubernetesComponent)(unsafe.Pointer(&in.KubernetesComponents))
	out.PluginComponents = *(*[]v1beta1.DisabledPluginComponent)(unsafe.Pointer(&in.PluginComponents))
	return nil
}

// Convert_v1alpha1_DisableComponents_To_v1beta1_DisableComponents is an autogenerated conversion function.
func Convert_v1alpha1_DisableComponents_To_v1beta1_DisableComponents(in *DisableComponents, out *v1beta1.DisableComponents, s conversion.Scope) error {
	return autoConvert_v1alpha1_DisableComponents_To_v1beta1_DisableComponents(in, out, s)
}

func autoConvert_v1beta1_DisableComponents_To_v1alpha1_DisableComponents(in *v1beta1.DisableComponents, out *DisableComponents, s conversion.Scope) error {
	out.KubernetesComponents = *(*[]DisabledKubernetesComponent)(unsafe.Pointer(&in.KubernetesComponents))
	out.PluginComponents = *(*[]DisabledPluginComponent)(unsafe.Pointer(&in.PluginComponents))
	return nil
}

// Convert_v1beta1_DisableComponents_To_v1alpha1_DisableComponents is an autogenerated conversion function.
func Convert_v1beta1_DisableComponents_To_v1alpha1_DisableComponents(in *v1beta1.DisableComponents, out *DisableComponents, s conversion.Scope) error {
	return autoConvert_v1beta1_DisableComponents_To_v1alpha1_DisableComponents(in, out, s)
}

func autoConvert_v1alpha1_EtcdBackupConfig_To_v1beta1_EtcdBackupConfig(in *EtcdBackupConfig, out *v1beta1.EtcdBackupConfig, s conversion.Scope) error {
	out.DisableAutomaticSnapshots = (*bool)(unsafe.Pointer(in.DisableAutomaticSnapshots))
	out.SnapshotName = in.SnapshotName
	out.ScheduleCron = in.ScheduleCron
	out.Retention = in.Retention
	out.Directory = in.Directory
	out.S3 = (*v1beta1.EtcdS3)(unsafe.Pointer(in.S3))
	return nil
}

// Convert_v1alpha1_EtcdBackupConfig_To_v1beta1_EtcdBackupConfig is an autogenerated conversion function.
func Convert_v1alpha1_EtcdBackupConfig_To_v1beta1_EtcdBackupConfig(in *EtcdBackupConfig, out *v1beta1.EtcdBackupConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_EtcdBackupConfig_To_v1beta1_EtcdBackupConfig(in, out, s)
}

func autoConvert_v1beta1_EtcdBackupConfig_To_v1alpha1_EtcdBackupConfig(in *v1beta1.EtcdBackupConfig, out *EtcdBackupConfig, s conversion.Scope) error {
	out.DisableAutomaticSnapshots = (*bool)(unsafe.Pointer(in.DisableAutomaticSnapshots))
	out.SnapshotName = in.SnapshotName
	out.ScheduleCron = in.ScheduleCron
	out.Retention = in.Retention
	out.Directory = in.Directory
	out.S3 = (*EtcdS3)(unsafe.Pointer(in.S3))
	return nil
}

// Convert_v1beta1_EtcdBackupConfig_To_v1alpha1_EtcdBackupConfig is an autogenerated conversion function.
func Convert_v1beta1_EtcdBackupConfig_To_v1alpha1_EtcdBackupConfig(in *v1beta1.EtcdBackupConfig, out *EtcdBackupConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdBackupConfig_To_v1alpha1_EtcdBackupConfig(in, out, s)
}

func autoConvert_v1alpha1_EtcdConfig_To_v1beta1_EtcdConfig(in *EtcdConfig, out *v1beta1.EtcdConfig, s conversion.Scope) error {
	out.ExposeMetrics = in.ExposeMetrics
	if err := Convert_v1alpha1_EtcdBackupConfig_To_v1beta1_EtcdBackupConfig(&in.BackupConfig, &out.BackupConfig, s); err != nil {
		return err
	}
	out.CustomConfig = (*apiv1beta1.ComponentConfig)(unsafe.Pointer(in.CustomConfig))
	return nil
}

// Convert_v1alpha1_EtcdConfig_To_v1beta1_EtcdConfig is an autogenerated conversion function.
func Convert_v1alpha1_EtcdConfig_To_v1beta1_EtcdConfig(in *EtcdConfig, out *v1beta1.EtcdConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_EtcdConfig_To_v1beta1_EtcdConfig(in, out, s)
}

func autoConvert_v1beta1_EtcdConfig_To_v1alpha1_EtcdConfig(in *v1beta1.EtcdConfig, out *EtcdConfig, s conversion.Scope) error {
	out.ExposeMetrics = in.ExposeMetrics
	if err := Convert_v1beta1_EtcdBackupConfig_To_v1alpha1_EtcdBackupConfig(&in.BackupConfig, &out.BackupConfig, s); err != nil {
		return err
	}
	out.CustomConfig = (*apiv1alpha1.ComponentConfig)(unsafe.Pointer(in.CustomConfig))
	return nil
}

// Convert_v1beta1_EtcdConfig_To_v1alpha1_EtcdConfig is an autogenerated conversion function.
func Convert_v1beta1_EtcdConfig_To_v1alpha1_EtcdConfig(in *v1beta1.EtcdConfig, out *EtcdConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdConfig_To_v1alpha1_EtcdConfig(in, out, s)
}

func autoConvert_v1alpha1_EtcdS3_To_v1beta1_EtcdS3(in *EtcdS3, out *v1beta1.EtcdS3, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.EndpointCASecret = (*v1.ObjectReference)(unsafe.Pointer(in.EndpointCASecret))
	out.EnforceSSLVerify = in.EnforceSSLVerify
	out.S3CredentialSecret = in.S3CredentialSecret
	out.Bucket = in.Bucket
	out.Region = in.Region
	out.Folder = in.Folder
	return nil
}

// Convert_v1alpha1_EtcdS3_To_v1beta1_EtcdS3 is an autogenerated conversion function.
func Convert_v1alpha1_EtcdS3_To_v1beta1_EtcdS3(in *EtcdS3, out *v1beta1.EtcdS3, s conversion.Scope) error {
	return autoConvert_v1alpha1_EtcdS3_To_v1beta1_EtcdS3(in, out, s)
}

func autoConvert_v1beta1_EtcdS3_To_v1alpha1_EtcdS3(in *v1beta1.EtcdS3, out *EtcdS3, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.EndpointCASecret = (*v1.ObjectReference)(unsafe.Pointer(in.EndpointCASecret))
	out.EnforceSSLVerify = in.EnforceSSLVerify
	out.S3CredentialSecret = in.S3CredentialSecret
	out.Bucket = in.Bucket
	out.Region = in.Region
	out.Folder = in.Folder
	return nil
}

// Convert_v1beta1_EtcdS3_To_v1alpha1_EtcdS3 is an autogenerated conversion function.
func Convert_v1beta1_EtcdS3_To_v1alpha1_EtcdS3(in *v1beta1.EtcdS3, out *EtcdS3, s conversion.Scope) error {
	return autoConvert_v1beta1_EtcdS3_To_v1alpha1_EtcdS3(in, out, s)
}

func autoConvert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane(in *RKE2ControlPlane, out *v1beta1.RKE2ControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_RKE2ControlPlaneSpec_To_v1beta1_RKE2ControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_RKE2ControlPlaneStatus_To_v1beta1_RKE2ControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane is an autogenerated conversion function.
func Convert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane(in *RKE2ControlPlane, out *v1beta1.RKE2ControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane(in, out, s)
}

func autoConvert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane(in *v1beta1.RKE2ControlPlane, out *RKE2ControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_RKE2ControlPlaneSpec_To_v1alpha1_RKE2ControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane is an autogenerated conversion function.
func Convert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane(in *v1beta1.RKE2ControlPlane, out *RKE2ControlPlane, s conversion.Scope) error {
	return autoConvert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane(in, out, s)
}

func autoConvert_v1alpha1_RKE2ControlPlaneList_To_v1beta1_RKE2ControlPlaneList(in *RKE2ControlPlaneList, out *v1beta1.RKE2ControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.RKE2ControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_RKE2ControlPlane_To_v1beta1_RKE2ControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_RKE2ControlPlaneList_To_v1beta1_RKE2ControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha1_RKE2ControlPlaneList_To_v1beta1_RKE2ControlPlaneList(in *RKE2ControlPlaneList, out *v1beta1.RKE2ControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlaneList_To_v1beta1_RKE2ControlPlaneList(in, out, s)
}

func autoConvert_v1beta1_RKE2ControlPlaneList_To_v1alpha1_RKE2ControlPlaneList(in *v1beta1.RKE2ControlPlaneList, out *RKE2ControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKE2ControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_RKE2ControlPlane_To_v1alpha1_RKE2ControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_RKE2ControlPlaneList_To_v1alpha1_RKE2ControlPlaneList is an autogenerated conversion function.
func Convert_v1beta1_RKE2ControlPlaneList_To_v1alpha1_RKE2ControlPlaneList(in *v1beta1.RKE2ControlPlaneList, out *RKE2ControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1beta1_RKE2ControlPlaneList_To_v1alpha1_RKE2ControlPlaneList(in, out, s)
}

func autoConvert_v1alpha1_RKE2ControlPlaneSpec_To_v1beta1_RKE2ControlPlaneSpec(in *RKE2ControlPlaneSpec, out *v1beta1.RKE2ControlPlaneSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_RKE2ConfigSpec_To_v1beta1_RKE2ConfigSpec(&in.RKE2ConfigSpec, &out.RKE2ConfigSpec, s); err != nil {
		return err
	}
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	if err := Convert_v1alpha1_RKE2ServerConfig_To_v1beta1_RKE2ServerConfig(&in.ServerConfig, &out.ServerConfig, s); err != nil {
		return err
	}
	out.ManifestsConfigMapReference = in.ManifestsConfigMapReference
	out.InfrastructureRef = in.InfrastructureRef
	out.NodeDrainTimeout = (*metav1.Duration)(unsafe.Pointer(in.NodeDrainTimeout))
	out.RegistrationMethod = v1beta1.RegistrationMethod(in.RegistrationMethod)
	out.RegistrationAddress = in.RegistrationAddress
	out.RolloutStrategy = (*v1beta1.RolloutStrategy)(unsafe.Pointer(in.RolloutStrategy))
	return nil
}

// Convert_v1alpha1_RKE2ControlPlaneSpec_To_v1beta1_RKE2ControlPlaneSpec is an autogenerated conversion function.
func Convert_v1alpha1_RKE2ControlPlaneSpec_To_v1beta1_RKE2ControlPlaneSpec(in *RKE2ControlPlaneSpec, out *v1beta1.RKE2ControlPlaneSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlaneSpec_To_v1beta1_RKE2ControlPlaneSpec(in, out, s)
}

func autoConvert_v1beta1_RKE2ControlPlaneSpec_To_v1alpha1_RKE2ControlPlaneSpec(in *v1beta1.RKE2ControlPlaneSpec, out *RKE2ControlPlaneSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_RKE2ConfigSpec_To_v1alpha1_RKE2ConfigSpec(&in.RKE2ConfigSpec, &out.RKE2ConfigSpec, s); err != nil {
		return err
	}
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	// WARNING: in.Version requires manual conversion: does not exist in peer-type
	// WARNING: in.MachineTemplate requires manual conversion: does not exist in peer-type
	if err := Convert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(&in.ServerConfig, &out.ServerConfig, s); err != nil {
		return err
	}
	out.ManifestsConfigMapReference = in.ManifestsConfigMapReference
	out.InfrastructureRef = in.InfrastructureRef
	out.NodeDrainTimeout = (*metav1.Duration)(unsafe.Pointer(in.NodeDrainTimeout))
	out.RegistrationMethod = RegistrationMethod(in.RegistrationMethod)
	out.RegistrationAddress = in.RegistrationAddress
	out.RolloutStrategy = (*RolloutStrategy)(unsafe.Pointer(in.RolloutStrategy))
	return nil
}

func autoConvert_v1alpha1_RKE2ControlPlaneStatus_To_v1beta1_RKE2ControlPlaneStatus(in *RKE2ControlPlaneStatus, out *v1beta1.RKE2ControlPlaneStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Initialized = in.Initialized
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.FailureReason = in.FailureReason
	out.FailureMessage = in.FailureMessage
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*clusterapiapiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.AvailableServerIPs = *(*[]string)(unsafe.Pointer(&in.AvailableServerIPs))
	return nil
}

func autoConvert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneStatus(in *v1beta1.RKE2ControlPlaneStatus, out *RKE2ControlPlaneStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Initialized = in.Initialized
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.FailureReason = in.FailureReason
	out.FailureMessage = in.FailureMessage
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*clusterapiapiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	out.Replicas = in.Replicas
	// WARNING: in.Version requires manual conversion: does not exist in peer-type
	out.ReadyReplicas = in.ReadyReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.AvailableServerIPs = *(*[]string)(unsafe.Pointer(&in.AvailableServerIPs))
	return nil
}

func autoConvert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate(in *RKE2ControlPlaneTemplate, out *v1beta1.RKE2ControlPlaneTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_RKE2ControlPlaneTemplateSpec_To_v1beta1_RKE2ControlPlaneTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_RKE2ControlPlaneTemplateStatus_To_v1beta1_RKE2ControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate is an autogenerated conversion function.
func Convert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate(in *RKE2ControlPlaneTemplate, out *v1beta1.RKE2ControlPlaneTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate(in, out, s)
}

func autoConvert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate(in *v1beta1.RKE2ControlPlaneTemplate, out *RKE2ControlPlaneTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_RKE2ControlPlaneTemplateSpec_To_v1alpha1_RKE2ControlPlaneTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_RKE2ControlPlaneStatus_To_v1alpha1_RKE2ControlPlaneTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate is an autogenerated conversion function.
func Convert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate(in *v1beta1.RKE2ControlPlaneTemplate, out *RKE2ControlPlaneTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate(in, out, s)
}

func autoConvert_v1alpha1_RKE2ControlPlaneTemplateList_To_v1beta1_RKE2ControlPlaneTemplateList(in *RKE2ControlPlaneTemplateList, out *v1beta1.RKE2ControlPlaneTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.RKE2ControlPlaneTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_RKE2ControlPlaneTemplate_To_v1beta1_RKE2ControlPlaneTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_RKE2ControlPlaneTemplateList_To_v1beta1_RKE2ControlPlaneTemplateList is an autogenerated conversion function.
func Convert_v1alpha1_RKE2ControlPlaneTemplateList_To_v1beta1_RKE2ControlPlaneTemplateList(in *RKE2ControlPlaneTemplateList, out *v1beta1.RKE2ControlPlaneTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ControlPlaneTemplateList_To_v1beta1_RKE2ControlPlaneTemplateList(in, out, s)
}

func autoConvert_v1beta1_RKE2ControlPlaneTemplateList_To_v1alpha1_RKE2ControlPlaneTemplateList(in *v1beta1.RKE2ControlPlaneTemplateList, out *RKE2ControlPlaneTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKE2ControlPlaneTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_RKE2ControlPlaneTemplate_To_v1alpha1_RKE2ControlPlaneTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_RKE2ControlPlaneTemplateList_To_v1alpha1_RKE2ControlPlaneTemplateList is an autogenerated conversion function.
func Convert_v1beta1_RKE2ControlPlaneTemplateList_To_v1alpha1_RKE2ControlPlaneTemplateList(in *v1beta1.RKE2ControlPlaneTemplateList, out *RKE2ControlPlaneTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_RKE2ControlPlaneTemplateList_To_v1alpha1_RKE2ControlPlaneTemplateList(in, out, s)
}

func autoConvert_v1alpha1_RKE2ControlPlaneTemplateSpec_To_v1beta1_RKE2ControlPlaneTemplateSpec(in *RKE2ControlPlaneTemplateSpec, out *v1beta1.RKE2ControlPlaneTemplateSpec, s conversion.Scope) error {
	return nil
}

func autoConvert_v1beta1_RKE2ControlPlaneTemplateSpec_To_v1alpha1_RKE2ControlPlaneTemplateSpec(in *v1beta1.RKE2ControlPlaneTemplateSpec, out *RKE2ControlPlaneTemplateSpec, s conversion.Scope) error {
	// WARNING: in.Template requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha1_RKE2ServerConfig_To_v1beta1_RKE2ServerConfig(in *RKE2ServerConfig, out *v1beta1.RKE2ServerConfig, s conversion.Scope) error {
	out.AuditPolicySecret = (*v1.ObjectReference)(unsafe.Pointer(in.AuditPolicySecret))
	out.BindAddress = in.BindAddress
	out.AdvertiseAddress = in.AdvertiseAddress
	out.TLSSan = *(*[]string)(unsafe.Pointer(&in.TLSSan))
	out.ServiceNodePortRange = in.ServiceNodePortRange
	out.ClusterDNS = in.ClusterDNS
	out.ClusterDomain = in.ClusterDomain
	if err := Convert_v1alpha1_DisableComponents_To_v1beta1_DisableComponents(&in.DisableComponents, &out.DisableComponents, s); err != nil {
		return err
	}
	out.CNI = v1beta1.CNI(in.CNI)
	out.CNIMultusEnable = in.CNIMultusEnable
	out.PauseImage = in.PauseImage
	if err := Convert_v1alpha1_EtcdConfig_To_v1beta1_EtcdConfig(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	out.KubeAPIServer = (*apiv1beta1.ComponentConfig)(unsafe.Pointer(in.KubeAPIServer))
	out.KubeControllerManager = (*apiv1beta1.ComponentConfig)(unsafe.Pointer(in.KubeControllerManager))
	out.KubeScheduler = (*apiv1beta1.ComponentConfig)(unsafe.Pointer(in.KubeScheduler))
	out.CloudControllerManager = (*apiv1beta1.ComponentConfig)(unsafe.Pointer(in.CloudControllerManager))
	out.CloudProviderName = in.CloudProviderName
	out.CloudProviderConfigMap = (*v1.ObjectReference)(unsafe.Pointer(in.CloudProviderConfigMap))
	return nil
}

// Convert_v1alpha1_RKE2ServerConfig_To_v1beta1_RKE2ServerConfig is an autogenerated conversion function.
func Convert_v1alpha1_RKE2ServerConfig_To_v1beta1_RKE2ServerConfig(in *RKE2ServerConfig, out *v1beta1.RKE2ServerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_RKE2ServerConfig_To_v1beta1_RKE2ServerConfig(in, out, s)
}

func autoConvert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(in *v1beta1.RKE2ServerConfig, out *RKE2ServerConfig, s conversion.Scope) error {
	out.AuditPolicySecret = (*v1.ObjectReference)(unsafe.Pointer(in.AuditPolicySecret))
	out.BindAddress = in.BindAddress
	out.AdvertiseAddress = in.AdvertiseAddress
	out.TLSSan = *(*[]string)(unsafe.Pointer(&in.TLSSan))
	out.ServiceNodePortRange = in.ServiceNodePortRange
	out.ClusterDNS = in.ClusterDNS
	out.ClusterDomain = in.ClusterDomain
	if err := Convert_v1beta1_DisableComponents_To_v1alpha1_DisableComponents(&in.DisableComponents, &out.DisableComponents, s); err != nil {
		return err
	}
	out.CNI = CNI(in.CNI)
	out.CNIMultusEnable = in.CNIMultusEnable
	out.PauseImage = in.PauseImage
	if err := Convert_v1beta1_EtcdConfig_To_v1alpha1_EtcdConfig(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	out.KubeAPIServer = (*apiv1alpha1.ComponentConfig)(unsafe.Pointer(in.KubeAPIServer))
	out.KubeControllerManager = (*apiv1alpha1.ComponentConfig)(unsafe.Pointer(in.KubeControllerManager))
	out.KubeScheduler = (*apiv1alpha1.ComponentConfig)(unsafe.Pointer(in.KubeScheduler))
	out.CloudControllerManager = (*apiv1alpha1.ComponentConfig)(unsafe.Pointer(in.CloudControllerManager))
	out.CloudProviderName = in.CloudProviderName
	out.CloudProviderConfigMap = (*v1.ObjectReference)(unsafe.Pointer(in.CloudProviderConfigMap))
	return nil
}

// Convert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig is an autogenerated conversion function.
func Convert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(in *v1beta1.RKE2ServerConfig, out *RKE2ServerConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_RKE2ServerConfig_To_v1alpha1_RKE2ServerConfig(in, out, s)
}

func autoConvert_v1alpha1_RollingUpdate_To_v1beta1_RollingUpdate(in *RollingUpdate, out *v1beta1.RollingUpdate, s conversion.Scope) error {
	out.MaxSurge = (*intstr.IntOrString)(unsafe.Pointer(in.MaxSurge))
	return nil
}

// Convert_v1alpha1_RollingUpdate_To_v1beta1_RollingUpdate is an autogenerated conversion function.
func Convert_v1alpha1_RollingUpdate_To_v1beta1_RollingUpdate(in *RollingUpdate, out *v1beta1.RollingUpdate, s conversion.Scope) error {
	return autoConvert_v1alpha1_RollingUpdate_To_v1beta1_RollingUpdate(in, out, s)
}

func autoConvert_v1beta1_RollingUpdate_To_v1alpha1_RollingUpdate(in *v1beta1.RollingUpdate, out *RollingUpdate, s conversion.Scope) error {
	out.MaxSurge = (*intstr.IntOrString)(unsafe.Pointer(in.MaxSurge))
	return nil
}

// Convert_v1beta1_RollingUpdate_To_v1alpha1_RollingUpdate is an autogenerated conversion function.
func Convert_v1beta1_RollingUpdate_To_v1alpha1_RollingUpdate(in *v1beta1.RollingUpdate, out *RollingUpdate, s conversion.Scope) error {
	return autoConvert_v1beta1_RollingUpdate_To_v1alpha1_RollingUpdate(in, out, s)
}

func autoConvert_v1alpha1_RolloutStrategy_To_v1beta1_RolloutStrategy(in *RolloutStrategy, out *v1beta1.RolloutStrategy, s conversion.Scope) error {
	out.Type = v1beta1.RolloutStrategyType(in.Type)
	out.RollingUpdate = (*v1beta1.RollingUpdate)(unsafe.Pointer(in.RollingUpdate))
	return nil
}

// Convert_v1alpha1_RolloutStrategy_To_v1beta1_RolloutStrategy is an autogenerated conversion function.
func Convert_v1alpha1_RolloutStrategy_To_v1beta1_RolloutStrategy(in *RolloutStrategy, out *v1beta1.RolloutStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_RolloutStrategy_To_v1beta1_RolloutStrategy(in, out, s)
}

func autoConvert_v1beta1_RolloutStrategy_To_v1alpha1_RolloutStrategy(in *v1beta1.RolloutStrategy, out *RolloutStrategy, s conversion.Scope) error {
	out.Type = RolloutStrategyType(in.Type)
	out.RollingUpdate = (*RollingUpdate)(unsafe.Pointer(in.RollingUpdate))
	return nil
}

// Convert_v1beta1_RolloutStrategy_To_v1alpha1_RolloutStrategy is an autogenerated conversion function.
func Convert_v1beta1_RolloutStrategy_To_v1alpha1_RolloutStrategy(in *v1beta1.RolloutStrategy, out *RolloutStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_RolloutStrategy_To_v1alpha1_RolloutStrategy(in, out, s)
}