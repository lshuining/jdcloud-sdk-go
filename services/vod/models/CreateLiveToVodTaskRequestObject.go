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


type CreateLiveToVodTaskRequestObject struct {

    /* 视频标题  */
    Title string `json:"title"`

    /* 文件名称  */
    FileName string `json:"fileName"`

    /* 文件大小 (Optional) */
    FileSize int64 `json:"fileSize"`

    /* 封面地址 (Optional) */
    CoverUrl string `json:"coverUrl"`

    /* 视频描述 (Optional) */
    Description string `json:"description"`

    /* 分类ID (Optional) */
    CategoryId int64 `json:"categoryId"`

    /* 视频标签集合 (Optional) */
    Tags []string `json:"tags"`

    /* 转码模板ID集合 (Optional) */
    TranscodeTemplateIds []int64 `json:"transcodeTemplateIds"`

    /* 水印ID集合 (Optional) */
    WatermarkIds []int64 `json:"watermarkIds"`

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 流名称  */
    StreamName string `json:"streamName"`

    /* 录制时间段集合
- 支持自定义1-10个时间段,拼接成一个文件
- 每个时间段不小于10s
- 总跨度不超过12小时
- 时间段按升序排列且无重叠
  */
    RecordTimes []RecordTime `json:"recordTimes"`

    /* 录制文件类型:
- 取值: ts, flv, mp4
- 不区分大小写
  */
    RecordFileType string `json:"recordFileType"`

    /* 直播录制任务外键 (Optional) */
    TaskExternalId string `json:"taskExternalId"`

    /* 任务优先级:
- 取值: low, medium, high
- 不区分大小写
 (Optional) */
    Priority string `json:"priority"`
}
