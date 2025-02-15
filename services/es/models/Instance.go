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

type Instance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例版本，目前支持5.6.9和6.5.4两个版本 (Optional) */
    InstanceVersion string `json:"instanceVersion"`

    /* 实例状态，running：运行，error：错误，creating：创建中，changing：变配中，stop：已停止，processing：处理中 (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例的配置信息 (Optional) */
    InstanceClass InstanceClass `json:"instanceClass"`

    /* 创建时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    CreateTime string `json:"createTime"`

    /* AZ信息，各AZ编码详见：https://docs.jdcloud.com/cn/jcs-for-elasticsearch/restrictions (Optional) */
    AzId string `json:"azId"`

    /* 所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 内网地址 (Optional) */
    InternalEndpoint InternalEndpoint `json:"internalEndpoint"`

    /* deprecated，见internalEndpoint (Optional) */
    Endpoint string `json:"endpoint"`

    /* kibana地址 (Optional) */
    KibanaUrl string `json:"kibanaUrl"`

    /* head地址 (Optional) */
    HeadUrl string `json:"headUrl"`

    /* 值为v4&v6时支持ipv6和ipv4, 值为空时仅支持ipv4 (Optional) */
    IpVersion string `json:"ipVersion"`

    /* Tag信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* 是否开启了专用主节点，true为开启，false为不开启 (Optional) */
    DedicatedMaster bool `json:"dedicatedMaster"`

    /* 是否开启了协调节点，true为开启，false为不开启 (Optional) */
    Coordinating bool `json:"coordinating"`

    /* kibana floatIp地址 (Optional) */
    KibanaFiUrl string `json:"kibanaFiUrl"`
}
