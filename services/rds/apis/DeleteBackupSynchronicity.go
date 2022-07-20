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

type DeleteBackupSynchronicityRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 跨地域备份同步服务ID  */
    ServiceId string `json:"serviceId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param serviceId: 跨地域备份同步服务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteBackupSynchronicityRequest(
    regionId string,
    serviceId string,
) *DeleteBackupSynchronicityRequest {

	return &DeleteBackupSynchronicityRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupSynchronicities/{serviceId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceId: serviceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param serviceId: 跨地域备份同步服务ID (Required)
 */
func NewDeleteBackupSynchronicityRequestWithAllParams(
    regionId string,
    serviceId string,
) *DeleteBackupSynchronicityRequest {

    return &DeleteBackupSynchronicityRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupSynchronicities/{serviceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceId: serviceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteBackupSynchronicityRequestWithoutParam() *DeleteBackupSynchronicityRequest {

    return &DeleteBackupSynchronicityRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupSynchronicities/{serviceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DeleteBackupSynchronicityRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serviceId: 跨地域备份同步服务ID(Required) */
func (r *DeleteBackupSynchronicityRequest) SetServiceId(serviceId string) {
    r.ServiceId = serviceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteBackupSynchronicityRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteBackupSynchronicityResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteBackupSynchronicityResult `json:"result"`
}

type DeleteBackupSynchronicityResult struct {
}