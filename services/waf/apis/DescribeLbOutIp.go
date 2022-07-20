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

type DescribeLbOutIpRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLbOutIpRequest(
    regionId string,
) *DescribeLbOutIpRequest {

	return &DescribeLbOutIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/admin:lboutIp",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 */
func NewDescribeLbOutIpRequestWithAllParams(
    regionId string,
) *DescribeLbOutIpRequest {

    return &DescribeLbOutIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/admin:lboutIp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLbOutIpRequestWithoutParam() *DescribeLbOutIpRequest {

    return &DescribeLbOutIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/admin:lboutIp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *DescribeLbOutIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLbOutIpRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLbOutIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLbOutIpResult `json:"result"`
}

type DescribeLbOutIpResult struct {
    List interface{} `json:"list"`
}