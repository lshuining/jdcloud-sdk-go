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

type DescribeDataSourceTableListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 数据源 ID  */
    DataSourceId string `json:"dataSourceId"`

    /* 页码；默认为1  */
    PageNumber int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100]  */
    PageSize int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param pageNumber: 页码；默认为1 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDataSourceTableListRequest(
    regionId string,
    dataSourceId string,
    pageNumber int,
    pageSize int,
) *DescribeDataSourceTableListRequest {

	return &DescribeDataSourceTableListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceId: dataSourceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param pageNumber: 页码；默认为1 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Required)
 */
func NewDescribeDataSourceTableListRequestWithAllParams(
    regionId string,
    dataSourceId string,
    pageNumber int,
    pageSize int,
) *DescribeDataSourceTableListRequest {

    return &DescribeDataSourceTableListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDataSourceTableListRequestWithoutParam() *DescribeDataSourceTableListRequest {

    return &DescribeDataSourceTableListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeDataSourceTableListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源 ID(Required) */
func (r *DescribeDataSourceTableListRequest) SetDataSourceId(dataSourceId string) {
    r.DataSourceId = dataSourceId
}

/* param pageNumber: 页码；默认为1(Required) */
func (r *DescribeDataSourceTableListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Required) */
func (r *DescribeDataSourceTableListRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDataSourceTableListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDataSourceTableListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDataSourceTableListResult `json:"result"`
}

type DescribeDataSourceTableListResult struct {
    List []string `json:"list"`
    TotalCount int `json:"totalCount"`
}