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

type CreateDomainGroupRequest struct {

    core.JDCloudRequest

    /* 是否共享内存，共享缓存仅对中国境内加速域名生效 (Optional) */
    ShareCache *string `json:"shareCache"`

    /* 主域名,开启共享缓存时必传 (Optional) */
    PrimaryDomain *string `json:"primaryDomain"`

    /* 域名组名称 (Optional) */
    DomainGroupName *string `json:"domainGroupName"`

    /* 域名组内域名包含主域名 (Optional) */
    Domains []string `json:"domains"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDomainGroupRequest(
) *CreateDomainGroupRequest {

	return &CreateDomainGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domainGroup:create",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param shareCache: 是否共享内存，共享缓存仅对中国境内加速域名生效 (Optional)
 * param primaryDomain: 主域名,开启共享缓存时必传 (Optional)
 * param domainGroupName: 域名组名称 (Optional)
 * param domains: 域名组内域名包含主域名 (Optional)
 */
func NewCreateDomainGroupRequestWithAllParams(
    shareCache *string,
    primaryDomain *string,
    domainGroupName *string,
    domains []string,
) *CreateDomainGroupRequest {

    return &CreateDomainGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup:create",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ShareCache: shareCache,
        PrimaryDomain: primaryDomain,
        DomainGroupName: domainGroupName,
        Domains: domains,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDomainGroupRequestWithoutParam() *CreateDomainGroupRequest {

    return &CreateDomainGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup:create",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param shareCache: 是否共享内存，共享缓存仅对中国境内加速域名生效(Optional) */
func (r *CreateDomainGroupRequest) SetShareCache(shareCache string) {
    r.ShareCache = &shareCache
}

/* param primaryDomain: 主域名,开启共享缓存时必传(Optional) */
func (r *CreateDomainGroupRequest) SetPrimaryDomain(primaryDomain string) {
    r.PrimaryDomain = &primaryDomain
}

/* param domainGroupName: 域名组名称(Optional) */
func (r *CreateDomainGroupRequest) SetDomainGroupName(domainGroupName string) {
    r.DomainGroupName = &domainGroupName
}

/* param domains: 域名组内域名包含主域名(Optional) */
func (r *CreateDomainGroupRequest) SetDomains(domains []string) {
    r.Domains = domains
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDomainGroupRequest) GetRegionId() string {
    return ""
}

type CreateDomainGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDomainGroupResult `json:"result"`
}

type CreateDomainGroupResult struct {
}