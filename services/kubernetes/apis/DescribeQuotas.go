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
    kubernetes "github.com/lshuining/jdcloud-sdk-go/services/kubernetes/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeQuotasRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* resourceTypes - 资源类型，暂时只支持[kubernetes]
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
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
 * param regionId: Region ID (Required)
 * param filters: resourceTypes - 资源类型，暂时只支持[kubernetes]
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

/* param regionId: Region ID(Required) */
func (r *DescribeQuotasRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param filters: resourceTypes - 资源类型，暂时只支持[kubernetes]
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
    Quotas []kubernetes.Quota `json:"quotas"`
}