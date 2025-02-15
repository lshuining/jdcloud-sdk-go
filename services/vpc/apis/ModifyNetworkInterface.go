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

type ModifyNetworkInterfaceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkInterface ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 弹性网卡名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    NetworkInterfaceName *string `json:"networkInterfaceName"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 以覆盖原有安全组的方式更新的安全组。如果更新安全组ID列表，最多5个安全组 (Optional) */
    SecurityGroups []string `json:"securityGroups"`
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyNetworkInterfaceRequest(
    regionId string,
    networkInterfaceId string,
) *ModifyNetworkInterfaceRequest {

	return &ModifyNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 * param networkInterfaceName: 弹性网卡名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param securityGroups: 以覆盖原有安全组的方式更新的安全组。如果更新安全组ID列表，最多5个安全组 (Optional)
 */
func NewModifyNetworkInterfaceRequestWithAllParams(
    regionId string,
    networkInterfaceId string,
    networkInterfaceName *string,
    description *string,
    securityGroups []string,
) *ModifyNetworkInterfaceRequest {

    return &ModifyNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
        NetworkInterfaceName: networkInterfaceName,
        Description: description,
        SecurityGroups: securityGroups,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyNetworkInterfaceRequestWithoutParam() *ModifyNetworkInterfaceRequest {

    return &ModifyNetworkInterfaceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyNetworkInterfaceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkInterfaceId: networkInterface ID(Required) */
func (r *ModifyNetworkInterfaceRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}

/* param networkInterfaceName: 弹性网卡名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Optional) */
func (r *ModifyNetworkInterfaceRequest) SetNetworkInterfaceName(networkInterfaceName string) {
    r.NetworkInterfaceName = &networkInterfaceName
}

/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *ModifyNetworkInterfaceRequest) SetDescription(description string) {
    r.Description = &description
}

/* param securityGroups: 以覆盖原有安全组的方式更新的安全组。如果更新安全组ID列表，最多5个安全组(Optional) */
func (r *ModifyNetworkInterfaceRequest) SetSecurityGroups(securityGroups []string) {
    r.SecurityGroups = securityGroups
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyNetworkInterfaceRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyNetworkInterfaceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyNetworkInterfaceResult `json:"result"`
}

type ModifyNetworkInterfaceResult struct {
}