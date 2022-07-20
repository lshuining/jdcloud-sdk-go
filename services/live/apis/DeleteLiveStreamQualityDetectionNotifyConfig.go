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

type DeleteLiveStreamQualityDetectionNotifyConfigRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`
}

/*
 * param publishDomain: 推流域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveStreamQualityDetectionNotifyConfigRequest(
    publishDomain string,
) *DeleteLiveStreamQualityDetectionNotifyConfigRequest {

	return &DeleteLiveStreamQualityDetectionNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/qualityDetectionNotifys/{publishDomain}:config",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 */
func NewDeleteLiveStreamQualityDetectionNotifyConfigRequestWithAllParams(
    publishDomain string,
) *DeleteLiveStreamQualityDetectionNotifyConfigRequest {

    return &DeleteLiveStreamQualityDetectionNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionNotifys/{publishDomain}:config",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveStreamQualityDetectionNotifyConfigRequestWithoutParam() *DeleteLiveStreamQualityDetectionNotifyConfigRequest {

    return &DeleteLiveStreamQualityDetectionNotifyConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionNotifys/{publishDomain}:config",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *DeleteLiveStreamQualityDetectionNotifyConfigRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveStreamQualityDetectionNotifyConfigRequest) GetRegionId() string {
    return ""
}

type DeleteLiveStreamQualityDetectionNotifyConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveStreamQualityDetectionNotifyConfigResult `json:"result"`
}

type DeleteLiveStreamQualityDetectionNotifyConfigResult struct {
}