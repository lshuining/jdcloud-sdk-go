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

type SetLiveRestartAuthKeyRequest struct {

    core.JDCloudRequest

    /* 直播回看播放域名  */
    RestartDomain string `json:"restartDomain"`

    /* 直播回看播放鉴权状态
  on: 开启
  off: 关闭
- 当回看播放鉴权状态on(开启)时,authKey不能为空
 (Optional) */
    AuthStatus *string `json:"authStatus"`

    /* 直播回看播放鉴权key
- 取值: 支持大小写字母和数字 长度6-32位
 (Optional) */
    AuthKey *string `json:"authKey"`
}

/*
 * param restartDomain: 直播回看播放域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetLiveRestartAuthKeyRequest(
    restartDomain string,
) *SetLiveRestartAuthKeyRequest {

	return &SetLiveRestartAuthKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveRestartAuthKey",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RestartDomain: restartDomain,
	}
}

/*
 * param restartDomain: 直播回看播放域名 (Required)
 * param authStatus: 直播回看播放鉴权状态
  on: 开启
  off: 关闭
- 当回看播放鉴权状态on(开启)时,authKey不能为空
 (Optional)
 * param authKey: 直播回看播放鉴权key
- 取值: 支持大小写字母和数字 长度6-32位
 (Optional)
 */
func NewSetLiveRestartAuthKeyRequestWithAllParams(
    restartDomain string,
    authStatus *string,
    authKey *string,
) *SetLiveRestartAuthKeyRequest {

    return &SetLiveRestartAuthKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRestartAuthKey",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RestartDomain: restartDomain,
        AuthStatus: authStatus,
        AuthKey: authKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetLiveRestartAuthKeyRequestWithoutParam() *SetLiveRestartAuthKeyRequest {

    return &SetLiveRestartAuthKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRestartAuthKey",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param restartDomain: 直播回看播放域名(Required) */
func (r *SetLiveRestartAuthKeyRequest) SetRestartDomain(restartDomain string) {
    r.RestartDomain = restartDomain
}

/* param authStatus: 直播回看播放鉴权状态
  on: 开启
  off: 关闭
- 当回看播放鉴权状态on(开启)时,authKey不能为空
(Optional) */
func (r *SetLiveRestartAuthKeyRequest) SetAuthStatus(authStatus string) {
    r.AuthStatus = &authStatus
}

/* param authKey: 直播回看播放鉴权key
- 取值: 支持大小写字母和数字 长度6-32位
(Optional) */
func (r *SetLiveRestartAuthKeyRequest) SetAuthKey(authKey string) {
    r.AuthKey = &authKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetLiveRestartAuthKeyRequest) GetRegionId() string {
    return ""
}

type SetLiveRestartAuthKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetLiveRestartAuthKeyResult `json:"result"`
}

type SetLiveRestartAuthKeyResult struct {
}