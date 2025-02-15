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

type ModifyNodeGroupRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 工作节点组 ID  */
    NodeGroupId string `json:"nodeGroupId"`

    /* 工作节点组名称 (Optional) */
    Name *string `json:"name"`

    /* 工作节点组描述 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域 ID (Required)
 * param nodeGroupId: 工作节点组 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyNodeGroupRequest(
    regionId string,
    nodeGroupId string,
) *ModifyNodeGroupRequest {

	return &ModifyNodeGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/nodeGroups/{nodeGroupId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NodeGroupId: nodeGroupId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param nodeGroupId: 工作节点组 ID (Required)
 * param name: 工作节点组名称 (Optional)
 * param description: 工作节点组描述 (Optional)
 */
func NewModifyNodeGroupRequestWithAllParams(
    regionId string,
    nodeGroupId string,
    name *string,
    description *string,
) *ModifyNodeGroupRequest {

    return &ModifyNodeGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups/{nodeGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NodeGroupId: nodeGroupId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyNodeGroupRequestWithoutParam() *ModifyNodeGroupRequest {

    return &ModifyNodeGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups/{nodeGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *ModifyNodeGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param nodeGroupId: 工作节点组 ID(Required) */
func (r *ModifyNodeGroupRequest) SetNodeGroupId(nodeGroupId string) {
    r.NodeGroupId = nodeGroupId
}

/* param name: 工作节点组名称(Optional) */
func (r *ModifyNodeGroupRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 工作节点组描述(Optional) */
func (r *ModifyNodeGroupRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyNodeGroupRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyNodeGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyNodeGroupResult `json:"result"`
}

type ModifyNodeGroupResult struct {
}