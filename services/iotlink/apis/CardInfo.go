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
    iotlink "github.com/lshuining/jdcloud-sdk-go/services/iotlink/models"
)

type CardInfoRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡iccid  */
    Iccid string `json:"iccid"`
}

/*
 * param regionId: Region ID (Required)
 * param iccid: 物联网卡iccid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCardInfoRequest(
    regionId string,
    iccid string,
) *CardInfoRequest {

	return &CardInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cardInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Iccid: iccid,
	}
}

/*
 * param regionId: Region ID (Required)
 * param iccid: 物联网卡iccid (Required)
 */
func NewCardInfoRequestWithAllParams(
    regionId string,
    iccid string,
) *CardInfoRequest {

    return &CardInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cardInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Iccid: iccid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCardInfoRequestWithoutParam() *CardInfoRequest {

    return &CardInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cardInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CardInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param iccid: 物联网卡iccid(Required) */
func (r *CardInfoRequest) SetIccid(iccid string) {
    r.Iccid = iccid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CardInfoRequest) GetRegionId() string {
    return r.RegionId
}

type CardInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CardInfoResult `json:"result"`
}

type CardInfoResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result iotlink.CardInfoResp `json:"result"`
}