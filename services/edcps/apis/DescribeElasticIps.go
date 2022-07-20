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
    edcps "github.com/lshuining/jdcloud-sdk-go/services/edcps/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeElasticIpsRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 弹性公网IP状态，取值范围：associate、disassociate (Optional) */
    Status *string `json:"status"`

    /* 弹性公网IP是否加入共享带宽，取值范围：yes、no (Optional) */
    HasJoinBandwidthPackage *string `json:"hasJoinBandwidthPackage"`

    /* 支付模式，取值为：prepaid_by_duration表示预付费，postpaid_by_duration表示按配置后付费 (Optional) */
    ChargeMode *string `json:"chargeMode"`

    /* 实例Id (Optional) */
    InstanceId *string `json:"instanceId"`

    /* 子网Id (Optional) */
    SubnetId *string `json:"subnetId"`

    /* elasticIpId - 弹性公网IPID，精确匹配，支持多个<br/>
elasticIp - 弹性公网IP，精确匹配，支持多个<br/>
bandwidthPackageId - 共享带宽ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeElasticIpsRequest(
    regionId string,
) *DescribeElasticIpsRequest {

	return &DescribeElasticIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param status: 弹性公网IP状态，取值范围：associate、disassociate (Optional)
 * param hasJoinBandwidthPackage: 弹性公网IP是否加入共享带宽，取值范围：yes、no (Optional)
 * param chargeMode: 支付模式，取值为：prepaid_by_duration表示预付费，postpaid_by_duration表示按配置后付费 (Optional)
 * param instanceId: 实例Id (Optional)
 * param subnetId: 子网Id (Optional)
 * param filters: elasticIpId - 弹性公网IPID，精确匹配，支持多个<br/>
elasticIp - 弹性公网IP，精确匹配，支持多个<br/>
bandwidthPackageId - 共享带宽ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeElasticIpsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    status *string,
    hasJoinBandwidthPackage *string,
    chargeMode *string,
    instanceId *string,
    subnetId *string,
    filters []common.Filter,
) *DescribeElasticIpsRequest {

    return &DescribeElasticIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Status: status,
        HasJoinBandwidthPackage: hasJoinBandwidthPackage,
        ChargeMode: chargeMode,
        InstanceId: instanceId,
        SubnetId: subnetId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeElasticIpsRequestWithoutParam() *DescribeElasticIpsRequest {

    return &DescribeElasticIpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeElasticIpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeElasticIpsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeElasticIpsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param status: 弹性公网IP状态，取值范围：associate、disassociate(Optional) */
func (r *DescribeElasticIpsRequest) SetStatus(status string) {
    r.Status = &status
}

/* param hasJoinBandwidthPackage: 弹性公网IP是否加入共享带宽，取值范围：yes、no(Optional) */
func (r *DescribeElasticIpsRequest) SetHasJoinBandwidthPackage(hasJoinBandwidthPackage string) {
    r.HasJoinBandwidthPackage = &hasJoinBandwidthPackage
}

/* param chargeMode: 支付模式，取值为：prepaid_by_duration表示预付费，postpaid_by_duration表示按配置后付费(Optional) */
func (r *DescribeElasticIpsRequest) SetChargeMode(chargeMode string) {
    r.ChargeMode = &chargeMode
}

/* param instanceId: 实例Id(Optional) */
func (r *DescribeElasticIpsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}

/* param subnetId: 子网Id(Optional) */
func (r *DescribeElasticIpsRequest) SetSubnetId(subnetId string) {
    r.SubnetId = &subnetId
}

/* param filters: elasticIpId - 弹性公网IPID，精确匹配，支持多个<br/>
elasticIp - 弹性公网IP，精确匹配，支持多个<br/>
bandwidthPackageId - 共享带宽ID，精确匹配，支持多个
(Optional) */
func (r *DescribeElasticIpsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeElasticIpsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeElasticIpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeElasticIpsResult `json:"result"`
}

type DescribeElasticIpsResult struct {
    ElasticIps []edcps.ElasticIp `json:"elasticIps"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}