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
    dbs "github.com/lshuining/jdcloud-sdk-go/services/dbs/models"
)

type RestoreLogicalBackupRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`

    /* 用于恢复的备份Id，仅限于本备份计划生成的备份  */
    BackupId string `json:"backupId"`

    /* 备份实例的数据源信息  */
    SourceEndpoint *dbs.RestoreSourceEndpoint `json:"sourceEndpoint"`

    /* 是否是新建数据源；true：是；false：不是  */
    CreateNewEndpoint bool `json:"createNewEndpoint"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 * param backupId: 用于恢复的备份Id，仅限于本备份计划生成的备份 (Required)
 * param sourceEndpoint: 备份实例的数据源信息 (Required)
 * param createNewEndpoint: 是否是新建数据源；true：是；false：不是 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRestoreLogicalBackupRequest(
    regionId string,
    backupPlanId string,
    backupId string,
    sourceEndpoint *dbs.RestoreSourceEndpoint,
    createNewEndpoint bool,
) *RestoreLogicalBackupRequest {

	return &RestoreLogicalBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:restoreLogicalBackup",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        BackupId: backupId,
        SourceEndpoint: sourceEndpoint,
        CreateNewEndpoint: createNewEndpoint,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 * param backupId: 用于恢复的备份Id，仅限于本备份计划生成的备份 (Required)
 * param sourceEndpoint: 备份实例的数据源信息 (Required)
 * param createNewEndpoint: 是否是新建数据源；true：是；false：不是 (Required)
 */
func NewRestoreLogicalBackupRequestWithAllParams(
    regionId string,
    backupPlanId string,
    backupId string,
    sourceEndpoint *dbs.RestoreSourceEndpoint,
    createNewEndpoint bool,
) *RestoreLogicalBackupRequest {

    return &RestoreLogicalBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:restoreLogicalBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        BackupId: backupId,
        SourceEndpoint: sourceEndpoint,
        CreateNewEndpoint: createNewEndpoint,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRestoreLogicalBackupRequestWithoutParam() *RestoreLogicalBackupRequest {

    return &RestoreLogicalBackupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:restoreLogicalBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *RestoreLogicalBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *RestoreLogicalBackupRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

/* param backupId: 用于恢复的备份Id，仅限于本备份计划生成的备份(Required) */
func (r *RestoreLogicalBackupRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

/* param sourceEndpoint: 备份实例的数据源信息(Required) */
func (r *RestoreLogicalBackupRequest) SetSourceEndpoint(sourceEndpoint *dbs.RestoreSourceEndpoint) {
    r.SourceEndpoint = sourceEndpoint
}

/* param createNewEndpoint: 是否是新建数据源；true：是；false：不是(Required) */
func (r *RestoreLogicalBackupRequest) SetCreateNewEndpoint(createNewEndpoint bool) {
    r.CreateNewEndpoint = createNewEndpoint
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreLogicalBackupRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreLogicalBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreLogicalBackupResult `json:"result"`
}

type RestoreLogicalBackupResult struct {
}