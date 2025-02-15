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
    cdn "github.com/lshuining/jdcloud-sdk-go/services/cdn/models"
)

type QueryNetProtectionRulesRequest struct {

    core.JDCloudRequest
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryNetProtectionRulesRequest(
) *QueryNetProtectionRulesRequest {

	return &QueryNetProtectionRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/netProtectionRules",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 */
func NewQueryNetProtectionRulesRequestWithAllParams(
) *QueryNetProtectionRulesRequest {

    return &QueryNetProtectionRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/netProtectionRules",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryNetProtectionRulesRequestWithoutParam() *QueryNetProtectionRulesRequest {

    return &QueryNetProtectionRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/netProtectionRules",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryNetProtectionRulesRequest) GetRegionId() string {
    return ""
}

type QueryNetProtectionRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryNetProtectionRulesResult `json:"result"`
}

type QueryNetProtectionRulesResult struct {
    SwitchStatus string `json:"switchStatus"`
    SrcNewConnLimitEnable string `json:"srcNewConnLimitEnable"`
    DstNewConnLimitEnable string `json:"dstNewConnLimitEnable"`
    DatagramRangeMin int64 `json:"datagramRangeMin"`
    DatagramRangeMax int64 `json:"datagramRangeMax"`
    SrcNewConnLimitValue int64 `json:"srcNewConnLimitValue"`
    DstNewConnLimitValue int64 `json:"dstNewConnLimitValue"`
    GeoBlack []cdn.GeoArea `json:"geoBlack"`
    IpBlack []string `json:"ipBlack"`
    IpWhite []string `json:"ipWhite"`
}