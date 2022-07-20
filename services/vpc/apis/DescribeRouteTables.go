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
    vpc "github.com/lshuining/jdcloud-sdk-go/services/vpc/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeRouteTablesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* routeTableIds - 路由表ID列表，支持多个
routeTableNames - 路由表名称列表，支持多个
vpcId	- 路由表所属vpc Id，支持单个
azType - 路由表az类型，取值：all(全部类型)，standard(标准路由表)，edge(边缘路由表)，默认standard ，支持单个
azs - 可用区，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRouteTablesRequest(
    regionId string,
) *DescribeRouteTablesRequest {

	return &DescribeRouteTablesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/routeTables/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: routeTableIds - 路由表ID列表，支持多个
routeTableNames - 路由表名称列表，支持多个
vpcId	- 路由表所属vpc Id，支持单个
azType - 路由表az类型，取值：all(全部类型)，standard(标准路由表)，edge(边缘路由表)，默认standard ，支持单个
azs - 可用区，支持多个
 (Optional)
 */
func NewDescribeRouteTablesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeRouteTablesRequest {

    return &DescribeRouteTablesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/routeTables/",
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
func NewDescribeRouteTablesRequestWithoutParam() *DescribeRouteTablesRequest {

    return &DescribeRouteTablesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/routeTables/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeRouteTablesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeRouteTablesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeRouteTablesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: routeTableIds - 路由表ID列表，支持多个
routeTableNames - 路由表名称列表，支持多个
vpcId	- 路由表所属vpc Id，支持单个
azType - 路由表az类型，取值：all(全部类型)，standard(标准路由表)，edge(边缘路由表)，默认standard ，支持单个
azs - 可用区，支持多个
(Optional) */
func (r *DescribeRouteTablesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRouteTablesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeRouteTablesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRouteTablesResult `json:"result"`
}

type DescribeRouteTablesResult struct {
    RouteTables []vpc.RouteTable `json:"routeTables"`
    TotalCount int `json:"totalCount"`
}