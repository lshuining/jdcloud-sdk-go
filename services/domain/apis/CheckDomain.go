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
    domain "github.com/lshuining/jdcloud-sdk-go/services/domain/models"
)

type CheckDomainRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 要检查的域名  */
    DomainName string `json:"domainName"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainName: 要检查的域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckDomainRequest(
    regionId string,
    domainName string,
) *CheckDomainRequest {

	return &CheckDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain:check",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainName: domainName,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainName: 要检查的域名 (Required)
 */
func NewCheckDomainRequestWithAllParams(
    regionId string,
    domainName string,
) *CheckDomainRequest {

    return &CheckDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain:check",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainName: domainName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckDomainRequestWithoutParam() *CheckDomainRequest {

    return &CheckDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain:check",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *CheckDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainName: 要检查的域名(Required) */
func (r *CheckDomainRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckDomainRequest) GetRegionId() string {
    return r.RegionId
}

type CheckDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckDomainResult `json:"result"`
}

type CheckDomainResult struct {
    Data domain.CheckDomain `json:"data"`
}