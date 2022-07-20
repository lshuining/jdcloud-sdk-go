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
    dh "github.com/lshuining/jdcloud-sdk-go/services/dh/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeDedicatedHostsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* dedicatedHostId - 专有宿主机ID，精确匹配，支持多个
az - 可用区，精确匹配，支持多个
status - 专有宿主机状态，精确匹配，支持多个，<a href="http://docs.jdcloud.com/dedicated-hosts/api/dh_status">参考专有宿主机状态</a>
name - 专有宿主机名称，模糊匹配，支持单个
dedicatedPoolId - 专有宿主机池ID，精确匹配，支持多个
dedicatedHostType - 专有宿主机机型，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDedicatedHostsRequest(
    regionId string,
) *DescribeDedicatedHostsRequest {

	return &DescribeDedicatedHostsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dedicatedHosts",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param filters: dedicatedHostId - 专有宿主机ID，精确匹配，支持多个
az - 可用区，精确匹配，支持多个
status - 专有宿主机状态，精确匹配，支持多个，<a href="http://docs.jdcloud.com/dedicated-hosts/api/dh_status">参考专有宿主机状态</a>
name - 专有宿主机名称，模糊匹配，支持单个
dedicatedPoolId - 专有宿主机池ID，精确匹配，支持多个
dedicatedHostType - 专有宿主机机型，精确匹配，支持多个
 (Optional)
 */
func NewDescribeDedicatedHostsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeDedicatedHostsRequest {

    return &DescribeDedicatedHostsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedHosts",
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
func NewDescribeDedicatedHostsRequestWithoutParam() *DescribeDedicatedHostsRequest {

    return &DescribeDedicatedHostsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedHosts",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeDedicatedHostsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeDedicatedHostsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeDedicatedHostsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: dedicatedHostId - 专有宿主机ID，精确匹配，支持多个
az - 可用区，精确匹配，支持多个
status - 专有宿主机状态，精确匹配，支持多个，<a href="http://docs.jdcloud.com/dedicated-hosts/api/dh_status">参考专有宿主机状态</a>
name - 专有宿主机名称，模糊匹配，支持单个
dedicatedPoolId - 专有宿主机池ID，精确匹配，支持多个
dedicatedHostType - 专有宿主机机型，精确匹配，支持多个
(Optional) */
func (r *DescribeDedicatedHostsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDedicatedHostsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDedicatedHostsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDedicatedHostsResult `json:"result"`
}

type DescribeDedicatedHostsResult struct {
    DedicatedHosts []dh.DedicatedHost `json:"dedicatedHosts"`
    TotalCount int `json:"totalCount"`
}