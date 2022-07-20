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

type DescribeViewTreeRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 展示方式，暂时不使用 (Optional) */
    LoadMode *int `json:"loadMode"`

    /* 套餐ID，0->免费版 1->企业版 2->企业高级版  */
    PackId int `json:"packId"`

    /* view ID，默认为-1  */
    ViewId int `json:"viewId"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param packId: 套餐ID，0->免费版 1->企业版 2->企业高级版 (Required)
 * param viewId: view ID，默认为-1 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeViewTreeRequest(
    regionId string,
    domainId string,
    packId int,
    viewId int,
) *DescribeViewTreeRequest {

	return &DescribeViewTreeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/viewTree",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        PackId: packId,
        ViewId: viewId,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param loadMode: 展示方式，暂时不使用 (Optional)
 * param packId: 套餐ID，0->免费版 1->企业版 2->企业高级版 (Required)
 * param viewId: view ID，默认为-1 (Required)
 */
func NewDescribeViewTreeRequestWithAllParams(
    regionId string,
    domainId string,
    loadMode *int,
    packId int,
    viewId int,
) *DescribeViewTreeRequest {

    return &DescribeViewTreeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/viewTree",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        LoadMode: loadMode,
        PackId: packId,
        ViewId: viewId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeViewTreeRequestWithoutParam() *DescribeViewTreeRequest {

    return &DescribeViewTreeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/viewTree",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DescribeViewTreeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *DescribeViewTreeRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param loadMode: 展示方式，暂时不使用(Optional) */
func (r *DescribeViewTreeRequest) SetLoadMode(loadMode int) {
    r.LoadMode = &loadMode
}

/* param packId: 套餐ID，0->免费版 1->企业版 2->企业高级版(Required) */
func (r *DescribeViewTreeRequest) SetPackId(packId int) {
    r.PackId = packId
}

/* param viewId: view ID，默认为-1(Required) */
func (r *DescribeViewTreeRequest) SetViewId(viewId int) {
    r.ViewId = viewId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeViewTreeRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeViewTreeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeViewTreeResult `json:"result"`
}

type DescribeViewTreeResult struct {
    Data []domainservice.ViewTree `json:"data"`
}