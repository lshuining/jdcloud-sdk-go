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

type DescribePermissionsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码  */
    PageNumber int `json:"pageNumber"`

    /* 每页显示数目  */
    PageSize int `json:"pageSize"`

    /* 关键字 (Optional) */
    Keyword *string `json:"keyword"`

    /* 权限类型,0-全部，1-系统权限，2-自定义权限  */
    QueryType int `json:"queryType"`
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码 (Required)
 * param pageSize: 每页显示数目 (Required)
 * param queryType: 权限类型,0-全部，1-系统权限，2-自定义权限 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePermissionsRequest(
    regionId string,
    pageNumber int,
    pageSize int,
    queryType int,
) *DescribePermissionsRequest {

	return &DescribePermissionsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/permissions",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        QueryType: queryType,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码 (Required)
 * param pageSize: 每页显示数目 (Required)
 * param keyword: 关键字 (Optional)
 * param queryType: 权限类型,0-全部，1-系统权限，2-自定义权限 (Required)
 */
func NewDescribePermissionsRequestWithAllParams(
    regionId string,
    pageNumber int,
    pageSize int,
    keyword *string,
    queryType int,
) *DescribePermissionsRequest {

    return &DescribePermissionsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permissions",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Keyword: keyword,
        QueryType: queryType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePermissionsRequestWithoutParam() *DescribePermissionsRequest {

    return &DescribePermissionsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permissions",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribePermissionsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码(Required) */
func (r *DescribePermissionsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 每页显示数目(Required) */
func (r *DescribePermissionsRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param keyword: 关键字(Optional) */
func (r *DescribePermissionsRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}

/* param queryType: 权限类型,0-全部，1-系统权限，2-自定义权限(Required) */
func (r *DescribePermissionsRequest) SetQueryType(queryType int) {
    r.QueryType = queryType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePermissionsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePermissionsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePermissionsResult `json:"result"`
}

type DescribePermissionsResult struct {
    Total int `json:"total"`
    Permissions []iam.Permission `json:"permissions"`
}