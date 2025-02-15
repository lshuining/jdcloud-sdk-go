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
    ossopenapi "github.com/lshuining/jdcloud-sdk-go/services/ossopenapi/models"
)

type GetBackSourceConfigurationRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* Bucket名称  */
    BucketName string `json:"bucketName"`
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetBackSourceConfigurationRequest(
    regionId string,
    bucketName string,
) *GetBackSourceConfigurationRequest {

	return &GetBackSourceConfigurationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketName}/backSource",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BucketName: bucketName,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 */
func NewGetBackSourceConfigurationRequestWithAllParams(
    regionId string,
    bucketName string,
) *GetBackSourceConfigurationRequest {

    return &GetBackSourceConfigurationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/backSource",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BucketName: bucketName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetBackSourceConfigurationRequestWithoutParam() *GetBackSourceConfigurationRequest {

    return &GetBackSourceConfigurationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/backSource",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *GetBackSourceConfigurationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bucketName: Bucket名称(Required) */
func (r *GetBackSourceConfigurationRequest) SetBucketName(bucketName string) {
    r.BucketName = bucketName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetBackSourceConfigurationRequest) GetRegionId() string {
    return r.RegionId
}

type GetBackSourceConfigurationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetBackSourceConfigurationResult `json:"result"`
}

type GetBackSourceConfigurationResult struct {
    BackSourceRules []ossopenapi.BackSourceRule `json:"backSourceRules"`
}