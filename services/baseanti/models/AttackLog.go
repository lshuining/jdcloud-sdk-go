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


type AttackLog struct {

    /* 公网 IP 地址 (Optional) */
    Ip string `json:"ip"`

    /* 公网 IP 类型或绑定资源类型:
  0: 未知类型,
  1: 弹性公网 IP(IP 为弹性公网 IP, 绑定资源类型未知),
  10: 弹性公网 IP(IP 为弹性公网 IP, 但未绑定资源),
  11: 云主机,
  12: 负载均衡,
  13: 原生容器实例,
  14: 原生容器 Pod,
  2: 云物理服务器,
 (Optional) */
    ResourceType int `json:"resourceType"`

    /* 攻击记录 ID (Optional) */
    AttackLogId string `json:"attackLogId"`

    /* 攻击开始时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    StartTime string `json:"startTime"`

    /* 攻击结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    EndTime string `json:"endTime"`

    /* 触发原因:
0: 未知,
1: 四层,
2: 七层,
3: 四层和七层
 (Optional) */
    Cause int `json:"cause"`

    /* 状态, 0: 清洗完成, 1: 清洗中, 2: 黑洞中 (Optional) */
    Status int `json:"status"`

    /* 是否黑洞 (Optional) */
    BlackHole bool `json:"blackHole"`

    /* 攻击流量峰值 (Optional) */
    Peak int `json:"peak"`

    /* 攻击流量峰值单位 (Optional) */
    Unit string `json:"unit"`

    /* 攻击类型 (Optional) */
    AttackType []string `json:"attackType"`
}
