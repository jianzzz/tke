// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "tkestack.io/tke/pkg/gateway/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Auth)(nil), (*config.Auth)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Auth_To_config_Auth(a.(*Auth), b.(*config.Auth), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Auth)(nil), (*Auth)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Auth_To_v1_Auth(a.(*config.Auth), b.(*Auth), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Component)(nil), (*config.Component)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Component_To_config_Component(a.(*Component), b.(*config.Component), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Component)(nil), (*Component)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Component_To_v1_Component(a.(*config.Component), b.(*Component), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Components)(nil), (*config.Components)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Components_To_config_Components(a.(*Components), b.(*config.Components), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Components)(nil), (*Components)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Components_To_v1_Components(a.(*config.Components), b.(*Components), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FrontProxyComponent)(nil), (*config.FrontProxyComponent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_FrontProxyComponent_To_config_FrontProxyComponent(a.(*FrontProxyComponent), b.(*config.FrontProxyComponent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.FrontProxyComponent)(nil), (*FrontProxyComponent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_FrontProxyComponent_To_v1_FrontProxyComponent(a.(*config.FrontProxyComponent), b.(*FrontProxyComponent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*GatewayConfiguration)(nil), (*config.GatewayConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GatewayConfiguration_To_config_GatewayConfiguration(a.(*GatewayConfiguration), b.(*config.GatewayConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GatewayConfiguration)(nil), (*GatewayConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GatewayConfiguration_To_v1_GatewayConfiguration(a.(*config.GatewayConfiguration), b.(*GatewayConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PassthroughComponent)(nil), (*config.PassthroughComponent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PassthroughComponent_To_config_PassthroughComponent(a.(*PassthroughComponent), b.(*config.PassthroughComponent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PassthroughComponent)(nil), (*PassthroughComponent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PassthroughComponent_To_v1_PassthroughComponent(a.(*config.PassthroughComponent), b.(*PassthroughComponent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Registry)(nil), (*config.Registry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Registry_To_config_Registry(a.(*Registry), b.(*config.Registry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Registry)(nil), (*Registry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Registry_To_v1_Registry(a.(*config.Registry), b.(*Registry), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Auth_To_config_Auth(in *Auth, out *config.Auth, s conversion.Scope) error {
	out.DefaultTenant = in.DefaultTenant
	return nil
}

// Convert_v1_Auth_To_config_Auth is an autogenerated conversion function.
func Convert_v1_Auth_To_config_Auth(in *Auth, out *config.Auth, s conversion.Scope) error {
	return autoConvert_v1_Auth_To_config_Auth(in, out, s)
}

func autoConvert_config_Auth_To_v1_Auth(in *config.Auth, out *Auth, s conversion.Scope) error {
	out.DefaultTenant = in.DefaultTenant
	return nil
}

// Convert_config_Auth_To_v1_Auth is an autogenerated conversion function.
func Convert_config_Auth_To_v1_Auth(in *config.Auth, out *Auth, s conversion.Scope) error {
	return autoConvert_config_Auth_To_v1_Auth(in, out, s)
}

func autoConvert_v1_Component_To_config_Component(in *Component, out *config.Component, s conversion.Scope) error {
	out.Address = in.Address
	out.FrontProxy = (*config.FrontProxyComponent)(unsafe.Pointer(in.FrontProxy))
	out.Passthrough = (*config.PassthroughComponent)(unsafe.Pointer(in.Passthrough))
	return nil
}

// Convert_v1_Component_To_config_Component is an autogenerated conversion function.
func Convert_v1_Component_To_config_Component(in *Component, out *config.Component, s conversion.Scope) error {
	return autoConvert_v1_Component_To_config_Component(in, out, s)
}

func autoConvert_config_Component_To_v1_Component(in *config.Component, out *Component, s conversion.Scope) error {
	out.Address = in.Address
	out.FrontProxy = (*FrontProxyComponent)(unsafe.Pointer(in.FrontProxy))
	out.Passthrough = (*PassthroughComponent)(unsafe.Pointer(in.Passthrough))
	return nil
}

// Convert_config_Component_To_v1_Component is an autogenerated conversion function.
func Convert_config_Component_To_v1_Component(in *config.Component, out *Component, s conversion.Scope) error {
	return autoConvert_config_Component_To_v1_Component(in, out, s)
}

