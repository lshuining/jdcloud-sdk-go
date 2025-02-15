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

package models

import charge "github.com/lshuining/jdcloud-sdk-go/services/charge/models"

type Ipv6Address struct {

    /* 地域代码, 如cn-north-1 (Optional) */
    Region string `json:"region"`

    /* 公网IPv6地址ID (Optional) */
    Ipv6AddressId string `json:"ipv6AddressId"`

    /* IPv6地址 (Optional) */
    Ipv6Address string `json:"ipv6Address"`

    /* IPv6网关ID (Optional) */
    Ipv6GatewayId string `json:"ipv6GatewayId"`

    /* 带宽, 单位Mbps (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 私有网络ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 私有网络名称 (Optional) */
    VpcName string `json:"vpcName"`

    /* 关联的实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 关联的实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 关联的实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`
}
