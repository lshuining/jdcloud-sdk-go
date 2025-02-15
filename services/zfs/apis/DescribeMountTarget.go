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
    zfs "github.com/lshuining/jdcloud-sdk-go/services/zfs/models"
)

type DescribeMountTargetRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 挂载目标ID  */
    MountTargetId string `json:"mountTargetId"`
}

/*
 * param regionId: 地域ID (Required)
 * param mountTargetId: 挂载目标ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMountTargetRequest(
    regionId string,
    mountTargetId string,
) *DescribeMountTargetRequest {

	return &DescribeMountTargetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/mountTargets/{mountTargetId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        MountTargetId: mountTargetId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param mountTargetId: 挂载目标ID (Required)
 */
func NewDescribeMountTargetRequestWithAllParams(
    regionId string,
    mountTargetId string,
) *DescribeMountTargetRequest {

    return &DescribeMountTargetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/mountTargets/{mountTargetId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        MountTargetId: mountTargetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMountTargetRequestWithoutParam() *DescribeMountTargetRequest {

    return &DescribeMountTargetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/mountTargets/{mountTargetId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeMountTargetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param mountTargetId: 挂载目标ID(Required) */
func (r *DescribeMountTargetRequest) SetMountTargetId(mountTargetId string) {
    r.MountTargetId = mountTargetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMountTargetRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMountTargetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMountTargetResult `json:"result"`
}

type DescribeMountTargetResult struct {
    MountTarget zfs.MountTarget `json:"mountTarget"`
}