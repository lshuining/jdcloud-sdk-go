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

type CloseLiveRestartRequest struct {

    core.JDCloudRequest

    /* 回看的播放域名  */
    RestartDomain string `json:"restartDomain"`
}

/*
 * param restartDomain: 回看的播放域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCloseLiveRestartRequest(
    restartDomain string,
) *CloseLiveRestartRequest {

	return &CloseLiveRestartRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveRestart:close",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RestartDomain: restartDomain,
	}
}

/*
 * param restartDomain: 回看的播放域名 (Required)
 */
func NewCloseLiveRestartRequestWithAllParams(
    restartDomain string,
) *CloseLiveRestartRequest {

    return &CloseLiveRestartRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRestart:close",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RestartDomain: restartDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCloseLiveRestartRequestWithoutParam() *CloseLiveRestartRequest {

    return &CloseLiveRestartRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRestart:close",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param restartDomain: 回看的播放域名(Required) */
func (r *CloseLiveRestartRequest) SetRestartDomain(restartDomain string) {
    r.RestartDomain = restartDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CloseLiveRestartRequest) GetRegionId() string {
    return ""
}

type CloseLiveRestartResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CloseLiveRestartResult `json:"result"`
}

type CloseLiveRestartResult struct {
}