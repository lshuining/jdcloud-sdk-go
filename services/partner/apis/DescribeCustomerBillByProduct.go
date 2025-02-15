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
    partner "github.com/lshuining/jdcloud-sdk-go/services/partner/models"
)

type DescribeCustomerBillByProductRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* pin (Optional) */
    Pin *string `json:"pin"`

    /* 按月查询开始时间（yyyy-MM-dd）,不可跨月  */
    StartTime string `json:"startTime"`

    /* 按月查询结束时间（yyyy-MM-dd）,不可跨月  */
    EndTime string `json:"endTime"`

    /* 每页条数,不超过100  */
    PageSize int `json:"pageSize"`

    /* 第几页  */
    PageIndex int `json:"pageIndex"`
}

/*
 * param regionId:  (Required)
 * param startTime: 按月查询开始时间（yyyy-MM-dd）,不可跨月 (Required)
 * param endTime: 按月查询结束时间（yyyy-MM-dd）,不可跨月 (Required)
 * param pageSize: 每页条数,不超过100 (Required)
 * param pageIndex: 第几页 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomerBillByProductRequest(
    regionId string,
    startTime string,
    endTime string,
    pageSize int,
    pageIndex int,
) *DescribeCustomerBillByProductRequest {

	return &DescribeCustomerBillByProductRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accountingBill:describeCustomerBillByProduct",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        PageSize: pageSize,
        PageIndex: pageIndex,
	}
}

/*
 * param regionId:  (Required)
 * param pin: pin (Optional)
 * param startTime: 按月查询开始时间（yyyy-MM-dd）,不可跨月 (Required)
 * param endTime: 按月查询结束时间（yyyy-MM-dd）,不可跨月 (Required)
 * param pageSize: 每页条数,不超过100 (Required)
 * param pageIndex: 第几页 (Required)
 */
func NewDescribeCustomerBillByProductRequestWithAllParams(
    regionId string,
    pin *string,
    startTime string,
    endTime string,
    pageSize int,
    pageIndex int,
) *DescribeCustomerBillByProductRequest {

    return &DescribeCustomerBillByProductRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accountingBill:describeCustomerBillByProduct",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        StartTime: startTime,
        EndTime: endTime,
        PageSize: pageSize,
        PageIndex: pageIndex,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCustomerBillByProductRequestWithoutParam() *DescribeCustomerBillByProductRequest {

    return &DescribeCustomerBillByProductRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accountingBill:describeCustomerBillByProduct",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *DescribeCustomerBillByProductRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: pin(Optional) */
func (r *DescribeCustomerBillByProductRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param startTime: 按月查询开始时间（yyyy-MM-dd）,不可跨月(Required) */
func (r *DescribeCustomerBillByProductRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 按月查询结束时间（yyyy-MM-dd）,不可跨月(Required) */
func (r *DescribeCustomerBillByProductRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param pageSize: 每页条数,不超过100(Required) */
func (r *DescribeCustomerBillByProductRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param pageIndex: 第几页(Required) */
func (r *DescribeCustomerBillByProductRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = pageIndex
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomerBillByProductRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCustomerBillByProductResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomerBillByProductResult `json:"result"`
}

type DescribeCustomerBillByProductResult struct {
    Pagination partner.Pagination `json:"pagination"`
    Result []partner.ServiceCodeBill `json:"result"`
}