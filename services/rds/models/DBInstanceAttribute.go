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

type DBInstanceAttribute struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例类型，例如主实例，只读实例等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceType string `json:"instanceType"`

    /* 实例引擎类型，如MySQL或SQL Server等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    Engine string `json:"engine"`

    /* 实例引擎版本，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 实例引擎的小版本 (Optional) */
    MinorVersion string `json:"minorVersion"`

    /* 实例规格代码 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 存储类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStorageType string `json:"instanceStorageType"`

    /* 实例数据加密. false：不加密; true：加密 (Optional) */
    StorageEncrypted bool `json:"storageEncrypted"`

    /* 磁盘，单位GB (Optional) */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* CPU核数 (Optional) */
    InstanceCPU int `json:"instanceCPU"`

    /* 内存大小，单位MB (Optional) */
    InstanceMemoryMB int `json:"instanceMemoryMB"`

    /* 地域ID，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    RegionId string `json:"regionId"`

    /* 可用区ID，第一个为主实例在的可用区，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    AzId []string `json:"azId"`

    /* VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 参数组的ID<br>- 仅支持MySQL (Optional) */
    ParameterGroupId string `json:"parameterGroupId"`

    /* 参数组的名称<br>- 仅支持MySQL (Optional) */
    ParameterGroupName string `json:"parameterGroupName"`

    /* 参数的状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>- 仅支持MySQL (Optional) */
    ParameterStatus string `json:"parameterStatus"`

    /* 实例内网域名 (Optional) */
    InternalDomainName string `json:"internalDomainName"`

    /* 实例公网域名 (Optional) */
    PublicDomainName string `json:"publicDomainName"`

    /* 应用访问端口 (Optional) */
    InstancePort string `json:"instancePort"`

    /* 访问模式，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>- 仅支持MySQL (Optional) */
    ConnectionMode string `json:"connectionMode"`

    /* 审计状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)<br>- 仅支持MySQL (Optional) */
    AuditStatus string `json:"auditStatus"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* MySQL只读实例对应的主实例ID<br>- 仅支持MySQL (Optional) */
    SourceInstanceId string `json:"sourceInstanceId"`

    /* 只读实例ID列表<br>- 仅支持MySQL (Optional) */
    RoInstanceIds []string `json:"roInstanceIds"`

    /* 高可用集群中主节点的信息<br>- 仅支持SQL Server (Optional) */
    PrimaryNode DBInstanceNode `json:"primaryNode"`

    /* 高可用集群中从节点的信息<br>- 仅支持SQL Server (Optional) */
    SecondaryNode DBInstanceNode `json:"secondaryNode"`

    /* 标签信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* 对接的目录服务的相关信息<br>仅支SQL Server (Optional) */
    ActiveDirectory ADService `json:"activeDirectory"`

    /* 只读代理服务 ID (Optional) */
    RoInstanceProxyID string `json:"roInstanceProxyID"`

    /* 只读代理服务 名称 (Optional) */
    RoInstanceProxyName string `json:"roInstanceProxyName"`

    /* 读写分离代理服务 ID (Optional) */
    ReadWriteProxyId string `json:"readWriteProxyId"`

    /* 实例关联的数据同步任务Id；未关联数据同步任务时返回空 (Optional) */
    SyncTaskId string `json:"syncTaskId"`

    /* 实例内网域名解析到的内网IP地址<br>仅支持MySQL, Percona, MariaDB (Optional) */
    InstanceVip string `json:"instanceVip"`
}
