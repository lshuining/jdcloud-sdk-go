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

type SetAppStatusRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 硬件版本  */
    HardwareId string `json:"hardwareId"`

    /* OSID  */
    OsId string `json:"osId"`

    /* App名称  */
    AppName string `json:"appName"`

    /* APP状态(0停止 1启动)  */
    Status int `json:"status"`

    /* Edge名称  */
    EdgeName string `json:"edgeName"`

    /* App版本  */
    AppVersion string `json:"appVersion"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param hardwareId: 硬件版本 (Required)
 * param osId: OSID (Required)
 * param appName: App名称 (Required)
 * param status: APP状态(0停止 1启动) (Required)
 * param edgeName: Edge名称 (Required)
 * param appVersion: App版本 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetAppStatusRequest(
    instanceId string,
    regionId string,
    hardwareId string,
    osId string,
    appName string,
    status int,
    edgeName string,
    appVersion string,
) *SetAppStatusRequest {

	return &SetAppStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:setAppStatus",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
        HardwareId: hardwareId,
        OsId: osId,
        AppName: appName,
        Status: status,
        EdgeName: edgeName,
        AppVersion: appVersion,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param hardwareId: 硬件版本 (Required)
 * param osId: OSID (Required)
 * param appName: App名称 (Required)
 * param status: APP状态(0停止 1启动) (Required)
 * param edgeName: Edge名称 (Required)
 * param appVersion: App版本 (Required)
 */
func NewSetAppStatusRequestWithAllParams(
    instanceId string,
    regionId string,
    hardwareId string,
    osId string,
    appName string,
    status int,
    edgeName string,
    appVersion string,
) *SetAppStatusRequest {

    return &SetAppStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:setAppStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        HardwareId: hardwareId,
        OsId: osId,
        AppName: appName,
        Status: status,
        EdgeName: edgeName,
        AppVersion: appVersion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetAppStatusRequestWithoutParam() *SetAppStatusRequest {

    return &SetAppStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:setAppStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *SetAppStatusRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *SetAppStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param hardwareId: 硬件版本(Required) */
func (r *SetAppStatusRequest) SetHardwareId(hardwareId string) {
    r.HardwareId = hardwareId
}

/* param osId: OSID(Required) */
func (r *SetAppStatusRequest) SetOsId(osId string) {
    r.OsId = osId
}

/* param appName: App名称(Required) */
func (r *SetAppStatusRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param status: APP状态(0停止 1启动)(Required) */
func (r *SetAppStatusRequest) SetStatus(status int) {
    r.Status = status
}

/* param edgeName: Edge名称(Required) */
func (r *SetAppStatusRequest) SetEdgeName(edgeName string) {
    r.EdgeName = edgeName
}

/* param appVersion: App版本(Required) */
func (r *SetAppStatusRequest) SetAppVersion(appVersion string) {
    r.AppVersion = appVersion
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetAppStatusRequest) GetRegionId() string {
    return r.RegionId
}

type SetAppStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetAppStatusResult `json:"result"`
}

type SetAppStatusResult struct {
    AppStatus int `json:"appStatus"`
}