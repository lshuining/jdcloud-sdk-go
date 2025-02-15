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


type Zone struct {

    /*  (Optional) */
    Plan_pending Plan_pending `json:"plan_pending"`

    /* 切换到星盾时的DNS主机 (Optional) */
    Original_dnshost string `json:"original_dnshost"`

    /* 当前用户在域上的可用权限 (Optional) */
    Permissions []string `json:"permissions"`

    /* 域的开发模式过期（正整数）或上次过期（负整数）的时间间隔（秒）。如果从未启用过开发模式，则此值为0。
 (Optional) */
    Development_mode int `json:"development_mode"`

    /*  (Optional) */
    Verification_key string `json:"verification_key"`

    /*  (Optional) */
    Plan Plan `json:"plan"`

    /* 迁移到星盾之前的原始域名服务器 (Optional) */
    Original_name_servers []string `json:"original_name_servers"`

    /* 域名 (Optional) */
    Name string `json:"name"`

    /*  (Optional) */
    Account Account `json:"account"`

    /* 最后一次检测到所有权证明和该域激活的时间 (Optional) */
    Activated_on string `json:"activated_on"`

    /* 指示域是否仅使用星盾DNS服务。如果值为真，则表示该域将不会获得安全或性能方面的好处。 (Optional) */
    Paused bool `json:"paused"`

    /* 域的状态 (Optional) */
    Status string `json:"status"`

    /*  (Optional) */
    Owner Owner `json:"owner"`

    /* 上次修改域的时间 (Optional) */
    Modified_on string `json:"modified_on"`

    /* 创建域的时间 (Optional) */
    Created_on string `json:"created_on"`

    /* 全接入的域意味着DNS由星盾托管。半接入的域通常是合作伙伴托管的域或CNAME设置。 (Optional) */
    Ty_pe string `json:"ty_pe"`

    /* 域标识符标签 (Optional) */
    Id string `json:"id"`

    /* 星盾分配的域名服务器。这仅适用于使用星盾DNS的域 (Optional) */
    Name_servers []string `json:"name_servers"`

    /* 切换到星盾时的域注册商 (Optional) */
    Original_registrar string `json:"original_registrar"`
}
