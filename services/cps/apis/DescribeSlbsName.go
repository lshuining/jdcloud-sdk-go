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
)

type DescribeSlbsNameRequest struct {

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
func NewDescribeSlbsNameRequest(
    regionId string,
    loadBalancerId string,
) *DescribeSlbsNameRequest {

	return &DescribeSlbsNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/slbs/{loadBalancerId}:describeSlbsName",
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
func NewDescribeSlbsNameRequestWithAllParams(
    regionId string,
    loadBalancerId string,
) *DescribeSlbsNameRequest {

    return &DescribeSlbsNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs/{loadBalancerId}:describeSlbsName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSlbsNameRequestWithoutParam() *DescribeSlbsNameRequest {

    return &DescribeSlbsNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs/{loadBalancerId}:describeSlbsName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *DescribeSlbsNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerId: 负载均衡实例ID(Required) */
func (r *DescribeSlbsNameRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSlbsNameRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSlbsNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSlbsNameResult `json:"result"`
}

type DescribeSlbsNameResult struct {
    Name string `json:"name"`
}