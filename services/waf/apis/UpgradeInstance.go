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

type UpgradeInstanceRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 请求  */
    Req *waf.OrderReq `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpgradeInstanceRequest(
    regionId string,
    wafInstanceId string,
    req *waf.OrderReq,
) *UpgradeInstanceRequest {

	return &UpgradeInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/billing:upgradeInstance",
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
func NewUpgradeInstanceRequestWithAllParams(
    regionId string,
    wafInstanceId string,
    req *waf.OrderReq,
) *UpgradeInstanceRequest {

    return &UpgradeInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/billing:upgradeInstance",
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
func NewUpgradeInstanceRequestWithoutParam() *UpgradeInstanceRequest {

    return &UpgradeInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/billing:upgradeInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *UpgradeInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param wafInstanceId: 实例Id(Required) */
func (r *UpgradeInstanceRequest) SetWafInstanceId(wafInstanceId string) {
    r.WafInstanceId = wafInstanceId
}

/* param req: 请求(Required) */
func (r *UpgradeInstanceRequest) SetReq(req *waf.OrderReq) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpgradeInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type UpgradeInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpgradeInstanceResult `json:"result"`
}

type UpgradeInstanceResult struct {
    BuyId string `json:"buyId"`
    WafInstanceId string `json:"wafInstanceId"`
    OrderNumber string `json:"orderNumber"`
}