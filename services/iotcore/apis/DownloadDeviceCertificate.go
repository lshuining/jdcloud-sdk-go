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

type DownloadDeviceCertificateRequest struct {

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
func NewDownloadDeviceCertificateRequest(
    regionId string,
    instanceId string,
    deviceId string,
) *DownloadDeviceCertificateRequest {

	return &DownloadDeviceCertificateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:downloadCertificate",
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
func NewDownloadDeviceCertificateRequestWithAllParams(
    regionId string,
    instanceId string,
    deviceId string,
) *DownloadDeviceCertificateRequest {

    return &DownloadDeviceCertificateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:downloadCertificate",
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
func NewDownloadDeviceCertificateRequestWithoutParam() *DownloadDeviceCertificateRequest {

    return &DownloadDeviceCertificateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/device:downloadCertificate",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *DownloadDeviceCertificateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例Id(Required) */
func (r *DownloadDeviceCertificateRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param deviceId: 设备ID(Required) */
func (r *DownloadDeviceCertificateRequest) SetDeviceId(deviceId string) {
    r.DeviceId = deviceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DownloadDeviceCertificateRequest) GetRegionId() string {
    return r.RegionId
}

type DownloadDeviceCertificateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DownloadDeviceCertificateResult `json:"result"`
}

type DownloadDeviceCertificateResult struct {
    DeviceId string `json:"deviceId"`
    DeviceCertUrl string `json:"deviceCertUrl"`
}