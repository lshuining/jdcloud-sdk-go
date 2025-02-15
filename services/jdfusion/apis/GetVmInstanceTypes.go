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
    jdfusion "github.com/lshuining/jdcloud-sdk-go/services/jdfusion/models"
)

type GetVmInstanceTypesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 可用区 (Optional) */
    Az *string `json:"az"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetVmInstanceTypesRequest(
    regionId string,
) *GetVmInstanceTypesRequest {

	return &GetVmInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vm_instanceTypes",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param az: 可用区 (Optional)
 */
func NewGetVmInstanceTypesRequestWithAllParams(
    regionId string,
    az *string,
) *GetVmInstanceTypesRequest {

    return &GetVmInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_instanceTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Az: az,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetVmInstanceTypesRequestWithoutParam() *GetVmInstanceTypesRequest {

    return &GetVmInstanceTypesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_instanceTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetVmInstanceTypesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param az: 可用区(Optional) */
func (r *GetVmInstanceTypesRequest) SetAz(az string) {
    r.Az = &az
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetVmInstanceTypesRequest) GetRegionId() string {
    return r.RegionId
}

type GetVmInstanceTypesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetVmInstanceTypesResult `json:"result"`
}

type GetVmInstanceTypesResult struct {
    InstanceTypes []jdfusion.InstanceTypeInfo `json:"instanceTypes"`
}