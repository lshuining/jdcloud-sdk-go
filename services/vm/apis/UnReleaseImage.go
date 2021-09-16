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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type UnReleaseImageRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 镜像ID。  */
    ImageId string `json:"imageId"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUnReleaseImageRequest(
    regionId string,
    imageId string,
) *UnReleaseImageRequest {

	return &UnReleaseImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images/{imageId}:unrelease",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageId: imageId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 */
func NewUnReleaseImageRequestWithAllParams(
    regionId string,
    imageId string,
) *UnReleaseImageRequest {

    return &UnReleaseImageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}:unrelease",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageId: imageId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUnReleaseImageRequestWithoutParam() *UnReleaseImageRequest {

    return &UnReleaseImageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}:unrelease",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *UnReleaseImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageId: 镜像ID。(Required) */
func (r *UnReleaseImageRequest) SetImageId(imageId string) {
    r.ImageId = imageId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UnReleaseImageRequest) GetRegionId() string {
    return r.RegionId
}

type UnReleaseImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UnReleaseImageResult `json:"result"`
}

type UnReleaseImageResult struct {
}