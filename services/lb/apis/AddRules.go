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
    lb "github.com/lshuining/jdcloud-sdk-go/services/lb/models"
)

type AddRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 转发规则组Id  */
    UrlMapId string `json:"urlMapId"`

    /* 转发规则的信息  */
    RuleSpecs []lb.RuleSpec `json:"ruleSpecs"`
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组Id (Required)
 * param ruleSpecs: 转发规则的信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddRulesRequest(
    regionId string,
    urlMapId string,
    ruleSpecs []lb.RuleSpec,
) *AddRulesRequest {

	return &AddRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/urlMaps/{urlMapId}:addRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UrlMapId: urlMapId,
        RuleSpecs: ruleSpecs,
	}
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组Id (Required)
 * param ruleSpecs: 转发规则的信息 (Required)
 */
func NewAddRulesRequestWithAllParams(
    regionId string,
    urlMapId string,
    ruleSpecs []lb.RuleSpec,
) *AddRulesRequest {

    return &AddRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}:addRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UrlMapId: urlMapId,
        RuleSpecs: ruleSpecs,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddRulesRequestWithoutParam() *AddRulesRequest {

    return &AddRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}:addRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param urlMapId: 转发规则组Id(Required) */
func (r *AddRulesRequest) SetUrlMapId(urlMapId string) {
    r.UrlMapId = urlMapId
}

/* param ruleSpecs: 转发规则的信息(Required) */
func (r *AddRulesRequest) SetRuleSpecs(ruleSpecs []lb.RuleSpec) {
    r.RuleSpecs = ruleSpecs
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddRulesRequest) GetRegionId() string {
    return r.RegionId
}

type AddRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddRulesResult `json:"result"`
}

type AddRulesResult struct {
}