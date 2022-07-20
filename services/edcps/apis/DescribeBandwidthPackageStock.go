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

type DescribeBandwidthPackageStockRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBandwidthPackageStockRequest(
    regionId string,
) *DescribeBandwidthPackageStockRequest {

	return &DescribeBandwidthPackageStockRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bandwidthPackageStock",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 */
func NewDescribeBandwidthPackageStockRequestWithAllParams(
    regionId string,
) *DescribeBandwidthPackageStockRequest {

    return &DescribeBandwidthPackageStockRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackageStock",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBandwidthPackageStockRequestWithoutParam() *DescribeBandwidthPackageStockRequest {

    return &DescribeBandwidthPackageStockRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackageStock",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeBandwidthPackageStockRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBandwidthPackageStockRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBandwidthPackageStockResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBandwidthPackageStockResult `json:"result"`
}

type DescribeBandwidthPackageStockResult struct {
    Region string `json:"region"`
    AvailableBandwidth int `json:"availableBandwidth"`
    AvailableExtraUplinkBandwidth int `json:"availableExtraUplinkBandwidth"`
}