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

type GetExportDataIdRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 导出数据的查询sql (Optional) */
    Sql *string `json:"sql"`

    /* 数据库名称 (Optional) */
    DbName *string `json:"dbName"`

    /* 导出数量 (Optional) */
    Count *int `json:"count"`

    /* 是否跳过检测 (Optional) */
    IgnoreCheck *bool `json:"ignoreCheck"`

    /* 跳过行数检测原因 (Optional) */
    Reason *string `json:"reason"`

    /* 导出文件字符编码 (Optional) */
    Charset *string `json:"charset"`

    /* 导出方式，SYNC("SYNC", 0), ASYNC("ASYNC", 1)，当前只支持SYNC导出; (Optional) */
    ExportTypeEnum *string `json:"exportTypeEnum"`

    /* 导出文件格式，CSV("CSV", 0), SQL("SQL", 1); (Optional) */
    ExportFileTypeEnum *string `json:"exportFileTypeEnum"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetExportDataIdRequest(
    regionId string,
) *GetExportDataIdRequest {

	return &GetExportDataIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/console:getExportDataId",
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
 * param sql: 导出数据的查询sql (Optional)
 * param dbName: 数据库名称 (Optional)
 * param count: 导出数量 (Optional)
 * param ignoreCheck: 是否跳过检测 (Optional)
 * param reason: 跳过行数检测原因 (Optional)
 * param charset: 导出文件字符编码 (Optional)
 * param exportTypeEnum: 导出方式，SYNC("SYNC", 0), ASYNC("ASYNC", 1)，当前只支持SYNC导出; (Optional)
 * param exportFileTypeEnum: 导出文件格式，CSV("CSV", 0), SQL("SQL", 1); (Optional)
 */
func NewGetExportDataIdRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    sql *string,
    dbName *string,
    count *int,
    ignoreCheck *bool,
    reason *string,
    charset *string,
    exportTypeEnum *string,
    exportFileTypeEnum *string,
) *GetExportDataIdRequest {

    return &GetExportDataIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/console:getExportDataId",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        Sql: sql,
        DbName: dbName,
        Count: count,
        IgnoreCheck: ignoreCheck,
        Reason: reason,
        Charset: charset,
        ExportTypeEnum: exportTypeEnum,
        ExportFileTypeEnum: exportFileTypeEnum,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetExportDataIdRequestWithoutParam() *GetExportDataIdRequest {

    return &GetExportDataIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/console:getExportDataId",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *GetExportDataIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源id(Optional) */
func (r *GetExportDataIdRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}

/* param sql: 导出数据的查询sql(Optional) */
func (r *GetExportDataIdRequest) SetSql(sql string) {
    r.Sql = &sql
}

/* param dbName: 数据库名称(Optional) */
func (r *GetExportDataIdRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}

/* param count: 导出数量(Optional) */
func (r *GetExportDataIdRequest) SetCount(count int) {
    r.Count = &count
}

/* param ignoreCheck: 是否跳过检测(Optional) */
func (r *GetExportDataIdRequest) SetIgnoreCheck(ignoreCheck bool) {
    r.IgnoreCheck = &ignoreCheck
}

/* param reason: 跳过行数检测原因(Optional) */
func (r *GetExportDataIdRequest) SetReason(reason string) {
    r.Reason = &reason
}

/* param charset: 导出文件字符编码(Optional) */
func (r *GetExportDataIdRequest) SetCharset(charset string) {
    r.Charset = &charset
}

/* param exportTypeEnum: 导出方式，SYNC("SYNC", 0), ASYNC("ASYNC", 1)，当前只支持SYNC导出;(Optional) */
func (r *GetExportDataIdRequest) SetExportTypeEnum(exportTypeEnum string) {
    r.ExportTypeEnum = &exportTypeEnum
}

/* param exportFileTypeEnum: 导出文件格式，CSV("CSV", 0), SQL("SQL", 1);(Optional) */
func (r *GetExportDataIdRequest) SetExportFileTypeEnum(exportFileTypeEnum string) {
    r.ExportFileTypeEnum = &exportFileTypeEnum
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetExportDataIdRequest) GetRegionId() string {
    return r.RegionId
}

type GetExportDataIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetExportDataIdResult `json:"result"`
}

type GetExportDataIdResult struct {
    StrResult string `json:"strResult"`
}