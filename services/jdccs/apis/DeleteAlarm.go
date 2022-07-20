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

type DeleteAlarmRequest struct {

    core.JDCloudRequest

    /* 报警规则ID  */
    AlarmId string `json:"alarmId"`
}

/*
 * param alarmId: 报警规则ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteAlarmRequest(
    alarmId string,
) *DeleteAlarmRequest {

	return &DeleteAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarms/{alarmId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        AlarmId: alarmId,
	}
}

/*
 * param alarmId: 报警规则ID (Required)
 */
func NewDeleteAlarmRequestWithAllParams(
    alarmId string,
) *DeleteAlarmRequest {

    return &DeleteAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarms/{alarmId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        AlarmId: alarmId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteAlarmRequestWithoutParam() *DeleteAlarmRequest {

    return &DeleteAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarms/{alarmId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param alarmId: 报警规则ID(Required) */
func (r *DeleteAlarmRequest) SetAlarmId(alarmId string) {
    r.AlarmId = alarmId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAlarmRequest) GetRegionId() string {
    return ""
}

type DeleteAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteAlarmResult `json:"result"`
}

type DeleteAlarmResult struct {
    Success bool `json:"success"`
}