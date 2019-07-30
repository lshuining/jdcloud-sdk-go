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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type DescribeIpSetsRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小, 默认为10, 取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpSetsRequest(
    regionId string,
    instanceId string,
) *DescribeIpSetsRequest {

	return &DescribeIpSetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/ipSets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小, 默认为10, 取值范围[10, 100] (Optional)
 */
func NewDescribeIpSetsRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
) *DescribeIpSetsRequest {

    return &DescribeIpSetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ipSets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpSetsRequestWithoutParam() *DescribeIpSetsRequest {

    return &DescribeIpSetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ipSets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeIpSetsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DescribeIpSetsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeIpSetsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小, 默认为10, 取值范围[10, 100](Optional) */
func (r *DescribeIpSetsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpSetsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpSetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpSetsResult `json:"result"`
}

type DescribeIpSetsResult struct {
    DataList []ipanti.IpSet `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}