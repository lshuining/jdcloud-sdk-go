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
    baseanti "github.com/jdcloud-api/jdcloud-sdk-go/services/baseanti/models"
)

type DescribeElasticIpResourcesRequest struct {

    core.JDCloudRequest

    /* 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州
  */
    RegionId string `json:"regionId"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeElasticIpResourcesRequest(
    regionId string,
) *DescribeElasticIpResourcesRequest {

	return &DescribeElasticIpResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIpResources",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州
 (Required)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 */
func NewDescribeElasticIpResourcesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
) *DescribeElasticIpResourcesRequest {

    return &DescribeElasticIpResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIpResources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeElasticIpResourcesRequestWithoutParam() *DescribeElasticIpResourcesRequest {

    return &DescribeElasticIpResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIpResources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州
(Required) */
func (r *DescribeElasticIpResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeElasticIpResourcesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *DescribeElasticIpResourcesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeElasticIpResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeElasticIpResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeElasticIpResourcesResult `json:"result"`
}

type DescribeElasticIpResourcesResult struct {
    DataList []baseanti.IpResource `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}