func autoConvert_v1_Components_To_config_Components(in *Components, out *config.Components, s conversion.Scope) error {
	out.Platform = (*config.Component)(unsafe.Pointer(in.Platform))
	out.Business = (*config.Component)(unsafe.Pointer(in.Business))
	out.Notify = (*config.Component)(unsafe.Pointer(in.Notify))
	out.Monitor = (*config.Component)(unsafe.Pointer(in.Monitor))
	out.Auth = (*config.Component)(unsafe.Pointer(in.Auth))
	out.Registry = (*config.Component)(unsafe.Pointer(in.Registry))
	out.LogAgent = (*config.Component)(unsafe.Pointer(in.LogAgent))
	out.Audit = (*config.Component)(unsafe.Pointer(in.Audit))
	out.Application = (*config.Component)(unsafe.Pointer(in.Application))
	out.Mesh = (*config.Component)(unsafe.Pointer(in.Mesh))
	return nil
}

// Convert_v1_Components_To_config_Components is an autogenerated conversion function.
func Convert_v1_Components_To_config_Components(in *Components, out *config.Components, s conversion.Scope) error {
	return autoConvert_v1_Components_To_config_Components(in, out, s)
}

func autoConvert_config_Components_To_v1_Components(in *config.Components, out *Components, s conversion.Scope) error {
	out.Platform = (*Component)(unsafe.Pointer(in.Platform))
	out.Business = (*Component)(unsafe.Pointer(in.Business))
	out.Notify = (*Component)(unsafe.Pointer(in.Notify))
	out.Monitor = (*Component)(unsafe.Pointer(in.Monitor))
	out.Auth = (*Component)(unsafe.Pointer(in.Auth))
	out.Registry = (*Component)(unsafe.Pointer(in.Registry))
	out.LogAgent = (*Component)(unsafe.Pointer(in.LogAgent))
	out.Audit = (*Component)(unsafe.Pointer(in.Audit))
	out.Application = (*Component)(unsafe.Pointer(in.Application))
	out.Mesh = (*Component)(unsafe.Pointer(in.Mesh))
	return nil
}

// Convert_config_Components_To_v1_Components is an autogenerated conversion function.
func Convert_config_Components_To_v1_Components(in *config.Components, out *Components, s conversion.Scope) error {
	return autoConvert_config_Components_To_v1_Components(in, out, s)
}

func autoConvert_v1_FrontProxyComponent_To_config_FrontProxyComponent(in *FrontProxyComponent, out *config.FrontProxyComponent, s conversion.Scope) error {
	out.CAFile = in.CAFile
	out.ClientCertFile = in.ClientCertFile
	out.ClientKeyFile = in.ClientKeyFile
	out.UsernameHeader = in.UsernameHeader
	out.GroupsHeader = in.GroupsHeader
	out.ExtraPrefixHeader = in.ExtraPrefixHeader
	return nil
}

// Convert_v1_FrontProxyComponent_To_config_FrontProxyComponent is an autogenerated conversion function.
func Convert_v1_FrontProxyComponent_To_config_FrontProxyComponent(in *FrontProxyComponent, out *config.FrontProxyComponent, s conversion.Scope) error {
	return autoConvert_v1_FrontProxyComponent_To_config_FrontProxyComponent(in, out, s)
}

func autoConvert_config_FrontProxyComponent_To_v1_FrontProxyComponent(in *config.FrontProxyComponent, out *FrontProxyComponent, s conversion.Scope) error {
	out.CAFile = in.CAFile
	out.ClientCertFile = in.ClientCertFile
	out.ClientKeyFile = in.ClientKeyFile
	out.UsernameHeader = in.UsernameHeader
	out.GroupsHeader = in.GroupsHeader
	out.ExtraPrefixHeader = in.ExtraPrefixHeader
	return nil
}

// Convert_config_FrontProxyComponent_To_v1_FrontProxyComponent is an autogenerated conversion function.
func Convert_config_FrontProxyComponent_To_v1_FrontProxyComponent(in *config.FrontProxyComponent, out *FrontProxyComponent, s conversion.Scope) error {
	return autoConvert_config_FrontProxyComponent_To_v1_FrontProxyComponent(in, out, s)
}

