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
    starshield "github.com/lshuining/jdcloud-sdk-go/services/starshield/models"
)

type DescribePackageRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 套餐类型  */
    PackType int `json:"packType"`
}

/*
 * param regionId: 地域ID (Required)
 * param packType: 套餐类型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePackageRequest(
    regionId string,
    packType int,
) *DescribePackageRequest {

	return &DescribePackageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/packages/{packType}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PackType: packType,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param packType: 套餐类型 (Required)
 */
func NewDescribePackageRequestWithAllParams(
    regionId string,
    packType int,
) *DescribePackageRequest {

    return &DescribePackageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/packages/{packType}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PackType: packType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePackageRequestWithoutParam() *DescribePackageRequest {

    return &DescribePackageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/packages/{packType}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribePackageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param packType: 套餐类型(Required) */
func (r *DescribePackageRequest) SetPackType(packType int) {
    r.PackType = packType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePackageRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePackageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePackageResult `json:"result"`
}

type DescribePackageResult struct {
    DescribePackRes starshield.DescribePackRes `json:"describePackRes"`
}