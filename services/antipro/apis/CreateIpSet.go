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
    antipro "github.com/lshuining/jdcloud-sdk-go/services/antipro/models"
)

type CreateIpSetRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 创建实例的 IP 库请求参数  */
    IpSetSpec *antipro.IpSetSpec `json:"ipSetSpec"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param ipSetSpec: 创建实例的 IP 库请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateIpSetRequest(
    regionId string,
    instanceId string,
    ipSetSpec *antipro.IpSetSpec,
) *CreateIpSetRequest {

	return &CreateIpSetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/ipSets",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        IpSetSpec: ipSetSpec,
	}
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param ipSetSpec: 创建实例的 IP 库请求参数 (Required)
 */
func NewCreateIpSetRequestWithAllParams(
    regionId string,
    instanceId string,
    ipSetSpec *antipro.IpSetSpec,
) *CreateIpSetRequest {

    return &CreateIpSetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ipSets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        IpSetSpec: ipSetSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateIpSetRequestWithoutParam() *CreateIpSetRequest {

    return &CreateIpSetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ipSets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *CreateIpSetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *CreateIpSetRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param ipSetSpec: 创建实例的 IP 库请求参数(Required) */
func (r *CreateIpSetRequest) SetIpSetSpec(ipSetSpec *antipro.IpSetSpec) {
    r.IpSetSpec = ipSetSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateIpSetRequest) GetRegionId() string {
    return r.RegionId
}

type CreateIpSetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateIpSetResult `json:"result"`
}

type CreateIpSetResult struct {
    IpSetId string `json:"ipSetId"`
}