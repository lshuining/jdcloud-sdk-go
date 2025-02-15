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
    dh "github.com/lshuining/jdcloud-sdk-go/services/dh/models"
)

type AllocDedicatedHostsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 描述专有宿主机配置
  */
    DedicatedHostSpec *dh.DedicatedHostSpec `json:"dedicatedHostSpec"`

    /* 是否支持AZ内专有宿主机强制均衡，默认为preferred--非强制，取值[preferred--非强制,required--强制]
 (Optional) */
    DeployPolicy *string `json:"deployPolicy"`

    /* 购买云主机的数量；取值范围：[1,100]，默认为1。
 (Optional) */
    MaxCount *int `json:"maxCount"`

    /* 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。
 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: 地域ID (Required)
 * param dedicatedHostSpec: 描述专有宿主机配置
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAllocDedicatedHostsRequest(
    regionId string,
    dedicatedHostSpec *dh.DedicatedHostSpec,
) *AllocDedicatedHostsRequest {

	return &AllocDedicatedHostsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dedicatedHosts",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DedicatedHostSpec: dedicatedHostSpec,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param dedicatedHostSpec: 描述专有宿主机配置
 (Required)
 * param deployPolicy: 是否支持AZ内专有宿主机强制均衡，默认为preferred--非强制，取值[preferred--非强制,required--强制]
 (Optional)
 * param maxCount: 购买云主机的数量；取值范围：[1,100]，默认为1。
 (Optional)
 * param clientToken: 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。
 (Optional)
 */
func NewAllocDedicatedHostsRequestWithAllParams(
    regionId string,
    dedicatedHostSpec *dh.DedicatedHostSpec,
    deployPolicy *string,
    maxCount *int,
    clientToken *string,
) *AllocDedicatedHostsRequest {

    return &AllocDedicatedHostsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedHosts",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DedicatedHostSpec: dedicatedHostSpec,
        DeployPolicy: deployPolicy,
        MaxCount: maxCount,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAllocDedicatedHostsRequestWithoutParam() *AllocDedicatedHostsRequest {

    return &AllocDedicatedHostsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedHosts",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *AllocDedicatedHostsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dedicatedHostSpec: 描述专有宿主机配置
(Required) */
func (r *AllocDedicatedHostsRequest) SetDedicatedHostSpec(dedicatedHostSpec *dh.DedicatedHostSpec) {
    r.DedicatedHostSpec = dedicatedHostSpec
}

/* param deployPolicy: 是否支持AZ内专有宿主机强制均衡，默认为preferred--非强制，取值[preferred--非强制,required--强制]
(Optional) */
func (r *AllocDedicatedHostsRequest) SetDeployPolicy(deployPolicy string) {
    r.DeployPolicy = &deployPolicy
}

/* param maxCount: 购买云主机的数量；取值范围：[1,100]，默认为1。
(Optional) */
func (r *AllocDedicatedHostsRequest) SetMaxCount(maxCount int) {
    r.MaxCount = &maxCount
}

/* param clientToken: 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。
(Optional) */
func (r *AllocDedicatedHostsRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AllocDedicatedHostsRequest) GetRegionId() string {
    return r.RegionId
}

type AllocDedicatedHostsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AllocDedicatedHostsResult `json:"result"`
}

type AllocDedicatedHostsResult struct {
    DedicatedHostIds []string `json:"dedicatedHostIds"`
}