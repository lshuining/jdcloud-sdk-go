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
    dcap "github.com/lshuining/jdcloud-sdk-go/services/dcap/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeInstanceListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1  */
    PageNumber int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100]  */
    PageSize int `json:"pageSize"`

    /* "instanceId: 实例ID，字符串数组，精确匹配"
"instanceName: 实例名称，字符串，模糊匹配"
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码；默认为1 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceListRequest(
    regionId string,
    pageNumber int,
    pageSize int,
) *DescribeInstanceListRequest {

	return &DescribeInstanceListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码；默认为1 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Required)
 * param filters: "instanceId: 实例ID，字符串数组，精确匹配"
"instanceName: 实例名称，字符串，模糊匹配"
 (Optional)
 */
func NewDescribeInstanceListRequestWithAllParams(
    regionId string,
    pageNumber int,
    pageSize int,
    filters []common.Filter,
) *DescribeInstanceListRequest {

    return &DescribeInstanceListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceListRequestWithoutParam() *DescribeInstanceListRequest {

    return &DescribeInstanceListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeInstanceListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Required) */
func (r *DescribeInstanceListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Required) */
func (r *DescribeInstanceListRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param filters: "instanceId: 实例ID，字符串数组，精确匹配"
"instanceName: 实例名称，字符串，模糊匹配"
(Optional) */
func (r *DescribeInstanceListRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceListResult `json:"result"`
}

type DescribeInstanceListResult struct {
    List []dcap.InstanceDesc `json:"list"`
    TotalCount int `json:"totalCount"`
}