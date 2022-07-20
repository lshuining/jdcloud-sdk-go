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
    privatezone "github.com/lshuining/jdcloud-sdk-go/services/privatezone/models"
)

type DescribeActionLogsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页容量，默认10，取值范围(1 - 100) (Optional) */
    PageSize *int `json:"pageSize"`

    /* 页序号，默认值1，不能小于1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 起始时间，格式：UTC时间例如2017-11-10T23:00:00Z  */
    Start string `json:"start"`

    /* 结束时间，格式：UTC时间例如2017-11-10T23:00:00Z  */
    End string `json:"end"`

    /* 日志模糊匹配的关键词 (Optional) */
    KeyWord *string `json:"keyWord"`

    /* 操作结果 false->失败 true->成功 (Optional) */
    Success *bool `json:"success"`

    /* 日志类型，支持的类型有：CREATE_ZONE、DELETE_ZONE、LOCK_ZONE、CREATE_RR、MODIFY_RR、DELETE_RR、SET_RR_STATUS、RETRY_RECURSE_ZONE (Optional) */
    ActionType *string `json:"actionType"`
}

/*
 * param regionId: 地域ID (Required)
 * param start: 起始时间，格式：UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 结束时间，格式：UTC时间例如2017-11-10T23:00:00Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeActionLogsRequest(
    regionId string,
    start string,
    end string,
) *DescribeActionLogsRequest {

	return &DescribeActionLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/actionLogs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Start: start,
        End: end,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageSize: 页容量，默认10，取值范围(1 - 100) (Optional)
 * param pageNumber: 页序号，默认值1，不能小于1 (Optional)
 * param start: 起始时间，格式：UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 结束时间，格式：UTC时间例如2017-11-10T23:00:00Z (Required)
 * param keyWord: 日志模糊匹配的关键词 (Optional)
 * param success: 操作结果 false->失败 true->成功 (Optional)
 * param actionType: 日志类型，支持的类型有：CREATE_ZONE、DELETE_ZONE、LOCK_ZONE、CREATE_RR、MODIFY_RR、DELETE_RR、SET_RR_STATUS、RETRY_RECURSE_ZONE (Optional)
 */
func NewDescribeActionLogsRequestWithAllParams(
    regionId string,
    pageSize *int,
    pageNumber *int,
    start string,
    end string,
    keyWord *string,
    success *bool,
    actionType *string,
) *DescribeActionLogsRequest {

    return &DescribeActionLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/actionLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageSize: pageSize,
        PageNumber: pageNumber,
        Start: start,
        End: end,
        KeyWord: keyWord,
        Success: success,
        ActionType: actionType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeActionLogsRequestWithoutParam() *DescribeActionLogsRequest {

    return &DescribeActionLogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/actionLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeActionLogsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageSize: 页容量，默认10，取值范围(1 - 100)(Optional) */
func (r *DescribeActionLogsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageNumber: 页序号，默认值1，不能小于1(Optional) */
func (r *DescribeActionLogsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param start: 起始时间，格式：UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *DescribeActionLogsRequest) SetStart(start string) {
    r.Start = start
}

/* param end: 结束时间，格式：UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *DescribeActionLogsRequest) SetEnd(end string) {
    r.End = end
}

/* param keyWord: 日志模糊匹配的关键词(Optional) */
func (r *DescribeActionLogsRequest) SetKeyWord(keyWord string) {
    r.KeyWord = &keyWord
}

/* param success: 操作结果 false->失败 true->成功(Optional) */
func (r *DescribeActionLogsRequest) SetSuccess(success bool) {
    r.Success = &success
}

/* param actionType: 日志类型，支持的类型有：CREATE_ZONE、DELETE_ZONE、LOCK_ZONE、CREATE_RR、MODIFY_RR、DELETE_RR、SET_RR_STATUS、RETRY_RECURSE_ZONE(Optional) */
func (r *DescribeActionLogsRequest) SetActionType(actionType string) {
    r.ActionType = &actionType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeActionLogsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeActionLogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeActionLogsResult `json:"result"`
}

type DescribeActionLogsResult struct {
    DataList []privatezone.DescribeActionLogsRes `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}