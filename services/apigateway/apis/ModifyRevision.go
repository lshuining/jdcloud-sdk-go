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

type ModifyRevisionRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本ID  */
    RevisionId string `json:"revisionId"`

    /* 修订备注 (Optional) */
    RevisionNote *string `json:"revisionNote"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revisionId: 版本ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyRevisionRequest(
    regionId string,
    apiGroupId string,
    revisionId string,
) *ModifyRevisionRequest {

	return &ModifyRevisionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revisionId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        RevisionId: revisionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revisionId: 版本ID (Required)
 * param revisionNote: 修订备注 (Optional)
 */
func NewModifyRevisionRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revisionId string,
    revisionNote *string,
) *ModifyRevisionRequest {

    return &ModifyRevisionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revisionId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        RevisionId: revisionId,
        RevisionNote: revisionNote,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyRevisionRequestWithoutParam() *ModifyRevisionRequest {

    return &ModifyRevisionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revisionId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyRevisionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *ModifyRevisionRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revisionId: 版本ID(Required) */
func (r *ModifyRevisionRequest) SetRevisionId(revisionId string) {
    r.RevisionId = revisionId
}

/* param revisionNote: 修订备注(Optional) */
func (r *ModifyRevisionRequest) SetRevisionNote(revisionNote string) {
    r.RevisionNote = &revisionNote
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyRevisionRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyRevisionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyRevisionResult `json:"result"`
}

type ModifyRevisionResult struct {
    UpdateRevision bool `json:"updateRevision"`
}