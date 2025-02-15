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

type DescribeUserViewRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 自定义线路ID  */
    ViewId int `json:"viewId"`

    /* 自定义线路名称, 最多64个字节，允许：数字、字母、下划线，-，中文 (Optional) */
    ViewName *string `json:"viewName"`

    /* 分页参数，页的序号  */
    PageNumber int `json:"pageNumber"`

    /* 分页参数，每页含有的结果的数目  */
    PageSize int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param viewId: 自定义线路ID (Required)
 * param pageNumber: 分页参数，页的序号 (Required)
 * param pageSize: 分页参数，每页含有的结果的数目 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUserViewRequest(
    regionId string,
    domainId string,
    viewId int,
    pageNumber int,
    pageSize int,
) *DescribeUserViewRequest {

	return &DescribeUserViewRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/UserView",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        ViewId: viewId,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param viewId: 自定义线路ID (Required)
 * param viewName: 自定义线路名称, 最多64个字节，允许：数字、字母、下划线，-，中文 (Optional)
 * param pageNumber: 分页参数，页的序号 (Required)
 * param pageSize: 分页参数，每页含有的结果的数目 (Required)
 */
func NewDescribeUserViewRequestWithAllParams(
    regionId string,
    domainId string,
    viewId int,
    viewName *string,
    pageNumber int,
    pageSize int,
) *DescribeUserViewRequest {

    return &DescribeUserViewRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/UserView",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        ViewId: viewId,
        ViewName: viewName,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUserViewRequestWithoutParam() *DescribeUserViewRequest {

    return &DescribeUserViewRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/UserView",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeUserViewRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *DescribeUserViewRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param viewId: 自定义线路ID(Required) */
func (r *DescribeUserViewRequest) SetViewId(viewId int) {
    r.ViewId = viewId
}

/* param viewName: 自定义线路名称, 最多64个字节，允许：数字、字母、下划线，-，中文(Optional) */
func (r *DescribeUserViewRequest) SetViewName(viewName string) {
    r.ViewName = &viewName
}

/* param pageNumber: 分页参数，页的序号(Required) */
func (r *DescribeUserViewRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页参数，每页含有的结果的数目(Required) */
func (r *DescribeUserViewRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUserViewRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeUserViewResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUserViewResult `json:"result"`
}

type DescribeUserViewResult struct {
    DataList []domainservice.UserViewInput `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}