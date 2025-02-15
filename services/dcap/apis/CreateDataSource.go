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
    dcap "github.com/lshuining/jdcloud-sdk-go/services/dcap/models"
)

type CreateDataSourceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 添加数据源需要的信息  */
    DataSourceSpec *dcap.DataSourceSpec `json:"dataSourceSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceSpec: 添加数据源需要的信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDataSourceRequest(
    regionId string,
    dataSourceSpec *dcap.DataSourceSpec,
) *CreateDataSourceRequest {

	return &CreateDataSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceSpec: dataSourceSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceSpec: 添加数据源需要的信息 (Required)
 */
func NewCreateDataSourceRequestWithAllParams(
    regionId string,
    dataSourceSpec *dcap.DataSourceSpec,
) *CreateDataSourceRequest {

    return &CreateDataSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceSpec: dataSourceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDataSourceRequestWithoutParam() *CreateDataSourceRequest {

    return &CreateDataSourceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateDataSourceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceSpec: 添加数据源需要的信息(Required) */
func (r *CreateDataSourceRequest) SetDataSourceSpec(dataSourceSpec *dcap.DataSourceSpec) {
    r.DataSourceSpec = dataSourceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDataSourceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateDataSourceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDataSourceResult `json:"result"`
}

type CreateDataSourceResult struct {
    DataSourceId string `json:"dataSourceId"`
}