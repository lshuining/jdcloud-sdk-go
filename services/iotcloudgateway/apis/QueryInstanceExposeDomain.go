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

type QueryInstanceExposeDomainRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryInstanceExposeDomainRequest(
    regionId string,
    instanceId string,
) *QueryInstanceExposeDomainRequest {

	return &QueryInstanceExposeDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:queryInstanceExposeDomain",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 */
func NewQueryInstanceExposeDomainRequestWithAllParams(
    regionId string,
    instanceId string,
) *QueryInstanceExposeDomainRequest {

    return &QueryInstanceExposeDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:queryInstanceExposeDomain",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryInstanceExposeDomainRequestWithoutParam() *QueryInstanceExposeDomainRequest {

    return &QueryInstanceExposeDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:queryInstanceExposeDomain",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *QueryInstanceExposeDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *QueryInstanceExposeDomainRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryInstanceExposeDomainRequest) GetRegionId() string {
    return r.RegionId
}

type QueryInstanceExposeDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryInstanceExposeDomainResult `json:"result"`
}

type QueryInstanceExposeDomainResult struct {
    Iotgwd string `json:"iotgwd"`
    Iotgwu string `json:"iotgwu"`
}