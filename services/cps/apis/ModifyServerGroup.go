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

type ModifyServerGroupRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 服务器组ID  */
    ServerGroupId string `json:"serverGroupId"`

    /* 名称 (Optional) */
    Name *string `json:"name"`
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param serverGroupId: 服务器组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyServerGroupRequest(
    regionId string,
    serverGroupId string,
) *ModifyServerGroupRequest {

	return &ModifyServerGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serverGroups/{serverGroupId}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServerGroupId: serverGroupId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param serverGroupId: 服务器组ID (Required)
 * param name: 名称 (Optional)
 */
func NewModifyServerGroupRequestWithAllParams(
    regionId string,
    serverGroupId string,
    name *string,
) *ModifyServerGroupRequest {

    return &ModifyServerGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serverGroups/{serverGroupId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServerGroupId: serverGroupId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyServerGroupRequestWithoutParam() *ModifyServerGroupRequest {

    return &ModifyServerGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serverGroups/{serverGroupId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *ModifyServerGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serverGroupId: 服务器组ID(Required) */
func (r *ModifyServerGroupRequest) SetServerGroupId(serverGroupId string) {
    r.ServerGroupId = serverGroupId
}

/* param name: 名称(Optional) */
func (r *ModifyServerGroupRequest) SetName(name string) {
    r.Name = &name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyServerGroupRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyServerGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyServerGroupResult `json:"result"`
}

type ModifyServerGroupResult struct {
    LoadBalancerId string `json:"loadBalancerId"`
    ServerGroupId string `json:"serverGroupId"`
    Name string `json:"name"`
}