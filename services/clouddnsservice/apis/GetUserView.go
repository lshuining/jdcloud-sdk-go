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
    clouddnsservice "github.com/lshuining/jdcloud-sdk-go/services/clouddnsservice/models"
)

type GetUserViewRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 主域名ID  */
    DomainId int `json:"domainId"`

    /* 自定义线路ID  */
    ViewId int `json:"viewId"`

    /* 自定义线路名称, 最多64个字符 (Optional) */
    ViewName *string `json:"viewName"`

    /* 分页参数，页的序号  */
    PageNumber int `json:"pageNumber"`

    /* 分页参数，每页含有的结果的数目  */
    PageSize int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 * param domainId: 主域名ID (Required)
 * param viewId: 自定义线路ID (Required)
 * param pageNumber: 分页参数，页的序号 (Required)
 * param pageSize: 分页参数，每页含有的结果的数目 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetUserViewRequest(
    regionId string,
    domainId int,
    viewId int,
    pageNumber int,
    pageSize int,
) *GetUserViewRequest {

	return &GetUserViewRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/userview/getUserView",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
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
 * param domainId: 主域名ID (Required)
 * param viewId: 自定义线路ID (Required)
 * param viewName: 自定义线路名称, 最多64个字符 (Optional)
 * param pageNumber: 分页参数，页的序号 (Required)
 * param pageSize: 分页参数，每页含有的结果的数目 (Required)
 */
func NewGetUserViewRequestWithAllParams(
    regionId string,
    domainId int,
    viewId int,
    viewName *string,
    pageNumber int,
    pageSize int,
) *GetUserViewRequest {

    return &GetUserViewRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userview/getUserView",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
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
func NewGetUserViewRequestWithoutParam() *GetUserViewRequest {

    return &GetUserViewRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userview/getUserView",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetUserViewRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 主域名ID(Required) */
func (r *GetUserViewRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

/* param viewId: 自定义线路ID(Required) */
func (r *GetUserViewRequest) SetViewId(viewId int) {
    r.ViewId = viewId
}

/* param viewName: 自定义线路名称, 最多64个字符(Optional) */
func (r *GetUserViewRequest) SetViewName(viewName string) {
    r.ViewName = &viewName
}

/* param pageNumber: 分页参数，页的序号(Required) */
func (r *GetUserViewRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页参数，每页含有的结果的数目(Required) */
func (r *GetUserViewRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetUserViewRequest) GetRegionId() string {
    return r.RegionId
}

type GetUserViewResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetUserViewResult `json:"result"`
}

type GetUserViewResult struct {
    DataList []clouddnsservice.UserViewInput `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}