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

type AssignIpv6CidrRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 子网IPv6 CIDR (Optional) */
    Ipv6Cidr *string `json:"ipv6Cidr"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param subnetId: 子网ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAssignIpv6CidrRequest(
    regionId string,
    subnetId string,
) *AssignIpv6CidrRequest {

	return &AssignIpv6CidrRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets/{subnetId}:assignIpv6Cidr",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param subnetId: 子网ID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param ipv6Cidr: 子网IPv6 CIDR (Optional)
 */
func NewAssignIpv6CidrRequestWithAllParams(
    regionId string,
    subnetId string,
    clientToken *string,
    ipv6Cidr *string,
) *AssignIpv6CidrRequest {

    return &AssignIpv6CidrRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}:assignIpv6Cidr",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SubnetId: subnetId,
        ClientToken: clientToken,
        Ipv6Cidr: ipv6Cidr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAssignIpv6CidrRequestWithoutParam() *AssignIpv6CidrRequest {

    return &AssignIpv6CidrRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}:assignIpv6Cidr",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *AssignIpv6CidrRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param subnetId: 子网ID(Required) */
func (r *AssignIpv6CidrRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *AssignIpv6CidrRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param ipv6Cidr: 子网IPv6 CIDR(Optional) */
func (r *AssignIpv6CidrRequest) SetIpv6Cidr(ipv6Cidr string) {
    r.Ipv6Cidr = &ipv6Cidr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AssignIpv6CidrRequest) GetRegionId() string {
    return r.RegionId
}

type AssignIpv6CidrResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AssignIpv6CidrResult `json:"result"`
}

type AssignIpv6CidrResult struct {
    Success bool `json:"success"`
}