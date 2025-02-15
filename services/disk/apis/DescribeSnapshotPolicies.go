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
    disk "github.com/lshuining/jdcloud-sdk-go/services/disk/models"
)

type DescribeSnapshotPoliciesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 策略名称,默认为空 (Optional) */
    Name *string `json:"name"`

    /* 策略ID,默认为空 (Optional) */
    PolicyId []string `json:"policyId"`

    /* 策略状态。1: 启用 2：禁用 (Optional) */
    Status []int `json:"status"`

    /* 排序字段，只支持create_time和update_time字段 (Optional) */
    Order *disk.OrderItem `json:"order"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSnapshotPoliciesRequest(
    regionId string,
) *DescribeSnapshotPoliciesRequest {

	return &DescribeSnapshotPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/snapshotPolicies:describe",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param name: 策略名称,默认为空 (Optional)
 * param policyId: 策略ID,默认为空 (Optional)
 * param status: 策略状态。1: 启用 2：禁用 (Optional)
 * param order: 排序字段，只支持create_time和update_time字段 (Optional)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 */
func NewDescribeSnapshotPoliciesRequestWithAllParams(
    regionId string,
    name *string,
    policyId []string,
    status []int,
    order *disk.OrderItem,
    pageNumber *int,
    pageSize *int,
) *DescribeSnapshotPoliciesRequest {

    return &DescribeSnapshotPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshotPolicies:describe",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        PolicyId: policyId,
        Status: status,
        Order: order,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSnapshotPoliciesRequestWithoutParam() *DescribeSnapshotPoliciesRequest {

    return &DescribeSnapshotPoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshotPolicies:describe",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeSnapshotPoliciesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 策略名称,默认为空(Optional) */
func (r *DescribeSnapshotPoliciesRequest) SetName(name string) {
    r.Name = &name
}

/* param policyId: 策略ID,默认为空(Optional) */
func (r *DescribeSnapshotPoliciesRequest) SetPolicyId(policyId []string) {
    r.PolicyId = policyId
}

/* param status: 策略状态。1: 启用 2：禁用(Optional) */
func (r *DescribeSnapshotPoliciesRequest) SetStatus(status []int) {
    r.Status = status
}

/* param order: 排序字段，只支持create_time和update_time字段(Optional) */
func (r *DescribeSnapshotPoliciesRequest) SetOrder(order *disk.OrderItem) {
    r.Order = order
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *DescribeSnapshotPoliciesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeSnapshotPoliciesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSnapshotPoliciesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSnapshotPoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSnapshotPoliciesResult `json:"result"`
}

type DescribeSnapshotPoliciesResult struct {
    Policies []disk.SnapshotPolicy `json:"policies"`
    TotalCount int `json:"totalCount"`
}