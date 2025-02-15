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

type BatchSetExtraCacheTimeRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 状态码和过期时间，多个用英文分号分隔,如404:3;500:10;异常状态码 ["4xx","400", "401",  "402", "404", "405", "406", "407", "408", "409", "410", "411", "412", "413", "414", "415", "416", "417",  "5xx","500", "501", "502", "503", "504", "505"]中的其中一个,缓存时间(单位:秒) (Optional) */
    Content *string `json:"content"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchSetExtraCacheTimeRequest(
    domain string,
) *BatchSetExtraCacheTimeRequest {

	return &BatchSetExtraCacheTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/extraCacheTime:batchSet",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param content: 状态码和过期时间，多个用英文分号分隔,如404:3;500:10;异常状态码 ["4xx","400", "401",  "402", "404", "405", "406", "407", "408", "409", "410", "411", "412", "413", "414", "415", "416", "417",  "5xx","500", "501", "502", "503", "504", "505"]中的其中一个,缓存时间(单位:秒) (Optional)
 */
func NewBatchSetExtraCacheTimeRequestWithAllParams(
    domain string,
    content *string,
) *BatchSetExtraCacheTimeRequest {

    return &BatchSetExtraCacheTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/extraCacheTime:batchSet",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Content: content,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchSetExtraCacheTimeRequestWithoutParam() *BatchSetExtraCacheTimeRequest {

    return &BatchSetExtraCacheTimeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/extraCacheTime:batchSet",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *BatchSetExtraCacheTimeRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param content: 状态码和过期时间，多个用英文分号分隔,如404:3;500:10;异常状态码 ["4xx","400", "401",  "402", "404", "405", "406", "407", "408", "409", "410", "411", "412", "413", "414", "415", "416", "417",  "5xx","500", "501", "502", "503", "504", "505"]中的其中一个,缓存时间(单位:秒)(Optional) */
func (r *BatchSetExtraCacheTimeRequest) SetContent(content string) {
    r.Content = &content
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchSetExtraCacheTimeRequest) GetRegionId() string {
    return ""
}

type BatchSetExtraCacheTimeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchSetExtraCacheTimeResult `json:"result"`
}

type BatchSetExtraCacheTimeResult struct {
}