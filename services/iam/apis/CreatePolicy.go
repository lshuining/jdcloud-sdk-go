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
    iam "github.com/lshuining/jdcloud-sdk-go/services/iam/models"
)

type CreatePolicyRequest struct {

    core.JDCloudRequest

    /* 策略信息  */
    CreatePolicyInfo *iam.CreatePolicyInfo `json:"createPolicyInfo"`
}

/*
 * param createPolicyInfo: 策略信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreatePolicyRequest(
    createPolicyInfo *iam.CreatePolicyInfo,
) *CreatePolicyRequest {

	return &CreatePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/policy",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CreatePolicyInfo: createPolicyInfo,
	}
}

/*
 * param createPolicyInfo: 策略信息 (Required)
 */
func NewCreatePolicyRequestWithAllParams(
    createPolicyInfo *iam.CreatePolicyInfo,
) *CreatePolicyRequest {

    return &CreatePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/policy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CreatePolicyInfo: createPolicyInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreatePolicyRequestWithoutParam() *CreatePolicyRequest {

    return &CreatePolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/policy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param createPolicyInfo: 策略信息(Required) */
func (r *CreatePolicyRequest) SetCreatePolicyInfo(createPolicyInfo *iam.CreatePolicyInfo) {
    r.CreatePolicyInfo = createPolicyInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreatePolicyRequest) GetRegionId() string {
    return ""
}

type CreatePolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreatePolicyResult `json:"result"`
}

type CreatePolicyResult struct {
    Policy iam.Policy `json:"policy"`
}