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

type OpenIotCardRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡号码列表(单次提交最多不超过200个号码)  */
    Iccids []string `json:"iccids"`
}

/*
 * param regionId: Region ID (Required)
 * param iccids: 物联网卡号码列表(单次提交最多不超过200个号码) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOpenIotCardRequest(
    regionId string,
    iccids []string,
) *OpenIotCardRequest {

	return &OpenIotCardRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/openIotCard",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Iccids: iccids,
	}
}

/*
 * param regionId: Region ID (Required)
 * param iccids: 物联网卡号码列表(单次提交最多不超过200个号码) (Required)
 */
func NewOpenIotCardRequestWithAllParams(
    regionId string,
    iccids []string,
) *OpenIotCardRequest {

    return &OpenIotCardRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/openIotCard",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Iccids: iccids,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOpenIotCardRequestWithoutParam() *OpenIotCardRequest {

    return &OpenIotCardRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/openIotCard",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *OpenIotCardRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param iccids: 物联网卡号码列表(单次提交最多不超过200个号码)(Required) */
func (r *OpenIotCardRequest) SetIccids(iccids []string) {
    r.Iccids = iccids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OpenIotCardRequest) GetRegionId() string {
    return r.RegionId
}

type OpenIotCardResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OpenIotCardResult `json:"result"`
}

type OpenIotCardResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result []iotlink.OperationIotlinkResp `json:"result"`
}