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

type EnableSecretVersionRequest struct {

    core.JDCloudRequest

    /* 机密ID  */
    SecretId string `json:"secretId"`

    /* 机密版本  */
    Version string `json:"version"`
}

/*
 * param secretId: 机密ID (Required)
 * param version: 机密版本 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableSecretVersionRequest(
    secretId string,
    version string,
) *EnableSecretVersionRequest {

	return &EnableSecretVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/secret/{secretId}/version/{version}:enable",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        SecretId: secretId,
        Version: version,
	}
}

/*
 * param secretId: 机密ID (Required)
 * param version: 机密版本 (Required)
 */
func NewEnableSecretVersionRequestWithAllParams(
    secretId string,
    version string,
) *EnableSecretVersionRequest {

    return &EnableSecretVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}/version/{version}:enable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        SecretId: secretId,
        Version: version,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableSecretVersionRequestWithoutParam() *EnableSecretVersionRequest {

    return &EnableSecretVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}/version/{version}:enable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param secretId: 机密ID(Required) */
func (r *EnableSecretVersionRequest) SetSecretId(secretId string) {
    r.SecretId = secretId
}

/* param version: 机密版本(Required) */
func (r *EnableSecretVersionRequest) SetVersion(version string) {
    r.Version = version
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableSecretVersionRequest) GetRegionId() string {
    return ""
}

type EnableSecretVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableSecretVersionResult `json:"result"`
}

type EnableSecretVersionResult struct {
}