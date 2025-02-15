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

type DescribeAttackTypeCountRequest struct {

    core.JDCloudRequest

    /* 开始时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    EndTime string `json:"endTime"`

    /* 基础防护已防护的公网 IP, ip 不为空时, 查询 ip 对应的各类型攻击次数, ip 为空时, 查询用户所有公网 IP 的各类型攻击次数. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网IP 和 弹性公网 IP. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP (Optional) */
    Ip []string `json:"ip"`
}

/*
 * param startTime: 开始时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAttackTypeCountRequest(
    startTime string,
    endTime string,
) *DescribeAttackTypeCountRequest {

	return &DescribeAttackTypeCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeAttackTypeCount",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param startTime: 开始时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param ip: 基础防护已防护的公网 IP, ip 不为空时, 查询 ip 对应的各类型攻击次数, ip 为空时, 查询用户所有公网 IP 的各类型攻击次数. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网IP 和 弹性公网 IP. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP (Optional)
 */
func NewDescribeAttackTypeCountRequestWithAllParams(
    startTime string,
    endTime string,
    ip []string,
) *DescribeAttackTypeCountRequest {

    return &DescribeAttackTypeCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeAttackTypeCount",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAttackTypeCountRequestWithoutParam() *DescribeAttackTypeCountRequest {

    return &DescribeAttackTypeCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeAttackTypeCount",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 开始时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeAttackTypeCountRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeAttackTypeCountRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param ip: 基础防护已防护的公网 IP, ip 不为空时, 查询 ip 对应的各类型攻击次数, ip 为空时, 查询用户所有公网 IP 的各类型攻击次数. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a> 接口查询基础防护已防护的私有网络弹性公网 IP. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口查询基础防护已防护的云物理服务器公网IP 和 弹性公网 IP. <br>- 使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeccsipresources'>describeCcsIpResources</a> 接口查询基础防护已防护的托管区公网 IP(Optional) */
func (r *DescribeAttackTypeCountRequest) SetIp(ip []string) {
    r.Ip = ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAttackTypeCountRequest) GetRegionId() string {
    return ""
}

type DescribeAttackTypeCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAttackTypeCountResult `json:"result"`
}

type DescribeAttackTypeCountResult struct {
    DataList []baseanti.AttackTypeCount `json:"dataList"`
}