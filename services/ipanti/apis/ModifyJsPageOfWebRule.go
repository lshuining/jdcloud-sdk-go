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
    ipanti "github.com/lshuining/jdcloud-sdk-go/services/ipanti/models"
)

type ModifyJsPageOfWebRuleRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId string `json:"webRuleId"`

    /* 支持插入JS指纹的页面 Id  */
    JsPageId string `json:"jsPageId"`

    /* 修改网站类规则允许插入 JS 指纹的页面请求参数  */
    JsPageSpec *ipanti.JsPageSpec `json:"jsPageSpec"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param jsPageId: 支持插入JS指纹的页面 Id (Required)
 * param jsPageSpec: 修改网站类规则允许插入 JS 指纹的页面请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyJsPageOfWebRuleRequest(
    regionId string,
    instanceId string,
    webRuleId string,
    jsPageId string,
    jsPageSpec *ipanti.JsPageSpec,
) *ModifyJsPageOfWebRuleRequest {

	return &ModifyJsPageOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/jsPages/{jsPageId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        JsPageId: jsPageId,
        JsPageSpec: jsPageSpec,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param jsPageId: 支持插入JS指纹的页面 Id (Required)
 * param jsPageSpec: 修改网站类规则允许插入 JS 指纹的页面请求参数 (Required)
 */
func NewModifyJsPageOfWebRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    webRuleId string,
    jsPageId string,
    jsPageSpec *ipanti.JsPageSpec,
) *ModifyJsPageOfWebRuleRequest {

    return &ModifyJsPageOfWebRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/jsPages/{jsPageId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        JsPageId: jsPageId,
        JsPageSpec: jsPageSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyJsPageOfWebRuleRequestWithoutParam() *ModifyJsPageOfWebRuleRequest {

    return &ModifyJsPageOfWebRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}/jsPages/{jsPageId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *ModifyJsPageOfWebRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *ModifyJsPageOfWebRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *ModifyJsPageOfWebRuleRequest) SetWebRuleId(webRuleId string) {
    r.WebRuleId = webRuleId
}

/* param jsPageId: 支持插入JS指纹的页面 Id(Required) */
func (r *ModifyJsPageOfWebRuleRequest) SetJsPageId(jsPageId string) {
    r.JsPageId = jsPageId
}

/* param jsPageSpec: 修改网站类规则允许插入 JS 指纹的页面请求参数(Required) */
func (r *ModifyJsPageOfWebRuleRequest) SetJsPageSpec(jsPageSpec *ipanti.JsPageSpec) {
    r.JsPageSpec = jsPageSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyJsPageOfWebRuleRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyJsPageOfWebRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyJsPageOfWebRuleResult `json:"result"`
}

type ModifyJsPageOfWebRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}