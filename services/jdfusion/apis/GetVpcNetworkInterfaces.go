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

type GetVpcNetworkInterfacesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云主机id (Optional) */
    VmId *string `json:"vmId"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetVpcNetworkInterfacesRequest(
    regionId string,
) *GetVpcNetworkInterfacesRequest {

	return &GetVpcNetworkInterfacesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpc_networkInterfaces",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param vmId: 云主机id (Optional)
 */
func NewGetVpcNetworkInterfacesRequestWithAllParams(
    regionId string,
    vmId *string,
) *GetVpcNetworkInterfacesRequest {

    return &GetVpcNetworkInterfacesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_networkInterfaces",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VmId: vmId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetVpcNetworkInterfacesRequestWithoutParam() *GetVpcNetworkInterfacesRequest {

    return &GetVpcNetworkInterfacesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_networkInterfaces",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetVpcNetworkInterfacesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vmId: 云主机id(Optional) */
func (r *GetVpcNetworkInterfacesRequest) SetVmId(vmId string) {
    r.VmId = &vmId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetVpcNetworkInterfacesRequest) GetRegionId() string {
    return r.RegionId
}

type GetVpcNetworkInterfacesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetVpcNetworkInterfacesResult `json:"result"`
}

type GetVpcNetworkInterfacesResult struct {
    NetInterfaces []jdfusion.NetInterfaceInfo `json:"netInterfaces"`
}