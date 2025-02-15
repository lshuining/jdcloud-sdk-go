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
    kms "github.com/lshuining/jdcloud-sdk-go/services/kms/models"
)

type UpdateSecretRequest struct {

    core.JDCloudRequest

    /* 机密ID  */
    SecretId string `json:"secretId"`

    /*   */
    SecretDescCfg *kms.SecretDescCfg `json:"secretDescCfg"`
}

/*
 * param secretId: 机密ID (Required)
 * param secretDescCfg:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateSecretRequest(
    secretId string,
    secretDescCfg *kms.SecretDescCfg,
) *UpdateSecretRequest {

	return &UpdateSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/secret/{secretId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        SecretId: secretId,
        SecretDescCfg: secretDescCfg,
	}
}

/*
 * param secretId: 机密ID (Required)
 * param secretDescCfg:  (Required)
 */
func NewUpdateSecretRequestWithAllParams(
    secretId string,
    secretDescCfg *kms.SecretDescCfg,
) *UpdateSecretRequest {

    return &UpdateSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        SecretId: secretId,
        SecretDescCfg: secretDescCfg,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateSecretRequestWithoutParam() *UpdateSecretRequest {

    return &UpdateSecretRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param secretId: 机密ID(Required) */
func (r *UpdateSecretRequest) SetSecretId(secretId string) {
    r.SecretId = secretId
}

/* param secretDescCfg: (Required) */
func (r *UpdateSecretRequest) SetSecretDescCfg(secretDescCfg *kms.SecretDescCfg) {
    r.SecretDescCfg = secretDescCfg
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateSecretRequest) GetRegionId() string {
    return ""
}

type UpdateSecretResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateSecretResult `json:"result"`
}

type UpdateSecretResult struct {
}