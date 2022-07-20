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
    iam "github.com/lshuining/jdcloud-sdk-go/services/iam/models"
)

type DescribeRolesRequest struct {

    core.JDCloudRequest

    /* 页码，默认1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认50，取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 角色名称关键词 (Optional) */
    RoleName *string `json:"roleName"`

    /* 角色类型，默认查找所有类型，2-服务相关角色，3-服务角色，4-用户角色 (Optional) */
    Type *int `json:"type"`

    /* 排序策略,0-按创建时间顺序排序,1-按创建时间倒序排序 (Optional) */
    Sort *int `json:"sort"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRolesRequest(
) *DescribeRolesRequest {

	return &DescribeRolesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/roles",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码，默认1 (Optional)
 * param pageSize: 分页大小，默认50，取值范围[10, 100] (Optional)
 * param roleName: 角色名称关键词 (Optional)
 * param type_: 角色类型，默认查找所有类型，2-服务相关角色，3-服务角色，4-用户角色 (Optional)
 * param sort: 排序策略,0-按创建时间顺序排序,1-按创建时间倒序排序 (Optional)
 */
func NewDescribeRolesRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    roleName *string,
    type_ *int,
    sort *int,
) *DescribeRolesRequest {

    return &DescribeRolesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/roles",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        RoleName: roleName,
        Type: type_,
        Sort: sort,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRolesRequestWithoutParam() *DescribeRolesRequest {

    return &DescribeRolesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/roles",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码，默认1(Optional) */
func (r *DescribeRolesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认50，取值范围[10, 100](Optional) */
func (r *DescribeRolesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param roleName: 角色名称关键词(Optional) */
func (r *DescribeRolesRequest) SetRoleName(roleName string) {
    r.RoleName = &roleName
}

/* param type_: 角色类型，默认查找所有类型，2-服务相关角色，3-服务角色，4-用户角色(Optional) */
func (r *DescribeRolesRequest) SetType(type_ int) {
    r.Type = &type_
}

/* param sort: 排序策略,0-按创建时间顺序排序,1-按创建时间倒序排序(Optional) */
func (r *DescribeRolesRequest) SetSort(sort int) {
    r.Sort = &sort
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRolesRequest) GetRegionId() string {
    return ""
}

type DescribeRolesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRolesResult `json:"result"`
}

type DescribeRolesResult struct {
    Total int `json:"total"`
    Roles []iam.ListRoleInfo `json:"roles"`
}