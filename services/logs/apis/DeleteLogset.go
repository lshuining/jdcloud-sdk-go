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

type DeleteLogsetRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集ID，多个日志集ID以逗号分割  */
    LogsetUIDs string `json:"logsetUIDs"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUIDs: 日志集ID，多个日志集ID以逗号分割 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLogsetRequest(
    regionId string,
    logsetUIDs string,
) *DeleteLogsetRequest {

	return &DeleteLogsetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUIDs}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUIDs: logsetUIDs,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUIDs: 日志集ID，多个日志集ID以逗号分割 (Required)
 */
func NewDeleteLogsetRequestWithAllParams(
    regionId string,
    logsetUIDs string,
) *DeleteLogsetRequest {

    return &DeleteLogsetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUIDs}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUIDs: logsetUIDs,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLogsetRequestWithoutParam() *DeleteLogsetRequest {

    return &DeleteLogsetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUIDs}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DeleteLogsetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUIDs: 日志集ID，多个日志集ID以逗号分割(Required) */
func (r *DeleteLogsetRequest) SetLogsetUIDs(logsetUIDs string) {
    r.LogsetUIDs = logsetUIDs
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLogsetRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteLogsetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLogsetResult `json:"result"`
}

type DeleteLogsetResult struct {
}