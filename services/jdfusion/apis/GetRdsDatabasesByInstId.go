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

type GetRdsDatabasesByInstIdRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* RDS实例ID  */
    InstId string `json:"instId"`
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetRdsDatabasesByInstIdRequest(
    regionId string,
    instId string,
) *GetRdsDatabasesByInstIdRequest {

	return &GetRdsDatabasesByInstIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rds_instances/{instId}/databases",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstId: instId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 */
func NewGetRdsDatabasesByInstIdRequestWithAllParams(
    regionId string,
    instId string,
) *GetRdsDatabasesByInstIdRequest {

    return &GetRdsDatabasesByInstIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/databases",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstId: instId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetRdsDatabasesByInstIdRequestWithoutParam() *GetRdsDatabasesByInstIdRequest {

    return &GetRdsDatabasesByInstIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/databases",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetRdsDatabasesByInstIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instId: RDS实例ID(Required) */
func (r *GetRdsDatabasesByInstIdRequest) SetInstId(instId string) {
    r.InstId = instId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetRdsDatabasesByInstIdRequest) GetRegionId() string {
    return r.RegionId
}

type GetRdsDatabasesByInstIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetRdsDatabasesByInstIdResult `json:"result"`
}

type GetRdsDatabasesByInstIdResult struct {
    Dbs []jdfusion.RdsDBInfo `json:"dbs"`
}