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

type EnableWhiteListRuleOfWebRuleRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId string `json:"webRuleId"`

    /* 网站类规则的白名单规则 Id  */
    WebWhiteListRuleId string `json:"webWhiteListRuleId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param webWhiteListRuleId: 网站类规则的白名单规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableWhiteListRuleOfWebRuleRequest(
    regionId string,
    instanceId string,
    webRuleId string,
    webWhiteListRuleId string,
) *EnableWhiteListRuleOfWebRuleRequest {

	return &EnableWhiteListRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webWhiteListRules/{webWhiteListRuleId}:enable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        WebWhiteListRuleId: webWhiteListRuleId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param webWhiteListRuleId: 网站类规则的白名单规则 Id (Required)
 */
func NewEnableWhiteListRuleOfWebRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    webRuleId string,
    webWhiteListRuleId string,
) *EnableWhiteListRuleOfWebRuleRequest {

    return &EnableWhiteListRuleOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webWhiteListRules/{webWhiteListRuleId}:enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        WebWhiteListRuleId: webWhiteListRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableWhiteListRuleOfWebRuleRequestWithoutParam() *EnableWhiteListRuleOfWebRuleRequest {

    return &EnableWhiteListRuleOfWebRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/webWhiteListRules/{webWhiteListRuleId}:enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *EnableWhiteListRuleOfWebRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *EnableWhiteListRuleOfWebRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *EnableWhiteListRuleOfWebRuleRequest) SetWebRuleId(webRuleId string) {
    r.WebRuleId = webRuleId
}

/* param webWhiteListRuleId: 网站类规则的白名单规则 Id(Required) */
func (r *EnableWhiteListRuleOfWebRuleRequest) SetWebWhiteListRuleId(webWhiteListRuleId string) {
    r.WebWhiteListRuleId = webWhiteListRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableWhiteListRuleOfWebRuleRequest) GetRegionId() string {
    return r.RegionId
}

type EnableWhiteListRuleOfWebRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableWhiteListRuleOfWebRuleResult `json:"result"`
}

type EnableWhiteListRuleOfWebRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}