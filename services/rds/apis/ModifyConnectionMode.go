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

type ModifyConnectionModeRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 连接模式<br> - standard：标准模式(缺省)，响应时间短，但没有 SQL 审计和拦截的能力 <br>- security：高安全模式，具备一定的 SQL注入拦截能力，并可开启 SQL 审计，但会增加一定的响应时间  */
    ConnectionMode string `json:"connectionMode"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param connectionMode: 连接模式<br> - standard：标准模式(缺省)，响应时间短，但没有 SQL 审计和拦截的能力 <br>- security：高安全模式，具备一定的 SQL注入拦截能力，并可开启 SQL 审计，但会增加一定的响应时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyConnectionModeRequest(
    regionId string,
    instanceId string,
    connectionMode string,
) *ModifyConnectionModeRequest {

	return &ModifyConnectionModeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyConnectionMode",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ConnectionMode: connectionMode,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param connectionMode: 连接模式<br> - standard：标准模式(缺省)，响应时间短，但没有 SQL 审计和拦截的能力 <br>- security：高安全模式，具备一定的 SQL注入拦截能力，并可开启 SQL 审计，但会增加一定的响应时间 (Required)
 */
func NewModifyConnectionModeRequestWithAllParams(
    regionId string,
    instanceId string,
    connectionMode string,
) *ModifyConnectionModeRequest {

    return &ModifyConnectionModeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyConnectionMode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ConnectionMode: connectionMode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyConnectionModeRequestWithoutParam() *ModifyConnectionModeRequest {

    return &ModifyConnectionModeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyConnectionMode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ModifyConnectionModeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *ModifyConnectionModeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param connectionMode: 连接模式<br> - standard：标准模式(缺省)，响应时间短，但没有 SQL 审计和拦截的能力 <br>- security：高安全模式，具备一定的 SQL注入拦截能力，并可开启 SQL 审计，但会增加一定的响应时间(Required) */
func (r *ModifyConnectionModeRequest) SetConnectionMode(connectionMode string) {
    r.ConnectionMode = connectionMode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyConnectionModeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyConnectionModeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyConnectionModeResult `json:"result"`
}

type ModifyConnectionModeResult struct {
}