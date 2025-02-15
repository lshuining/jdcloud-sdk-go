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

type DedicatedHostSpec struct {

    /* 专有宿主机池ID，创建专有宿主机必须指定专有宿主机池ID  */
    DedicatedPoolId string `json:"dedicatedPoolId"`

    /* 专有宿主机所属的可用区，不传入该参数时可用区属性从专有宿主机池中继承；指定的可用区必须是对应专有宿主机池中设置的可用区的子集 (Optional) */
    Az *string `json:"az"`

    /* 专有宿主机名称  */
    Name string `json:"name"`

    /* 专有宿主机描述 (Optional) */
    Description *string `json:"description"`

    /* 计费配置。
专有宿主机不支持按用量方式计费，默认为按配置计费。 (Optional) */
    Charge *charge.ChargeSpec `json:"charge"`
}
