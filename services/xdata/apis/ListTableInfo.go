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

type ListTableInfoRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`

    /* 数据库名称  */
    DatabaseName string `json:"databaseName"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceName: 实例名称 (Required)
 * param databaseName: 数据库名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListTableInfoRequest(
    regionId string,
    instanceName string,
    databaseName string,
) *ListTableInfoRequest {

	return &ListTableInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwTable",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceName: instanceName,
        DatabaseName: databaseName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceName: 实例名称 (Required)
 * param databaseName: 数据库名称 (Required)
 */
func NewListTableInfoRequestWithAllParams(
    regionId string,
    instanceName string,
    databaseName string,
) *ListTableInfoRequest {

    return &ListTableInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwTable",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceName: instanceName,
        DatabaseName: databaseName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListTableInfoRequestWithoutParam() *ListTableInfoRequest {

    return &ListTableInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwTable",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ListTableInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceName: 实例名称(Required) */
func (r *ListTableInfoRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

/* param databaseName: 数据库名称(Required) */
func (r *ListTableInfoRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = databaseName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListTableInfoRequest) GetRegionId() string {
    return r.RegionId
}

type ListTableInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListTableInfoResult `json:"result"`
}

type ListTableInfoResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data []xdata.DwTable `json:"data"`
}