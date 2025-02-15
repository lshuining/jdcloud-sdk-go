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
    ydsms "github.com/lshuining/jdcloud-sdk-go/services/ydsms/models"
)

type GetSmsSignByIdUsingGETRequest struct {

    core.JDCloudRequest

    /* 签名id  */
    SignId string `json:"signId"`
}

/*
 * param signId: 签名id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSmsSignByIdUsingGETRequest(
    signId string,
) *GetSmsSignByIdUsingGETRequest {

	return &GetSmsSignByIdUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsSigns/{signId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        SignId: signId,
	}
}

/*
 * param signId: 签名id (Required)
 */
func NewGetSmsSignByIdUsingGETRequestWithAllParams(
    signId string,
) *GetSmsSignByIdUsingGETRequest {

    return &GetSmsSignByIdUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsSigns/{signId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        SignId: signId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSmsSignByIdUsingGETRequestWithoutParam() *GetSmsSignByIdUsingGETRequest {

    return &GetSmsSignByIdUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsSigns/{signId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param signId: 签名id(Required) */
func (r *GetSmsSignByIdUsingGETRequest) SetSignId(signId string) {
    r.SignId = signId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSmsSignByIdUsingGETRequest) GetRegionId() string {
    return ""
}

type GetSmsSignByIdUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSmsSignByIdUsingGETResult `json:"result"`
}

type GetSmsSignByIdUsingGETResult struct {
    SmsSignVO ydsms.SmsSignVO `json:"smsSignVO"`
}