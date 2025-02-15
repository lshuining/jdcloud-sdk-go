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

type UpdateAgRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 高可用组 ID  */
    AgId string `json:"agId"`

    /* 描述，长度不超过 256 字符 (Optional) */
    Description *string `json:"description"`

    /* 高可用组名称，只支持中文、数字、大小写字母、英文下划线 “_” 及中划线 “-”，且不能超过 32 字符 (Optional) */
    Name *string `json:"name"`
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAgRequest(
    regionId string,
    agId string,
) *UpdateAgRequest {

	return &UpdateAgRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/availabilityGroups/{agId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
	}
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 * param description: 描述，长度不超过 256 字符 (Optional)
 * param name: 高可用组名称，只支持中文、数字、大小写字母、英文下划线 “_” 及中划线 “-”，且不能超过 32 字符 (Optional)
 */
func NewUpdateAgRequestWithAllParams(
    regionId string,
    agId string,
    description *string,
    name *string,
) *UpdateAgRequest {

    return &UpdateAgRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
        Description: description,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAgRequestWithoutParam() *UpdateAgRequest {

    return &UpdateAgRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *UpdateAgRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agId: 高可用组 ID(Required) */
func (r *UpdateAgRequest) SetAgId(agId string) {
    r.AgId = agId
}

/* param description: 描述，长度不超过 256 字符(Optional) */
func (r *UpdateAgRequest) SetDescription(description string) {
    r.Description = &description
}

/* param name: 高可用组名称，只支持中文、数字、大小写字母、英文下划线 “_” 及中划线 “-”，且不能超过 32 字符(Optional) */
func (r *UpdateAgRequest) SetName(name string) {
    r.Name = &name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAgRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateAgResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAgResult `json:"result"`
}

type UpdateAgResult struct {
}