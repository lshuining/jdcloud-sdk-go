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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type CreateForwardRuleRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 添加非网站类规则请求参数  */
    ForwardRuleSpec *ipanti.ForwardRuleSpec `json:"forwardRuleSpec"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleSpec: 添加非网站类规则请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateForwardRuleRequest(
    regionId string,
    instanceId string,
    forwardRuleSpec *ipanti.ForwardRuleSpec,
) *CreateForwardRuleRequest {

	return &CreateForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleSpec: forwardRuleSpec,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleSpec: 添加非网站类规则请求参数 (Required)
 */
func NewCreateForwardRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    forwardRuleSpec *ipanti.ForwardRuleSpec,
) *CreateForwardRuleRequest {

    return &CreateForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleSpec: forwardRuleSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateForwardRuleRequestWithoutParam() *CreateForwardRuleRequest {

    return &CreateForwardRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *CreateForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *CreateForwardRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param forwardRuleSpec: 添加非网站类规则请求参数(Required) */
func (r *CreateForwardRuleRequest) SetForwardRuleSpec(forwardRuleSpec *ipanti.ForwardRuleSpec) {
    r.ForwardRuleSpec = forwardRuleSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type CreateForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateForwardRuleResult `json:"result"`
}

type CreateForwardRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}