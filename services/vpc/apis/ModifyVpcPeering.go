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

type ModifyVpcPeeringRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* vpcPeeringId ID  */
    VpcPeeringId string `json:"vpcPeeringId"`

    /* VpcPeering的名字,不为空。名称取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Optional) */
    VpcPeeringName *string `json:"vpcPeeringName"`

    /* VpcPeering 描述，取值范围：0-256个中文、英文大小写的字母、数字和下划线分隔符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param vpcPeeringId: vpcPeeringId ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyVpcPeeringRequest(
    regionId string,
    vpcPeeringId string,
) *ModifyVpcPeeringRequest {

	return &ModifyVpcPeeringRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcPeerings/{vpcPeeringId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcPeeringId: vpcPeeringId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param vpcPeeringId: vpcPeeringId ID (Required)
 * param vpcPeeringName: VpcPeering的名字,不为空。名称取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Optional)
 * param description: VpcPeering 描述，取值范围：0-256个中文、英文大小写的字母、数字和下划线分隔符 (Optional)
 */
func NewModifyVpcPeeringRequestWithAllParams(
    regionId string,
    vpcPeeringId string,
    vpcPeeringName *string,
    description *string,
) *ModifyVpcPeeringRequest {

    return &ModifyVpcPeeringRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcPeerings/{vpcPeeringId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcPeeringId: vpcPeeringId,
        VpcPeeringName: vpcPeeringName,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyVpcPeeringRequestWithoutParam() *ModifyVpcPeeringRequest {

    return &ModifyVpcPeeringRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcPeerings/{vpcPeeringId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyVpcPeeringRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcPeeringId: vpcPeeringId ID(Required) */
func (r *ModifyVpcPeeringRequest) SetVpcPeeringId(vpcPeeringId string) {
    r.VpcPeeringId = vpcPeeringId
}

/* param vpcPeeringName: VpcPeering的名字,不为空。名称取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符(Optional) */
func (r *ModifyVpcPeeringRequest) SetVpcPeeringName(vpcPeeringName string) {
    r.VpcPeeringName = &vpcPeeringName
}

/* param description: VpcPeering 描述，取值范围：0-256个中文、英文大小写的字母、数字和下划线分隔符(Optional) */
func (r *ModifyVpcPeeringRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyVpcPeeringRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyVpcPeeringResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyVpcPeeringResult `json:"result"`
}

type ModifyVpcPeeringResult struct {
}