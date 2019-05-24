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
)

type CheckDeviceIdRequest struct {

    core.JDCloudRequest

    /* Device 唯一标识  */
    DeviceId string `json:"deviceId"`
}

/*
 * param deviceId: Device 唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckDeviceIdRequest(
    deviceId string,
) *CheckDeviceIdRequest {

	return &CheckDeviceIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/device/{deviceId}:check",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        DeviceId: deviceId,
	}
}

/*
 * param deviceId: Device 唯一标识 (Required)
 */
func NewCheckDeviceIdRequestWithAllParams(
    deviceId string,
) *CheckDeviceIdRequest {

    return &CheckDeviceIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/device/{deviceId}:check",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DeviceId: deviceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckDeviceIdRequestWithoutParam() *CheckDeviceIdRequest {

    return &CheckDeviceIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/device/{deviceId}:check",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param deviceId: Device 唯一标识(Required) */
func (r *CheckDeviceIdRequest) SetDeviceId(deviceId string) {
    r.DeviceId = deviceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckDeviceIdRequest) GetRegionId() string {
    return ""
}

type CheckDeviceIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckDeviceIdResult `json:"result"`
}

type CheckDeviceIdResult struct {
}