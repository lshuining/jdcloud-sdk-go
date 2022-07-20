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

type QueryRemainingUsingGETRequest struct {

    core.JDCloudRequest

    /* 通道类型，1通道短信，2官方短信  */
    PackageType int `json:"packageType"`
}

/*
 * param packageType: 通道类型，1通道短信，2官方短信 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRemainingUsingGETRequest(
    packageType int,
) *QueryRemainingUsingGETRequest {

	return &QueryRemainingUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsUsers/{packageType}:queryRemaining",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PackageType: packageType,
	}
}

/*
 * param packageType: 通道类型，1通道短信，2官方短信 (Required)
 */
func NewQueryRemainingUsingGETRequestWithAllParams(
    packageType int,
) *QueryRemainingUsingGETRequest {

    return &QueryRemainingUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsUsers/{packageType}:queryRemaining",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PackageType: packageType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRemainingUsingGETRequestWithoutParam() *QueryRemainingUsingGETRequest {

    return &QueryRemainingUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsUsers/{packageType}:queryRemaining",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param packageType: 通道类型，1通道短信，2官方短信(Required) */
func (r *QueryRemainingUsingGETRequest) SetPackageType(packageType int) {
    r.PackageType = packageType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRemainingUsingGETRequest) GetRegionId() string {
    return ""
}

type QueryRemainingUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRemainingUsingGETResult `json:"result"`
}

type QueryRemainingUsingGETResult struct {
    PackageType int `json:"packageType"`
    Remaining int64 `json:"remaining"`
    TotalCount int64 `json:"totalCount"`
    RemainingRate string `json:"remainingRate"`
}