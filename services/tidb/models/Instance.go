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

    /* 实例引擎版本 (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 实例引擎版本的详细版本号 (Optional) */
    MinorVersion string `json:"minorVersion"`

    /* 集群中节点的总数 (Optional) */
    TotalNodeNum int `json:"totalNodeNum"`

    /* 整个集群总的CPU核数 (Optional) */
    TotalCPU int `json:"totalCPU"`

    /* 整个集群总的内存大小，单位GB (Optional) */
    TotalMemoryGB int `json:"totalMemoryGB"`

    /* 整个集群总的存储空间大小，单位GB (Optional) */
    TotalStorageGB int `json:"totalStorageGB"`

    /* 地域ID (Optional) */
    RegionId string `json:"regionId"`

    /* 可用区ID，目前仅支持单可用区 (Optional) */
    AzId []string `json:"azId"`

    /* VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例创建时间, UTC 时间格式 (Optional) */
    CreateTime string `json:"createTime"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 标签信息 (Optional) */
    Tags []Tag `json:"tags"`
}
