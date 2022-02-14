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
    logs "github.com/jdcloud-api/jdcloud-sdk-go/services/logs/models"
)

type SearchRequest struct {

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

    /* 页数。 最小为1，最大为99 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页个数。默认为10，最大100 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 返回排序,不填或者为空，默认为desc，"asc":按照时间正序返回结果，"desc":按照时间倒序返回结果 (Optional) */
    Sort *string `json:"sort"`

    /* 指定返回字段，只对系统日志生效，不填默认按照产品线配置返回字段，Name支持：key，Values填入返回字段 (Optional) */
    Filters []logs.Filter `json:"filters"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集ID (Required)
 * param logtopicUID: 日志主题ID (Required)
 * param action: "preview"表示预览, "fulltext"表示全文检索, "advance"表示按照搜索语句检索 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSearchRequest(
    regionId string,
    logsetUID string,
    logtopicUID string,
    action string,
) *SearchRequest {

	return &SearchRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/search",
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
 * param pageNumber: 页数。 最小为1，最大为99 (Optional)
 * param pageSize: 每页个数。默认为10，最大100 (Optional)
 * param sort: 返回排序,不填或者为空，默认为desc，"asc":按照时间正序返回结果，"desc":按照时间倒序返回结果 (Optional)
 * param filters: 指定返回字段，只对系统日志生效，不填默认按照产品线配置返回字段，Name支持：key，Values填入返回字段 (Optional)
 */
func NewSearchRequestWithAllParams(
    regionId string,
    logsetUID string,
    logtopicUID string,
    action string,
    expr *string,
    caseSensitive *bool,
    startTime *string,
    endTime *string,
    pageNumber *int,
    pageSize *int,
    sort *string,
    filters []logs.Filter,
) *SearchRequest {

    return &SearchRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/search",
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
        PageNumber: pageNumber,
        PageSize: pageSize,
        Sort: sort,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSearchRequestWithoutParam() *SearchRequest {

    return &SearchRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/search",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *SearchRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUID: 日志集ID(Required) */
func (r *SearchRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}

/* param logtopicUID: 日志主题ID(Required) */
func (r *SearchRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}

/* param action: "preview"表示预览, "fulltext"表示全文检索, "advance"表示按照搜索语句检索(Required) */
func (r *SearchRequest) SetAction(action string) {
    r.Action = action
}

/* param expr: Base64编码的搜索表达式,(Optional) */
func (r *SearchRequest) SetExpr(expr string) {
    r.Expr = &expr
}

/* param caseSensitive: 搜索关键字大小写敏感， 默认false(Optional) */
func (r *SearchRequest) SetCaseSensitive(caseSensitive bool) {
    r.CaseSensitive = &caseSensitive
}

/* param startTime: 开始时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填(Optional) */
func (r *SearchRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间。格式 “YYYY-MM-DDThh:mm:ssTZD”, 比如 “2018-11-09T15:34:46+0800”.当action != preview时，必填(Optional) */
func (r *SearchRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param pageNumber: 页数。 最小为1，最大为99(Optional) */
func (r *SearchRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页个数。默认为10，最大100(Optional) */
func (r *SearchRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param sort: 返回排序,不填或者为空，默认为desc，"asc":按照时间正序返回结果，"desc":按照时间倒序返回结果(Optional) */
func (r *SearchRequest) SetSort(sort string) {
    r.Sort = &sort
}

/* param filters: 指定返回字段，只对系统日志生效，不填默认按照产品线配置返回字段，Name支持：key，Values填入返回字段(Optional) */
func (r *SearchRequest) SetFilters(filters []logs.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SearchRequest) GetRegionId() string {
    return r.RegionId
}

type SearchResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SearchResult `json:"result"`
}

type SearchResult struct {
    Data []interface{} `json:"data"`
    SearchFields logs.SearchFields `json:"searchFields"`
    Total int64 `json:"total"`
}