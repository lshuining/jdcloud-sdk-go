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
)

type DisableBlackListRuleOfWebRuleRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId string `json:"webRuleId"`

    /* 网站类规则的黑名单规则 Id  */
    WebBlackListRuleId string `json:"webBlackListRuleId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param webBlackListRuleId: 网站类规则的黑名单规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableBlackListRuleOfWebRuleRequest(
    regionId string,
    instanceId string,
    webRuleId string,
    webBlackListRuleId string,
) *DisableBlackListRuleOfWebRuleRequest {

	return &DisableBlackListRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webBlackListRules/{webBlackListRuleId}:disable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        WebBlackListRuleId: webBlackListRuleId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param webBlackListRuleId: 网站类规则的黑名单规则 Id (Required)
 */
func NewDisableBlackListRuleOfWebRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    webRuleId string,
    webBlackListRuleId string,
) *DisableBlackListRuleOfWebRuleRequest {

    return &DisableBlackListRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webBlackListRules/{webBlackListRuleId}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        WebBlackListRuleId: webBlackListRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableBlackListRuleOfWebRuleRequestWithoutParam() *DisableBlackListRuleOfWebRuleRequest {

    return &DisableBlackListRuleOfWebRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webBlackListRules/{webBlackListRuleId}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DisableBlackListRuleOfWebRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DisableBlackListRuleOfWebRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *DisableBlackListRuleOfWebRuleRequest) SetWebRuleId(webRuleId string) {
    r.WebRuleId = webRuleId
}

/* param webBlackListRuleId: 网站类规则的黑名单规则 Id(Required) */
func (r *DisableBlackListRuleOfWebRuleRequest) SetWebBlackListRuleId(webBlackListRuleId string) {
    r.WebBlackListRuleId = webBlackListRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableBlackListRuleOfWebRuleRequest) GetRegionId() string {
    return r.RegionId
}

type DisableBlackListRuleOfWebRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableBlackListRuleOfWebRuleResult `json:"result"`
}

type DisableBlackListRuleOfWebRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}