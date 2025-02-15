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
    partner "github.com/lshuining/jdcloud-sdk-go/services/partner/models"
)

type GetTotalConsumptionRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 按月查询开始时间（yyyy/MM/dd）  */
    StartTime string `json:"startTime"`

    /* 按月查询结束时间（yyyy/MM/dd）  */
    EndTime string `json:"endTime"`
}

/*
 * param regionId:  (Required)
 * param startTime: 按月查询开始时间（yyyy/MM/dd） (Required)
 * param endTime: 按月查询结束时间（yyyy/MM/dd） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTotalConsumptionRequest(
    regionId string,
    startTime string,
    endTime string,
) *GetTotalConsumptionRequest {

	return &GetTotalConsumptionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accountingBill:getTotalConsumption",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId:  (Required)
 * param startTime: 按月查询开始时间（yyyy/MM/dd） (Required)
 * param endTime: 按月查询结束时间（yyyy/MM/dd） (Required)
 */
func NewGetTotalConsumptionRequestWithAllParams(
    regionId string,
    startTime string,
    endTime string,
) *GetTotalConsumptionRequest {

    return &GetTotalConsumptionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accountingBill:getTotalConsumption",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTotalConsumptionRequestWithoutParam() *GetTotalConsumptionRequest {

    return &GetTotalConsumptionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accountingBill:getTotalConsumption",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *GetTotalConsumptionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 按月查询开始时间（yyyy/MM/dd）(Required) */
func (r *GetTotalConsumptionRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 按月查询结束时间（yyyy/MM/dd）(Required) */
func (r *GetTotalConsumptionRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTotalConsumptionRequest) GetRegionId() string {
    return r.RegionId
}

type GetTotalConsumptionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTotalConsumptionResult `json:"result"`
}

type GetTotalConsumptionResult struct {
    Result []partner.SummaryBill `json:"result"`
}