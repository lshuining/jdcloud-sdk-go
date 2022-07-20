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
    iothub "github.com/lshuining/jdcloud-sdk-go/services/iothub/models"
)

type UpdateDeviceRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 设备Id  */
    DeviceId string `json:"deviceId"`

    /* 设备型号 (Optional) */
    Model *string `json:"model"`

    /* 设备厂商 (Optional) */
    Manufacturer *string `json:"manufacturer"`

    /* 设备描述 (Optional) */
    Description *string `json:"description"`

    /* 设备状态 (Optional) */
    Status *int `json:"status"`
}

/*
 * param regionId: 设备归属的实例所在区域 (Required)
 * param deviceId: 设备Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateDeviceRequest(
    regionId string,
    deviceId string,
) *UpdateDeviceRequest {

	return &UpdateDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/device/{deviceId}:update",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DeviceId: deviceId,
	}
}

/*
 * param regionId: 设备归属的实例所在区域 (Required)
 * param deviceId: 设备Id (Required)
 * param model: 设备型号 (Optional)
 * param manufacturer: 设备厂商 (Optional)
 * param description: 设备描述 (Optional)
 * param status: 设备状态 (Optional)
 */
func NewUpdateDeviceRequestWithAllParams(
    regionId string,
    deviceId string,
    model *string,
    manufacturer *string,
    description *string,
    status *int,
) *UpdateDeviceRequest {

    return &UpdateDeviceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/device/{deviceId}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DeviceId: deviceId,
        Model: model,
        Manufacturer: manufacturer,
        Description: description,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateDeviceRequestWithoutParam() *UpdateDeviceRequest {

    return &UpdateDeviceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/device/{deviceId}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *UpdateDeviceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param deviceId: 设备Id(Required) */
func (r *UpdateDeviceRequest) SetDeviceId(deviceId string) {
    r.DeviceId = deviceId
}

/* param model: 设备型号(Optional) */
func (r *UpdateDeviceRequest) SetModel(model string) {
    r.Model = &model
}

/* param manufacturer: 设备厂商(Optional) */
func (r *UpdateDeviceRequest) SetManufacturer(manufacturer string) {
    r.Manufacturer = &manufacturer
}

/* param description: 设备描述(Optional) */
func (r *UpdateDeviceRequest) SetDescription(description string) {
    r.Description = &description
}

/* param status: 设备状态(Optional) */
func (r *UpdateDeviceRequest) SetStatus(status int) {
    r.Status = &status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateDeviceRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateDeviceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateDeviceResult `json:"result"`
}

type UpdateDeviceResult struct {
    Device iothub.DeviceVO `json:"device"`
}