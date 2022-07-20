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
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DeletePortSetRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 端口库 Id  */
    PortSetId string `json:"portSetId"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param portSetId: 端口库 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeletePortSetRequest(
    regionId string,
    instanceId string,
    portSetId string,
) *DeletePortSetRequest {

	return &DeletePortSetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/portSets/{portSetId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        PortSetId: portSetId,
	}
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param portSetId: 端口库 Id (Required)
 */
func NewDeletePortSetRequestWithAllParams(
    regionId string,
    instanceId string,
    portSetId string,
) *DeletePortSetRequest {

    return &DeletePortSetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/portSets/{portSetId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PortSetId: portSetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeletePortSetRequestWithoutParam() *DeletePortSetRequest {

    return &DeletePortSetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/portSets/{portSetId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *DeletePortSetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *DeletePortSetRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param portSetId: 端口库 Id(Required) */
func (r *DeletePortSetRequest) SetPortSetId(portSetId string) {
    r.PortSetId = portSetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeletePortSetRequest) GetRegionId() string {
    return r.RegionId
}

type DeletePortSetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeletePortSetResult `json:"result"`
}

type DeletePortSetResult struct {
    Data common.BatchResultDetail `json:"data"`
}