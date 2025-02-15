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

type DescribeLiveStreamRecordNotifyConfigRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`
}

/*
 * param publishDomain: 推流域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveStreamRecordNotifyConfigRequest(
    publishDomain string,
) *DescribeLiveStreamRecordNotifyConfigRequest {

	return &DescribeLiveStreamRecordNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/recordNotifys/{publishDomain}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 */
func NewDescribeLiveStreamRecordNotifyConfigRequestWithAllParams(
    publishDomain string,
) *DescribeLiveStreamRecordNotifyConfigRequest {

    return &DescribeLiveStreamRecordNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordNotifys/{publishDomain}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveStreamRecordNotifyConfigRequestWithoutParam() *DescribeLiveStreamRecordNotifyConfigRequest {

    return &DescribeLiveStreamRecordNotifyConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordNotifys/{publishDomain}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *DescribeLiveStreamRecordNotifyConfigRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveStreamRecordNotifyConfigRequest) GetRegionId() string {
    return ""
}

type DescribeLiveStreamRecordNotifyConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveStreamRecordNotifyConfigResult `json:"result"`
}

type DescribeLiveStreamRecordNotifyConfigResult struct {
    PublishDomain string `json:"publishDomain"`
    NotifyUrl string `json:"notifyUrl"`
}