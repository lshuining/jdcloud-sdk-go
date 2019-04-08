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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
)

type DescribeLiveTranscodeStreamPlayerUserNumRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    DomainName string `json:"domainName"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 运营商
 (Optional) */
    IspName *string `json:"ispName"`

    /* 查询的区域，如beijing,shanghai。多个用逗号分隔
 (Optional) */
    LocationName *string `json:"locationName"`

    /* 查询的流协议类型，取值范围："rtmp,hdl,hls"，多个时以逗号分隔
 (Optional) */
    ProtocolType *string `json:"protocolType"`

    /* 查询周期，当前取值范围：“oneMin,fiveMin,halfHour,hour,twoHour,sixHour,day,followTime”，分别表示1min，5min，半小时，1小时，2小时，6小时，1天，跟随时间。默认为空，表示fiveMin。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
 (Optional) */
    Period *string `json:"period"`

    /* 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
  */
    StartTime string `json:"startTime"`

    /* 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，为空时默认为当前时间
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param domainName: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveTranscodeStreamPlayerUserNumRequest(
    domainName string,
    appName string,
    startTime string,
) *DescribeLiveTranscodeStreamPlayerUserNumRequest {

	return &DescribeLiveTranscodeStreamPlayerUserNumRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeLiveTranscodeStreamPlayerUserNum",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        DomainName: domainName,
        AppName: appName,
        StartTime: startTime,
	}
}

/*
 * param domainName: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param ispName: 运营商
 (Optional)
 * param locationName: 查询的区域，如beijing,shanghai。多个用逗号分隔
 (Optional)
 * param protocolType: 查询的流协议类型，取值范围："rtmp,hdl,hls"，多个时以逗号分隔
 (Optional)
 * param period: 查询周期，当前取值范围：“oneMin,fiveMin,halfHour,hour,twoHour,sixHour,day,followTime”，分别表示1min，5min，半小时，1小时，2小时，6小时，1天，跟随时间。默认为空，表示fiveMin。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
 (Optional)
 * param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 * param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，为空时默认为当前时间
 (Optional)
 */
func NewDescribeLiveTranscodeStreamPlayerUserNumRequestWithAllParams(
    domainName string,
    appName string,
    ispName *string,
    locationName *string,
    protocolType *string,
    period *string,
    startTime string,
    endTime *string,
) *DescribeLiveTranscodeStreamPlayerUserNumRequest {

    return &DescribeLiveTranscodeStreamPlayerUserNumRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveTranscodeStreamPlayerUserNum",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
        AppName: appName,
        IspName: ispName,
        LocationName: locationName,
        ProtocolType: protocolType,
        Period: period,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveTranscodeStreamPlayerUserNumRequestWithoutParam() *DescribeLiveTranscodeStreamPlayerUserNumRequest {

    return &DescribeLiveTranscodeStreamPlayerUserNumRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveTranscodeStreamPlayerUserNum",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 推流域名(Required) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param appName: 应用名称(Required) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param ispName: 运营商
(Optional) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetIspName(ispName string) {
    r.IspName = &ispName
}

/* param locationName: 查询的区域，如beijing,shanghai。多个用逗号分隔
(Optional) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetLocationName(locationName string) {
    r.LocationName = &locationName
}

/* param protocolType: 查询的流协议类型，取值范围："rtmp,hdl,hls"，多个时以逗号分隔
(Optional) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetProtocolType(protocolType string) {
    r.ProtocolType = &protocolType
}

/* param period: 查询周期，当前取值范围：“oneMin,fiveMin,halfHour,hour,twoHour,sixHour,day,followTime”，分别表示1min，5min，半小时，1小时，2小时，6小时，1天，跟随时间。默认为空，表示fiveMin。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
(Optional) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
(Required) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，为空时默认为当前时间
(Optional) */
func (r *DescribeLiveTranscodeStreamPlayerUserNumRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveTranscodeStreamPlayerUserNumRequest) GetRegionId() string {
    return ""
}

type DescribeLiveTranscodeStreamPlayerUserNumResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveTranscodeStreamPlayerUserNumResult `json:"result"`
}

type DescribeLiveTranscodeStreamPlayerUserNumResult struct {
    DataList []live.LiveStreamUserNumResult `json:"dataList"`
}