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

type StopMcuTranscodeRequest struct {

    core.JDCloudRequest

    /* 应用ID (Optional) */
    AppId *string `json:"appId"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional) */
    UserRoomId *string `json:"userRoomId"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStopMcuTranscodeRequest(
) *StopMcuTranscodeRequest {

	return &StopMcuTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/stopMcuTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param appId: 应用ID (Optional)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional)
 */
func NewStopMcuTranscodeRequestWithAllParams(
    appId *string,
    userRoomId *string,
) *StopMcuTranscodeRequest {

    return &StopMcuTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/stopMcuTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserRoomId: userRoomId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStopMcuTranscodeRequestWithoutParam() *StopMcuTranscodeRequest {

    return &StopMcuTranscodeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/stopMcuTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Optional) */
func (r *StopMcuTranscodeRequest) SetAppId(appId string) {
    r.AppId = &appId
}

/* param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号(Optional) */
func (r *StopMcuTranscodeRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = &userRoomId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StopMcuTranscodeRequest) GetRegionId() string {
    return ""
}

type StopMcuTranscodeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StopMcuTranscodeResult `json:"result"`
}

type StopMcuTranscodeResult struct {
}