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

type GetJobRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Job uuid  */
    Id string `json:"id"`
}

/*
 * param regionId: Region ID (Required)
 * param id: Job uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetJobRequest(
    regionId string,
    id string,
) *GetJobRequest {

	return &GetJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/jobs/{id}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: Region ID (Required)
 * param id: Job uuid (Required)
 */
func NewGetJobRequestWithAllParams(
    regionId string,
    id string,
) *GetJobRequest {

    return &GetJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobs/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetJobRequestWithoutParam() *GetJobRequest {

    return &GetJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobs/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: Job uuid(Required) */
func (r *GetJobRequest) SetId(id string) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetJobRequest) GetRegionId() string {
    return r.RegionId
}

type GetJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetJobResult `json:"result"`
}

type GetJobResult struct {
    Job compile.Job `json:"job"`
}