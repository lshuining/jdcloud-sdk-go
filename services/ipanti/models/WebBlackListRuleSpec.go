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


type WebBlackListRuleSpec struct {

    /* 黑名单规则名称  */
    Name string `json:"name"`

    /* 模式:<br>- 0: uri<br>- 1: ip<br>- 2: cookie<br>- 3: geo<br>- 4: header  */
    Mode int `json:"mode"`

    /* 匹配 key, mode 为 cookie 和 header 时必传. <br>- mode 为 cookie 时, 传 cookie 的 name<br>- mode 为 header 时, 传 header 的 key (Optional) */
    Key *string `json:"key"`

    /* 匹配 value. <br>- mode 为 uri 时, 传要匹配的 uri<br>- mode 为 ip 时, 传引用的 ip 黑白名单 Id<br>- mode 为 cookie 时, 传 cookie 的 value<br>- mode 为 geo 时, 传 geo 区域编码以 ',' 分隔的字符串. 查询 <a href='http://docs.jdcloud.com/anti-ddos-pro/api/describewebruleblacklistgeoareas'>describeWebRuleBlackListGeoAreas</a> 接口获取可设置的地域编码列表<br>- mode 为 header 时, 传 header 的 value  */
    Value string `json:"value"`

    /* 匹配规则, mode 为 uri, cookie 和 header 时必传. 支持以下匹配规则: <br>- 0: 完全匹配<br>- 1: 前缀匹配<br>- 2: 包含<br>- 3: 正则匹配<br>- 4: 后缀匹配 (Optional) */
    Pattern *int `json:"pattern"`

    /* 命中后处理动作. <br>- 0: 封禁并返回自定义页面<br>- 1: 跳转<br>- 2: 验证码  */
    Action int `json:"action"`

    /* 命中后处理值, 命中后处理动作为跳转时有效, 表示跳转路径 (Optional) */
    ActionValue *string `json:"actionValue"`

    /* 规则状态. <br>- 0: 关闭<br>- 1: 开启  */
    Status int `json:"status"`

    /* 关联的自定义页面id, 命中后处理动作为封禁并返回自定义页面时有效, 为空时表示默认页面 (Optional) */
    PageId *string `json:"pageId"`
}
