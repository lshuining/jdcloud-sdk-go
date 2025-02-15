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
    baseanti "github.com/lshuining/jdcloud-sdk-go/services/baseanti/models"
)

type DescribeIpResourceInfoRequest struct {

    core.JDCloudRequest

    /* 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州  */
    RegionId string `json:"regionId"`

    /* 基础防护已防护的公网 IP, 仅支持 ipv4 格式. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网 IP 和 弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP  */
    Ip string `json:"ip"`
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州 (Required)
 * param ip: 基础防护已防护的公网 IP, 仅支持 ipv4 格式. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网 IP 和 弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpResourceInfoRequest(
    regionId string,
    ip string,
) *DescribeIpResourceInfoRequest {

	return &DescribeIpResourceInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipResources/{ip}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ip: ip,
	}
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州 (Required)
 * param ip: 基础防护已防护的公网 IP, 仅支持 ipv4 格式. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网 IP 和 弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP (Required)
 */
func NewDescribeIpResourceInfoRequestWithAllParams(
    regionId string,
    ip string,
) *DescribeIpResourceInfoRequest {

    return &DescribeIpResourceInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpResourceInfoRequestWithoutParam() *DescribeIpResourceInfoRequest {

    return &DescribeIpResourceInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州(Required) */
func (r *DescribeIpResourceInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ip: 基础防护已防护的公网 IP, 仅支持 ipv4 格式. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网 IP 和 弹性公网 IP<br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP(Required) */
func (r *DescribeIpResourceInfoRequest) SetIp(ip string) {
    r.Ip = ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpResourceInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpResourceInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpResourceInfoResult `json:"result"`
}

type DescribeIpResourceInfoResult struct {
    Data baseanti.IpResourceInfo `json:"data"`
}