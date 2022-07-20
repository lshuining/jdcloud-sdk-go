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

type DeleteLoadBalancerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* LB ID  */
    LoadBalancerId string `json:"loadBalancerId"`
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerId: LB ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLoadBalancerRequest(
    regionId string,
    loadBalancerId string,
) *DeleteLoadBalancerRequest {

	return &DeleteLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerId: LB ID (Required)
 */
func NewDeleteLoadBalancerRequestWithAllParams(
    regionId string,
    loadBalancerId string,
) *DeleteLoadBalancerRequest {

    return &DeleteLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLoadBalancerRequestWithoutParam() *DeleteLoadBalancerRequest {

    return &DeleteLoadBalancerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteLoadBalancerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerId: LB ID(Required) */
func (r *DeleteLoadBalancerRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLoadBalancerRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteLoadBalancerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLoadBalancerResult `json:"result"`
}

type DeleteLoadBalancerResult struct {
}