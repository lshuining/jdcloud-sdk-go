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

type DescribeServicesRequest struct {

    core.JDCloudRequest

    /* 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则 (Optional) */
    Filters []monitor.Filter `json:"filters"`

    /* 要查询的产品线类型   0：all    1：资源监控   2：其它   默认：1。若指定了查询的serviceCode，则忽略该参数 (Optional) */
    ProductType *int `json:"productType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeServicesRequest(
) *DescribeServicesRequest {

	return &DescribeServicesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/services",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param filters: 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则 (Optional)
 * param productType: 要查询的产品线类型   0：all    1：资源监控   2：其它   默认：1。若指定了查询的serviceCode，则忽略该参数 (Optional)
 */
func NewDescribeServicesRequestWithAllParams(
    filters []monitor.Filter,
    productType *int,
) *DescribeServicesRequest {

    return &DescribeServicesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/services",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        Filters: filters,
        ProductType: productType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeServicesRequestWithoutParam() *DescribeServicesRequest {

    return &DescribeServicesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/services",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param filters: 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则(Optional) */
func (r *DescribeServicesRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

/* param productType: 要查询的产品线类型   0：all    1：资源监控   2：其它   默认：1。若指定了查询的serviceCode，则忽略该参数(Optional) */
func (r *DescribeServicesRequest) SetProductType(productType int) {
    r.ProductType = &productType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeServicesRequest) GetRegionId() string {
    return ""
}

type DescribeServicesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeServicesResult `json:"result"`
}

type DescribeServicesResult struct {
    Services []monitor.ServiceInfoV2 `json:"services"`
}