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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type SetCacheRulesRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /*  (Optional) */
    CacheRules []cdn.CacheRuleVo `json:"cacheRules"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetCacheRulesRequest(
    domain string,
) *SetCacheRulesRequest {

	return &SetCacheRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/cacheRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param cacheRules:  (Optional)
 */
func NewSetCacheRulesRequestWithAllParams(
    domain string,
    cacheRules []cdn.CacheRuleVo,
) *SetCacheRulesRequest {

    return &SetCacheRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/cacheRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        CacheRules: cacheRules,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetCacheRulesRequestWithoutParam() *SetCacheRulesRequest {

    return &SetCacheRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/cacheRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetCacheRulesRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param cacheRules: (Optional) */
func (r *SetCacheRulesRequest) SetCacheRules(cacheRules []cdn.CacheRuleVo) {
    r.CacheRules = cacheRules
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetCacheRulesRequest) GetRegionId() string {
    return ""
}

type SetCacheRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetCacheRulesResult `json:"result"`
}

type SetCacheRulesResult struct {
    Data []cdn.CacheVo `json:"data"`
}