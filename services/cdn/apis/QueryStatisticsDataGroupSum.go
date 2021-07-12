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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryStatisticsDataGroupSumRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可 (Optional) */
    SubDomain *string `json:"subDomain"`

    /* 需要查询的字段 (Optional) */
    Fields *string `json:"fields"`

    /* 查询的区域，如beijing,shanghai。多个用逗号分隔 (Optional) */
    Area *string `json:"area"`

    /* 查询的运营商，cmcc,cnc,ct，表示移动、联通、电信。多个用逗号分隔 (Optional) */
    Isp *string `json:"isp"`

    /* 是否查询回源统计信息。取值为true和false，默认为false。注意，如果查询回源信息，Fields的取值当前只支持oribandwidth，oripv，oricodestat三个，其余Fields忽略。 (Optional) */
    Origin *bool `json:"origin"`

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`

    /* 分组依据，可选值：[area,isp,domain,scheme] (Optional) */
    GroupBy *string `json:"groupBy"`

    /* true 代表查询境外数据，默认false查询境内数据 (Optional) */
    Abroad *bool `json:"abroad"`

    /* 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间 (Optional) */
    CacheType *string `json:"cacheType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryStatisticsDataGroupSumRequest(
) *QueryStatisticsDataGroupSumRequest {

	return &QueryStatisticsDataGroupSumRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/vodStatistics:groupSum",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 * param subDomain: 查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可 (Optional)
 * param fields: 需要查询的字段 (Optional)
 * param area: 查询的区域，如beijing,shanghai。多个用逗号分隔 (Optional)
 * param isp: 查询的运营商，cmcc,cnc,ct，表示移动、联通、电信。多个用逗号分隔 (Optional)
 * param origin: 是否查询回源统计信息。取值为true和false，默认为false。注意，如果查询回源信息，Fields的取值当前只支持oribandwidth，oripv，oricodestat三个，其余Fields忽略。 (Optional)
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 * param groupBy: 分组依据，可选值：[area,isp,domain,scheme] (Optional)
 * param abroad: true 代表查询境外数据，默认false查询境内数据 (Optional)
 * param cacheType: 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间 (Optional)
 */
func NewQueryStatisticsDataGroupSumRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    subDomain *string,
    fields *string,
    area *string,
    isp *string,
    origin *bool,
    period *string,
    groupBy *string,
    abroad *bool,
    cacheType *string,
) *QueryStatisticsDataGroupSumRequest {

    return &QueryStatisticsDataGroupSumRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/vodStatistics:groupSum",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        SubDomain: subDomain,
        Fields: fields,
        Area: area,
        Isp: isp,
        Origin: origin,
        Period: period,
        GroupBy: groupBy,
        Abroad: abroad,
        CacheType: cacheType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryStatisticsDataGroupSumRequestWithoutParam() *QueryStatisticsDataGroupSumRequest {

    return &QueryStatisticsDataGroupSumRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/vodStatistics:groupSum",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param subDomain: 查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetSubDomain(subDomain string) {
    r.SubDomain = &subDomain
}

/* param fields: 需要查询的字段(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: 查询的区域，如beijing,shanghai。多个用逗号分隔(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: 查询的运营商，cmcc,cnc,ct，表示移动、联通、电信。多个用逗号分隔(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param origin: 是否查询回源统计信息。取值为true和false，默认为false。注意，如果查询回源信息，Fields的取值当前只支持oribandwidth，oripv，oricodestat三个，其余Fields忽略。(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetOrigin(origin bool) {
    r.Origin = &origin
}

/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param groupBy: 分组依据，可选值：[area,isp,domain,scheme](Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetGroupBy(groupBy string) {
    r.GroupBy = &groupBy
}

/* param abroad: true 代表查询境外数据，默认false查询境内数据(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetAbroad(abroad bool) {
    r.Abroad = &abroad
}

/* param cacheType: 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间(Optional) */
func (r *QueryStatisticsDataGroupSumRequest) SetCacheType(cacheType string) {
    r.CacheType = &cacheType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryStatisticsDataGroupSumRequest) GetRegionId() string {
    return ""
}

type QueryStatisticsDataGroupSumResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryStatisticsDataGroupSumResult `json:"result"`
}

type QueryStatisticsDataGroupSumResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    Statistics []cdn.StatisticsGroupSumDataItem `json:"statistics"`
}