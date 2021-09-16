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


type InstanceTemplateCustomData struct {

    /* 模板ID (Optional) */
    Id string `json:"id"`

    /* 用户自定义元数据。
以key-value键值对形式指定，可在实例系统内通过元数据服务查询获取。最多支持40对键值对，且key不超过256字符，value不超过16KB，不区分大小写。
注意：key不要以连字符(-)结尾，否则此key不生效。
 (Optional) */
    Metadata []Metadata `json:"metadata"`

    /* 自定义脚本。
目前仅支持启动脚本，即 `launch-script`，须 `base64` 编码且编码前数据长度不能超过16KB。
**linux系统**：支持 `bash` 和 `python`，编码前须分别以 `#!/bin/bash` 和 `#!/usr/bin/env python` 作为内容首行。
**Windows系统**：支持 `bat` 和 `powershell`，编码前须分别以 `<cmd></cmd>和<powershell></powershell>` 作为内容首、尾行。 (Optional) */
    Userdata []Userdata `json:"userdata"`
}
