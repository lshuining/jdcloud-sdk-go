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

type DeleteInstanceVoucherRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 实例抵扣券 ID  */
    InstanceVoucherId string `json:"instanceVoucherId"`
}

/*
 * param regionId: 地域 ID (Required)
 * param instanceVoucherId: 实例抵扣券 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteInstanceVoucherRequest(
    regionId string,
    instanceVoucherId string,
) *DeleteInstanceVoucherRequest {

	return &DeleteInstanceVoucherRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceVouchers/{instanceVoucherId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceVoucherId: instanceVoucherId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param instanceVoucherId: 实例抵扣券 ID (Required)
 */
func NewDeleteInstanceVoucherRequestWithAllParams(
    regionId string,
    instanceVoucherId string,
) *DeleteInstanceVoucherRequest {

    return &DeleteInstanceVoucherRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceVouchers/{instanceVoucherId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceVoucherId: instanceVoucherId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteInstanceVoucherRequestWithoutParam() *DeleteInstanceVoucherRequest {

    return &DeleteInstanceVoucherRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceVouchers/{instanceVoucherId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DeleteInstanceVoucherRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceVoucherId: 实例抵扣券 ID(Required) */
func (r *DeleteInstanceVoucherRequest) SetInstanceVoucherId(instanceVoucherId string) {
    r.InstanceVoucherId = instanceVoucherId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteInstanceVoucherRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteInstanceVoucherResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteInstanceVoucherResult `json:"result"`
}

type DeleteInstanceVoucherResult struct {
}