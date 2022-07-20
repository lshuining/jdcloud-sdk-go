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

type RemoveNetworkAclRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkAclId ID  */
    NetworkAclId string `json:"networkAclId"`

    /* networkAcl规则ID列表  */
    RuleIds []string `json:"ruleIds"`
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param ruleIds: networkAcl规则ID列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveNetworkAclRulesRequest(
    regionId string,
    networkAclId string,
    ruleIds []string,
) *RemoveNetworkAclRulesRequest {

	return &RemoveNetworkAclRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkAcls/{networkAclId}:removeNetworkAclRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkAclId: networkAclId,
        RuleIds: ruleIds,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param ruleIds: networkAcl规则ID列表 (Required)
 */
func NewRemoveNetworkAclRulesRequestWithAllParams(
    regionId string,
    networkAclId string,
    ruleIds []string,
) *RemoveNetworkAclRulesRequest {

    return &RemoveNetworkAclRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:removeNetworkAclRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkAclId: networkAclId,
        RuleIds: ruleIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveNetworkAclRulesRequestWithoutParam() *RemoveNetworkAclRulesRequest {

    return &RemoveNetworkAclRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:removeNetworkAclRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *RemoveNetworkAclRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkAclId: networkAclId ID(Required) */
func (r *RemoveNetworkAclRulesRequest) SetNetworkAclId(networkAclId string) {
    r.NetworkAclId = networkAclId
}

/* param ruleIds: networkAcl规则ID列表(Required) */
func (r *RemoveNetworkAclRulesRequest) SetRuleIds(ruleIds []string) {
    r.RuleIds = ruleIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveNetworkAclRulesRequest) GetRegionId() string {
    return r.RegionId
}

type RemoveNetworkAclRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveNetworkAclRulesResult `json:"result"`
}

type RemoveNetworkAclRulesResult struct {
}