// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/lshuining/jdcloud-sdk-go/core"
)

type DeployRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 发布的修订版本号  */
    Revision string `json:"revision"`

    /* 环境：test、preview、online  */
    Environment string `json:"environment"`

    /* 后端服务类型：mock、unique、vpc (Optional) */
    BackendServiceType *string `json:"backendServiceType"`

    /* 后端地址 (Optional) */
    BackendUrl *string `json:"backendUrl"`

    /* 描述 (Optional) */
    Description *string `json:"description"`

    /* 微服务网关名称 (Optional) */
    JdsfName *string `json:"jdsfName"`

    /* 微服务注册中心ID (Optional) */
    JdsfRegistryName *string `json:"jdsfRegistryName"`

    /* 微服务ID (Optional) */
    JdsfId *string `json:"jdsfId"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 发布的修订版本号 (Required)
 * param environment: 环境：test、preview、online (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeployRequest(
    regionId string,
    apiGroupId string,
    revision string,
    environment string,
) *DeployRequest {

	return &DeployRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        Environment: environment,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 发布的修订版本号 (Required)
 * param environment: 环境：test、preview、online (Required)
 * param backendServiceType: 后端服务类型：mock、unique、vpc (Optional)
 * param backendUrl: 后端地址 (Optional)
 * param description: 描述 (Optional)
 * param jdsfName: 微服务网关名称 (Optional)
 * param jdsfRegistryName: 微服务注册中心ID (Optional)
 * param jdsfId: 微服务ID (Optional)
 */
func NewDeployRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    environment string,
    backendServiceType *string,
    backendUrl *string,
    description *string,
    jdsfName *string,
    jdsfRegistryName *string,
    jdsfId *string,
) *DeployRequest {

    return &DeployRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        Environment: environment,
        BackendServiceType: backendServiceType,
        BackendUrl: backendUrl,
        Description: description,
        JdsfName: jdsfName,
        JdsfRegistryName: jdsfRegistryName,
        JdsfId: jdsfId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeployRequestWithoutParam() *DeployRequest {

    return &DeployRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeployRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *DeployRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 发布的修订版本号(Required) */
func (r *DeployRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param environment: 环境：test、preview、online(Required) */
func (r *DeployRequest) SetEnvironment(environment string) {
    r.Environment = environment
}

/* param backendServiceType: 后端服务类型：mock、unique、vpc(Optional) */
func (r *DeployRequest) SetBackendServiceType(backendServiceType string) {
    r.BackendServiceType = &backendServiceType
}

/* param backendUrl: 后端地址(Optional) */
func (r *DeployRequest) SetBackendUrl(backendUrl string) {
    r.BackendUrl = &backendUrl
}

/* param description: 描述(Optional) */
func (r *DeployRequest) SetDescription(description string) {
    r.Description = &description
}

/* param jdsfName: 微服务网关名称(Optional) */
func (r *DeployRequest) SetJdsfName(jdsfName string) {
    r.JdsfName = &jdsfName
}

/* param jdsfRegistryName: 微服务注册中心ID(Optional) */
func (r *DeployRequest) SetJdsfRegistryName(jdsfRegistryName string) {
    r.JdsfRegistryName = &jdsfRegistryName
}

/* param jdsfId: 微服务ID(Optional) */
func (r *DeployRequest) SetJdsfId(jdsfId string) {
    r.JdsfId = &jdsfId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeployRequest) GetRegionId() string {
    return r.RegionId
}

type DeployResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeployResult `json:"result"`
}

type DeployResult struct {
}