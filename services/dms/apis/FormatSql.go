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

type FormatSqlRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 需要格式化的sql (Optional) */
    SqlStr *string `json:"sqlStr"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewFormatSqlRequest(
    regionId string,
) *FormatSqlRequest {

	return &FormatSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/sql:format",
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
 * param sqlStr: 需要格式化的sql (Optional)
 */
func NewFormatSqlRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    sqlStr *string,
) *FormatSqlRequest {

    return &FormatSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sql:format",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        SqlStr: sqlStr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewFormatSqlRequestWithoutParam() *FormatSqlRequest {

    return &FormatSqlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sql:format",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *FormatSqlRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源id(Optional) */
func (r *FormatSqlRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}

/* param sqlStr: 需要格式化的sql(Optional) */
func (r *FormatSqlRequest) SetSqlStr(sqlStr string) {
    r.SqlStr = &sqlStr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r FormatSqlRequest) GetRegionId() string {
    return r.RegionId
}

type FormatSqlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result FormatSqlResult `json:"result"`
}

type FormatSqlResult struct {
    StrResult string `json:"strResult"`
}