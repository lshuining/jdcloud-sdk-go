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
    domainservice "github.com/lshuining/jdcloud-sdk-go/services/domainservice/models"
)

type DescribeDomainsRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 分页查询时查询的每页的序号，起始值为1，默认为1  */
    PageNumber int `json:"pageNumber"`

    /* 分页查询时设置的每页行数，默认为10  */
    PageSize int `json:"pageSize"`

    /* 关键字，按照”%domainName%”模式匹配主域名 (Optional) */
    DomainName *string `json:"domainName"`

    /* 域名ID。不为0时，只查此domainId的域名 (Optional) */
    DomainId *int `json:"domainId"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param pageNumber: 分页查询时查询的每页的序号，起始值为1，默认为1 (Required)
 * param pageSize: 分页查询时设置的每页行数，默认为10 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDomainsRequest(
    regionId string,
    pageNumber int,
    pageSize int,
) *DescribeDomainsRequest {

	return &DescribeDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param pageNumber: 分页查询时查询的每页的序号，起始值为1，默认为1 (Required)
 * param pageSize: 分页查询时设置的每页行数，默认为10 (Required)
 * param domainName: 关键字，按照”%domainName%”模式匹配主域名 (Optional)
 * param domainId: 域名ID。不为0时，只查此domainId的域名 (Optional)
 */
func NewDescribeDomainsRequestWithAllParams(
    regionId string,
    pageNumber int,
    pageSize int,
    domainName *string,
    domainId *int,
) *DescribeDomainsRequest {

    return &DescribeDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        DomainName: domainName,
        DomainId: domainId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDomainsRequestWithoutParam() *DescribeDomainsRequest {

    return &DescribeDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DescribeDomainsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 分页查询时查询的每页的序号，起始值为1，默认为1(Required) */
func (r *DescribeDomainsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页查询时设置的每页行数，默认为10(Required) */
func (r *DescribeDomainsRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param domainName: 关键字，按照”%domainName%”模式匹配主域名(Optional) */
func (r *DescribeDomainsRequest) SetDomainName(domainName string) {
    r.DomainName = &domainName
}

/* param domainId: 域名ID。不为0时，只查此domainId的域名(Optional) */
func (r *DescribeDomainsRequest) SetDomainId(domainId int) {
    r.DomainId = &domainId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDomainsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDomainsResult `json:"result"`
}

type DescribeDomainsResult struct {
    DataList []domainservice.DomainInfo `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}