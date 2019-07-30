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


type PolicyDetail struct {

    /* 策略id (Optional) */
    PolicyId string `json:"policyId"`

    /* 策略名称 (Optional) */
    Name string `json:"name"`

    /* 京东云资源标识(jrn) (Optional) */
    Jrn string `json:"jrn"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 策略类型 (Optional) */
    PolicyType string `json:"policyType"`

    /* 策略版本号 (Optional) */
    Version string `json:"version"`

    /* 默认策略文档版本 (Optional) */
    DefaultEdition int `json:"defaultEdition"`

    /* 策略文档 (Optional) */
    Content string `json:"content"`

    /* 策略创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 策略更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
