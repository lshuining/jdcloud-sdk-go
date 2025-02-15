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

type StartJobRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Job uuid  */
    Id string `json:"id"`

    /* 类型branch/commit/tag (Optional) */
    Category *string `json:"category"`

    /* 类型对应的值 (Optional) */
    Branch *string `json:"branch"`
}

/*
 * param regionId: Region ID (Required)
 * param id: Job uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStartJobRequest(
    regionId string,
    id string,
) *StartJobRequest {

	return &StartJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/jobs/{id}:start",
			Method:  "POST",
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
 * param category: 类型branch/commit/tag (Optional)
 * param branch: 类型对应的值 (Optional)
 */
func NewStartJobRequestWithAllParams(
    regionId string,
    id string,
    category *string,
    branch *string,
) *StartJobRequest {

    return &StartJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobs/{id}:start",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Category: category,
        Branch: branch,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStartJobRequestWithoutParam() *StartJobRequest {

    return &StartJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobs/{id}:start",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *StartJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: Job uuid(Required) */
func (r *StartJobRequest) SetId(id string) {
    r.Id = id
}

/* param category: 类型branch/commit/tag(Optional) */
func (r *StartJobRequest) SetCategory(category string) {
    r.Category = &category
}

/* param branch: 类型对应的值(Optional) */
func (r *StartJobRequest) SetBranch(branch string) {
    r.Branch = &branch
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StartJobRequest) GetRegionId() string {
    return r.RegionId
}

type StartJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StartJobResult `json:"result"`
}

type StartJobResult struct {
    Commitresult bool `json:"commitresult"`
    BuildUuid string `json:"buildUuid"`
}