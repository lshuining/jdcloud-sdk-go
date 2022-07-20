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
    apigateway "github.com/lshuining/jdcloud-sdk-go/services/apigateway/models"
)

type QueryAccessAuthRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 访问授权ID  */
    AccessAuthId string `json:"accessAuthId"`
}

/*
 * param regionId: 地域ID (Required)
 * param accessAuthId: 访问授权ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAccessAuthRequest(
    regionId string,
    accessAuthId string,
) *QueryAccessAuthRequest {

	return &QueryAccessAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accessAuths/{accessAuthId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AccessAuthId: accessAuthId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param accessAuthId: 访问授权ID (Required)
 */
func NewQueryAccessAuthRequestWithAllParams(
    regionId string,
    accessAuthId string,
) *QueryAccessAuthRequest {

    return &QueryAccessAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessAuths/{accessAuthId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AccessAuthId: accessAuthId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAccessAuthRequestWithoutParam() *QueryAccessAuthRequest {

    return &QueryAccessAuthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessAuths/{accessAuthId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QueryAccessAuthRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param accessAuthId: 访问授权ID(Required) */
func (r *QueryAccessAuthRequest) SetAccessAuthId(accessAuthId string) {
    r.AccessAuthId = accessAuthId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAccessAuthRequest) GetRegionId() string {
    return r.RegionId
}

type QueryAccessAuthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAccessAuthResult `json:"result"`
}

type QueryAccessAuthResult struct {
    AccessAuth apigateway.AccessAuth `json:"accessAuth"`
}