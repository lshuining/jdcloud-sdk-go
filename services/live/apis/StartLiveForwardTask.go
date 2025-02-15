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

type StartLiveForwardTaskRequest struct {

    core.JDCloudRequest

    /* 任务ID,批量用,分隔
  */
    TaskIds string `json:"taskIds"`
}

/*
 * param taskIds: 任务ID,批量用,分隔
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStartLiveForwardTaskRequest(
    taskIds string,
) *StartLiveForwardTaskRequest {

	return &StartLiveForwardTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/LiveForwardTask:start",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TaskIds: taskIds,
	}
}

/*
 * param taskIds: 任务ID,批量用,分隔
 (Required)
 */
func NewStartLiveForwardTaskRequestWithAllParams(
    taskIds string,
) *StartLiveForwardTaskRequest {

    return &StartLiveForwardTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/LiveForwardTask:start",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TaskIds: taskIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStartLiveForwardTaskRequestWithoutParam() *StartLiveForwardTaskRequest {

    return &StartLiveForwardTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/LiveForwardTask:start",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param taskIds: 任务ID,批量用,分隔
(Required) */
func (r *StartLiveForwardTaskRequest) SetTaskIds(taskIds string) {
    r.TaskIds = taskIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StartLiveForwardTaskRequest) GetRegionId() string {
    return ""
}

type StartLiveForwardTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StartLiveForwardTaskResult `json:"result"`
}

type StartLiveForwardTaskResult struct {
}