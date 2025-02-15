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

type UpdateAccessAuthRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 访问授权ID  */
    AccessAuthId string `json:"accessAuthId"`

    /* 访问授权详情 (Optional) */
    AccessAuthView *apigateway.AccessAuthView `json:"accessAuthView"`
}

/*
 * param regionId: 地域ID (Required)
 * param accessAuthId: 访问授权ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAccessAuthRequest(
    regionId string,
    accessAuthId string,
) *UpdateAccessAuthRequest {

	return &UpdateAccessAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accessAuths/{accessAuthId}",
			Method:  "PATCH",
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
 * param accessAuthView: 访问授权详情 (Optional)
 */
func NewUpdateAccessAuthRequestWithAllParams(
    regionId string,
    accessAuthId string,
    accessAuthView *apigateway.AccessAuthView,
) *UpdateAccessAuthRequest {

    return &UpdateAccessAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessAuths/{accessAuthId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AccessAuthId: accessAuthId,
        AccessAuthView: accessAuthView,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAccessAuthRequestWithoutParam() *UpdateAccessAuthRequest {

    return &UpdateAccessAuthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessAuths/{accessAuthId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UpdateAccessAuthRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param accessAuthId: 访问授权ID(Required) */
func (r *UpdateAccessAuthRequest) SetAccessAuthId(accessAuthId string) {
    r.AccessAuthId = accessAuthId
}

/* param accessAuthView: 访问授权详情(Optional) */
func (r *UpdateAccessAuthRequest) SetAccessAuthView(accessAuthView *apigateway.AccessAuthView) {
    r.AccessAuthView = accessAuthView
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAccessAuthRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateAccessAuthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAccessAuthResult `json:"result"`
}

type UpdateAccessAuthResult struct {
    AccessAuthId string `json:"accessAuthId"`
}