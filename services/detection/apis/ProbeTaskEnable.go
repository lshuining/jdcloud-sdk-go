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

type ProbeTaskEnableRequest struct {

    core.JDCloudRequest

    /* 默认：禁用； true：启用，false：禁用 (Optional) */
    Enabled *bool `json:"enabled"`

    /* 要启用或禁用的探测任务的task_id列表，列表长度[1，100)  */
    TaskId []string `json:"taskId"`
}

/*
 * param taskId: 要启用或禁用的探测任务的task_id列表，列表长度[1，100) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewProbeTaskEnableRequest(
    taskId []string,
) *ProbeTaskEnableRequest {

	return &ProbeTaskEnableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/probeTask:switch",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        TaskId: taskId,
	}
}

/*
 * param enabled: 默认：禁用； true：启用，false：禁用 (Optional)
 * param taskId: 要启用或禁用的探测任务的task_id列表，列表长度[1，100) (Required)
 */
func NewProbeTaskEnableRequestWithAllParams(
    enabled *bool,
    taskId []string,
) *ProbeTaskEnableRequest {

    return &ProbeTaskEnableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask:switch",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        Enabled: enabled,
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewProbeTaskEnableRequestWithoutParam() *ProbeTaskEnableRequest {

    return &ProbeTaskEnableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask:switch",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param enabled: 默认：禁用； true：启用，false：禁用(Optional) */
func (r *ProbeTaskEnableRequest) SetEnabled(enabled bool) {
    r.Enabled = &enabled
}

/* param taskId: 要启用或禁用的探测任务的task_id列表，列表长度[1，100)(Required) */
func (r *ProbeTaskEnableRequest) SetTaskId(taskId []string) {
    r.TaskId = taskId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ProbeTaskEnableRequest) GetRegionId() string {
    return ""
}

type ProbeTaskEnableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ProbeTaskEnableResult `json:"result"`
}

type ProbeTaskEnableResult struct {
    Suc bool `json:"suc"`
}