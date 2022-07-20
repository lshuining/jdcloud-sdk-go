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
    waf "github.com/lshuining/jdcloud-sdk-go/services/waf/models"
)

type DelWafConditionRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 请求  */
    Req *waf.DelRulesReq `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDelWafConditionRequest(
    regionId string,
    wafInstanceId string,
    req *waf.DelRulesReq,
) *DelWafConditionRequest {

	return &DelWafConditionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/waf:delCondition",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 */
func NewDelWafConditionRequestWithAllParams(
    regionId string,
    wafInstanceId string,
    req *waf.DelRulesReq,
) *DelWafConditionRequest {

    return &DelWafConditionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/waf:delCondition",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDelWafConditionRequestWithoutParam() *DelWafConditionRequest {

    return &DelWafConditionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/waf:delCondition",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DelWafConditionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param wafInstanceId: 实例Id(Required) */
func (r *DelWafConditionRequest) SetWafInstanceId(wafInstanceId string) {
    r.WafInstanceId = wafInstanceId
}

/* param req: 请求(Required) */
func (r *DelWafConditionRequest) SetReq(req *waf.DelRulesReq) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DelWafConditionRequest) GetRegionId() string {
    return r.RegionId
}

type DelWafConditionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DelWafConditionResult `json:"result"`
}

type DelWafConditionResult struct {
}