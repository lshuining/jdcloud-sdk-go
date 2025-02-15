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

type LoginRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* rds，drds的实例id。  */
    InstanceId string `json:"instanceId"`

    /* 数据源源类型，CDS("CDS", 1), MYSQL("MYSQL", 2), ORACLE("ORACLE", 3), SQLSERVER("SQLSERVER", 4), CDSMYSQL("CDSMYSQL", 5), CDSORACLE("CDSORACLE", 6), CDSSQLSERVER("CDSSQLSERVER", 7), DATACENTER("DATACENTER", 8), HBASE("Hbase",9),MONGODB("MongoDb",10),ES("ES",11), MYSQL_INS("MYSQL_INS",12), DRDS_INS("DRDS_INS",13)。 (Optional) */
    DbType *int `json:"dbType"`

    /* CLASSIC("CLASSIC", 0), RDS("RDS", 1), ECS("ECS", 2), VPC("VPC", 3), 当前只支持rds模式。 (Optional) */
    AddrMode *int `json:"addrMode"`

    /* 数据库用户名。 (Optional) */
    DbUser *string `json:"dbUser"`

    /* 数据库用户密码。 (Optional) */
    DbPassword *string `json:"dbPassword"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: rds，drds的实例id。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewLoginRequest(
    regionId string,
    instanceId string,
) *LoginRequest {

	return &LoginRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:login",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: rds，drds的实例id。 (Required)
 * param dbType: 数据源源类型，CDS("CDS", 1), MYSQL("MYSQL", 2), ORACLE("ORACLE", 3), SQLSERVER("SQLSERVER", 4), CDSMYSQL("CDSMYSQL", 5), CDSORACLE("CDSORACLE", 6), CDSSQLSERVER("CDSSQLSERVER", 7), DATACENTER("DATACENTER", 8), HBASE("Hbase",9),MONGODB("MongoDb",10),ES("ES",11), MYSQL_INS("MYSQL_INS",12), DRDS_INS("DRDS_INS",13)。 (Optional)
 * param addrMode: CLASSIC("CLASSIC", 0), RDS("RDS", 1), ECS("ECS", 2), VPC("VPC", 3), 当前只支持rds模式。 (Optional)
 * param dbUser: 数据库用户名。 (Optional)
 * param dbPassword: 数据库用户密码。 (Optional)
 */
func NewLoginRequestWithAllParams(
    regionId string,
    instanceId string,
    dbType *int,
    addrMode *int,
    dbUser *string,
    dbPassword *string,
) *LoginRequest {

    return &LoginRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:login",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DbType: dbType,
        AddrMode: addrMode,
        DbUser: dbUser,
        DbPassword: dbPassword,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewLoginRequestWithoutParam() *LoginRequest {

    return &LoginRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:login",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *LoginRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: rds，drds的实例id。(Required) */
func (r *LoginRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param dbType: 数据源源类型，CDS("CDS", 1), MYSQL("MYSQL", 2), ORACLE("ORACLE", 3), SQLSERVER("SQLSERVER", 4), CDSMYSQL("CDSMYSQL", 5), CDSORACLE("CDSORACLE", 6), CDSSQLSERVER("CDSSQLSERVER", 7), DATACENTER("DATACENTER", 8), HBASE("Hbase",9),MONGODB("MongoDb",10),ES("ES",11), MYSQL_INS("MYSQL_INS",12), DRDS_INS("DRDS_INS",13)。(Optional) */
func (r *LoginRequest) SetDbType(dbType int) {
    r.DbType = &dbType
}

/* param addrMode: CLASSIC("CLASSIC", 0), RDS("RDS", 1), ECS("ECS", 2), VPC("VPC", 3), 当前只支持rds模式。(Optional) */
func (r *LoginRequest) SetAddrMode(addrMode int) {
    r.AddrMode = &addrMode
}

/* param dbUser: 数据库用户名。(Optional) */
func (r *LoginRequest) SetDbUser(dbUser string) {
    r.DbUser = &dbUser
}

/* param dbPassword: 数据库用户密码。(Optional) */
func (r *LoginRequest) SetDbPassword(dbPassword string) {
    r.DbPassword = &dbPassword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r LoginRequest) GetRegionId() string {
    return r.RegionId
}

type LoginResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result LoginResult `json:"result"`
}

type LoginResult struct {
}