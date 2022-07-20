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

type CreateVpcEipRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分配弹性公网ip  */
    Allocate *jdfusion.AllocateEipAddress `json:"allocate"`
}

/*
 * param regionId: 地域ID (Required)
 * param allocate: 分配弹性公网ip (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVpcEipRequest(
    regionId string,
    allocate *jdfusion.AllocateEipAddress,
) *CreateVpcEipRequest {

	return &CreateVpcEipRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpc_eips",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Allocate: allocate,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param allocate: 分配弹性公网ip (Required)
 */
func NewCreateVpcEipRequestWithAllParams(
    regionId string,
    allocate *jdfusion.AllocateEipAddress,
) *CreateVpcEipRequest {

    return &CreateVpcEipRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_eips",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Allocate: allocate,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVpcEipRequestWithoutParam() *CreateVpcEipRequest {

    return &CreateVpcEipRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_eips",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateVpcEipRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param allocate: 分配弹性公网ip(Required) */
func (r *CreateVpcEipRequest) SetAllocate(allocate *jdfusion.AllocateEipAddress) {
    r.Allocate = allocate
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVpcEipRequest) GetRegionId() string {
    return r.RegionId
}

type CreateVpcEipResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVpcEipResult `json:"result"`
}

type CreateVpcEipResult struct {
    Task jdfusion.ResourceTFInfo `json:"task"`
}