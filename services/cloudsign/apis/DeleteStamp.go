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

type DeleteStampRequest struct {

    core.JDCloudRequest

    /* 印章ID  */
    StampId string `json:"stampId"`
}

/*
 * param stampId: 印章ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteStampRequest(
    stampId string,
) *DeleteStampRequest {

	return &DeleteStampRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/stamp/{stampId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        StampId: stampId,
	}
}

/*
 * param stampId: 印章ID (Required)
 */
func NewDeleteStampRequestWithAllParams(
    stampId string,
) *DeleteStampRequest {

    return &DeleteStampRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/stamp/{stampId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        StampId: stampId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteStampRequestWithoutParam() *DeleteStampRequest {

    return &DeleteStampRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/stamp/{stampId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param stampId: 印章ID(Required) */
func (r *DeleteStampRequest) SetStampId(stampId string) {
    r.StampId = stampId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteStampRequest) GetRegionId() string {
    return ""
}

type DeleteStampResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteStampResult `json:"result"`
}

type DeleteStampResult struct {
}