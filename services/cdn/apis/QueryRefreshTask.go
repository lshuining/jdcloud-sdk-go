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
    cdn "github.com/lshuining/jdcloud-sdk-go/services/cdn/models"
)

type QueryRefreshTaskRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* url或者目录的模糊查询关键字 (Optional) */
    Keyword *string `json:"keyword"`

    /* 任务id (Optional) */
    TaskId *string `json:"taskId"`

    /* null (Optional) */
    TaskStatus *string `json:"taskStatus"`

    /* null (Optional) */
    TaskType *string `json:"taskType"`

    /* 分页页数,默认值1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页页面大小,默认值50 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 查询的账号范围 (Optional) */
    AccountType *string `json:"accountType"`

    /* 查询的子账号，多个用逗号隔开 (Optional) */
    SubUsers *string `json:"subUsers"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRefreshTaskRequest(
) *QueryRefreshTaskRequest {

	return &QueryRefreshTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/task",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param keyword: url或者目录的模糊查询关键字 (Optional)
 * param taskId: 任务id (Optional)
 * param taskStatus: null (Optional)
 * param taskType: null (Optional)
 * param pageNumber: 分页页数,默认值1 (Optional)
 * param pageSize: 分页页面大小,默认值50 (Optional)
 * param accountType: 查询的账号范围 (Optional)
 * param subUsers: 查询的子账号，多个用逗号隔开 (Optional)
 */
func NewQueryRefreshTaskRequestWithAllParams(
    startTime *string,
    endTime *string,
    keyword *string,
    taskId *string,
    taskStatus *string,
    taskType *string,
    pageNumber *int,
    pageSize *int,
    accountType *string,
    subUsers *string,
) *QueryRefreshTaskRequest {

    return &QueryRefreshTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/task",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Keyword: keyword,
        TaskId: taskId,
        TaskStatus: taskStatus,
        TaskType: taskType,
        PageNumber: pageNumber,
        PageSize: pageSize,
        AccountType: accountType,
        SubUsers: subUsers,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRefreshTaskRequestWithoutParam() *QueryRefreshTaskRequest {

    return &QueryRefreshTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/task",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryRefreshTaskRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryRefreshTaskRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param keyword: url或者目录的模糊查询关键字(Optional) */
func (r *QueryRefreshTaskRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}

/* param taskId: 任务id(Optional) */
func (r *QueryRefreshTaskRequest) SetTaskId(taskId string) {
    r.TaskId = &taskId
}

/* param taskStatus: null(Optional) */
func (r *QueryRefreshTaskRequest) SetTaskStatus(taskStatus string) {
    r.TaskStatus = &taskStatus
}

/* param taskType: null(Optional) */
func (r *QueryRefreshTaskRequest) SetTaskType(taskType string) {
    r.TaskType = &taskType
}

/* param pageNumber: 分页页数,默认值1(Optional) */
func (r *QueryRefreshTaskRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页页面大小,默认值50(Optional) */
func (r *QueryRefreshTaskRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param accountType: 查询的账号范围(Optional) */
func (r *QueryRefreshTaskRequest) SetAccountType(accountType string) {
    r.AccountType = &accountType
}

/* param subUsers: 查询的子账号，多个用逗号隔开(Optional) */
func (r *QueryRefreshTaskRequest) SetSubUsers(subUsers string) {
    r.SubUsers = &subUsers
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRefreshTaskRequest) GetRegionId() string {
    return ""
}

type QueryRefreshTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRefreshTaskResult `json:"result"`
}

type QueryRefreshTaskResult struct {
    Total int `json:"total"`
    Tasks []cdn.RefreshTask `json:"tasks"`
}