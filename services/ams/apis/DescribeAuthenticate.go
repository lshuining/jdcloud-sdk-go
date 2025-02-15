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

type DescribeAuthenticateRequest struct {

    core.JDCloudRequest

    /* PinId  */
    PId string `json:"pId"`

    /* 版本 (Optional) */
    Ver *int `json:"ver"`
}

/*
 * param pId: PinId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAuthenticateRequest(
    pId string,
) *DescribeAuthenticateRequest {

	return &DescribeAuthenticateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/appManager/{pId}/authenticates",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PId: pId,
	}
}

/*
 * param pId: PinId (Required)
 * param ver: 版本 (Optional)
 */
func NewDescribeAuthenticateRequestWithAllParams(
    pId string,
    ver *int,
) *DescribeAuthenticateRequest {

    return &DescribeAuthenticateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/appManager/{pId}/authenticates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PId: pId,
        Ver: ver,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAuthenticateRequestWithoutParam() *DescribeAuthenticateRequest {

    return &DescribeAuthenticateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/appManager/{pId}/authenticates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pId: PinId(Required) */
func (r *DescribeAuthenticateRequest) SetPId(pId string) {
    r.PId = pId
}

/* param ver: 版本(Optional) */
func (r *DescribeAuthenticateRequest) SetVer(ver int) {
    r.Ver = &ver
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAuthenticateRequest) GetRegionId() string {
    return ""
}

type DescribeAuthenticateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAuthenticateResult `json:"result"`
}

type DescribeAuthenticateResult struct {
    PId string `json:"pId"`
    Ver int `json:"ver"`
    Blacklist int `json:"blacklist"`
    Status int `json:"status"`
    License string `json:"license"`
}