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

type QueryBandRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 地域 (Optional) */
    Area *string `json:"area"`

    /* 运营商 (Optional) */
    Isp *string `json:"isp"`

    /* 查询周期 (Optional) */
    Period *string `json:"period"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryBandRequest(
) *QueryBandRequest {

	return &QueryBandRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/bandQuery",
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
 * param area: 地域 (Optional)
 * param isp: 运营商 (Optional)
 * param period: 查询周期 (Optional)
 */
func NewQueryBandRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    area *string,
    isp *string,
    period *string,
) *QueryBandRequest {

    return &QueryBandRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/bandQuery",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Area: area,
        Isp: isp,
        Period: period,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryBandRequestWithoutParam() *QueryBandRequest {

    return &QueryBandRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/bandQuery",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryBandRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryBandRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryBandRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param area: 地域(Optional) */
func (r *QueryBandRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: 运营商(Optional) */
func (r *QueryBandRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param period: 查询周期(Optional) */
func (r *QueryBandRequest) SetPeriod(period string) {
    r.Period = &period
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryBandRequest) GetRegionId() string {
    return ""
}

type QueryBandResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryBandResult `json:"result"`
}

type QueryBandResult struct {
    ResultList []cdn.BandTrafficDataItem `json:"resultList"`
}