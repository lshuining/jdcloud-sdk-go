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

type DeleteSecretRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Secret Name  */
    Name string `json:"name"`
}

/*
 * param regionId: Region ID (Required)
 * param name: Secret Name (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteSecretRequest(
    regionId string,
    name string,
) *DeleteSecretRequest {

	return &DeleteSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/secrets/{name}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: Secret Name (Required)
 */
func NewDeleteSecretRequestWithAllParams(
    regionId string,
    name string,
) *DeleteSecretRequest {

    return &DeleteSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/secrets/{name}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteSecretRequestWithoutParam() *DeleteSecretRequest {

    return &DeleteSecretRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/secrets/{name}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteSecretRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: Secret Name(Required) */
func (r *DeleteSecretRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteSecretRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteSecretResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteSecretResult `json:"result"`
}

type DeleteSecretResult struct {
}