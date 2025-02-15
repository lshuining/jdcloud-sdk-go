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

type ModifySourceEndpointRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`

    /* 源数据库的账号 (Optional) */
    AccountName *string `json:"accountName"`

    /* 源数据库的密码 (Optional) */
    Password *string `json:"password"`

    /* 不同数据库引擎独有的配置参数 (Optional) */
    EngineRelatedConfig *dbs.EngineRelatedConfig `json:"engineRelatedConfig"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifySourceEndpointRequest(
    regionId string,
    backupPlanId string,
) *ModifySourceEndpointRequest {

	return &ModifySourceEndpointRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:modifySourceEndpoint",
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
 * param accountName: 源数据库的账号 (Optional)
 * param password: 源数据库的密码 (Optional)
 * param engineRelatedConfig: 不同数据库引擎独有的配置参数 (Optional)
 */
func NewModifySourceEndpointRequestWithAllParams(
    regionId string,
    backupPlanId string,
    accountName *string,
    password *string,
    engineRelatedConfig *dbs.EngineRelatedConfig,
) *ModifySourceEndpointRequest {

    return &ModifySourceEndpointRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:modifySourceEndpoint",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        AccountName: accountName,
        Password: password,
        EngineRelatedConfig: engineRelatedConfig,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifySourceEndpointRequestWithoutParam() *ModifySourceEndpointRequest {

    return &ModifySourceEndpointRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:modifySourceEndpoint",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *ModifySourceEndpointRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *ModifySourceEndpointRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

/* param accountName: 源数据库的账号(Optional) */
func (r *ModifySourceEndpointRequest) SetAccountName(accountName string) {
    r.AccountName = &accountName
}

/* param password: 源数据库的密码(Optional) */
func (r *ModifySourceEndpointRequest) SetPassword(password string) {
    r.Password = &password
}

/* param engineRelatedConfig: 不同数据库引擎独有的配置参数(Optional) */
func (r *ModifySourceEndpointRequest) SetEngineRelatedConfig(engineRelatedConfig *dbs.EngineRelatedConfig) {
    r.EngineRelatedConfig = engineRelatedConfig
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifySourceEndpointRequest) GetRegionId() string {
    return r.RegionId
}

type ModifySourceEndpointResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifySourceEndpointResult `json:"result"`
}

type ModifySourceEndpointResult struct {
}