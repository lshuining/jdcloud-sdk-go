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
    function "github.com/lshuining/jdcloud-sdk-go/services/function/models"
)

type InvokeRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 函数名称  */
    FunctionName string `json:"functionName"`

    /* 版本名称  */
    VersionName string `json:"versionName"`

    /* 执行函数的输入事件  */
    Event string `json:"event"`
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 版本名称 (Required)
 * param event: 执行函数的输入事件 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewInvokeRequest(
    regionId string,
    functionName string,
    versionName string,
    event string,
) *InvokeRequest {

	return &InvokeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:invoke",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FunctionName: functionName,
        VersionName: versionName,
        Event: event,
	}
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 版本名称 (Required)
 * param event: 执行函数的输入事件 (Required)
 */
func NewInvokeRequestWithAllParams(
    regionId string,
    functionName string,
    versionName string,
    event string,
) *InvokeRequest {

    return &InvokeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:invoke",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FunctionName: functionName,
        VersionName: versionName,
        Event: event,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewInvokeRequestWithoutParam() *InvokeRequest {

    return &InvokeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:invoke",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *InvokeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param functionName: 函数名称(Required) */
func (r *InvokeRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}

/* param versionName: 版本名称(Required) */
func (r *InvokeRequest) SetVersionName(versionName string) {
    r.VersionName = versionName
}

/* param event: 执行函数的输入事件(Required) */
func (r *InvokeRequest) SetEvent(event string) {
    r.Event = event
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r InvokeRequest) GetRegionId() string {
    return r.RegionId
}

type InvokeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result InvokeResult `json:"result"`
}

type InvokeResult struct {
    Data function.FunctionInvokeResult `json:"data"`
}