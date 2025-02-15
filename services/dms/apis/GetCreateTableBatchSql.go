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
    dms "github.com/lshuining/jdcloud-sdk-go/services/dms/models"
)

type GetCreateTableBatchSqlRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 数据库名称 (Optional) */
    DbName *string `json:"dbName"`

    /* 查询结果。 (Optional) */
    CreateTableInfos []dms.CreateTableInfo `json:"createTableInfos"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetCreateTableBatchSqlRequest(
    regionId string,
) *GetCreateTableBatchSqlRequest {

	return &GetCreateTableBatchSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/console:getCreateTableBatchSql",
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
 * param dbName: 数据库名称 (Optional)
 * param createTableInfos: 查询结果。 (Optional)
 */
func NewGetCreateTableBatchSqlRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    dbName *string,
    createTableInfos []dms.CreateTableInfo,
) *GetCreateTableBatchSqlRequest {

    return &GetCreateTableBatchSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/console:getCreateTableBatchSql",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        DbName: dbName,
        CreateTableInfos: createTableInfos,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetCreateTableBatchSqlRequestWithoutParam() *GetCreateTableBatchSqlRequest {

    return &GetCreateTableBatchSqlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/console:getCreateTableBatchSql",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *GetCreateTableBatchSqlRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源id(Optional) */
func (r *GetCreateTableBatchSqlRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}

/* param dbName: 数据库名称(Optional) */
func (r *GetCreateTableBatchSqlRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}

/* param createTableInfos: 查询结果。(Optional) */
func (r *GetCreateTableBatchSqlRequest) SetCreateTableInfos(createTableInfos []dms.CreateTableInfo) {
    r.CreateTableInfos = createTableInfos
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetCreateTableBatchSqlRequest) GetRegionId() string {
    return r.RegionId
}

type GetCreateTableBatchSqlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetCreateTableBatchSqlResult `json:"result"`
}

type GetCreateTableBatchSqlResult struct {
    StrResult string `json:"strResult"`
}