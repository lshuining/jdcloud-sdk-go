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


type CCAttackLog struct {

    /* CC 攻击记录Id (Optional) */
    AttackId string `json:"attackId"`

    /* 攻击流量大小 (Optional) */
    AttackTraffic float64 `json:"attackTraffic"`

    /* 是否触发黑洞，0否 1是 (Optional) */
    BlackHole int `json:"blackHole"`

    /* 攻击开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 攻击结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 流量单位，bps、Kbps、Mbps、Gbps (Optional) */
    Unit string `json:"unit"`

    /* 高防实例id (Optional) */
    InstanceId string `json:"instanceId"`

    /* 高防实例名称 (Optional) */
    Name string `json:"name"`
}
