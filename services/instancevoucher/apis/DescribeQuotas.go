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
    instancevoucher "github.com/lshuining/jdcloud-sdk-go/services/instancevoucher/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeQuotasRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* Filter names: (仅支持eq)
resourceType - 产品类型，精确匹配，支持多个 支持[vm, nativecontainer, pod]
reservedType - 资源分配方式，精确匹配，支持多个 支持[nonReserved]
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeQuotasRequest(
    regionId string,
) *DescribeQuotasRequest {

	return &DescribeQuotasRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quotas",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param filters: Filter names: (仅支持eq)
resourceType - 产品类型，精确匹配，支持多个 支持[vm, nativecontainer, pod]
reservedType - 资源分配方式，精确匹配，支持多个 支持[nonReserved]
 (Optional)
 */
func NewDescribeQuotasRequestWithAllParams(
    regionId string,
    filters []common.Filter,
) *DescribeQuotasRequest {

    return &DescribeQuotasRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeQuotasRequestWithoutParam() *DescribeQuotasRequest {

    return &DescribeQuotasRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DescribeQuotasRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param filters: Filter names: (仅支持eq)
resourceType - 产品类型，精确匹配，支持多个 支持[vm, nativecontainer, pod]
reservedType - 资源分配方式，精确匹配，支持多个 支持[nonReserved]
(Optional) */
func (r *DescribeQuotasRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeQuotasRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeQuotasResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeQuotasResult `json:"result"`
}

type DescribeQuotasResult struct {
    Quotas []instancevoucher.Quota `json:"quotas"`
}