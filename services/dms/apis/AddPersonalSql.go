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

type AddPersonalSqlRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 收藏的sql语句。 (Optional) */
    SqlStr *string `json:"sqlStr"`

    /* 备注。 (Optional) */
    Hint *string `json:"hint"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddPersonalSqlRequest(
    regionId string,
) *AddPersonalSqlRequest {

	return &AddPersonalSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/personalSql:add",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dataSourceId: 数据源id (Optional)
 * param sqlStr: 收藏的sql语句。 (Optional)
 * param hint: 备注。 (Optional)
 */
func NewAddPersonalSqlRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    sqlStr *string,
    hint *string,
) *AddPersonalSqlRequest {

    return &AddPersonalSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/personalSql:add",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        SqlStr: sqlStr,
        Hint: hint,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddPersonalSqlRequestWithoutParam() *AddPersonalSqlRequest {

    return &AddPersonalSqlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/personalSql:add",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *AddPersonalSqlRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源id(Optional) */
func (r *AddPersonalSqlRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}

/* param sqlStr: 收藏的sql语句。(Optional) */
func (r *AddPersonalSqlRequest) SetSqlStr(sqlStr string) {
    r.SqlStr = &sqlStr
}

/* param hint: 备注。(Optional) */
func (r *AddPersonalSqlRequest) SetHint(hint string) {
    r.Hint = &hint
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddPersonalSqlRequest) GetRegionId() string {
    return r.RegionId
}

type AddPersonalSqlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddPersonalSqlResult `json:"result"`
}

type AddPersonalSqlResult struct {
    StrResult string `json:"strResult"`
}