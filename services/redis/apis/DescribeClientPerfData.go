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

type DescribeClientPerfDataRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 客户端ip  */
    Ip string `json:"ip"`

    /* 客户端uuid  */
    Uuid string `json:"uuid"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param ip: 客户端ip (Required)
 * param uuid: 客户端uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeClientPerfDataRequest(
    regionId string,
    cacheInstanceId string,
    ip string,
    uuid string,
) *DescribeClientPerfDataRequest {

	return &DescribeClientPerfDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/clientPerfData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        Ip: ip,
        Uuid: uuid,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param ip: 客户端ip (Required)
 * param uuid: 客户端uuid (Required)
 */
func NewDescribeClientPerfDataRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    ip string,
    uuid string,
) *DescribeClientPerfDataRequest {

    return &DescribeClientPerfDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/clientPerfData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        Ip: ip,
        Uuid: uuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeClientPerfDataRequestWithoutParam() *DescribeClientPerfDataRequest {

    return &DescribeClientPerfDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/clientPerfData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *DescribeClientPerfDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *DescribeClientPerfDataRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param ip: 客户端ip(Required) */
func (r *DescribeClientPerfDataRequest) SetIp(ip string) {
    r.Ip = ip
}

/* param uuid: 客户端uuid(Required) */
func (r *DescribeClientPerfDataRequest) SetUuid(uuid string) {
    r.Uuid = uuid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeClientPerfDataRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeClientPerfDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeClientPerfDataResult `json:"result"`
}

type DescribeClientPerfDataResult struct {
    PerformanceData []redis.PerformanceDataMsg `json:"performanceData"`
    DefaultFilterValue redis.DefaultFilterValue `json:"defaultFilterValue"`
}