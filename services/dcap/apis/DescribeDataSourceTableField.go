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

type DescribeDataSourceTableFieldRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 数据源 ID  */
    DataSourceId string `json:"dataSourceId"`

    /* 表-名称  */
    TableName string `json:"tableName"`

    /* 字段-名称  */
    FieldName string `json:"fieldName"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param tableName: 表-名称 (Required)
 * param fieldName: 字段-名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDataSourceTableFieldRequest(
    regionId string,
    dataSourceId string,
    tableName string,
    fieldName string,
) *DescribeDataSourceTableFieldRequest {

	return &DescribeDataSourceTableFieldRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables/{tableName}/fields/{fieldName}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceId: dataSourceId,
        TableName: tableName,
        FieldName: fieldName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param tableName: 表-名称 (Required)
 * param fieldName: 字段-名称 (Required)
 */
func NewDescribeDataSourceTableFieldRequestWithAllParams(
    regionId string,
    dataSourceId string,
    tableName string,
    fieldName string,
) *DescribeDataSourceTableFieldRequest {

    return &DescribeDataSourceTableFieldRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables/{tableName}/fields/{fieldName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        TableName: tableName,
        FieldName: fieldName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDataSourceTableFieldRequestWithoutParam() *DescribeDataSourceTableFieldRequest {

    return &DescribeDataSourceTableFieldRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables/{tableName}/fields/{fieldName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeDataSourceTableFieldRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源 ID(Required) */
func (r *DescribeDataSourceTableFieldRequest) SetDataSourceId(dataSourceId string) {
    r.DataSourceId = dataSourceId
}

/* param tableName: 表-名称(Required) */
func (r *DescribeDataSourceTableFieldRequest) SetTableName(tableName string) {
    r.TableName = tableName
}

/* param fieldName: 字段-名称(Required) */
func (r *DescribeDataSourceTableFieldRequest) SetFieldName(fieldName string) {
    r.FieldName = fieldName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDataSourceTableFieldRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDataSourceTableFieldResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDataSourceTableFieldResult `json:"result"`
}

type DescribeDataSourceTableFieldResult struct {
    FieldName string `json:"fieldName"`
    FieldType int `json:"fieldType"`
    FieldLength int `json:"fieldLength"`
    FieldAttr string `json:"fieldAttr"`
    FieldComment string `json:"fieldComment"`
    EncryptField string `json:"encryptField"`
    IndexField string `json:"indexField"`
    KeepPlainText bool `json:"keepPlainText"`
    FieldTag string `json:"fieldTag"`
    FieldLevel string `json:"fieldLevel"`
}