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

type ExecGetExitCodeRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Container ID  */
    ContainerId string `json:"containerId"`

    /* exec ID  */
    ExecId string `json:"execId"`
}

/*
 * param regionId: Region ID (Required)
 * param containerId: Container ID (Required)
 * param execId: exec ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExecGetExitCodeRequest(
    regionId string,
    containerId string,
    execId string,
) *ExecGetExitCodeRequest {

	return &ExecGetExitCodeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/containers/{containerId}:execGetExitCode",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ContainerId: containerId,
        ExecId: execId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param containerId: Container ID (Required)
 * param execId: exec ID (Required)
 */
func NewExecGetExitCodeRequestWithAllParams(
    regionId string,
    containerId string,
    execId string,
) *ExecGetExitCodeRequest {

    return &ExecGetExitCodeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/containers/{containerId}:execGetExitCode",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ContainerId: containerId,
        ExecId: execId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExecGetExitCodeRequestWithoutParam() *ExecGetExitCodeRequest {

    return &ExecGetExitCodeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/containers/{containerId}:execGetExitCode",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ExecGetExitCodeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param containerId: Container ID(Required) */
func (r *ExecGetExitCodeRequest) SetContainerId(containerId string) {
    r.ContainerId = containerId
}

/* param execId: exec ID(Required) */
func (r *ExecGetExitCodeRequest) SetExecId(execId string) {
    r.ExecId = execId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExecGetExitCodeRequest) GetRegionId() string {
    return r.RegionId
}

type ExecGetExitCodeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExecGetExitCodeResult `json:"result"`
}

type ExecGetExitCodeResult struct {
    ExitCode int `json:"exitCode"`
}