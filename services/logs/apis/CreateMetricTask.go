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

type CreateMetricTaskRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集 UID  */
    LogsetUID string `json:"logsetUID"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`

    /* 聚合函数,支持 count sum max min avg; 配置方式(SettingType) 为 空或visual 时，必填； (Optional) */
    Aggregate *string `json:"aggregate"`

    /* 自定义单位  */
    CustomUnit string `json:"customUnit"`

    /* 查询字段,支持 英文字母 数字 下划线 中划线 点（中文日志原文和各产品线的key）; 配置方式(SettingType) 为 空或visual 时，必填； (Optional) */
    DataField *string `json:"dataField"`

    /* 过滤语法，可以为空 (Optional) */
    FilterContent *string `json:"filterContent"`

    /* 是否打开过滤; 配置方式(SettingType) 为 空或visual 时，必填； (Optional) */
    FilterOpen *string `json:"filterOpen"`

    /* 过滤类型，只能是fulltext和 advance; 配置方式(SettingType) 为 空或visual 时，必填； (Optional) */
    FilterType *string `json:"filterType"`

    /* 时间周期，固定60s  */
    Interval int64 `json:"interval"`

    /* 监控项 , 支持大小写英文字母 下划线 数字 点，且不超过255byte（不支持中划线）; 配置方式(SettingType) 为 空或visual 时，必填； (Optional) */
    Metric *string `json:"metric"`

    /* 监控任务名称,同一个日志主题下唯一，支持中文 大小写英文字母 下划线 中划线 数字，且不超过32字符  */
    Name string `json:"name"`

    /* 配置方式: 可选参数；枚举值 visual，sql；分别代表可视化配置及sql配置方式，传空表示可视化配置； (Optional) */
    SettingType *string `json:"settingType"`

    /*  (Optional) */
    SqlSpec *logs.MetricTaskSqlSpec `json:"sqlSpec"`

    /* 单位  */
    Unit string `json:"unit"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param customUnit: 自定义单位 (Required)
 * param interval: 时间周期，固定60s (Required)
 * param name: 监控任务名称,同一个日志主题下唯一，支持中文 大小写英文字母 下划线 中划线 数字，且不超过32字符 (Required)
 * param unit: 单位 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateMetricTaskRequest(
    regionId string,
    logsetUID string,
    logtopicUID string,
    customUnit string,
    interval int64,
    name string,
    unit string,
) *CreateMetricTaskRequest {

	return &CreateMetricTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/metrictasks",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        CustomUnit: customUnit,
        Interval: interval,
        Name: name,
        Unit: unit,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param aggregate: 聚合函数,支持 count sum max min avg; 配置方式(SettingType) 为 空或visual 时，必填； (Optional)
 * param customUnit: 自定义单位 (Required)
 * param dataField: 查询字段,支持 英文字母 数字 下划线 中划线 点（中文日志原文和各产品线的key）; 配置方式(SettingType) 为 空或visual 时，必填； (Optional)
 * param filterContent: 过滤语法，可以为空 (Optional)
 * param filterOpen: 是否打开过滤; 配置方式(SettingType) 为 空或visual 时，必填； (Optional)
 * param filterType: 过滤类型，只能是fulltext和 advance; 配置方式(SettingType) 为 空或visual 时，必填； (Optional)
 * param interval: 时间周期，固定60s (Required)
 * param metric: 监控项 , 支持大小写英文字母 下划线 数字 点，且不超过255byte（不支持中划线）; 配置方式(SettingType) 为 空或visual 时，必填； (Optional)
 * param name: 监控任务名称,同一个日志主题下唯一，支持中文 大小写英文字母 下划线 中划线 数字，且不超过32字符 (Required)
 * param settingType: 配置方式: 可选参数；枚举值 visual，sql；分别代表可视化配置及sql配置方式，传空表示可视化配置； (Optional)
 * param sqlSpec:  (Optional)
 * param unit: 单位 (Required)
 */
func NewCreateMetricTaskRequestWithAllParams(
    regionId string,
    logsetUID string,
    logtopicUID string,
    aggregate *string,
    customUnit string,
    dataField *string,
    filterContent *string,
    filterOpen *string,
    filterType *string,
    interval int64,
    metric *string,
    name string,
    settingType *string,
    sqlSpec *logs.MetricTaskSqlSpec,
    unit string,
) *CreateMetricTaskRequest {

    return &CreateMetricTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/metrictasks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        Aggregate: aggregate,
        CustomUnit: customUnit,
        DataField: dataField,
        FilterContent: filterContent,
        FilterOpen: filterOpen,
        FilterType: filterType,
        Interval: interval,
        Metric: metric,
        Name: name,
        SettingType: settingType,
        SqlSpec: sqlSpec,
        Unit: unit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateMetricTaskRequestWithoutParam() *CreateMetricTaskRequest {

    return &CreateMetricTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/metrictasks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *CreateMetricTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUID: 日志集 UID(Required) */
func (r *CreateMetricTaskRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}

/* param logtopicUID: 日志主题 UID(Required) */
func (r *CreateMetricTaskRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}

/* param aggregate: 聚合函数,支持 count sum max min avg; 配置方式(SettingType) 为 空或visual 时，必填；(Optional) */
func (r *CreateMetricTaskRequest) SetAggregate(aggregate string) {
    r.Aggregate = &aggregate
}

/* param customUnit: 自定义单位(Required) */
func (r *CreateMetricTaskRequest) SetCustomUnit(customUnit string) {
    r.CustomUnit = customUnit
}

/* param dataField: 查询字段,支持 英文字母 数字 下划线 中划线 点（中文日志原文和各产品线的key）; 配置方式(SettingType) 为 空或visual 时，必填；(Optional) */
func (r *CreateMetricTaskRequest) SetDataField(dataField string) {
    r.DataField = &dataField
}

/* param filterContent: 过滤语法，可以为空(Optional) */
func (r *CreateMetricTaskRequest) SetFilterContent(filterContent string) {
    r.FilterContent = &filterContent
}

/* param filterOpen: 是否打开过滤; 配置方式(SettingType) 为 空或visual 时，必填；(Optional) */
func (r *CreateMetricTaskRequest) SetFilterOpen(filterOpen string) {
    r.FilterOpen = &filterOpen
}

/* param filterType: 过滤类型，只能是fulltext和 advance; 配置方式(SettingType) 为 空或visual 时，必填；(Optional) */
func (r *CreateMetricTaskRequest) SetFilterType(filterType string) {
    r.FilterType = &filterType
}

/* param interval: 时间周期，固定60s(Required) */
func (r *CreateMetricTaskRequest) SetInterval(interval int64) {
    r.Interval = interval
}

/* param metric: 监控项 , 支持大小写英文字母 下划线 数字 点，且不超过255byte（不支持中划线）; 配置方式(SettingType) 为 空或visual 时，必填；(Optional) */
func (r *CreateMetricTaskRequest) SetMetric(metric string) {
    r.Metric = &metric
}

/* param name: 监控任务名称,同一个日志主题下唯一，支持中文 大小写英文字母 下划线 中划线 数字，且不超过32字符(Required) */
func (r *CreateMetricTaskRequest) SetName(name string) {
    r.Name = name
}

/* param settingType: 配置方式: 可选参数；枚举值 visual，sql；分别代表可视化配置及sql配置方式，传空表示可视化配置；(Optional) */
func (r *CreateMetricTaskRequest) SetSettingType(settingType string) {
    r.SettingType = &settingType
}

/* param sqlSpec: (Optional) */
func (r *CreateMetricTaskRequest) SetSqlSpec(sqlSpec *logs.MetricTaskSqlSpec) {
    r.SqlSpec = sqlSpec
}

/* param unit: 单位(Required) */
func (r *CreateMetricTaskRequest) SetUnit(unit string) {
    r.Unit = unit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateMetricTaskRequest) GetRegionId() string {
    return r.RegionId
}

type CreateMetricTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateMetricTaskResult `json:"result"`
}

type CreateMetricTaskResult struct {
    Id string `json:"id"`
    Suc string `json:"suc"`
}