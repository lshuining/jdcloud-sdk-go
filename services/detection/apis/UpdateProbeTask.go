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
    detection "github.com/lshuining/jdcloud-sdk-go/services/detection/models"
)

type UpdateProbeTaskRequest struct {

    core.JDCloudRequest

    /* 探测任务的task_id  */
    ProbeTaskID string `json:"probeTaskID"`

    /* http body：选择探测类型为1=http时有效，最长不超过1024字节 (Optional) */
    HttpBody *string `json:"httpBody"`

    /* http cookie：选择探测类型为1=http时有效，最大允许20个key、value对，最长不超过1024字节 (Optional) */
    HttpCookie []detection.KeyValue `json:"httpCookie"`

    /* http header：选择探测类型为1=http时有效，最大允许20个key、value对，最长不超过1024字节 (Optional) */
    HttpHeader []detection.KeyValue `json:"httpHeader"`

    /* http探测方法,可选值：1:get、2:post、3:head (Optional) */
    HttpType *int64 `json:"httpType"`

    /* task名称，不允许重复，长度不超过32字符，只允许中英文、数字、下划线_、中划线-, [0-9][a-z] [A-Z] [- _ ] (Optional) */
    Name *string `json:"name"`

    /* 探测源（发起对探测目标探测的云主机，需安装相应的agent才能探测）  */
    Probes []detection.Probe `json:"probes"`
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 * param probes: 探测源（发起对探测目标探测的云主机，需安装相应的agent才能探测） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateProbeTaskRequest(
    probeTaskID string,
    probes []detection.Probe,
) *UpdateProbeTaskRequest {

	return &UpdateProbeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/probeTask/{probeTaskID}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v2",
		},
        ProbeTaskID: probeTaskID,
        Probes: probes,
	}
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 * param httpBody: http body：选择探测类型为1=http时有效，最长不超过1024字节 (Optional)
 * param httpCookie: http cookie：选择探测类型为1=http时有效，最大允许20个key、value对，最长不超过1024字节 (Optional)
 * param httpHeader: http header：选择探测类型为1=http时有效，最大允许20个key、value对，最长不超过1024字节 (Optional)
 * param httpType: http探测方法,可选值：1:get、2:post、3:head (Optional)
 * param name: task名称，不允许重复，长度不超过32字符，只允许中英文、数字、下划线_、中划线-, [0-9][a-z] [A-Z] [- _ ] (Optional)
 * param probes: 探测源（发起对探测目标探测的云主机，需安装相应的agent才能探测） (Required)
 */
func NewUpdateProbeTaskRequestWithAllParams(
    probeTaskID string,
    httpBody *string,
    httpCookie []detection.KeyValue,
    httpHeader []detection.KeyValue,
    httpType *int64,
    name *string,
    probes []detection.Probe,
) *UpdateProbeTaskRequest {

    return &UpdateProbeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask/{probeTaskID}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v2",
        },
        ProbeTaskID: probeTaskID,
        HttpBody: httpBody,
        HttpCookie: httpCookie,
        HttpHeader: httpHeader,
        HttpType: httpType,
        Name: name,
        Probes: probes,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateProbeTaskRequestWithoutParam() *UpdateProbeTaskRequest {

    return &UpdateProbeTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask/{probeTaskID}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param probeTaskID: 探测任务的task_id(Required) */
func (r *UpdateProbeTaskRequest) SetProbeTaskID(probeTaskID string) {
    r.ProbeTaskID = probeTaskID
}

/* param httpBody: http body：选择探测类型为1=http时有效，最长不超过1024字节(Optional) */
func (r *UpdateProbeTaskRequest) SetHttpBody(httpBody string) {
    r.HttpBody = &httpBody
}

/* param httpCookie: http cookie：选择探测类型为1=http时有效，最大允许20个key、value对，最长不超过1024字节(Optional) */
func (r *UpdateProbeTaskRequest) SetHttpCookie(httpCookie []detection.KeyValue) {
    r.HttpCookie = httpCookie
}

/* param httpHeader: http header：选择探测类型为1=http时有效，最大允许20个key、value对，最长不超过1024字节(Optional) */
func (r *UpdateProbeTaskRequest) SetHttpHeader(httpHeader []detection.KeyValue) {
    r.HttpHeader = httpHeader
}

/* param httpType: http探测方法,可选值：1:get、2:post、3:head(Optional) */
func (r *UpdateProbeTaskRequest) SetHttpType(httpType int64) {
    r.HttpType = &httpType
}

/* param name: task名称，不允许重复，长度不超过32字符，只允许中英文、数字、下划线_、中划线-, [0-9][a-z] [A-Z] [- _ ](Optional) */
func (r *UpdateProbeTaskRequest) SetName(name string) {
    r.Name = &name
}

/* param probes: 探测源（发起对探测目标探测的云主机，需安装相应的agent才能探测）(Required) */
func (r *UpdateProbeTaskRequest) SetProbes(probes []detection.Probe) {
    r.Probes = probes
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateProbeTaskRequest) GetRegionId() string {
    return ""
}

type UpdateProbeTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateProbeTaskResult `json:"result"`
}

type UpdateProbeTaskResult struct {
    Suc bool `json:"suc"`
}