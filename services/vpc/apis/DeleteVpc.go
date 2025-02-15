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
)

type DeleteVpcRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Vpc ID  */
    VpcId string `json:"vpcId"`
}

/*
 * param regionId: Region ID (Required)
 * param vpcId: Vpc ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteVpcRequest(
    regionId string,
    vpcId string,
) *DeleteVpcRequest {

	return &DeleteVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs/{vpcId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcId: vpcId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param vpcId: Vpc ID (Required)
 */
func NewDeleteVpcRequestWithAllParams(
    regionId string,
    vpcId string,
) *DeleteVpcRequest {

    return &DeleteVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcId: vpcId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteVpcRequestWithoutParam() *DeleteVpcRequest {

    return &DeleteVpcRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteVpcRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcId: Vpc ID(Required) */
func (r *DeleteVpcRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteVpcRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteVpcResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteVpcResult `json:"result"`
}

type DeleteVpcResult struct {
}