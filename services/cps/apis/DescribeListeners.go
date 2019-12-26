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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeListenersRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 名称 (Optional) */
    Name *string `json:"name"`

    /* 负载均衡实例ID，精确匹配 (Optional) */
    LoadBalancerId *string `json:"loadBalancerId"`

    /* listenerId - 监听器ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeListenersRequest(
    regionId string,
) *DescribeListenersRequest {

	return &DescribeListenersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param name: 名称 (Optional)
 * param loadBalancerId: 负载均衡实例ID，精确匹配 (Optional)
 * param filters: listenerId - 监听器ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeListenersRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    name *string,
    loadBalancerId *string,
    filters []common.Filter,
) *DescribeListenersRequest {

    return &DescribeListenersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Name: name,
        LoadBalancerId: loadBalancerId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeListenersRequestWithoutParam() *DescribeListenersRequest {

    return &DescribeListenersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *DescribeListenersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeListenersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeListenersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param name: 名称(Optional) */
func (r *DescribeListenersRequest) SetName(name string) {
    r.Name = &name
}

/* param loadBalancerId: 负载均衡实例ID，精确匹配(Optional) */
func (r *DescribeListenersRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = &loadBalancerId
}

/* param filters: listenerId - 监听器ID，精确匹配，支持多个
(Optional) */
func (r *DescribeListenersRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeListenersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeListenersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeListenersResult `json:"result"`
}

type DescribeListenersResult struct {
    Listeners []cps.Listener `json:"listeners"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}