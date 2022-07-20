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
    compile "github.com/lshuining/jdcloud-sdk-go/services/compile/models"
)

type CreateDemoRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDemoRequest(
    regionId string,
) *CreateDemoRequest {

	return &CreateDemoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/demo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 */
func NewCreateDemoRequestWithAllParams(
    regionId string,
) *CreateDemoRequest {

    return &CreateDemoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/demo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDemoRequestWithoutParam() *CreateDemoRequest {

    return &CreateDemoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/demo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateDemoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDemoRequest) GetRegionId() string {
    return r.RegionId
}

type CreateDemoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDemoResult `json:"result"`
}

type CreateDemoResult struct {
    Job compile.Job `json:"job"`
}