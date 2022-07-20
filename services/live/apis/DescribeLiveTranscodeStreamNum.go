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

type DescribeLiveTranscodeStreamNumRequest struct {

    core.JDCloudRequest

    /* 推流域名 (Optional) */
    DomainName *string `json:"domainName"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveTranscodeStreamNumRequest(
) *DescribeLiveTranscodeStreamNumRequest {

	return &DescribeLiveTranscodeStreamNumRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeLiveTranscodeStreamNum",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param domainName: 推流域名 (Optional)
 */
func NewDescribeLiveTranscodeStreamNumRequestWithAllParams(
    domainName *string,
) *DescribeLiveTranscodeStreamNumRequest {

    return &DescribeLiveTranscodeStreamNumRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveTranscodeStreamNum",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveTranscodeStreamNumRequestWithoutParam() *DescribeLiveTranscodeStreamNumRequest {

    return &DescribeLiveTranscodeStreamNumRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveTranscodeStreamNum",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 推流域名(Optional) */
func (r *DescribeLiveTranscodeStreamNumRequest) SetDomainName(domainName string) {
    r.DomainName = &domainName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveTranscodeStreamNumRequest) GetRegionId() string {
    return ""
}

type DescribeLiveTranscodeStreamNumResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveTranscodeStreamNumResult `json:"result"`
}

type DescribeLiveTranscodeStreamNumResult struct {
    Datetime string `json:"datetime"`
    StreamCount int `json:"streamCount"`
}