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
    domainservice "github.com/lshuining/jdcloud-sdk-go/services/domainservice/models"
)

type CreateResourceRecordRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* RR参数  */
    Req *domainservice.AddRR `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param req: RR参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateResourceRecordRequest(
    regionId string,
    domainId string,
    req *domainservice.AddRR,
) *CreateResourceRecordRequest {

	return &CreateResourceRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param req: RR参数 (Required)
 */
func NewCreateResourceRecordRequestWithAllParams(
    regionId string,
    domainId string,
    req *domainservice.AddRR,
) *CreateResourceRecordRequest {

    return &CreateResourceRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateResourceRecordRequestWithoutParam() *CreateResourceRecordRequest {

    return &CreateResourceRecordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *CreateResourceRecordRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *CreateResourceRecordRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param req: RR参数(Required) */
func (r *CreateResourceRecordRequest) SetReq(req *domainservice.AddRR) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateResourceRecordRequest) GetRegionId() string {
    return r.RegionId
}

type CreateResourceRecordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateResourceRecordResult `json:"result"`
}

type CreateResourceRecordResult struct {
    DataList domainservice.RR `json:"dataList"`
}