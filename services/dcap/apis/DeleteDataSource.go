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

type DeleteDataSourceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 数据源 ID  */
    DataSourceId string `json:"dataSourceId"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteDataSourceRequest(
    regionId string,
    dataSourceId string,
) *DeleteDataSourceRequest {

	return &DeleteDataSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources/{dataSourceId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceId: dataSourceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 */
func NewDeleteDataSourceRequestWithAllParams(
    regionId string,
    dataSourceId string,
) *DeleteDataSourceRequest {

    return &DeleteDataSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteDataSourceRequestWithoutParam() *DeleteDataSourceRequest {

    return &DeleteDataSourceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteDataSourceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源 ID(Required) */
func (r *DeleteDataSourceRequest) SetDataSourceId(dataSourceId string) {
    r.DataSourceId = dataSourceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteDataSourceRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteDataSourceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteDataSourceResult `json:"result"`
}

type DeleteDataSourceResult struct {
}