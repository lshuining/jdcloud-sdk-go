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

type SetLiveRestartDomainCertificateRequest struct {

    core.JDCloudRequest

    /* 直播回看域名  */
    RestartDomain string `json:"restartDomain"`

    /* 直播回看证书状态
  on: 开启
  off: 关闭
- 当播放证书状态on(开启)时,cert和key不能为空
  */
    CertStatus string `json:"certStatus"`

    /* 直播回看证书
- 取值: 最大支持4098
- 当播放证书状态on(开启)时,cert不能为空
 (Optional) */
    Cert *string `json:"cert"`

    /* 直播回看证书key
- 取值: 最大支持2048
- 当播放证书状态on(开启)时,key不能为空
 (Optional) */
    Key *string `json:"key"`

    /* 直播回看证书别名
- 取值: 支持大小写字母和数字 长度最大256
 (Optional) */
    Title *string `json:"title"`
}

/*
 * param restartDomain: 直播回看域名 (Required)
 * param certStatus: 直播回看证书状态
  on: 开启
  off: 关闭
- 当播放证书状态on(开启)时,cert和key不能为空
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetLiveRestartDomainCertificateRequest(
    restartDomain string,
    certStatus string,
) *SetLiveRestartDomainCertificateRequest {

	return &SetLiveRestartDomainCertificateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveRestartDomainCertificate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RestartDomain: restartDomain,
        CertStatus: certStatus,
	}
}

/*
 * param restartDomain: 直播回看域名 (Required)
 * param certStatus: 直播回看证书状态
  on: 开启
  off: 关闭
- 当播放证书状态on(开启)时,cert和key不能为空
 (Required)
 * param cert: 直播回看证书
- 取值: 最大支持4098
- 当播放证书状态on(开启)时,cert不能为空
 (Optional)
 * param key: 直播回看证书key
- 取值: 最大支持2048
- 当播放证书状态on(开启)时,key不能为空
 (Optional)
 * param title: 直播回看证书别名
- 取值: 支持大小写字母和数字 长度最大256
 (Optional)
 */
func NewSetLiveRestartDomainCertificateRequestWithAllParams(
    restartDomain string,
    certStatus string,
    cert *string,
    key *string,
    title *string,
) *SetLiveRestartDomainCertificateRequest {

    return &SetLiveRestartDomainCertificateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRestartDomainCertificate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RestartDomain: restartDomain,
        CertStatus: certStatus,
        Cert: cert,
        Key: key,
        Title: title,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetLiveRestartDomainCertificateRequestWithoutParam() *SetLiveRestartDomainCertificateRequest {

    return &SetLiveRestartDomainCertificateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRestartDomainCertificate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param restartDomain: 直播回看域名(Required) */
func (r *SetLiveRestartDomainCertificateRequest) SetRestartDomain(restartDomain string) {
    r.RestartDomain = restartDomain
}

/* param certStatus: 直播回看证书状态
  on: 开启
  off: 关闭
- 当播放证书状态on(开启)时,cert和key不能为空
(Required) */
func (r *SetLiveRestartDomainCertificateRequest) SetCertStatus(certStatus string) {
    r.CertStatus = certStatus
}

/* param cert: 直播回看证书
- 取值: 最大支持4098
- 当播放证书状态on(开启)时,cert不能为空
(Optional) */
func (r *SetLiveRestartDomainCertificateRequest) SetCert(cert string) {
    r.Cert = &cert
}

/* param key: 直播回看证书key
- 取值: 最大支持2048
- 当播放证书状态on(开启)时,key不能为空
(Optional) */
func (r *SetLiveRestartDomainCertificateRequest) SetKey(key string) {
    r.Key = &key
}

/* param title: 直播回看证书别名
- 取值: 支持大小写字母和数字 长度最大256
(Optional) */
func (r *SetLiveRestartDomainCertificateRequest) SetTitle(title string) {
    r.Title = &title
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetLiveRestartDomainCertificateRequest) GetRegionId() string {
    return ""
}

type SetLiveRestartDomainCertificateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetLiveRestartDomainCertificateResult `json:"result"`
}

type SetLiveRestartDomainCertificateResult struct {
}