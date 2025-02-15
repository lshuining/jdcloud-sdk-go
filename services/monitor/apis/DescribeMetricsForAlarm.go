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
    monitor "github.com/lshuining/jdcloud-sdk-go/services/monitor/models"
)

type DescribeMetricsForAlarmRequest struct {

    core.JDCloudRequest

    /* 产品线 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 产品类型，如redis2.8cluster(集群)\redis2.8MS(主从)。当serviceCode与product同时指定时，product优先级更高 (Optional) */
    Product *string `json:"product"`

    /* 产品维度，必须指定serviceCode或product才生效。 (Optional) */
    Dimension *string `json:"dimension"`

    /* metric类型，取值0、1；默认值：0（常规指标，用于控制台创建报警规则）、1（其它） (Optional) */
    MetricType *int `json:"metricType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricsForAlarmRequest(
) *DescribeMetricsForAlarmRequest {

	return &DescribeMetricsForAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/groupAlarms:metrics",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param serviceCode: 产品线 (Optional)
 * param product: 产品类型，如redis2.8cluster(集群)\redis2.8MS(主从)。当serviceCode与product同时指定时，product优先级更高 (Optional)
 * param dimension: 产品维度，必须指定serviceCode或product才生效。 (Optional)
 * param metricType: metric类型，取值0、1；默认值：0（常规指标，用于控制台创建报警规则）、1（其它） (Optional)
 */
func NewDescribeMetricsForAlarmRequestWithAllParams(
    serviceCode *string,
    product *string,
    dimension *string,
    metricType *int,
) *DescribeMetricsForAlarmRequest {

    return &DescribeMetricsForAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarms:metrics",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        ServiceCode: serviceCode,
        Product: product,
        Dimension: dimension,
        MetricType: metricType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricsForAlarmRequestWithoutParam() *DescribeMetricsForAlarmRequest {

    return &DescribeMetricsForAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarms:metrics",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param serviceCode: 产品线(Optional) */
func (r *DescribeMetricsForAlarmRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param product: 产品类型，如redis2.8cluster(集群)\redis2.8MS(主从)。当serviceCode与product同时指定时，product优先级更高(Optional) */
func (r *DescribeMetricsForAlarmRequest) SetProduct(product string) {
    r.Product = &product
}

/* param dimension: 产品维度，必须指定serviceCode或product才生效。(Optional) */
func (r *DescribeMetricsForAlarmRequest) SetDimension(dimension string) {
    r.Dimension = &dimension
}

/* param metricType: metric类型，取值0、1；默认值：0（常规指标，用于控制台创建报警规则）、1（其它）(Optional) */
func (r *DescribeMetricsForAlarmRequest) SetMetricType(metricType int) {
    r.MetricType = &metricType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricsForAlarmRequest) GetRegionId() string {
    return ""
}

type DescribeMetricsForAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricsForAlarmResult `json:"result"`
}

type DescribeMetricsForAlarmResult struct {
    Metrics []monitor.RuleMetricDetail `json:"metrics"`
}