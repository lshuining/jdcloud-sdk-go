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
import resourcetag "github.com/lshuining/jdcloud-sdk-go/services/resourcetag/models"

type InstanceSpec struct {

    /* 可用区, 如 cn-north-1  */
    Az string `json:"az"`

    /* 实例类型, 如 cps.c.normal  */
    DeviceType string `json:"deviceType"`

    /* 主机名 (Optional) */
    Hostname *string `json:"hostname"`

    /* 镜像类型, 取值范围：standard  */
    ImageType string `json:"imageType"`

    /* 操作系统类型ID  */
    OsTypeId string `json:"osTypeId"`

    /* 系统盘RAID类型ID  */
    SysRaidTypeId string `json:"sysRaidTypeId"`

    /* 数据盘RAID类型ID (Optional) */
    DataRaidTypeId *string `json:"dataRaidTypeId"`

    /* 子网编号 (Optional) */
    SubnetId *string `json:"subnetId"`

    /* 是否启用外网，取值范围：yes、no (Optional) */
    EnableInternet *string `json:"enableInternet"`

    /* 启用外网时弹性公网IP的计费模式，取值范围：prepaid_by_duration、postpaid_by_duration (Optional) */
    InternetChargeMode *string `json:"internetChargeMode"`

    /* 是否启用IPv6，取值范围：yes、no (Optional) */
    EnableIpv6 *string `json:"enableIpv6"`

    /* IPv6地址 (Optional) */
    Ipv6Address *string `json:"ipv6Address"`

    /* 网络类型，取值范围：basic（基础网络）、vpc（私有网络）、retail（零售中台网络）  */
    NetworkType string `json:"networkType"`

    /* 网络CIDR (Optional) */
    Cidr *string `json:"cidr"`

    /* 内网IP (Optional) */
    PrivateIp *string `json:"privateIp"`

    /* 内网添加的别名IP范围 (Optional) */
    AliasIps []AliasIpInfo `json:"aliasIps"`

    /* 外网链路类型, 目前只支持bgp (Optional) */
    LineType *string `json:"lineType"`

    /* 外网带宽, 范围[1,200] 单位Mbps (Optional) */
    Bandwidth *int `json:"bandwidth"`

    /* 云物理服务器名称  */
    Name string `json:"name"`

    /* 云物理服务器描述 (Optional) */
    Description *string `json:"description"`

    /* 密码，不传值会随机生成密码 (Optional) */
    Password *string `json:"password"`

    /* 购买数量  */
    Count int `json:"count"`

    /* 可执行脚本Base64编码后的内容，支持shell和python脚本 (Optional) */
    UserData *string `json:"userData"`

    /* 密钥对id (Optional) */
    KeypairId *string `json:"keypairId"`

    /* 计费配置  */
    Charge *charge.ChargeSpec `json:"charge"`

    /* 网络接口模式，取值：bond（网口bond）、dual（双网口） (Optional) */
    InterfaceMode *string `json:"interfaceMode"`

    /* 辅网口是否启用IPv6，取值范围：yes、no (Optional) */
    ExtensionEnableIpv6 *string `json:"extensionEnableIpv6"`

    /* 辅网口IPv6地址 (Optional) */
    ExtensionIpv6Address *string `json:"extensionIpv6Address"`

    /* 辅网口子网ID (Optional) */
    ExtensionSubnetId *string `json:"extensionSubnetId"`

    /* 辅网口手动分配的内网ip (Optional) */
    ExtensionPrivateIp *string `json:"extensionPrivateIp"`

    /* 辅网口内网添加的别名IP范围 (Optional) */
    ExtensionAliasIps []AliasIpInfo `json:"extensionAliasIps"`

    /* 辅网口是否启用外网，取值范围：yes、no (Optional) */
    ExtensionEnableInternet *string `json:"extensionEnableInternet"`

    /* 辅网口链路类型, 目前支持BGP (Optional) */
    ExtensionLineType *string `json:"extensionLineType"`

    /* 辅网口外网带宽，范围[1,200] 单位Mbps (Optional) */
    ExtensionBandwidth *int `json:"extensionBandwidth"`

    /* 辅网口启用外网时弹性公网IP的计费模式，取值范围：prepaid_by_duration、postpaid_by_duration (Optional) */
    ExtensionInternetChargeMode *string `json:"extensionInternetChargeMode"`

    /* 标签 (Optional) */
    ResourceTags []resourcetag.Tag `json:"resourceTags"`
}
