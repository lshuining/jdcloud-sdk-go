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
)

type SubmitViJobRequest struct {

    core.JDCloudRequest

    /* 视频审查模板ID  */
    TemplateId string `json:"templateId"`

    /* 对象存储区域，输入和输入同区域  */
    Region string `json:"region"`

    /* 输入空间 (Optional) */
    InputBucket *string `json:"inputBucket"`

    /* 输入文件 (Optional) */
    InputFileKey *string `json:"inputFileKey"`

    /* 输入空间 (Optional) */
    OutputBucket *string `json:"outputBucket"`

    /* 输入路径 (Optional) */
    OutputFilePath *string `json:"outputFilePath"`
}

/*
 * param templateId: 视频审查模板ID (Required)
 * param region: 对象存储区域，输入和输入同区域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSubmitViJobRequest(
    templateId string,
    region string,
) *SubmitViJobRequest {

	return &SubmitViJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/viJobs:submit",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        TemplateId: templateId,
        Region: region,
	}
}

/*
 * param templateId: 视频审查模板ID (Required)
 * param region: 对象存储区域，输入和输入同区域 (Required)
 * param inputBucket: 输入空间 (Optional)
 * param inputFileKey: 输入文件 (Optional)
 * param outputBucket: 输入空间 (Optional)
 * param outputFilePath: 输入路径 (Optional)
 */
func NewSubmitViJobRequestWithAllParams(
    templateId string,
    region string,
    inputBucket *string,
    inputFileKey *string,
    outputBucket *string,
    outputFilePath *string,
) *SubmitViJobRequest {

    return &SubmitViJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/viJobs:submit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TemplateId: templateId,
        Region: region,
        InputBucket: inputBucket,
        InputFileKey: inputFileKey,
        OutputBucket: outputBucket,
        OutputFilePath: outputFilePath,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSubmitViJobRequestWithoutParam() *SubmitViJobRequest {

    return &SubmitViJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/viJobs:submit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param templateId: 视频审查模板ID(Required) */
func (r *SubmitViJobRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

/* param region: 对象存储区域，输入和输入同区域(Required) */
func (r *SubmitViJobRequest) SetRegion(region string) {
    r.Region = region
}

/* param inputBucket: 输入空间(Optional) */
func (r *SubmitViJobRequest) SetInputBucket(inputBucket string) {
    r.InputBucket = &inputBucket
}

/* param inputFileKey: 输入文件(Optional) */
func (r *SubmitViJobRequest) SetInputFileKey(inputFileKey string) {
    r.InputFileKey = &inputFileKey
}

/* param outputBucket: 输入空间(Optional) */
func (r *SubmitViJobRequest) SetOutputBucket(outputBucket string) {
    r.OutputBucket = &outputBucket
}

/* param outputFilePath: 输入路径(Optional) */
func (r *SubmitViJobRequest) SetOutputFilePath(outputFilePath string) {
    r.OutputFilePath = &outputFilePath
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SubmitViJobRequest) GetRegionId() string {
    return ""
}

type SubmitViJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SubmitViJobResult `json:"result"`
}

type SubmitViJobResult struct {
    JobId string `json:"jobId"`
    TemplateId string `json:"templateId"`
    Region string `json:"region"`
    InputBucket string `json:"inputBucket"`
    InputFileKey string `json:"inputFileKey"`
    OutputBucket string `json:"outputBucket"`
    OutputFilePath string `json:"outputFilePath"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}