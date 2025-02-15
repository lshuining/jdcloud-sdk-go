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
    resourcetag "github.com/lshuining/jdcloud-sdk-go/services/resourcetag/models"
)

type UnTagResourcesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 解绑标签参数  */
    UnTagResources *resourcetag.UnTagResourcesReqVo `json:"unTagResources"`
}

/*
 * param regionId: Region ID (Required)
 * param unTagResources: 解绑标签参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUnTagResourcesRequest(
    regionId string,
    unTagResources *resourcetag.UnTagResourcesReqVo,
) *UnTagResourcesRequest {

	return &UnTagResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/tags:unTagResources",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UnTagResources: unTagResources,
	}
}

/*
 * param regionId: Region ID (Required)
 * param unTagResources: 解绑标签参数 (Required)
 */
func NewUnTagResourcesRequestWithAllParams(
    regionId string,
    unTagResources *resourcetag.UnTagResourcesReqVo,
) *UnTagResourcesRequest {

    return &UnTagResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tags:unTagResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UnTagResources: unTagResources,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUnTagResourcesRequestWithoutParam() *UnTagResourcesRequest {

    return &UnTagResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tags:unTagResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UnTagResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param unTagResources: 解绑标签参数(Required) */
func (r *UnTagResourcesRequest) SetUnTagResources(unTagResources *resourcetag.UnTagResourcesReqVo) {
    r.UnTagResources = unTagResources
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UnTagResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type UnTagResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UnTagResourcesResult `json:"result"`
}

type UnTagResourcesResult struct {
    Data resourcetag.UnTagResourcesResVo `json:"data"`
}