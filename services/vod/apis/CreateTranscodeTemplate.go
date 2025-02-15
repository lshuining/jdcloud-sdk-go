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

package apis

import (
    "github.com/lshuining/jdcloud-sdk-go/core"
    vod "github.com/lshuining/jdcloud-sdk-go/services/vod/models"
)

type CreateTranscodeTemplateRequest struct {

    core.JDCloudRequest

    /* 模板名称。长度不超过128个字符。UTF-8编码。
  */
    Name string `json:"name"`

    /* 视频参数配置  */
    Video *vod.Video `json:"video"`

    /* 音频参数配置  */
    Audio *vod.Audio `json:"audio"`

    /* 封装配置  */
    Encapsulation *vod.Encapsulation `json:"encapsulation"`

    /* 输出文件配置 (Optional) */
    OutFile *vod.OutFile `json:"outFile"`

    /* 清晰度规格标记。取值范围：
  SD - 标清
  HD - 高清
  FHD - 超清
  2K
  4K
  */
    Definition string `json:"definition"`

    /* 模板类型。取值范围：
  jdchd - 京享超清
  jdchs - 极速转码
 (Optional) */
    TemplateType *string `json:"templateType"`
}

/*
 * param name: 模板名称。长度不超过128个字符。UTF-8编码。
 (Required)
 * param video: 视频参数配置 (Required)
 * param audio: 音频参数配置 (Required)
 * param encapsulation: 封装配置 (Required)
 * param definition: 清晰度规格标记。取值范围：
  SD - 标清
  HD - 高清
  FHD - 超清
  2K
  4K
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTranscodeTemplateRequest(
    name string,
    video *vod.Video,
    audio *vod.Audio,
    encapsulation *vod.Encapsulation,
    definition string,
) *CreateTranscodeTemplateRequest {

	return &CreateTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeTemplates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Name: name,
        Video: video,
        Audio: audio,
        Encapsulation: encapsulation,
        Definition: definition,
	}
}

/*
 * param name: 模板名称。长度不超过128个字符。UTF-8编码。
 (Required)
 * param video: 视频参数配置 (Required)
 * param audio: 音频参数配置 (Required)
 * param encapsulation: 封装配置 (Required)
 * param outFile: 输出文件配置 (Optional)
 * param definition: 清晰度规格标记。取值范围：
  SD - 标清
  HD - 高清
  FHD - 超清
  2K
  4K
 (Required)
 * param templateType: 模板类型。取值范围：
  jdchd - 京享超清
  jdchs - 极速转码
 (Optional)
 */
func NewCreateTranscodeTemplateRequestWithAllParams(
    name string,
    video *vod.Video,
    audio *vod.Audio,
    encapsulation *vod.Encapsulation,
    outFile *vod.OutFile,
    definition string,
    templateType *string,
) *CreateTranscodeTemplateRequest {

    return &CreateTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeTemplates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        Video: video,
        Audio: audio,
        Encapsulation: encapsulation,
        OutFile: outFile,
        Definition: definition,
        TemplateType: templateType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTranscodeTemplateRequestWithoutParam() *CreateTranscodeTemplateRequest {

    return &CreateTranscodeTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeTemplates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: 模板名称。长度不超过128个字符。UTF-8编码。
(Required) */
func (r *CreateTranscodeTemplateRequest) SetName(name string) {
    r.Name = name
}

/* param video: 视频参数配置(Required) */
func (r *CreateTranscodeTemplateRequest) SetVideo(video *vod.Video) {
    r.Video = video
}

/* param audio: 音频参数配置(Required) */
func (r *CreateTranscodeTemplateRequest) SetAudio(audio *vod.Audio) {
    r.Audio = audio
}

/* param encapsulation: 封装配置(Required) */
func (r *CreateTranscodeTemplateRequest) SetEncapsulation(encapsulation *vod.Encapsulation) {
    r.Encapsulation = encapsulation
}

/* param outFile: 输出文件配置(Optional) */
func (r *CreateTranscodeTemplateRequest) SetOutFile(outFile *vod.OutFile) {
    r.OutFile = outFile
}

/* param definition: 清晰度规格标记。取值范围：
  SD - 标清
  HD - 高清
  FHD - 超清
  2K
  4K
(Required) */
func (r *CreateTranscodeTemplateRequest) SetDefinition(definition string) {
    r.Definition = definition
}

/* param templateType: 模板类型。取值范围：
  jdchd - 京享超清
  jdchs - 极速转码
(Optional) */
func (r *CreateTranscodeTemplateRequest) SetTemplateType(templateType string) {
    r.TemplateType = &templateType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTranscodeTemplateRequest) GetRegionId() string {
    return ""
}

type CreateTranscodeTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTranscodeTemplateResult `json:"result"`
}

type CreateTranscodeTemplateResult struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Video vod.Video `json:"video"`
    Audio vod.Audio `json:"audio"`
    Encapsulation vod.Encapsulation `json:"encapsulation"`
    OutFile vod.OutFile `json:"outFile"`
    Definition string `json:"definition"`
    Source string `json:"source"`
    TemplateType string `json:"templateType"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}