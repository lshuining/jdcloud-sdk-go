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

type GetLogsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Container ID  */
    ContainerId string `json:"containerId"`

    /* 返回日志文件中倒数 tailLines 行，如不指定，默认从容器启动时或 sinceSeconds 指定的时间读取。
 (Optional) */
    TailLines *int `json:"tailLines"`

    /* 返回相对于当前时间之前sinceSeconds之内的日志。
 (Optional) */
    SinceSeconds *int `json:"sinceSeconds"`

    /* 限制返回的日志文件内容字节数，取值范围 [1-4]KB，最大 4KB.
 (Optional) */
    LimitBytes *int `json:"limitBytes"`
}

/*
 * param regionId: Region ID (Required)
 * param containerId: Container ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetLogsRequest(
    regionId string,
    containerId string,
) *GetLogsRequest {

	return &GetLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/containers/{containerId}:getLogs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ContainerId: containerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param containerId: Container ID (Required)
 * param tailLines: 返回日志文件中倒数 tailLines 行，如不指定，默认从容器启动时或 sinceSeconds 指定的时间读取。
 (Optional)
 * param sinceSeconds: 返回相对于当前时间之前sinceSeconds之内的日志。
 (Optional)
 * param limitBytes: 限制返回的日志文件内容字节数，取值范围 [1-4]KB，最大 4KB.
 (Optional)
 */
func NewGetLogsRequestWithAllParams(
    regionId string,
    containerId string,
    tailLines *int,
    sinceSeconds *int,
    limitBytes *int,
) *GetLogsRequest {

    return &GetLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/containers/{containerId}:getLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ContainerId: containerId,
        TailLines: tailLines,
        SinceSeconds: sinceSeconds,
        LimitBytes: limitBytes,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetLogsRequestWithoutParam() *GetLogsRequest {

    return &GetLogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/containers/{containerId}:getLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetLogsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param containerId: Container ID(Required) */
func (r *GetLogsRequest) SetContainerId(containerId string) {
    r.ContainerId = containerId
}

/* param tailLines: 返回日志文件中倒数 tailLines 行，如不指定，默认从容器启动时或 sinceSeconds 指定的时间读取。
(Optional) */
func (r *GetLogsRequest) SetTailLines(tailLines int) {
    r.TailLines = &tailLines
}

/* param sinceSeconds: 返回相对于当前时间之前sinceSeconds之内的日志。
(Optional) */
func (r *GetLogsRequest) SetSinceSeconds(sinceSeconds int) {
    r.SinceSeconds = &sinceSeconds
}

/* param limitBytes: 限制返回的日志文件内容字节数，取值范围 [1-4]KB，最大 4KB.
(Optional) */
func (r *GetLogsRequest) SetLimitBytes(limitBytes int) {
    r.LimitBytes = &limitBytes
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetLogsRequest) GetRegionId() string {
    return r.RegionId
}

type GetLogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetLogsResult `json:"result"`
}

type GetLogsResult struct {
    Logs interface{} `json:"logs"`
}