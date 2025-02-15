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

type DescribeTagsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 标签参数  */
    TagKeysVo *resourcetag.TagsReqVo `json:"tagKeysVo"`
}

/*
 * param regionId: Region ID (Required)
 * param tagKeysVo: 标签参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTagsRequest(
    regionId string,
    tagKeysVo *resourcetag.TagsReqVo,
) *DescribeTagsRequest {

	return &DescribeTagsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/describeTags",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TagKeysVo: tagKeysVo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param tagKeysVo: 标签参数 (Required)
 */
func NewDescribeTagsRequestWithAllParams(
    regionId string,
    tagKeysVo *resourcetag.TagsReqVo,
) *DescribeTagsRequest {

    return &DescribeTagsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeTags",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TagKeysVo: tagKeysVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTagsRequestWithoutParam() *DescribeTagsRequest {

    return &DescribeTagsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeTags",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeTagsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param tagKeysVo: 标签参数(Required) */
func (r *DescribeTagsRequest) SetTagKeysVo(tagKeysVo *resourcetag.TagsReqVo) {
    r.TagKeysVo = tagKeysVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTagsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTagsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTagsResult `json:"result"`
}

type DescribeTagsResult struct {
    Data resourcetag.TagsResVo `json:"data"`
}