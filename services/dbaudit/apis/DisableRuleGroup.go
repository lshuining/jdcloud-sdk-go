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

type DisableRuleGroupRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 审计实例ID  */
    InsId string `json:"insId"`

    /* 规则组ID  */
    RuleGroupId string `json:"ruleGroupId"`

    /* 数据库ID (Optional) */
    DbId *string `json:"dbId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param ruleGroupId: 规则组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableRuleGroupRequest(
    regionId string,
    insId string,
    ruleGroupId string,
) *DisableRuleGroupRequest {

	return &DisableRuleGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}:disable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InsId: insId,
        RuleGroupId: ruleGroupId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param ruleGroupId: 规则组ID (Required)
 * param dbId: 数据库ID (Optional)
 */
func NewDisableRuleGroupRequestWithAllParams(
    regionId string,
    insId string,
    ruleGroupId string,
    dbId *string,
) *DisableRuleGroupRequest {

    return &DisableRuleGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InsId: insId,
        RuleGroupId: ruleGroupId,
        DbId: dbId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableRuleGroupRequestWithoutParam() *DisableRuleGroupRequest {

    return &DisableRuleGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DisableRuleGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param insId: 审计实例ID(Required) */
func (r *DisableRuleGroupRequest) SetInsId(insId string) {
    r.InsId = insId
}

/* param ruleGroupId: 规则组ID(Required) */
func (r *DisableRuleGroupRequest) SetRuleGroupId(ruleGroupId string) {
    r.RuleGroupId = ruleGroupId
}

/* param dbId: 数据库ID(Optional) */
func (r *DisableRuleGroupRequest) SetDbId(dbId string) {
    r.DbId = &dbId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableRuleGroupRequest) GetRegionId() string {
    return r.RegionId
}

type DisableRuleGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableRuleGroupResult `json:"result"`
}

type DisableRuleGroupResult struct {
}