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
    iotcore "github.com/lshuining/jdcloud-sdk-go/services/iotcore/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type ListProductsWithAdminRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* IoT Engine实例ID信息  */
    InstanceId string `json:"instanceId"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为10，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* productName-产品名称，模糊匹配，支持单个
productKey-产品key，精确匹配，支持单个
productType-产品类型，精确匹配，支持单个
templateName-模板名称，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoT Engine实例ID信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListProductsWithAdminRequest(
    regionId string,
    instanceId string,
) *ListProductsWithAdminRequest {

	return &ListProductsWithAdminRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/products:describeAll",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoT Engine实例ID信息 (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为10，取值范围：[10,100] (Optional)
 * param filters: productName-产品名称，模糊匹配，支持单个
productKey-产品key，精确匹配，支持单个
productType-产品类型，精确匹配，支持单个
templateName-模板名称，精确匹配，支持多个
 (Optional)
 */
func NewListProductsWithAdminRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *ListProductsWithAdminRequest {

    return &ListProductsWithAdminRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/products:describeAll",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListProductsWithAdminRequestWithoutParam() *ListProductsWithAdminRequest {

    return &ListProductsWithAdminRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/products:describeAll",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ListProductsWithAdminRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: IoT Engine实例ID信息(Required) */
func (r *ListProductsWithAdminRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *ListProductsWithAdminRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为10，取值范围：[10,100](Optional) */
func (r *ListProductsWithAdminRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: productName-产品名称，模糊匹配，支持单个
productKey-产品key，精确匹配，支持单个
productType-产品类型，精确匹配，支持单个
templateName-模板名称，精确匹配，支持多个
(Optional) */
func (r *ListProductsWithAdminRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListProductsWithAdminRequest) GetRegionId() string {
    return r.RegionId
}

type ListProductsWithAdminResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListProductsWithAdminResult `json:"result"`
}

type ListProductsWithAdminResult struct {
    Page iotcore.PageinfoVO `json:"page"`
    Products []iotcore.Product `json:"products"`
}