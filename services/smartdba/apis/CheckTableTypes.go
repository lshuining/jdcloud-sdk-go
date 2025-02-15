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
    smartdba "github.com/lshuining/jdcloud-sdk-go/services/smartdba/models"
)

type CheckTableTypesRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceGid string `json:"instanceGid"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckTableTypesRequest(
    regionId string,
    instanceGid string,
) *CheckTableTypesRequest {

	return &CheckTableTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/checkTableTypes",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: 实例ID (Required)
 */
func NewCheckTableTypesRequestWithAllParams(
    regionId string,
    instanceGid string,
) *CheckTableTypesRequest {

    return &CheckTableTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/checkTableTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckTableTypesRequestWithoutParam() *CheckTableTypesRequest {

    return &CheckTableTypesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/checkTableTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *CheckTableTypesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: 实例ID(Required) */
func (r *CheckTableTypesRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckTableTypesRequest) GetRegionId() string {
    return r.RegionId
}

type CheckTableTypesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckTableTypesResult `json:"result"`
}

type CheckTableTypesResult struct {
    Data []smartdba.Table `json:"data"`
}