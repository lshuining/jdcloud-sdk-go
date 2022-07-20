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

type AddDeviceRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 设备名称 (Optional) */
    DeviceName *string `json:"deviceName"`

    /* 设备所归属的产品 (Optional) */
    ProductKey *string `json:"productKey"`

    /* 设备型号 (Optional) */
    Model *string `json:"model"`

    /* 厂商 (Optional) */
    Manufacturer *string `json:"manufacturer"`

    /* 设备描述 (Optional) */
    Description *string `json:"description"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddDeviceRequest(
    instanceId string,
    regionId string,
) *AddDeviceRequest {

	return &AddDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/device:add",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param deviceName: 设备名称 (Optional)
 * param productKey: 设备所归属的产品 (Optional)
 * param model: 设备型号 (Optional)
 * param manufacturer: 厂商 (Optional)
 * param description: 设备描述 (Optional)
 */
func NewAddDeviceRequestWithAllParams(
    instanceId string,
    regionId string,
    deviceName *string,
    productKey *string,
    model *string,
    manufacturer *string,
    description *string,
) *AddDeviceRequest {

    return &AddDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/device:add",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        DeviceName: deviceName,
        ProductKey: productKey,
        Model: model,
        Manufacturer: manufacturer,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddDeviceRequestWithoutParam() *AddDeviceRequest {

    return &AddDeviceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/device:add",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *AddDeviceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *AddDeviceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param deviceName: 设备名称(Optional) */
func (r *AddDeviceRequest) SetDeviceName(deviceName string) {
    r.DeviceName = &deviceName
}

/* param productKey: 设备所归属的产品(Optional) */
func (r *AddDeviceRequest) SetProductKey(productKey string) {
    r.ProductKey = &productKey
}

/* param model: 设备型号(Optional) */
func (r *AddDeviceRequest) SetModel(model string) {
    r.Model = &model
}

/* param manufacturer: 厂商(Optional) */
func (r *AddDeviceRequest) SetManufacturer(manufacturer string) {
    r.Manufacturer = &manufacturer
}

/* param description: 设备描述(Optional) */
func (r *AddDeviceRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddDeviceRequest) GetRegionId() string {
    return r.RegionId
}

type AddDeviceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddDeviceResult `json:"result"`
}

type AddDeviceResult struct {
    DeviceName string `json:"deviceName"`
    Identifier string `json:"identifier"`
    Secret string `json:"secret"`
}