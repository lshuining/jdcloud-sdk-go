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

type DescribeDomainQueryCountRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 查询的主域名，请使用describeDomains接口获取  */
    DomainName string `json:"domainName"`

    /* 查询时间段的起始时间, UTC时间格式，例如2017-11-10T23:00:00Z  */
    Start string `json:"start"`

    /* 查询时间段的终止时间, UTC时间格式，例如2017-11-10T23:00:00Z  */
    End string `json:"end"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param domainName: 查询的主域名，请使用describeDomains接口获取 (Required)
 * param start: 查询时间段的起始时间, UTC时间格式，例如2017-11-10T23:00:00Z (Required)
 * param end: 查询时间段的终止时间, UTC时间格式，例如2017-11-10T23:00:00Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDomainQueryCountRequest(
    regionId string,
    domainId string,
    domainName string,
    start string,
    end string,
) *DescribeDomainQueryCountRequest {

	return &DescribeDomainQueryCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/queryCount",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        DomainName: domainName,
        Start: start,
        End: end,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param domainName: 查询的主域名，请使用describeDomains接口获取 (Required)
 * param start: 查询时间段的起始时间, UTC时间格式，例如2017-11-10T23:00:00Z (Required)
 * param end: 查询时间段的终止时间, UTC时间格式，例如2017-11-10T23:00:00Z (Required)
 */
func NewDescribeDomainQueryCountRequestWithAllParams(
    regionId string,
    domainId string,
    domainName string,
    start string,
    end string,
) *DescribeDomainQueryCountRequest {

    return &DescribeDomainQueryCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/queryCount",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        DomainName: domainName,
        Start: start,
        End: end,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDomainQueryCountRequestWithoutParam() *DescribeDomainQueryCountRequest {

    return &DescribeDomainQueryCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/queryCount",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DescribeDomainQueryCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *DescribeDomainQueryCountRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param domainName: 查询的主域名，请使用describeDomains接口获取(Required) */
func (r *DescribeDomainQueryCountRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param start: 查询时间段的起始时间, UTC时间格式，例如2017-11-10T23:00:00Z(Required) */
func (r *DescribeDomainQueryCountRequest) SetStart(start string) {
    r.Start = start
}

/* param end: 查询时间段的终止时间, UTC时间格式，例如2017-11-10T23:00:00Z(Required) */
func (r *DescribeDomainQueryCountRequest) SetEnd(end string) {
    r.End = end
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDomainQueryCountRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDomainQueryCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDomainQueryCountResult `json:"result"`
}

type DescribeDomainQueryCountResult struct {
    Time []int64 `json:"time"`
    Traffic []int64 `json:"traffic"`
}