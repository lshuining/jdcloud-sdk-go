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

type CreateRdsAccountRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 账号名，在同一个RDS实例中，账号名不能重复。账号名的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md)  */
    AccountName string `json:"accountName"`

    /* 密码,密码的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md)  */
    AccountPassword string `json:"accountPassword"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param accountName: 账号名，在同一个RDS实例中，账号名不能重复。账号名的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Required)
 * param accountPassword: 密码,密码的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRdsAccountRequest(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *CreateRdsAccountRequest {

	return &CreateRdsAccountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/accounts",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param accountName: 账号名，在同一个RDS实例中，账号名不能重复。账号名的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Required)
 * param accountPassword: 密码,密码的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Required)
 */
func NewCreateRdsAccountRequestWithAllParams(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *CreateRdsAccountRequest {

    return &CreateRdsAccountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/accounts",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRdsAccountRequestWithoutParam() *CreateRdsAccountRequest {

    return &CreateRdsAccountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/accounts",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *CreateRdsAccountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *CreateRdsAccountRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param accountName: 账号名，在同一个RDS实例中，账号名不能重复。账号名的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md)(Required) */
func (r *CreateRdsAccountRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

/* param accountPassword: 密码,密码的具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md)(Required) */
func (r *CreateRdsAccountRequest) SetAccountPassword(accountPassword string) {
    r.AccountPassword = accountPassword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRdsAccountRequest) GetRegionId() string {
    return r.RegionId
}

type CreateRdsAccountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRdsAccountResult `json:"result"`
}

type CreateRdsAccountResult struct {
}