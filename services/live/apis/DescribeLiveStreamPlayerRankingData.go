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
    live "github.com/lshuining/jdcloud-sdk-go/services/live/models"
)

type DescribeLiveStreamPlayerRankingDataRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    DomainName string `json:"domainName"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 协议，取值范围："hdl,hls"
  */
    ProtocolType string `json:"protocolType"`

    /* 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
  */
    StartTime string `json:"startTime"`

    /* 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，时间跨度不不超过⼀一天
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param domainName: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param protocolType: 协议，取值范围："hdl,hls"
 (Required)
 * param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveStreamPlayerRankingDataRequest(
    domainName string,
    appName string,
    protocolType string,
    startTime string,
) *DescribeLiveStreamPlayerRankingDataRequest {

	return &DescribeLiveStreamPlayerRankingDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeLiveStreamPlayerRankingData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        DomainName: domainName,
        AppName: appName,
        ProtocolType: protocolType,
        StartTime: startTime,
	}
}

/*
 * param domainName: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param protocolType: 协议，取值范围："hdl,hls"
 (Required)
 * param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
 (Required)
 * param endTime: 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，时间跨度不不超过⼀一天
 (Optional)
 */
func NewDescribeLiveStreamPlayerRankingDataRequestWithAllParams(
    domainName string,
    appName string,
    protocolType string,
    startTime string,
    endTime *string,
) *DescribeLiveStreamPlayerRankingDataRequest {

    return &DescribeLiveStreamPlayerRankingDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveStreamPlayerRankingData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
        AppName: appName,
        ProtocolType: protocolType,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveStreamPlayerRankingDataRequestWithoutParam() *DescribeLiveStreamPlayerRankingDataRequest {

    return &DescribeLiveStreamPlayerRankingDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveStreamPlayerRankingData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 推流域名(Required) */
func (r *DescribeLiveStreamPlayerRankingDataRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param appName: 应用名称(Required) */
func (r *DescribeLiveStreamPlayerRankingDataRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param protocolType: 协议，取值范围："hdl,hls"
(Required) */
func (r *DescribeLiveStreamPlayerRankingDataRequest) SetProtocolType(protocolType string) {
    r.ProtocolType = protocolType
}

/* param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
(Required) */
func (r *DescribeLiveStreamPlayerRankingDataRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，时间跨度不不超过⼀一天
(Optional) */
func (r *DescribeLiveStreamPlayerRankingDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveStreamPlayerRankingDataRequest) GetRegionId() string {
    return ""
}

type DescribeLiveStreamPlayerRankingDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveStreamPlayerRankingDataResult `json:"result"`
}

type DescribeLiveStreamPlayerRankingDataResult struct {
    DataList []live.LiveStreamPlayerRankingResult `json:"dataList"`
}