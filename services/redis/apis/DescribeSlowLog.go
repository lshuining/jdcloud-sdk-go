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
    redis "github.com/lshuining/jdcloud-sdk-go/services/redis/models"
)

type DescribeSlowLogRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 分片id (Optional) */
    ShardId *string `json:"shardId"`

    /* 分片地址 (Optional) */
    ShardAddr *string `json:"shardAddr"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSlowLogRequest(
    regionId string,
    cacheInstanceId string,
) *DescribeSlowLogRequest {

	return &DescribeSlowLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/slowLog",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param shardId: 分片id (Optional)
 * param shardAddr: 分片地址 (Optional)
 */
func NewDescribeSlowLogRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    pageNumber *int,
    pageSize *int,
    startTime *string,
    endTime *string,
    shardId *string,
    shardAddr *string,
) *DescribeSlowLogRequest {

    return &DescribeSlowLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/slowLog",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        StartTime: startTime,
        EndTime: endTime,
        ShardId: shardId,
        ShardAddr: shardAddr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSlowLogRequestWithoutParam() *DescribeSlowLogRequest {

    return &DescribeSlowLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/slowLog",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *DescribeSlowLogRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *DescribeSlowLogRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeSlowLogRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeSlowLogRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param startTime: 开始时间(Optional) */
func (r *DescribeSlowLogRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *DescribeSlowLogRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param shardId: 分片id(Optional) */
func (r *DescribeSlowLogRequest) SetShardId(shardId string) {
    r.ShardId = &shardId
}

/* param shardAddr: 分片地址(Optional) */
func (r *DescribeSlowLogRequest) SetShardAddr(shardAddr string) {
    r.ShardAddr = &shardAddr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSlowLogRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSlowLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSlowLogResult `json:"result"`
}

type DescribeSlowLogResult struct {
    SlowLogs []redis.SlowLog `json:"slowLogs"`
    TotalCount int `json:"totalCount"`
}