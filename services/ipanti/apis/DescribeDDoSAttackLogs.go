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
    ipanti "github.com/lshuining/jdcloud-sdk-go/services/ipanti/models"
)

type DescribeDDoSAttackLogsRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小, 默认为10, 取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    EndTime *string `json:"endTime"`

    /* 高防实例 ID (Optional) */
    InstanceId []string `json:"instanceId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDDoSAttackLogsRequest(
    regionId string,
    startTime string,
) *DescribeDDoSAttackLogsRequest {

	return &DescribeDDoSAttackLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/attacklog:describeDDoSAttackLogs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小, 默认为10, 取值范围[10, 100] (Optional)
 * param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional)
 * param instanceId: 高防实例 ID (Optional)
 */
func NewDescribeDDoSAttackLogsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    startTime string,
    endTime *string,
    instanceId []string,
) *DescribeDDoSAttackLogsRequest {

    return &DescribeDDoSAttackLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/attacklog:describeDDoSAttackLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        StartTime: startTime,
        EndTime: endTime,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDDoSAttackLogsRequestWithoutParam() *DescribeDDoSAttackLogsRequest {

    return &DescribeDDoSAttackLogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/attacklog:describeDDoSAttackLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeDDoSAttackLogsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeDDoSAttackLogsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小, 默认为10, 取值范围[10, 100](Optional) */
func (r *DescribeDDoSAttackLogsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeDDoSAttackLogsRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Optional) */
func (r *DescribeDDoSAttackLogsRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param instanceId: 高防实例 ID(Optional) */
func (r *DescribeDDoSAttackLogsRequest) SetInstanceId(instanceId []string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDDoSAttackLogsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDDoSAttackLogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDDoSAttackLogsResult `json:"result"`
}

type DescribeDDoSAttackLogsResult struct {
    DataList []ipanti.DDoSAttackLog `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}