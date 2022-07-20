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

type GetTargetsRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用getDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 子域名  */
    SubDomainName string `json:"subDomainName"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param subDomainName: 子域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTargetsRequest(
    regionId string,
    domainId string,
    subDomainName string,
) *GetTargetsRequest {

	return &GetTargetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/getTargets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
        SubDomainName: subDomainName,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param subDomainName: 子域名 (Required)
 */
func NewGetTargetsRequestWithAllParams(
    regionId string,
    domainId string,
    subDomainName string,
) *GetTargetsRequest {

    return &GetTargetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/getTargets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        SubDomainName: subDomainName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTargetsRequestWithoutParam() *GetTargetsRequest {

    return &GetTargetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/getTargets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetTargetsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用getDomains接口获取。(Required) */
func (r *GetTargetsRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param subDomainName: 子域名(Required) */
func (r *GetTargetsRequest) SetSubDomainName(subDomainName string) {
    r.SubDomainName = subDomainName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTargetsRequest) GetRegionId() string {
    return r.RegionId
}

type GetTargetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTargetsResult `json:"result"`
}

type GetTargetsResult struct {
    Data []string `json:"data"`
}