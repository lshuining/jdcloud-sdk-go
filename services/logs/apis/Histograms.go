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
    logs "github.com/lshuining/jdcloud-sdk-go/services/logs/models"
)

type HistogramsRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集ID  */
    LogsetUID string `json:"logsetUID"`

    /* 日志主题ID  */
    LogtopicUID string `json:"logtopicUID"`

    /* "preview"表示预览, "fulltext"表示全文检索, "advance"表示按照搜索语句检索  */
    Action string `json:"action"`

    /* Base64编码的搜索表达式, (Optional) */
    Expr *string `json:"expr"`

    /* 搜索关键字大小写敏感， 默认false (Optional) */
    CaseSensitive *bool `json:"caseSensitive"`

    /* 开始时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集ID (Required)
 * param logtopicUID: 日志主题ID (Required)
 * param action: "preview"表示预览, "fulltext"表示全文检索, "advance"表示按照搜索语句检索 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewHistogramsRequest(
    regionId string,
    logsetUID string,
    logtopicUID string,
    action string,
) *HistogramsRequest {

	return &HistogramsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/histograms",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        Action: action,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集ID (Required)
 * param logtopicUID: 日志主题ID (Required)
 * param action: "preview"表示预览, "fulltext"表示全文检索, "advance"表示按照搜索语句检索 (Required)
 * param expr: Base64编码的搜索表达式, (Optional)
 * param caseSensitive: 搜索关键字大小写敏感， 默认false (Optional)
 * param startTime: 开始时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填 (Optional)
 * param endTime: 结束时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填 (Optional)
 */
func NewHistogramsRequestWithAllParams(
    regionId string,
    logsetUID string,
    logtopicUID string,
    action string,
    expr *string,
    caseSensitive *bool,
    startTime *string,
    endTime *string,
) *HistogramsRequest {

    return &HistogramsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/histograms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        Action: action,
        Expr: expr,
        CaseSensitive: caseSensitive,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewHistogramsRequestWithoutParam() *HistogramsRequest {

    return &HistogramsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/histograms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *HistogramsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUID: 日志集ID(Required) */
func (r *HistogramsRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}

/* param logtopicUID: 日志主题ID(Required) */
func (r *HistogramsRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}

/* param action: "preview"表示预览, "fulltext"表示全文检索, "advance"表示按照搜索语句检索(Required) */
func (r *HistogramsRequest) SetAction(action string) {
    r.Action = action
}

/* param expr: Base64编码的搜索表达式,(Optional) */
func (r *HistogramsRequest) SetExpr(expr string) {
    r.Expr = &expr
}

/* param caseSensitive: 搜索关键字大小写敏感， 默认false(Optional) */
func (r *HistogramsRequest) SetCaseSensitive(caseSensitive bool) {
    r.CaseSensitive = &caseSensitive
}

/* param startTime: 开始时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填(Optional) */
func (r *HistogramsRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填(Optional) */
func (r *HistogramsRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r HistogramsRequest) GetRegionId() string {
    return r.RegionId
}

type HistogramsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result HistogramsResult `json:"result"`
}

type HistogramsResult struct {
    Count int64 `json:"count"`
    Data []interface{} `json:"data"`
    Progress string `json:"progress"`
    SearchFields logs.SearchFields `json:"searchFields"`
    Total int64 `json:"total"`
}