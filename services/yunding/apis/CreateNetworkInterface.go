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

type CreateNetworkInterfaceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 可用区，用户的默认可用区，该参数无效，不建议使用 (Optional) */
    Az *string `json:"az"`

    /* 网卡名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    NetworkInterfaceName *string `json:"networkInterfaceName"`

    /* 网卡主IP，如果不指定，会自动从子网中分配 (Optional) */
    PrimaryIpAddress *string `json:"primaryIpAddress"`

    /* SecondaryIp列表 (Optional) */
    SecondaryIpAddresses []string `json:"secondaryIpAddresses"`

    /* 自动分配的SecondaryIp数量 (Optional) */
    SecondaryIpCount *int `json:"secondaryIpCount"`

    /* 要绑定的安全组ID列表，最多指定5个安全组 (Optional) */
    SecurityGroups []string `json:"securityGroups"`

    /* 源和目标IP地址校验，取值为0或者1,默认为1 (Optional) */
    SanityCheck *int `json:"sanityCheck"`

    /* 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param subnetId: 子网ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateNetworkInterfaceRequest(
    regionId string,
    subnetId string,
) *CreateNetworkInterfaceRequest {

	return &CreateNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydNetworkInterfaces",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param subnetId: 子网ID (Required)
 * param az: 可用区，用户的默认可用区，该参数无效，不建议使用 (Optional)
 * param networkInterfaceName: 网卡名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional)
 * param primaryIpAddress: 网卡主IP，如果不指定，会自动从子网中分配 (Optional)
 * param secondaryIpAddresses: SecondaryIp列表 (Optional)
 * param secondaryIpCount: 自动分配的SecondaryIp数量 (Optional)
 * param securityGroups: 要绑定的安全组ID列表，最多指定5个安全组 (Optional)
 * param sanityCheck: 源和目标IP地址校验，取值为0或者1,默认为1 (Optional)
 * param description: 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 */
func NewCreateNetworkInterfaceRequestWithAllParams(
    regionId string,
    subnetId string,
    az *string,
    networkInterfaceName *string,
    primaryIpAddress *string,
    secondaryIpAddresses []string,
    secondaryIpCount *int,
    securityGroups []string,
    sanityCheck *int,
    description *string,
) *CreateNetworkInterfaceRequest {

    return &CreateNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydNetworkInterfaces",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        SubnetId: subnetId,
        Az: az,
        NetworkInterfaceName: networkInterfaceName,
        PrimaryIpAddress: primaryIpAddress,
        SecondaryIpAddresses: secondaryIpAddresses,
        SecondaryIpCount: secondaryIpCount,
        SecurityGroups: securityGroups,
        SanityCheck: sanityCheck,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateNetworkInterfaceRequestWithoutParam() *CreateNetworkInterfaceRequest {

    return &CreateNetworkInterfaceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydNetworkInterfaces",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateNetworkInterfaceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param subnetId: 子网ID(Required) */
func (r *CreateNetworkInterfaceRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

/* param az: 可用区，用户的默认可用区，该参数无效，不建议使用(Optional) */
func (r *CreateNetworkInterfaceRequest) SetAz(az string) {
    r.Az = &az
}

/* param networkInterfaceName: 网卡名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。(Optional) */
func (r *CreateNetworkInterfaceRequest) SetNetworkInterfaceName(networkInterfaceName string) {
    r.NetworkInterfaceName = &networkInterfaceName
}

/* param primaryIpAddress: 网卡主IP，如果不指定，会自动从子网中分配(Optional) */
func (r *CreateNetworkInterfaceRequest) SetPrimaryIpAddress(primaryIpAddress string) {
    r.PrimaryIpAddress = &primaryIpAddress
}

/* param secondaryIpAddresses: SecondaryIp列表(Optional) */
func (r *CreateNetworkInterfaceRequest) SetSecondaryIpAddresses(secondaryIpAddresses []string) {
    r.SecondaryIpAddresses = secondaryIpAddresses
}

/* param secondaryIpCount: 自动分配的SecondaryIp数量(Optional) */
func (r *CreateNetworkInterfaceRequest) SetSecondaryIpCount(secondaryIpCount int) {
    r.SecondaryIpCount = &secondaryIpCount
}

/* param securityGroups: 要绑定的安全组ID列表，最多指定5个安全组(Optional) */
func (r *CreateNetworkInterfaceRequest) SetSecurityGroups(securityGroups []string) {
    r.SecurityGroups = securityGroups
}

/* param sanityCheck: 源和目标IP地址校验，取值为0或者1,默认为1(Optional) */
func (r *CreateNetworkInterfaceRequest) SetSanityCheck(sanityCheck int) {
    r.SanityCheck = &sanityCheck
}

/* param description: 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateNetworkInterfaceRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateNetworkInterfaceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateNetworkInterfaceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateNetworkInterfaceResult `json:"result"`
}

type CreateNetworkInterfaceResult struct {
    NetworkInterfaceId string `json:"networkInterfaceId"`
}