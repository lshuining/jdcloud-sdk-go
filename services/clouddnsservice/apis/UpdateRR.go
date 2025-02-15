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
    clouddnsservice "github.com/lshuining/jdcloud-sdk-go/services/clouddnsservice/models"
)

type UpdateRRRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用getDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* updateRR参数  */
    Req *clouddnsservice.UpdateRR `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param req: updateRR参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateRRRequest(
    regionId string,
    domainId string,
    req *clouddnsservice.UpdateRR,
) *UpdateRRRequest {

	return &UpdateRRRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/RRUpdate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param req: updateRR参数 (Required)
 */
func NewUpdateRRRequestWithAllParams(
    regionId string,
    domainId string,
    req *clouddnsservice.UpdateRR,
) *UpdateRRRequest {

    return &UpdateRRRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/RRUpdate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateRRRequestWithoutParam() *UpdateRRRequest {

    return &UpdateRRRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/RRUpdate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *UpdateRRRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用getDomains接口获取。(Required) */
func (r *UpdateRRRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param req: updateRR参数(Required) */
func (r *UpdateRRRequest) SetReq(req *clouddnsservice.UpdateRR) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateRRRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateRRResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateRRResult `json:"result"`
}

type UpdateRRResult struct {
}