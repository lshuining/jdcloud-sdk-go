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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type ReportPrecheckRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /*   */
    Plan interface{} `json:"plan"`

    /* 预检查任务项  */
    CheckItem string `json:"checkItem"`

    /* 是否成功  */
    Success bool `json:"success"`

    /* 错误信息，仅sucess为false时返回 (Optional) */
    ErrorMessages *string `json:"errorMessages"`

    /*   */
    StartTime string `json:"startTime"`

    /*   */
    EndTime string `json:"endTime"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param plan:  (Required)
 * param checkItem: 预检查任务项 (Required)
 * param success: 是否成功 (Required)
 * param startTime:  (Required)
 * param endTime:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewReportPrecheckRequest(
    regionId string,
    plan interface{},
    checkItem string,
    success bool,
    startTime string,
    endTime string,
) *ReportPrecheckRequest {

	return &ReportPrecheckRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/precheckPlan",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        Plan: plan,
        CheckItem: checkItem,
        Success: success,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param plan:  (Required)
 * param checkItem: 预检查任务项 (Required)
 * param success: 是否成功 (Required)
 * param errorMessages: 错误信息，仅sucess为false时返回 (Optional)
 * param startTime:  (Required)
 * param endTime:  (Required)
 */
func NewReportPrecheckRequestWithAllParams(
    regionId string,
    plan interface{},
    checkItem string,
    success bool,
    errorMessages *string,
    startTime string,
    endTime string,
) *ReportPrecheckRequest {

    return &ReportPrecheckRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/precheckPlan",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Plan: plan,
        CheckItem: checkItem,
        Success: success,
        ErrorMessages: errorMessages,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewReportPrecheckRequestWithoutParam() *ReportPrecheckRequest {

    return &ReportPrecheckRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/precheckPlan",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *ReportPrecheckRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param plan: (Required) */
func (r *ReportPrecheckRequest) SetPlan(plan interface{}) {
    r.Plan = plan
}

/* param checkItem: 预检查任务项(Required) */
func (r *ReportPrecheckRequest) SetCheckItem(checkItem string) {
    r.CheckItem = checkItem
}

/* param success: 是否成功(Required) */
func (r *ReportPrecheckRequest) SetSuccess(success bool) {
    r.Success = success
}

/* param errorMessages: 错误信息，仅sucess为false时返回(Optional) */
func (r *ReportPrecheckRequest) SetErrorMessages(errorMessages string) {
    r.ErrorMessages = &errorMessages
}

/* param startTime: (Required) */
func (r *ReportPrecheckRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: (Required) */
func (r *ReportPrecheckRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ReportPrecheckRequest) GetRegionId() string {
    return r.RegionId
}

type ReportPrecheckResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ReportPrecheckResult `json:"result"`
}

type ReportPrecheckResult struct {
    JobId string `json:"jobId"`
}