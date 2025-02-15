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
    smartdba "github.com/lshuining/jdcloud-sdk-go/services/smartdba/models"
)

type DescribeMetricValueByGidRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例。  */
    InstanceGid string `json:"instanceGid"`

    /* metric名称  */
    Metric string `json:"metric"`

    /* 查询开始时间，格式为：2006-01-02T15:04:05Z  */
    StartTime string `json:"startTime"`

    /* 查询截止时间，格式为：2006-01-02T15:04:05Z  */
    EndTime string `json:"endTime"`

    /* 实例角色，默认master (Optional) */
    Role *string `json:"role"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例。 (Required)
 * param metric: metric名称 (Required)
 * param startTime: 查询开始时间，格式为：2006-01-02T15:04:05Z (Required)
 * param endTime: 查询截止时间，格式为：2006-01-02T15:04:05Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricValueByGidRequest(
    regionId string,
    instanceGid string,
    metric string,
    startTime string,
    endTime string,
) *DescribeMetricValueByGidRequest {

	return &DescribeMetricValueByGidRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/metrics",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
        Metric: metric,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例。 (Required)
 * param metric: metric名称 (Required)
 * param startTime: 查询开始时间，格式为：2006-01-02T15:04:05Z (Required)
 * param endTime: 查询截止时间，格式为：2006-01-02T15:04:05Z (Required)
 * param role: 实例角色，默认master (Optional)
 */
func NewDescribeMetricValueByGidRequestWithAllParams(
    regionId string,
    instanceGid string,
    metric string,
    startTime string,
    endTime string,
    role *string,
) *DescribeMetricValueByGidRequest {

    return &DescribeMetricValueByGidRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/metrics",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
        Metric: metric,
        StartTime: startTime,
        EndTime: endTime,
        Role: role,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricValueByGidRequestWithoutParam() *DescribeMetricValueByGidRequest {

    return &DescribeMetricValueByGidRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/metrics",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeMetricValueByGidRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: RDS 实例ID，唯一标识一个RDS实例。(Required) */
func (r *DescribeMetricValueByGidRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

/* param metric: metric名称(Required) */
func (r *DescribeMetricValueByGidRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param startTime: 查询开始时间，格式为：2006-01-02T15:04:05Z(Required) */
func (r *DescribeMetricValueByGidRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询截止时间，格式为：2006-01-02T15:04:05Z(Required) */
func (r *DescribeMetricValueByGidRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param role: 实例角色，默认master(Optional) */
func (r *DescribeMetricValueByGidRequest) SetRole(role string) {
    r.Role = &role
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricValueByGidRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMetricValueByGidResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricValueByGidResult `json:"result"`
}

type DescribeMetricValueByGidResult struct {
    ItemData []smartdba.DataPoint `json:"itemData"`
    MetricInfo smartdba.Metric `json:"metricInfo"`
}