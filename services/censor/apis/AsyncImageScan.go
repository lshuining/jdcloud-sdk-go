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
    censor "github.com/lshuining/jdcloud-sdk-go/services/censor/models"
)

type AsyncImageScanRequest struct {

    core.JDCloudRequest

    /* 机审策略，eg: default (Optional) */
    BizType *string `json:"bizType"`

    /* 指定检测场景 (Optional) */
    Scenes []string `json:"scenes"`

    /* 检测任务列表，包含一个或多个元素。每个元素是个结构体，最多可添加10个元素，每个元素的具体结构描述见ImageTask。 (Optional) */
    Tasks []censor.ImageTask `json:"tasks"`

    /* 异步检测结果回调通知您的URL，支持HTTP/HTTPS。该字段为空时，您必须定时检索检测结果。 (Optional) */
    Callback *string `json:"callback"`

    /* 随机字符串，该值用于回调通知请求中的签名。当使用callback时，该字段必须提供。 (Optional) */
    Seed *string `json:"seed"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAsyncImageScanRequest(
) *AsyncImageScanRequest {

	return &AsyncImageScanRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/image:asyncscan",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param bizType: 机审策略，eg: default (Optional)
 * param scenes: 指定检测场景 (Optional)
 * param tasks: 检测任务列表，包含一个或多个元素。每个元素是个结构体，最多可添加10个元素，每个元素的具体结构描述见ImageTask。 (Optional)
 * param callback: 异步检测结果回调通知您的URL，支持HTTP/HTTPS。该字段为空时，您必须定时检索检测结果。 (Optional)
 * param seed: 随机字符串，该值用于回调通知请求中的签名。当使用callback时，该字段必须提供。 (Optional)
 */
func NewAsyncImageScanRequestWithAllParams(
    bizType *string,
    scenes []string,
    tasks []censor.ImageTask,
    callback *string,
    seed *string,
) *AsyncImageScanRequest {

    return &AsyncImageScanRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/image:asyncscan",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        BizType: bizType,
        Scenes: scenes,
        Tasks: tasks,
        Callback: callback,
        Seed: seed,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAsyncImageScanRequestWithoutParam() *AsyncImageScanRequest {

    return &AsyncImageScanRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/image:asyncscan",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param bizType: 机审策略，eg: default(Optional) */
func (r *AsyncImageScanRequest) SetBizType(bizType string) {
    r.BizType = &bizType
}

/* param scenes: 指定检测场景(Optional) */
func (r *AsyncImageScanRequest) SetScenes(scenes []string) {
    r.Scenes = scenes
}

/* param tasks: 检测任务列表，包含一个或多个元素。每个元素是个结构体，最多可添加10个元素，每个元素的具体结构描述见ImageTask。(Optional) */
func (r *AsyncImageScanRequest) SetTasks(tasks []censor.ImageTask) {
    r.Tasks = tasks
}

/* param callback: 异步检测结果回调通知您的URL，支持HTTP/HTTPS。该字段为空时，您必须定时检索检测结果。(Optional) */
func (r *AsyncImageScanRequest) SetCallback(callback string) {
    r.Callback = &callback
}

/* param seed: 随机字符串，该值用于回调通知请求中的签名。当使用callback时，该字段必须提供。(Optional) */
func (r *AsyncImageScanRequest) SetSeed(seed string) {
    r.Seed = &seed
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AsyncImageScanRequest) GetRegionId() string {
    return ""
}

type AsyncImageScanResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AsyncImageScanResult `json:"result"`
}

type AsyncImageScanResult struct {
    Data []censor.TaskData `json:"data"`
}