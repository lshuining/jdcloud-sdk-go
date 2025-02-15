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
    iotcore "github.com/lshuining/jdcloud-sdk-go/services/iotcore/models"
)

type DescribeDeviceRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    InstanceId string `json:"instanceId"`

    /* 设备ID  */
    DeviceId string `json:"deviceId"`
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceId: 设备ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeviceRequest(
    regionId string,
    instanceId string,
    deviceId string,
) *DescribeDeviceRequest {

	return &DescribeDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:describe",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceId: deviceId,
	}
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceId: 设备ID (Required)
 */
func NewDescribeDeviceRequestWithAllParams(
    regionId string,
    instanceId string,
    deviceId string,
) *DescribeDeviceRequest {

    return &DescribeDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:describe",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceId: deviceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeviceRequestWithoutParam() *DescribeDeviceRequest {

    return &DescribeDeviceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:describe",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *DescribeDeviceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例Id(Required) */
func (r *DescribeDeviceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param deviceId: 设备ID(Required) */
func (r *DescribeDeviceRequest) SetDeviceId(deviceId string) {
    r.DeviceId = deviceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeviceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeviceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeviceResult `json:"result"`
}

type DescribeDeviceResult struct {
    DeviceInfoVO iotcore.DeviceInfoVO `json:"deviceInfoVO"`
}