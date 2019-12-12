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


type BandwidthTraffic struct {

    /* 机房英文标识 (Optional) */
    Idc string `json:"idc"`

    /* 机房名称 (Optional) */
    IdcName string `json:"idcName"`

    /* 带宽实例ID (Optional) */
    BandwidthId string `json:"bandwidthId"`

    /* 带宽名称 (Optional) */
    BandwidthName string `json:"bandwidthName"`

    /* 总上行实时流量 (Optional) */
    TotalTrafficIn float64 `json:"totalTrafficIn"`

    /* 总下行实时流量 (Optional) */
    TotalTrafficOut float64 `json:"totalTrafficOut"`

    /* 总带宽 (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 线路类型 bgp:BGP telecom:电信单线 unicom:联通单线 mobile:移动单线 (Optional) */
    LineType string `json:"lineType"`

    /* 关联的公网IP (Optional) */
    RelatedIp []interface{} `json:"relatedIp"`

    /*  (Optional) */
    Switchboard []interface{} `json:"switchboard"`
}
