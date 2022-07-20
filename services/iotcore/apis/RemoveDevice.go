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

type RemoveDeviceRequest struct {

    core.JDCloudRequest

    /* 设备名称  */
    DeviceName string `json:"deviceName"`

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 产品Key  */
    ProductKey string `json:"productKey"`
}

/*
 * param deviceName: 设备名称 (Required)
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param productKey: 产品Key (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveDeviceRequest(
    deviceName string,
    instanceId string,
    regionId string,
    productKey string,
) *RemoveDeviceRequest {

	return &RemoveDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/device/{deviceName}:delete",
			Method:  "DELETE",
			Header:  nil,
			Version: "v2",
		},
        DeviceName: deviceName,
        InstanceId: instanceId,
        RegionId: regionId,
        ProductKey: productKey,
	}
}

/*
 * param deviceName: 设备名称 (Required)
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param productKey: 产品Key (Required)
 */
func NewRemoveDeviceRequestWithAllParams(
    deviceName string,
    instanceId string,
    regionId string,
    productKey string,
) *RemoveDeviceRequest {

    return &RemoveDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/device/{deviceName}:delete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
        DeviceName: deviceName,
        InstanceId: instanceId,
        RegionId: regionId,
        ProductKey: productKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveDeviceRequestWithoutParam() *RemoveDeviceRequest {

    return &RemoveDeviceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/device/{deviceName}:delete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param deviceName: 设备名称(Required) */
func (r *RemoveDeviceRequest) SetDeviceName(deviceName string) {
    r.DeviceName = deviceName
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *RemoveDeviceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *RemoveDeviceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productKey: 产品Key(Required) */
func (r *RemoveDeviceRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveDeviceRequest) GetRegionId() string {
    return r.RegionId
}

type RemoveDeviceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveDeviceResult `json:"result"`
}

type RemoveDeviceResult struct {
}