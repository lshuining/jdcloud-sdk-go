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
    jdccs "github.com/lshuining/jdcloud-sdk-go/services/jdccs/models"
)

type DescribeBandwidthTrafficRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 带宽（出口）实例ID  */
    BandwidthId string `json:"bandwidthId"`
}

/*
 * param idc: IDC机房ID (Required)
 * param bandwidthId: 带宽（出口）实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBandwidthTrafficRequest(
    idc string,
    bandwidthId string,
) *DescribeBandwidthTrafficRequest {

	return &DescribeBandwidthTrafficRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/bandwidthTraffics/{bandwidthId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
        BandwidthId: bandwidthId,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param bandwidthId: 带宽（出口）实例ID (Required)
 */
func NewDescribeBandwidthTrafficRequestWithAllParams(
    idc string,
    bandwidthId string,
) *DescribeBandwidthTrafficRequest {

    return &DescribeBandwidthTrafficRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/bandwidthTraffics/{bandwidthId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        BandwidthId: bandwidthId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBandwidthTrafficRequestWithoutParam() *DescribeBandwidthTrafficRequest {

    return &DescribeBandwidthTrafficRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/bandwidthTraffics/{bandwidthId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *DescribeBandwidthTrafficRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param bandwidthId: 带宽（出口）实例ID(Required) */
func (r *DescribeBandwidthTrafficRequest) SetBandwidthId(bandwidthId string) {
    r.BandwidthId = bandwidthId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBandwidthTrafficRequest) GetRegionId() string {
    return ""
}

type DescribeBandwidthTrafficResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBandwidthTrafficResult `json:"result"`
}

type DescribeBandwidthTrafficResult struct {
    BandwidthTraffic jdccs.BandwidthTraffic `json:"bandwidthTraffic"`
}