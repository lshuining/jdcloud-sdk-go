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

type CreateUserDomainRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 协议 (Optional) */
    Protocol *string `json:"protocol"`

    /* api分组id  */
    ApiGroupId string `json:"apiGroupId"`
}

/*
 * param regionId: 地域ID (Required)
 * param domain: 域名 (Required)
 * param apiGroupId: api分组id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateUserDomainRequest(
    regionId string,
    domain string,
    apiGroupId string,
) *CreateUserDomainRequest {

	return &CreateUserDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/userdomain",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Domain: domain,
        ApiGroupId: apiGroupId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param domain: 域名 (Required)
 * param protocol: 协议 (Optional)
 * param apiGroupId: api分组id (Required)
 */
func NewCreateUserDomainRequestWithAllParams(
    regionId string,
    domain string,
    protocol *string,
    apiGroupId string,
) *CreateUserDomainRequest {

    return &CreateUserDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userdomain",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Domain: domain,
        Protocol: protocol,
        ApiGroupId: apiGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateUserDomainRequestWithoutParam() *CreateUserDomainRequest {

    return &CreateUserDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userdomain",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateUserDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domain: 域名(Required) */
func (r *CreateUserDomainRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param protocol: 协议(Optional) */
func (r *CreateUserDomainRequest) SetProtocol(protocol string) {
    r.Protocol = &protocol
}

/* param apiGroupId: api分组id(Required) */
func (r *CreateUserDomainRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateUserDomainRequest) GetRegionId() string {
    return r.RegionId
}

type CreateUserDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateUserDomainResult `json:"result"`
}

type CreateUserDomainResult struct {
    DomainId string `json:"domainId"`
}