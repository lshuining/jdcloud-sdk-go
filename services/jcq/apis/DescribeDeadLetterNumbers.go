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
    jcq "github.com/lshuining/jdcloud-sdk-go/services/jcq/models"
)

type DescribeDeadLetterNumbersRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* consumerGroupId为空则显示该用户所有订阅关系里的死信数量 (Optional) */
    ConsumerGroupId *string `json:"consumerGroupId"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeadLetterNumbersRequest(
    regionId string,
) *DescribeDeadLetterNumbersRequest {

	return &DescribeDeadLetterNumbersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deadLetterNumbers",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param consumerGroupId: consumerGroupId为空则显示该用户所有订阅关系里的死信数量 (Optional)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 */
func NewDescribeDeadLetterNumbersRequestWithAllParams(
    regionId string,
    consumerGroupId *string,
    pageNumber *int,
    pageSize *int,
) *DescribeDeadLetterNumbersRequest {

    return &DescribeDeadLetterNumbersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deadLetterNumbers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ConsumerGroupId: consumerGroupId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeadLetterNumbersRequestWithoutParam() *DescribeDeadLetterNumbersRequest {

    return &DescribeDeadLetterNumbersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deadLetterNumbers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *DescribeDeadLetterNumbersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param consumerGroupId: consumerGroupId为空则显示该用户所有订阅关系里的死信数量(Optional) */
func (r *DescribeDeadLetterNumbersRequest) SetConsumerGroupId(consumerGroupId string) {
    r.ConsumerGroupId = &consumerGroupId
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeDeadLetterNumbersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeDeadLetterNumbersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeadLetterNumbersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeadLetterNumbersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeadLetterNumbersResult `json:"result"`
}

type DescribeDeadLetterNumbersResult struct {
    DeadLetterNumbers []jcq.DeadLetterNumber `json:"deadLetterNumbers"`
    TotalCount int `json:"totalCount"`
}