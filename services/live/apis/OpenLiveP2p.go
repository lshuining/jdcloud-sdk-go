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

type OpenLiveP2pRequest struct {

    core.JDCloudRequest

    /* 播放域名 (Optional) */
    PlayDomain *string `json:"playDomain"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOpenLiveP2pRequest(
) *OpenLiveP2pRequest {

	return &OpenLiveP2pRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveP2p:open",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param playDomain: 播放域名 (Optional)
 */
func NewOpenLiveP2pRequestWithAllParams(
    playDomain *string,
) *OpenLiveP2pRequest {

    return &OpenLiveP2pRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveP2p:open",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        PlayDomain: playDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOpenLiveP2pRequestWithoutParam() *OpenLiveP2pRequest {

    return &OpenLiveP2pRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveP2p:open",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param playDomain: 播放域名(Optional) */
func (r *OpenLiveP2pRequest) SetPlayDomain(playDomain string) {
    r.PlayDomain = &playDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OpenLiveP2pRequest) GetRegionId() string {
    return ""
}

type OpenLiveP2pResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OpenLiveP2pResult `json:"result"`
}

type OpenLiveP2pResult struct {
}