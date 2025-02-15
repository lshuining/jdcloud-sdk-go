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
    kms "github.com/lshuining/jdcloud-sdk-go/services/kms/models"
)

type DescribeKeyDetailRequest struct {

    core.JDCloudRequest

    /* 密钥ID  */
    KeyId string `json:"keyId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param keyId: 密钥ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeKeyDetailRequest(
    keyId string,
) *DescribeKeyDetailRequest {

	return &DescribeKeyDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/key/{keyId}:describeKeyDetail",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        KeyId: keyId,
	}
}

/*
 * param keyId: 密钥ID (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 */
func NewDescribeKeyDetailRequestWithAllParams(
    keyId string,
    pageNumber *int,
    pageSize *int,
) *DescribeKeyDetailRequest {

    return &DescribeKeyDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:describeKeyDetail",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        KeyId: keyId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeKeyDetailRequestWithoutParam() *DescribeKeyDetailRequest {

    return &DescribeKeyDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:describeKeyDetail",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyId: 密钥ID(Required) */
func (r *DescribeKeyDetailRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeKeyDetailRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeKeyDetailRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeKeyDetailRequest) GetRegionId() string {
    return ""
}

type DescribeKeyDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeKeyDetailResult `json:"result"`
}

type DescribeKeyDetailResult struct {
    KeyDetail kms.KeyDetail `json:"keyDetail"`
}