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

type BatchDisableAlarmsRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 告警规则的ID列表  */
    Ids []string `json:"ids"`
}

/*
 * param regionId: 地域 Id (Required)
 * param ids: 告警规则的ID列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchDisableAlarmsRequest(
    regionId string,
    ids []string,
) *BatchDisableAlarmsRequest {

	return &BatchDisableAlarmsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarms/disable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ids: ids,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param ids: 告警规则的ID列表 (Required)
 */
func NewBatchDisableAlarmsRequestWithAllParams(
    regionId string,
    ids []string,
) *BatchDisableAlarmsRequest {

    return &BatchDisableAlarmsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms/disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ids: ids,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchDisableAlarmsRequestWithoutParam() *BatchDisableAlarmsRequest {

    return &BatchDisableAlarmsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms/disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *BatchDisableAlarmsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ids: 告警规则的ID列表(Required) */
func (r *BatchDisableAlarmsRequest) SetIds(ids []string) {
    r.Ids = ids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchDisableAlarmsRequest) GetRegionId() string {
    return r.RegionId
}

type BatchDisableAlarmsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchDisableAlarmsResult `json:"result"`
}

type BatchDisableAlarmsResult struct {
}