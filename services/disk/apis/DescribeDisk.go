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
    disk "github.com/lshuining/jdcloud-sdk-go/services/disk/models"
)

type DescribeDiskRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDiskRequest(
    regionId string,
    diskId string,
) *DescribeDiskRequest {

	return &DescribeDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks/{diskId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DiskId: diskId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 */
func NewDescribeDiskRequestWithAllParams(
    regionId string,
    diskId string,
) *DescribeDiskRequest {

    return &DescribeDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DiskId: diskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDiskRequestWithoutParam() *DescribeDiskRequest {

    return &DescribeDiskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param diskId: 云硬盘ID(Required) */
func (r *DescribeDiskRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDiskRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDiskResult `json:"result"`
}

type DescribeDiskResult struct {
    Disk disk.Disk `json:"disk"`
}