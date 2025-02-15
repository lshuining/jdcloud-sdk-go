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
)

type UpdateAlarmRequest struct {

    core.JDCloudRequest

    /* 报警规则ID  */
    AlarmId string `json:"alarmId"`

    /* 规则名称 (Optional) */
    Name *string `json:"name"`

    /* 监控项，bandwidthTrafficIn:上行实时流量 bandwidthTrafficOut:下行实时流量 (Optional) */
    Metric *string `json:"metric"`

    /* 统计周期（单位：分钟） (Optional) */
    Period *int `json:"period"`

    /* 统计方法：平均值=avg、最大值=max、最小值=min (Optional) */
    StatisticMethod *string `json:"statisticMethod"`

    /* 计算方式 >=、>、<、<=、=、！= (Optional) */
    Operator *string `json:"operator"`

    /* 阈值 (Optional) */
    Threshold *float64 `json:"threshold"`

    /* 连续多少次后报警 (Optional) */
    Times *int `json:"times"`

    /* 通知周期 单位：小时 (Optional) */
    NoticePeriod *int `json:"noticePeriod"`

    /* 规则状态 disabled:禁用 enabled:启用 (Optional) */
    Status *string `json:"status"`

    /* 通知方式 all:全部 sms：短信 email:邮件 (Optional) */
    NoticeMethod *string `json:"noticeMethod"`

    /* 通知对象用户ID,若多个用逗号分隔 (Optional) */
    UserId *string `json:"userId"`

    /* 通知对象组ID (Optional) */
    GroupId *string `json:"groupId"`
}

/*
 * param alarmId: 报警规则ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAlarmRequest(
    alarmId string,
) *UpdateAlarmRequest {

	return &UpdateAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarms/{alarmId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        AlarmId: alarmId,
	}
}

/*
 * param alarmId: 报警规则ID (Required)
 * param name: 规则名称 (Optional)
 * param metric: 监控项，bandwidthTrafficIn:上行实时流量 bandwidthTrafficOut:下行实时流量 (Optional)
 * param period: 统计周期（单位：分钟） (Optional)
 * param statisticMethod: 统计方法：平均值=avg、最大值=max、最小值=min (Optional)
 * param operator: 计算方式 >=、>、<、<=、=、！= (Optional)
 * param threshold: 阈值 (Optional)
 * param times: 连续多少次后报警 (Optional)
 * param noticePeriod: 通知周期 单位：小时 (Optional)
 * param status: 规则状态 disabled:禁用 enabled:启用 (Optional)
 * param noticeMethod: 通知方式 all:全部 sms：短信 email:邮件 (Optional)
 * param userId: 通知对象用户ID,若多个用逗号分隔 (Optional)
 * param groupId: 通知对象组ID (Optional)
 */
func NewUpdateAlarmRequestWithAllParams(
    alarmId string,
    name *string,
    metric *string,
    period *int,
    statisticMethod *string,
    operator *string,
    threshold *float64,
    times *int,
    noticePeriod *int,
    status *string,
    noticeMethod *string,
    userId *string,
    groupId *string,
) *UpdateAlarmRequest {

    return &UpdateAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarms/{alarmId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        AlarmId: alarmId,
        Name: name,
        Metric: metric,
        Period: period,
        StatisticMethod: statisticMethod,
        Operator: operator,
        Threshold: threshold,
        Times: times,
        NoticePeriod: noticePeriod,
        Status: status,
        NoticeMethod: noticeMethod,
        UserId: userId,
        GroupId: groupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAlarmRequestWithoutParam() *UpdateAlarmRequest {

    return &UpdateAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarms/{alarmId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param alarmId: 报警规则ID(Required) */
func (r *UpdateAlarmRequest) SetAlarmId(alarmId string) {
    r.AlarmId = alarmId
}

/* param name: 规则名称(Optional) */
func (r *UpdateAlarmRequest) SetName(name string) {
    r.Name = &name
}

/* param metric: 监控项，bandwidthTrafficIn:上行实时流量 bandwidthTrafficOut:下行实时流量(Optional) */
func (r *UpdateAlarmRequest) SetMetric(metric string) {
    r.Metric = &metric
}

/* param period: 统计周期（单位：分钟）(Optional) */
func (r *UpdateAlarmRequest) SetPeriod(period int) {
    r.Period = &period
}

/* param statisticMethod: 统计方法：平均值=avg、最大值=max、最小值=min(Optional) */
func (r *UpdateAlarmRequest) SetStatisticMethod(statisticMethod string) {
    r.StatisticMethod = &statisticMethod
}

/* param operator: 计算方式 >=、>、<、<=、=、！=(Optional) */
func (r *UpdateAlarmRequest) SetOperator(operator string) {
    r.Operator = &operator
}

/* param threshold: 阈值(Optional) */
func (r *UpdateAlarmRequest) SetThreshold(threshold float64) {
    r.Threshold = &threshold
}

/* param times: 连续多少次后报警(Optional) */
func (r *UpdateAlarmRequest) SetTimes(times int) {
    r.Times = &times
}

/* param noticePeriod: 通知周期 单位：小时(Optional) */
func (r *UpdateAlarmRequest) SetNoticePeriod(noticePeriod int) {
    r.NoticePeriod = &noticePeriod
}

/* param status: 规则状态 disabled:禁用 enabled:启用(Optional) */
func (r *UpdateAlarmRequest) SetStatus(status string) {
    r.Status = &status
}

/* param noticeMethod: 通知方式 all:全部 sms：短信 email:邮件(Optional) */
func (r *UpdateAlarmRequest) SetNoticeMethod(noticeMethod string) {
    r.NoticeMethod = &noticeMethod
}

/* param userId: 通知对象用户ID,若多个用逗号分隔(Optional) */
func (r *UpdateAlarmRequest) SetUserId(userId string) {
    r.UserId = &userId
}

/* param groupId: 通知对象组ID(Optional) */
func (r *UpdateAlarmRequest) SetGroupId(groupId string) {
    r.GroupId = &groupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAlarmRequest) GetRegionId() string {
    return ""
}

type UpdateAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAlarmResult `json:"result"`
}

type UpdateAlarmResult struct {
    Success bool `json:"success"`
}