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

type BindMetricsRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 自定义的监控项id数组 (Optional) */
    MetricIds []string `json:"metricIds"`

    /* 展示类型，取值为： real_time 表示实时监控页面, monitor 表示性能趋势页面 market 表示监控大盘 (Optional) */
    PanelType *string `json:"panelType"`

    /* RDS 实例ID，唯一标识一个RDS实例。如果全部实例生效，可以传all (Optional) */
    Gid *string `json:"gid"`
}

/*
 * param regionId: 地域代码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBindMetricsRequest(
    regionId string,
) *BindMetricsRequest {

	return &BindMetricsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bindMetrics",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param metricIds: 自定义的监控项id数组 (Optional)
 * param panelType: 展示类型，取值为： real_time 表示实时监控页面, monitor 表示性能趋势页面 market 表示监控大盘 (Optional)
 * param gid: RDS 实例ID，唯一标识一个RDS实例。如果全部实例生效，可以传all (Optional)
 */
func NewBindMetricsRequestWithAllParams(
    regionId string,
    metricIds []string,
    panelType *string,
    gid *string,
) *BindMetricsRequest {

    return &BindMetricsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bindMetrics",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        MetricIds: metricIds,
        PanelType: panelType,
        Gid: gid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBindMetricsRequestWithoutParam() *BindMetricsRequest {

    return &BindMetricsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bindMetrics",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *BindMetricsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param metricIds: 自定义的监控项id数组(Optional) */
func (r *BindMetricsRequest) SetMetricIds(metricIds []string) {
    r.MetricIds = metricIds
}

/* param panelType: 展示类型，取值为： real_time 表示实时监控页面, monitor 表示性能趋势页面 market 表示监控大盘(Optional) */
func (r *BindMetricsRequest) SetPanelType(panelType string) {
    r.PanelType = &panelType
}

/* param gid: RDS 实例ID，唯一标识一个RDS实例。如果全部实例生效，可以传all(Optional) */
func (r *BindMetricsRequest) SetGid(gid string) {
    r.Gid = &gid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BindMetricsRequest) GetRegionId() string {
    return r.RegionId
}

type BindMetricsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BindMetricsResult `json:"result"`
}

type BindMetricsResult struct {
}