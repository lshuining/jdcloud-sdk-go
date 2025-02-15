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

type GetSslCertDetailRequest struct {

    core.JDCloudRequest

    /* 证书 Id  */
    SslCertId string `json:"sslCertId"`
}

/*
 * param sslCertId: 证书 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSslCertDetailRequest(
    sslCertId string,
) *GetSslCertDetailRequest {

	return &GetSslCertDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sslCert/{sslCertId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        SslCertId: sslCertId,
	}
}

/*
 * param sslCertId: 证书 Id (Required)
 */
func NewGetSslCertDetailRequestWithAllParams(
    sslCertId string,
) *GetSslCertDetailRequest {

    return &GetSslCertDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert/{sslCertId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        SslCertId: sslCertId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSslCertDetailRequestWithoutParam() *GetSslCertDetailRequest {

    return &GetSslCertDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert/{sslCertId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param sslCertId: 证书 Id(Required) */
func (r *GetSslCertDetailRequest) SetSslCertId(sslCertId string) {
    r.SslCertId = sslCertId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSslCertDetailRequest) GetRegionId() string {
    return ""
}

type GetSslCertDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSslCertDetailResult `json:"result"`
}

type GetSslCertDetailResult struct {
    SslCertId string `json:"sslCertId"`
    CertName string `json:"certName"`
    CommonName string `json:"commonName"`
    CertType string `json:"certType"`
    SslCertStartTime string `json:"sslCertStartTime"`
    SslCertEndTime string `json:"sslCertEndTime"`
    Digest string `json:"digest"`
    RelatedDomains []string `json:"relatedDomains"`
    BindResources []string `json:"bindResources"`
}