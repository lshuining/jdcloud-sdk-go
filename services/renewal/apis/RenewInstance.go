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
    renewal "github.com/lshuining/jdcloud-sdk-go/services/renewal/models"
)

type RenewInstanceRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /*   */
    RenewInstanceParam *renewal.RenewInstanceParam `json:"renewInstanceParam"`
}

/*
 * param regionId: 地域 (Required)
 * param renewInstanceParam:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRenewInstanceRequest(
    regionId string,
    renewInstanceParam *renewal.RenewInstanceParam,
) *RenewInstanceRequest {

	return &RenewInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:renew",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        RenewInstanceParam: renewInstanceParam,
	}
}

/*
 * param regionId: 地域 (Required)
 * param renewInstanceParam:  (Required)
 */
func NewRenewInstanceRequestWithAllParams(
    regionId string,
    renewInstanceParam *renewal.RenewInstanceParam,
) *RenewInstanceRequest {

    return &RenewInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:renew",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        RenewInstanceParam: renewInstanceParam,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRenewInstanceRequestWithoutParam() *RenewInstanceRequest {

    return &RenewInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:renew",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *RenewInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param renewInstanceParam: (Required) */
func (r *RenewInstanceRequest) SetRenewInstanceParam(renewInstanceParam *renewal.RenewInstanceParam) {
    r.RenewInstanceParam = renewInstanceParam
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RenewInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type RenewInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RenewInstanceResult `json:"result"`
}

type RenewInstanceResult struct {
    OrderNumber string `json:"orderNumber"`
}