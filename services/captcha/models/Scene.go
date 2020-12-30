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


type Scene struct {

    /* 场景id,更新时必填 (Optional) */
    SceneId int64 `json:"sceneId"`

    /* 场景密钥 (Optional) */
    SceneSecret string `json:"sceneSecret"`

    /* 所属应用id  */
    AppId int64 `json:"appId"`

    /* 所属应用名称 (Optional) */
    AppName string `json:"appName"`

    /* 场景名称  */
    SceneName string `json:"sceneName"`

    /* 场景类型：account：账号场景（登录、注册等）activity：活动场景（秒杀、领券等）content：内容场景（发帖评论、签到投票等）other：其它  */
    SceneType string `json:"sceneType"`

    /* 平均qps  */
    SceneAvgQps int `json:"sceneAvgQps"`

    /* 高峰期qps  */
    SceneMaxQps int `json:"sceneMaxQps"`

    /* 场景描述 (Optional) */
    Description string `json:"description"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}
