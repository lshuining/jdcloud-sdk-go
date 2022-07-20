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

type ModifyAnalysisTimeRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 自动缓存分析时间，设置为-表示关闭，否则为：HH:mm-HH:mm 时区，例如"01:00-02:00 +0800"，表示东八区的1点到2点  */
    AnalysisTime string `json:"analysisTime"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param analysisTime: 自动缓存分析时间，设置为-表示关闭，否则为：HH:mm-HH:mm 时区，例如"01:00-02:00 +0800"，表示东八区的1点到2点 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyAnalysisTimeRequest(
    regionId string,
    cacheInstanceId string,
    analysisTime string,
) *ModifyAnalysisTimeRequest {

	return &ModifyAnalysisTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/analysisTime",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        AnalysisTime: analysisTime,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param analysisTime: 自动缓存分析时间，设置为-表示关闭，否则为：HH:mm-HH:mm 时区，例如"01:00-02:00 +0800"，表示东八区的1点到2点 (Required)
 */
func NewModifyAnalysisTimeRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    analysisTime string,
) *ModifyAnalysisTimeRequest {

    return &ModifyAnalysisTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/analysisTime",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        AnalysisTime: analysisTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyAnalysisTimeRequestWithoutParam() *ModifyAnalysisTimeRequest {

    return &ModifyAnalysisTimeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/analysisTime",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *ModifyAnalysisTimeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *ModifyAnalysisTimeRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param analysisTime: 自动缓存分析时间，设置为-表示关闭，否则为：HH:mm-HH:mm 时区，例如"01:00-02:00 +0800"，表示东八区的1点到2点(Required) */
func (r *ModifyAnalysisTimeRequest) SetAnalysisTime(analysisTime string) {
    r.AnalysisTime = analysisTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyAnalysisTimeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyAnalysisTimeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyAnalysisTimeResult `json:"result"`
}

type ModifyAnalysisTimeResult struct {
}