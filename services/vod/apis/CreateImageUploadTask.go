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
)

type CreateImageUploadTaskRequest struct {

    core.JDCloudRequest

    /* HTTP 请求方法，取值范围：GET、POST、PUT、DELETE、HEAD、PATCH，默认值为 PUT (Optional) */
    HttpMethod *string `json:"httpMethod"`

    /* 文件名称 (Optional) */
    FileName *string `json:"fileName"`

    /* 文件大小 (Optional) */
    FileSize *int64 `json:"fileSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateImageUploadTaskRequest(
) *CreateImageUploadTaskRequest {

	return &CreateImageUploadTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/imageUploadTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param httpMethod: HTTP 请求方法，取值范围：GET、POST、PUT、DELETE、HEAD、PATCH，默认值为 PUT (Optional)
 * param fileName: 文件名称 (Optional)
 * param fileSize: 文件大小 (Optional)
 */
func NewCreateImageUploadTaskRequestWithAllParams(
    httpMethod *string,
    fileName *string,
    fileSize *int64,
) *CreateImageUploadTaskRequest {

    return &CreateImageUploadTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/imageUploadTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        HttpMethod: httpMethod,
        FileName: fileName,
        FileSize: fileSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateImageUploadTaskRequestWithoutParam() *CreateImageUploadTaskRequest {

    return &CreateImageUploadTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/imageUploadTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param httpMethod: HTTP 请求方法，取值范围：GET、POST、PUT、DELETE、HEAD、PATCH，默认值为 PUT(Optional) */
func (r *CreateImageUploadTaskRequest) SetHttpMethod(httpMethod string) {
    r.HttpMethod = &httpMethod
}

/* param fileName: 文件名称(Optional) */
func (r *CreateImageUploadTaskRequest) SetFileName(fileName string) {
    r.FileName = &fileName
}

/* param fileSize: 文件大小(Optional) */
func (r *CreateImageUploadTaskRequest) SetFileSize(fileSize int64) {
    r.FileSize = &fileSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateImageUploadTaskRequest) GetRegionId() string {
    return ""
}

type CreateImageUploadTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateImageUploadTaskResult `json:"result"`
}

type CreateImageUploadTaskResult struct {
    ImageId string `json:"imageId"`
    UploadUrl string `json:"uploadUrl"`
}