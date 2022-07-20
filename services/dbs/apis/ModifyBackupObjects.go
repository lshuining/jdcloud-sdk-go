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

type ModifyBackupObjectsRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`

    /* 数据库名称 (Optional) */
    Database *string `json:"database"`

    /* 表名, 如果不填或者为空，表示对整个数据库进行备份 (Optional) */
    Objects []string `json:"objects"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyBackupObjectsRequest(
    regionId string,
    backupPlanId string,
) *ModifyBackupObjectsRequest {

	return &ModifyBackupObjectsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:modifyBackupObjects",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BackupPlanId: backupPlanId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 * param database: 数据库名称 (Optional)
 * param objects: 表名, 如果不填或者为空，表示对整个数据库进行备份 (Optional)
 */
func NewModifyBackupObjectsRequestWithAllParams(
    regionId string,
    backupPlanId string,
    database *string,
    objects []string,
) *ModifyBackupObjectsRequest {

    return &ModifyBackupObjectsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:modifyBackupObjects",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        Database: database,
        Objects: objects,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyBackupObjectsRequestWithoutParam() *ModifyBackupObjectsRequest {

    return &ModifyBackupObjectsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:modifyBackupObjects",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *ModifyBackupObjectsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *ModifyBackupObjectsRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

/* param database: 数据库名称(Optional) */
func (r *ModifyBackupObjectsRequest) SetDatabase(database string) {
    r.Database = &database
}

/* param objects: 表名, 如果不填或者为空，表示对整个数据库进行备份(Optional) */
func (r *ModifyBackupObjectsRequest) SetObjects(objects []string) {
    r.Objects = objects
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyBackupObjectsRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyBackupObjectsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyBackupObjectsResult `json:"result"`
}

type ModifyBackupObjectsResult struct {
}