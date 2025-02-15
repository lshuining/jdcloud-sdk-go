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
    cps "github.com/lshuining/jdcloud-sdk-go/services/cps/models"
)

type DescribeLoadBalancerRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 负载均衡实例ID  */
    LoadBalancerId string `json:"loadBalancerId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param loadBalancerId: 负载均衡实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLoadBalancerRequest(
    regionId string,
    loadBalancerId string,
) *DescribeLoadBalancerRequest {

	return &DescribeLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/slbs/{loadBalancerId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param loadBalancerId: 负载均衡实例ID (Required)
 */
func NewDescribeLoadBalancerRequestWithAllParams(
    regionId string,
    loadBalancerId string,
) *DescribeLoadBalancerRequest {

    return &DescribeLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs/{loadBalancerId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLoadBalancerRequestWithoutParam() *DescribeLoadBalancerRequest {

    return &DescribeLoadBalancerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs/{loadBalancerId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *DescribeLoadBalancerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerId: 负载均衡实例ID(Required) */
func (r *DescribeLoadBalancerRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLoadBalancerRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLoadBalancerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLoadBalancerResult `json:"result"`
}

type DescribeLoadBalancerResult struct {
    LoadBalancer cps.LoadBalancer `json:"loadBalancer"`
}