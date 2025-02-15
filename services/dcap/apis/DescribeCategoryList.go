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
    dcap "github.com/lshuining/jdcloud-sdk-go/services/dcap/models"
)

type DescribeCategoryListRequest struct {

    core.JDCloudRequest
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCategoryListRequest(
) *DescribeCategoryListRequest {

	return &DescribeCategoryListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/category",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 */
func NewDescribeCategoryListRequestWithAllParams(
) *DescribeCategoryListRequest {

    return &DescribeCategoryListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/category",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCategoryListRequestWithoutParam() *DescribeCategoryListRequest {

    return &DescribeCategoryListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/category",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCategoryListRequest) GetRegionId() string {
    return ""
}

type DescribeCategoryListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCategoryListResult `json:"result"`
}

type DescribeCategoryListResult struct {
    List []dcap.CategoryDesc `json:"list"`
}