func autoConvert_v1_GatewayConfiguration_To_config_GatewayConfiguration(in *GatewayConfiguration, out *config.GatewayConfiguration, s conversion.Scope) error {
	out.DisableOIDCProxy = in.DisableOIDCProxy
	if err := Convert_v1_Components_To_config_Components(&in.Components, &out.Components, s); err != nil {
		return err
	}
	out.Registry = (*config.Registry)(unsafe.Pointer(in.Registry))
	out.Auth = (*config.Auth)(unsafe.Pointer(in.Auth))
	return nil
}

// Convert_v1_GatewayConfiguration_To_config_GatewayConfiguration is an autogenerated conversion function.
func Convert_v1_GatewayConfiguration_To_config_GatewayConfiguration(in *GatewayConfiguration, out *config.GatewayConfiguration, s conversion.Scope) error {
	return autoConvert_v1_GatewayConfiguration_To_config_GatewayConfiguration(in, out, s)
}

func autoConvert_config_GatewayConfiguration_To_v1_GatewayConfiguration(in *config.GatewayConfiguration, out *GatewayConfiguration, s conversion.Scope) error {
	out.DisableOIDCProxy = in.DisableOIDCProxy
	if err := Convert_config_Components_To_v1_Components(&in.Components, &out.Components, s); err != nil {
		return err
	}
	out.Registry = (*Registry)(unsafe.Pointer(in.Registry))
	out.Auth = (*Auth)(unsafe.Pointer(in.Auth))
	return nil
}

// Convert_config_GatewayConfiguration_To_v1_GatewayConfiguration is an autogenerated conversion function.
func Convert_config_GatewayConfiguration_To_v1_GatewayConfiguration(in *config.GatewayConfiguration, out *GatewayConfiguration, s conversion.Scope) error {
	return autoConvert_config_GatewayConfiguration_To_v1_GatewayConfiguration(in, out, s)
}

func autoConvert_v1_PassthroughComponent_To_config_PassthroughComponent(in *PassthroughComponent, out *config.PassthroughComponent, s conversion.Scope) error {
	out.CAFile = in.CAFile
	return nil
}

// Convert_v1_PassthroughComponent_To_config_PassthroughComponent is an autogenerated conversion function.
func Convert_v1_PassthroughComponent_To_config_PassthroughComponent(in *PassthroughComponent, out *config.PassthroughComponent, s conversion.Scope) error {
	return autoConvert_v1_PassthroughComponent_To_config_PassthroughComponent(in, out, s)
}

func autoConvert_config_PassthroughComponent_To_v1_PassthroughComponent(in *config.PassthroughComponent, out *PassthroughComponent, s conversion.Scope) error {
	out.CAFile = in.CAFile
	return nil
}

// Convert_config_PassthroughComponent_To_v1_PassthroughComponent is an autogenerated conversion function.
func Convert_config_PassthroughComponent_To_v1_PassthroughComponent(in *config.PassthroughComponent, out *PassthroughComponent, s conversion.Scope) error {
	return autoConvert_config_PassthroughComponent_To_v1_PassthroughComponent(in, out, s)
}

func autoConvert_v1_Registry_To_config_Registry(in *Registry, out *config.Registry, s conversion.Scope) error {
	out.DefaultTenant = in.DefaultTenant
	out.DomainSuffix = in.DomainSuffix
	return nil
}

// Convert_v1_Registry_To_config_Registry is an autogenerated conversion function.
func Convert_v1_Registry_To_config_Registry(in *Registry, out *config.Registry, s conversion.Scope) error {
	return autoConvert_v1_Registry_To_config_Registry(in, out, s)
}

func autoConvert_config_Registry_To_v1_Registry(in *config.Registry, out *Registry, s conversion.Scope) error {
	out.DefaultTenant = in.DefaultTenant
	out.DomainSuffix = in.DomainSuffix
	return nil
}

// Convert_config_Registry_To_v1_Registry is an autogenerated conversion function.
func Convert_config_Registry_To_v1_Registry(in *config.Registry, out *Registry, s conversion.Scope) error {
	return autoConvert_config_Registry_To_v1_Registry(in, out, s)
}
