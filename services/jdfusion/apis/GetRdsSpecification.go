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

type GetRdsSpecificationRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* RDS数据库引擎，目前只支持mysql  */
    Engine string `json:"engine"`
}

/*
 * param regionId: 地域ID (Required)
 * param engine: RDS数据库引擎，目前只支持mysql (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetRdsSpecificationRequest(
    regionId string,
    engine string,
) *GetRdsSpecificationRequest {

	return &GetRdsSpecificationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rds_specification/{engine}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Engine: engine,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param engine: RDS数据库引擎，目前只支持mysql (Required)
 */
func NewGetRdsSpecificationRequestWithAllParams(
    regionId string,
    engine string,
) *GetRdsSpecificationRequest {

    return &GetRdsSpecificationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_specification/{engine}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Engine: engine,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetRdsSpecificationRequestWithoutParam() *GetRdsSpecificationRequest {

    return &GetRdsSpecificationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_specification/{engine}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetRdsSpecificationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param engine: RDS数据库引擎，目前只支持mysql(Required) */
func (r *GetRdsSpecificationRequest) SetEngine(engine string) {
    r.Engine = engine
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetRdsSpecificationRequest) GetRegionId() string {
    return r.RegionId
}

type GetRdsSpecificationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetRdsSpecificationResult `json:"result"`
}

type GetRdsSpecificationResult struct {
    Specifications []jdfusion.RdsSpecification `json:"specifications"`
}