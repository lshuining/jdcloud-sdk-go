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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeOneDataPointRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 监控项英文标识(id)  */
    Metric string `json:"metric"`

    /* 资源的类型，取值vm, lb, ip, database 等。可用的serviceCode请使用describeServices接口查询  */
    ServiceCode string `json:"serviceCode"`

    /* 资源的维度。serviceCode下可用的dimension请使用describeServices接口查询 (Optional) */
    Dimension *string `json:"dimension"`

    /* 资源的uuid，支持多个resourceId批量查询，每个id用竖线分隔。 如：id1|id2|id3|id4  */
    ResourceId string `json:"resourceId"`

    /* 自定义标签 (Optional) */
    Tags []monitor.TagFilter `json:"tags"`

    /* 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（早于30d时，将被重置为30d）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional) */
    EndTime *string `json:"endTime"`

    /* 查询的时间间隔，最大不超过30天，支持分钟级别,小时级别，天级别，例如：1m、1h、1d (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* 聚合方式：max avg min等,用于不同维度之间聚合 (Optional) */
    AggrType *string `json:"aggrType"`

    /* 聚合方式：max avg min等,用于将维度内一个周期数据聚合为一个点的聚合方式,默认last (Optional) */
    DownAggrType *string `json:"downAggrType"`
}

/*
 * param regionId: 地域 Id (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等。可用的serviceCode请使用describeServices接口查询 (Required)
 * param resourceId: 资源的uuid，支持多个resourceId批量查询，每个id用竖线分隔。 如：id1|id2|id3|id4 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeOneDataPointRequest(
    regionId string,
    metric string,
    serviceCode string,
    resourceId string,
) *DescribeOneDataPointRequest {

	return &DescribeOneDataPointRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/metrics/{metric}/lastDownsample",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        Metric: metric,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等。可用的serviceCode请使用describeServices接口查询 (Required)
 * param dimension: 资源的维度。serviceCode下可用的dimension请使用describeServices接口查询 (Optional)
 * param resourceId: 资源的uuid，支持多个resourceId批量查询，每个id用竖线分隔。 如：id1|id2|id3|id4 (Required)
 * param tags: 自定义标签 (Optional)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（早于30d时，将被重置为30d）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800） (Optional)
 * param timeInterval: 查询的时间间隔，最大不超过30天，支持分钟级别,小时级别，天级别，例如：1m、1h、1d (Optional)
 * param aggrType: 聚合方式：max avg min等,用于不同维度之间聚合 (Optional)
 * param downAggrType: 聚合方式：max avg min等,用于将维度内一个周期数据聚合为一个点的聚合方式,默认last (Optional)
 */
func NewDescribeOneDataPointRequestWithAllParams(
    regionId string,
    metric string,
    serviceCode string,
    dimension *string,
    resourceId string,
    tags []monitor.TagFilter,
    startTime *string,
    endTime *string,
    timeInterval *string,
    aggrType *string,
    downAggrType *string,
) *DescribeOneDataPointRequest {

    return &DescribeOneDataPointRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/metrics/{metric}/lastDownsample",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Metric: metric,
        ServiceCode: serviceCode,
        Dimension: dimension,
        ResourceId: resourceId,
        Tags: tags,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
        AggrType: aggrType,
        DownAggrType: downAggrType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeOneDataPointRequestWithoutParam() *DescribeOneDataPointRequest {

    return &DescribeOneDataPointRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/metrics/{metric}/lastDownsample",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeOneDataPointRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param metric: 监控项英文标识(id)(Required) */
func (r *DescribeOneDataPointRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param serviceCode: 资源的类型，取值vm, lb, ip, database 等。可用的serviceCode请使用describeServices接口查询(Required) */
func (r *DescribeOneDataPointRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param dimension: 资源的维度。serviceCode下可用的dimension请使用describeServices接口查询(Optional) */
func (r *DescribeOneDataPointRequest) SetDimension(dimension string) {
    r.Dimension = &dimension
}

/* param resourceId: 资源的uuid，支持多个resourceId批量查询，每个id用竖线分隔。 如：id1|id2|id3|id4(Required) */
func (r *DescribeOneDataPointRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param tags: 自定义标签(Optional) */
func (r *DescribeOneDataPointRequest) SetTags(tags []monitor.TagFilter) {
    r.Tags = tags
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（早于30d时，将被重置为30d）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）(Optional) */
func (r *DescribeOneDataPointRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）(Optional) */
func (r *DescribeOneDataPointRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param timeInterval: 查询的时间间隔，最大不超过30天，支持分钟级别,小时级别，天级别，例如：1m、1h、1d(Optional) */
func (r *DescribeOneDataPointRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

/* param aggrType: 聚合方式：max avg min等,用于不同维度之间聚合(Optional) */
func (r *DescribeOneDataPointRequest) SetAggrType(aggrType string) {
    r.AggrType = &aggrType
}

/* param downAggrType: 聚合方式：max avg min等,用于将维度内一个周期数据聚合为一个点的聚合方式,默认last(Optional) */
func (r *DescribeOneDataPointRequest) SetDownAggrType(downAggrType string) {
    r.DownAggrType = &downAggrType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeOneDataPointRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeOneDataPointResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeOneDataPointResult `json:"result"`
}

type DescribeOneDataPointResult struct {
    Items []monitor.LastDownsampleRespItem `json:"items"`
}