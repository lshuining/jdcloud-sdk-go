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
    dbaudit "github.com/lshuining/jdcloud-sdk-go/services/dbaudit/models"
)

type AddRuleRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 审计实例ID  */
    InsId string `json:"insId"`

    /* 规则组ID  */
    RuleGroupId string `json:"ruleGroupId"`

    /* 数据库ID (Optional) */
    DbId *string `json:"dbId"`

    /* 规则详情  */
    Rule *dbaudit.Rule `json:"rule"`
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param ruleGroupId: 规则组ID (Required)
 * param rule: 规则详情 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddRuleRequest(
    regionId string,
    insId string,
    ruleGroupId string,
    rule *dbaudit.Rule,
) *AddRuleRequest {

	return &AddRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InsId: insId,
        RuleGroupId: ruleGroupId,
        Rule: rule,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param ruleGroupId: 规则组ID (Required)
 * param dbId: 数据库ID (Optional)
 * param rule: 规则详情 (Required)
 */
func NewAddRuleRequestWithAllParams(
    regionId string,
    insId string,
    ruleGroupId string,
    dbId *string,
    rule *dbaudit.Rule,
) *AddRuleRequest {

    return &AddRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InsId: insId,
        RuleGroupId: ruleGroupId,
        DbId: dbId,
        Rule: rule,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddRuleRequestWithoutParam() *AddRuleRequest {

    return &AddRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *AddRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param insId: 审计实例ID(Required) */
func (r *AddRuleRequest) SetInsId(insId string) {
    r.InsId = insId
}

/* param ruleGroupId: 规则组ID(Required) */
func (r *AddRuleRequest) SetRuleGroupId(ruleGroupId string) {
    r.RuleGroupId = ruleGroupId
}

/* param dbId: 数据库ID(Optional) */
func (r *AddRuleRequest) SetDbId(dbId string) {
    r.DbId = &dbId
}

/* param rule: 规则详情(Required) */
func (r *AddRuleRequest) SetRule(rule *dbaudit.Rule) {
    r.Rule = rule
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddRuleRequest) GetRegionId() string {
    return r.RegionId
}

type AddRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddRuleResult `json:"result"`
}

type AddRuleResult struct {
}