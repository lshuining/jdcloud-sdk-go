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


type CreateDomainResultObject struct {

    /* 域名ID (Optional) */
    Id string `json:"id"`

    /* 域名名称 (Optional) */
    Name string `json:"name"`

    /* 域名CNAME (Optional) */
    Cname string `json:"cname"`

    /* 域名来源：系统生成 | 用户自建 (Optional) */
    Source string `json:"source"`

    /* 是否默认域名 (Optional) */
    AsDefault bool `json:"asDefault"`

    /*  (Optional) */
    CreateTime string `json:"createTime"`

    /*  (Optional) */
    UpdateTime string `json:"updateTime"`
}
