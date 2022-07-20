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
    xdata "github.com/lshuining/jdcloud-sdk-go/services/xdata/models"
)

type CreateTableRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`

    /* 数据表描述信息  */
    DbModelDBTable *xdata.DwTableDesc `json:"dbModelDBTable"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceName: 实例名称 (Required)
 * param dbModelDBTable: 数据表描述信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTableRequest(
    regionId string,
    instanceName string,
    dbModelDBTable *xdata.DwTableDesc,
) *CreateTableRequest {

	return &CreateTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwTable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceName: instanceName,
        DbModelDBTable: dbModelDBTable,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceName: 实例名称 (Required)
 * param dbModelDBTable: 数据表描述信息 (Required)
 */
func NewCreateTableRequestWithAllParams(
    regionId string,
    instanceName string,
    dbModelDBTable *xdata.DwTableDesc,
) *CreateTableRequest {

    return &CreateTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwTable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceName: instanceName,
        DbModelDBTable: dbModelDBTable,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTableRequestWithoutParam() *CreateTableRequest {

    return &CreateTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwTable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateTableRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceName: 实例名称(Required) */
func (r *CreateTableRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

/* param dbModelDBTable: 数据表描述信息(Required) */
func (r *CreateTableRequest) SetDbModelDBTable(dbModelDBTable *xdata.DwTableDesc) {
    r.DbModelDBTable = dbModelDBTable
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTableRequest) GetRegionId() string {
    return r.RegionId
}

type CreateTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTableResult `json:"result"`
}

type CreateTableResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}