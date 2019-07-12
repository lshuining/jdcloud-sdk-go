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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    vod "github.com/jdcloud-api/jdcloud-sdk-go/services/vod/models"
)

type CreateTranscodeTemplateRequest struct {

    core.JDCloudRequest

    /* 模板名称 (Optional) */
    Name *string `json:"name"`

    /*  (Optional) */
    Video *vod.Video `json:"video"`

    /*  (Optional) */
    Audio *vod.Audio `json:"audio"`

    /*  (Optional) */
    Encapsulation *vod.Encapsulation `json:"encapsulation"`

    /* 清晰度规格 (Optional) */
    Definition *string `json:"definition"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTranscodeTemplateRequest(
) *CreateTranscodeTemplateRequest {

	return &CreateTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeTemplates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param name: 模板名称 (Optional)
 * param video:  (Optional)
 * param audio:  (Optional)
 * param encapsulation:  (Optional)
 * param definition: 清晰度规格 (Optional)
 */
func NewCreateTranscodeTemplateRequestWithAllParams(
    name *string,
    video *vod.Video,
    audio *vod.Audio,
    encapsulation *vod.Encapsulation,
    definition *string,
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

/* param name: 模板名称(Optional) */
func (r *CreateTranscodeTemplateRequest) SetName(name string) {
    r.Name = &name
}

/* param video: (Optional) */
func (r *CreateTranscodeTemplateRequest) SetVideo(video *vod.Video) {
    r.Video = video
}

/* param audio: (Optional) */
func (r *CreateTranscodeTemplateRequest) SetAudio(audio *vod.Audio) {
    r.Audio = audio
}

/* param encapsulation: (Optional) */
func (r *CreateTranscodeTemplateRequest) SetEncapsulation(encapsulation *vod.Encapsulation) {
    r.Encapsulation = encapsulation
}

/* param definition: 清晰度规格(Optional) */
func (r *CreateTranscodeTemplateRequest) SetDefinition(definition string) {
    r.Definition = &definition
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
    Definition string `json:"definition"`
    Source string `json:"source"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}