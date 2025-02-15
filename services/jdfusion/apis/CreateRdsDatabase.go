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
    jdfusion "github.com/lshuining/jdcloud-sdk-go/services/jdfusion/models"
)

type CreateRdsDatabaseRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* RDS实例ID  */
    InstId string `json:"instId"`

    /* 创建RDS实例的数据库信息  */
    Database *jdfusion.CreateRDSDB `json:"database"`
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 * param database: 创建RDS实例的数据库信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRdsDatabaseRequest(
    regionId string,
    instId string,
    database *jdfusion.CreateRDSDB,
) *CreateRdsDatabaseRequest {

	return &CreateRdsDatabaseRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rds_instances/{instId}/databases",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstId: instId,
        Database: database,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 * param database: 创建RDS实例的数据库信息 (Required)
 */
func NewCreateRdsDatabaseRequestWithAllParams(
    regionId string,
    instId string,
    database *jdfusion.CreateRDSDB,
) *CreateRdsDatabaseRequest {

    return &CreateRdsDatabaseRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/databases",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstId: instId,
        Database: database,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRdsDatabaseRequestWithoutParam() *CreateRdsDatabaseRequest {

    return &CreateRdsDatabaseRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/databases",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateRdsDatabaseRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instId: RDS实例ID(Required) */
func (r *CreateRdsDatabaseRequest) SetInstId(instId string) {
    r.InstId = instId
}

/* param database: 创建RDS实例的数据库信息(Required) */
func (r *CreateRdsDatabaseRequest) SetDatabase(database *jdfusion.CreateRDSDB) {
    r.Database = database
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRdsDatabaseRequest) GetRegionId() string {
    return r.RegionId
}

type CreateRdsDatabaseResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRdsDatabaseResult `json:"result"`
}

type CreateRdsDatabaseResult struct {
    Task jdfusion.ResourceTFInfo `json:"task"`
}