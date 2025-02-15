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

type DBInstance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例类别，例如主实例，只读实例等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceType string `json:"instanceType"`

    /* 存储类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>- 仅支持MySQL，Percona，MariaDB, SQL Server (Optional) */
    InstanceStorageType string `json:"instanceStorageType"`

    /* 实例数据加密. false：不加密; true：加密 (Optional) */
    StorageEncrypted bool `json:"storageEncrypted"`

    /* 实例引擎类型，如MySQL或SQL Server等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    Engine string `json:"engine"`

    /* 实例引擎版本，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 实例规格代码 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 磁盘，单位GB (Optional) */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* CPU核数 (Optional) */
    InstanceCPU int `json:"instanceCPU"`

    /* 内存，单位MB (Optional) */
    InstanceMemoryMB int `json:"instanceMemoryMB"`

    /* 地域ID，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    RegionId string `json:"regionId"`

    /* 可用区ID，第一个为主实例在的可用区，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    AzId []string `json:"azId"`

    /* VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例公网域名<br>- 仅支持MySQL (Optional) */
    PublicDomainName string `json:"publicDomainName"`

    /* 实例内网域名<br>- 仅支持MySQL (Optional) */
    InternalDomainName string `json:"internalDomainName"`

    /* 实例创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 实例跨地域备份服务开启相关信息 (Optional) */
    BackupSynchronicity []BackupSynchronicityAbstract `json:"backupSynchronicity"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 标签信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* MySQL、PostgreSQL只读实例对应的主实例ID (Optional) */
    SourceInstanceId string `json:"sourceInstanceId"`

    /* 应用访问端口<br>- 仅支持MySQL (Optional) */
    InstancePort string `json:"instancePort"`
}
