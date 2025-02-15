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

type GetTriggerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 函数名称  */
    FunctionName string `json:"functionName"`

    /* 版本名称  */
    VersionName string `json:"versionName"`

    /* 触发器Id  */
    TriggerId string `json:"triggerId"`
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 版本名称 (Required)
 * param triggerId: 触发器Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTriggerRequest(
    regionId string,
    functionName string,
    versionName string,
    triggerId string,
) *GetTriggerRequest {

	return &GetTriggerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:gettrigger",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FunctionName: functionName,
        VersionName: versionName,
        TriggerId: triggerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 版本名称 (Required)
 * param triggerId: 触发器Id (Required)
 */
func NewGetTriggerRequestWithAllParams(
    regionId string,
    functionName string,
    versionName string,
    triggerId string,
) *GetTriggerRequest {

    return &GetTriggerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:gettrigger",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FunctionName: functionName,
        VersionName: versionName,
        TriggerId: triggerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTriggerRequestWithoutParam() *GetTriggerRequest {

    return &GetTriggerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/versions/{versionName}:gettrigger",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetTriggerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param functionName: 函数名称(Required) */
func (r *GetTriggerRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}

/* param versionName: 版本名称(Required) */
func (r *GetTriggerRequest) SetVersionName(versionName string) {
    r.VersionName = versionName
}

/* param triggerId: 触发器Id(Required) */
func (r *GetTriggerRequest) SetTriggerId(triggerId string) {
    r.TriggerId = triggerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTriggerRequest) GetRegionId() string {
    return r.RegionId
}

type GetTriggerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTriggerResult `json:"result"`
}

type GetTriggerResult struct {
    Data function.Trigger `json:"data"`
}