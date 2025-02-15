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

type CheckGroupNameExistRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组名称  */
    GroupName string `json:"groupName"`
}

/*
 * param regionId: 地域ID (Required)
 * param groupName: 分组名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckGroupNameExistRequest(
    regionId string,
    groupName string,
) *CheckGroupNameExistRequest {

	return &CheckGroupNameExistRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups:checkGroupNameExist",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        GroupName: groupName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param groupName: 分组名称 (Required)
 */
func NewCheckGroupNameExistRequestWithAllParams(
    regionId string,
    groupName string,
) *CheckGroupNameExistRequest {

    return &CheckGroupNameExistRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups:checkGroupNameExist",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        GroupName: groupName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckGroupNameExistRequestWithoutParam() *CheckGroupNameExistRequest {

    return &CheckGroupNameExistRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups:checkGroupNameExist",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CheckGroupNameExistRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param groupName: 分组名称(Required) */
func (r *CheckGroupNameExistRequest) SetGroupName(groupName string) {
    r.GroupName = groupName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckGroupNameExistRequest) GetRegionId() string {
    return r.RegionId
}

type CheckGroupNameExistResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckGroupNameExistResult `json:"result"`
}

type CheckGroupNameExistResult struct {
    ApiGroupId string `json:"apiGroupId"`
}