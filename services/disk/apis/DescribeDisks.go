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
    disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeDisksRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* Tag筛选条件 (Optional) */
    Tags []disk.TagFilter `json:"tags"`

    /* diskId - 云硬盘ID，精确匹配，支持多个
diskType - 云硬盘类型，精确匹配，支持多个，取值为 ssd,premium-hdd,ssd.io1,ssd.gp1,hdd.std1
instanceId - 云硬盘所挂载主机的ID，精确匹配，支持多个
instanceType - 云硬盘所挂载主机的类型，精确匹配，支持多个
status - 云硬盘状态，精确匹配，支持多个 
az - 可用区，精确匹配，支持多个
name - 云硬盘名称，模糊匹配，支持单个
multiAttach - 云硬盘是否多点挂载，精确匹配，支持单个
encrypted - 云硬盘是否加密，精确匹配，支持单个
policyId - 绑定policyId的云硬盘，精确匹配，支持多个
notPolicyId - 未绑定policyId的云硬盘，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDisksRequest(
    regionId string,
) *DescribeDisksRequest {

	return &DescribeDisksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param tags: Tag筛选条件 (Optional)
 * param filters: diskId - 云硬盘ID，精确匹配，支持多个
diskType - 云硬盘类型，精确匹配，支持多个，取值为 ssd,premium-hdd,ssd.io1,ssd.gp1,hdd.std1
instanceId - 云硬盘所挂载主机的ID，精确匹配，支持多个
instanceType - 云硬盘所挂载主机的类型，精确匹配，支持多个
status - 云硬盘状态，精确匹配，支持多个 
az - 可用区，精确匹配，支持多个
name - 云硬盘名称，模糊匹配，支持单个
multiAttach - 云硬盘是否多点挂载，精确匹配，支持单个
encrypted - 云硬盘是否加密，精确匹配，支持单个
policyId - 绑定policyId的云硬盘，精确匹配，支持多个
notPolicyId - 未绑定policyId的云硬盘，精确匹配，支持多个
 (Optional)
 */
func NewDescribeDisksRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    tags []disk.TagFilter,
    filters []common.Filter,
) *DescribeDisksRequest {

    return &DescribeDisksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Tags: tags,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDisksRequestWithoutParam() *DescribeDisksRequest {

    return &DescribeDisksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeDisksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *DescribeDisksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeDisksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param tags: Tag筛选条件(Optional) */
func (r *DescribeDisksRequest) SetTags(tags []disk.TagFilter) {
    r.Tags = tags
}

/* param filters: diskId - 云硬盘ID，精确匹配，支持多个
diskType - 云硬盘类型，精确匹配，支持多个，取值为 ssd,premium-hdd,ssd.io1,ssd.gp1,hdd.std1
instanceId - 云硬盘所挂载主机的ID，精确匹配，支持多个
instanceType - 云硬盘所挂载主机的类型，精确匹配，支持多个
status - 云硬盘状态，精确匹配，支持多个 
az - 可用区，精确匹配，支持多个
name - 云硬盘名称，模糊匹配，支持单个
multiAttach - 云硬盘是否多点挂载，精确匹配，支持单个
encrypted - 云硬盘是否加密，精确匹配，支持单个
policyId - 绑定policyId的云硬盘，精确匹配，支持多个
notPolicyId - 未绑定policyId的云硬盘，精确匹配，支持多个
(Optional) */
func (r *DescribeDisksRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDisksRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDisksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDisksResult `json:"result"`
}

type DescribeDisksResult struct {
    Disks []disk.Disk `json:"disks"`
    TotalCount int `json:"totalCount"`
}