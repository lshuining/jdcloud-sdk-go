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
    cdn "github.com/lshuining/jdcloud-sdk-go/services/cdn/models"
)

type QueryJDBoxStatisticsDataWithGroupRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,时间戳 (Optional) */
    StartTime *int64 `json:"startTime"`

    /* 查询截止时间,时间戳 (Optional) */
    EndTime *int64 `json:"endTime"`

    /* 取值范围：area，isp，pin ,mac_addr，category，多个用,隔开,多个维度的查询，必须要限制较短的时间间隔 (Optional) */
    GroupBy *string `json:"groupBy"`

    /* 查询的字段，取值范围(avgbandwidth,pv,flow)。多个用逗号分隔。默认为空，表示查询带宽流量 (Optional) */
    Fields *string `json:"fields"`

    /* 区域 (Optional) */
    Area *string `json:"area"`

    /* 运营商 (Optional) */
    Isp *string `json:"isp"`

    /* 查询周期，当前取值范围：“oneMin,fiveMin”，分别表示1min，5min。默认为空，表示fiveMin (Optional) */
    Period *string `json:"period"`

    /* 业务类型 (Optional) */
    Category *string `json:"category"`

    /* 设备id (Optional) */
    MacAddr *string `json:"macAddr"`

    /* 插件pin (Optional) */
    PluginPin *string `json:"pluginPin"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryJDBoxStatisticsDataWithGroupRequest(
) *QueryJDBoxStatisticsDataWithGroupRequest {

	return &QueryJDBoxStatisticsDataWithGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/jdBoxStatisticsWithGroup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,时间戳 (Optional)
 * param endTime: 查询截止时间,时间戳 (Optional)
 * param groupBy: 取值范围：area，isp，pin ,mac_addr，category，多个用,隔开,多个维度的查询，必须要限制较短的时间间隔 (Optional)
 * param fields: 查询的字段，取值范围(avgbandwidth,pv,flow)。多个用逗号分隔。默认为空，表示查询带宽流量 (Optional)
 * param area: 区域 (Optional)
 * param isp: 运营商 (Optional)
 * param period: 查询周期，当前取值范围：“oneMin,fiveMin”，分别表示1min，5min。默认为空，表示fiveMin (Optional)
 * param category: 业务类型 (Optional)
 * param macAddr: 设备id (Optional)
 * param pluginPin: 插件pin (Optional)
 */
func NewQueryJDBoxStatisticsDataWithGroupRequestWithAllParams(
    startTime *int64,
    endTime *int64,
    groupBy *string,
    fields *string,
    area *string,
    isp *string,
    period *string,
    category *string,
    macAddr *string,
    pluginPin *string,
) *QueryJDBoxStatisticsDataWithGroupRequest {

    return &QueryJDBoxStatisticsDataWithGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/jdBoxStatisticsWithGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        GroupBy: groupBy,
        Fields: fields,
        Area: area,
        Isp: isp,
        Period: period,
        Category: category,
        MacAddr: macAddr,
        PluginPin: pluginPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryJDBoxStatisticsDataWithGroupRequestWithoutParam() *QueryJDBoxStatisticsDataWithGroupRequest {

    return &QueryJDBoxStatisticsDataWithGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/jdBoxStatisticsWithGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,时间戳(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetStartTime(startTime int64) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,时间戳(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetEndTime(endTime int64) {
    r.EndTime = &endTime
}

/* param groupBy: 取值范围：area，isp，pin ,mac_addr，category，多个用,隔开,多个维度的查询，必须要限制较短的时间间隔(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetGroupBy(groupBy string) {
    r.GroupBy = &groupBy
}

/* param fields: 查询的字段，取值范围(avgbandwidth,pv,flow)。多个用逗号分隔。默认为空，表示查询带宽流量(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: 区域(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: 运营商(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param period: 查询周期，当前取值范围：“oneMin,fiveMin”，分别表示1min，5min。默认为空，表示fiveMin(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param category: 业务类型(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetCategory(category string) {
    r.Category = &category
}

/* param macAddr: 设备id(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetMacAddr(macAddr string) {
    r.MacAddr = &macAddr
}

/* param pluginPin: 插件pin(Optional) */
func (r *QueryJDBoxStatisticsDataWithGroupRequest) SetPluginPin(pluginPin string) {
    r.PluginPin = &pluginPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryJDBoxStatisticsDataWithGroupRequest) GetRegionId() string {
    return ""
}

type QueryJDBoxStatisticsDataWithGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryJDBoxStatisticsDataWithGroupResult `json:"result"`
}

type QueryJDBoxStatisticsDataWithGroupResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Statistics []cdn.StatisticsDataItem `json:"statistics"`
}