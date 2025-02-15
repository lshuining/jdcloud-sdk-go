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

type QueryStatisticsTopIpRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 待查询的子域名,查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可 (Optional) */
    SubDomain *string `json:"subDomain"`

    /* 查询的topN的条数，取值范围：1-100，默认为20 (Optional) */
    Size *int `json:"size"`

    /* 排序依据,当前可选：pv,flow, 分别表示按pv、按流量topN ip，默认为"pv" (Optional) */
    TopBy *string `json:"topBy"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryStatisticsTopIpRequest(
) *QueryStatisticsTopIpRequest {

	return &QueryStatisticsTopIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/statistics:topIp",
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
 * param subDomain: 待查询的子域名,查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可 (Optional)
 * param size: 查询的topN的条数，取值范围：1-100，默认为20 (Optional)
 * param topBy: 排序依据,当前可选：pv,flow, 分别表示按pv、按流量topN ip，默认为"pv" (Optional)
 */
func NewQueryStatisticsTopIpRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    subDomain *string,
    size *int,
    topBy *string,
) *QueryStatisticsTopIpRequest {

    return &QueryStatisticsTopIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:topIp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        SubDomain: subDomain,
        Size: size,
        TopBy: topBy,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryStatisticsTopIpRequestWithoutParam() *QueryStatisticsTopIpRequest {

    return &QueryStatisticsTopIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:topIp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsTopIpRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsTopIpRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsTopIpRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param subDomain: 待查询的子域名,查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可(Optional) */
func (r *QueryStatisticsTopIpRequest) SetSubDomain(subDomain string) {
    r.SubDomain = &subDomain
}

/* param size: 查询的topN的条数，取值范围：1-100，默认为20(Optional) */
func (r *QueryStatisticsTopIpRequest) SetSize(size int) {
    r.Size = &size
}

/* param topBy: 排序依据,当前可选：pv,flow, 分别表示按pv、按流量topN ip，默认为"pv"(Optional) */
func (r *QueryStatisticsTopIpRequest) SetTopBy(topBy string) {
    r.TopBy = &topBy
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryStatisticsTopIpRequest) GetRegionId() string {
    return ""
}

type QueryStatisticsTopIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryStatisticsTopIpResult `json:"result"`
}

type QueryStatisticsTopIpResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    IpData []cdn.StatisticsTopIpData `json:"ipData"`
}