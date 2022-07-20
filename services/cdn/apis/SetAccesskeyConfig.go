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

type SetAccesskeyConfigRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 鉴权类型，0表示无鉴权，1表示参数鉴权，2表示路径鉴权 (Optional) */
    AccesskeyType *int `json:"accesskeyType"`

    /* 密码，长度为8到32 (Optional) */
    AccesskeyKey *string `json:"accesskeyKey"`

    /* 是否是回源鉴权 0表示是 1表示否 (Optional) */
    AccesskeyKeep *int `json:"accesskeyKeep"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetAccesskeyConfigRequest(
    domain string,
) *SetAccesskeyConfigRequest {

	return &SetAccesskeyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/accesskeyConfig",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param accesskeyType: 鉴权类型，0表示无鉴权，1表示参数鉴权，2表示路径鉴权 (Optional)
 * param accesskeyKey: 密码，长度为8到32 (Optional)
 * param accesskeyKeep: 是否是回源鉴权 0表示是 1表示否 (Optional)
 */
func NewSetAccesskeyConfigRequestWithAllParams(
    domain string,
    accesskeyType *int,
    accesskeyKey *string,
    accesskeyKeep *int,
) *SetAccesskeyConfigRequest {

    return &SetAccesskeyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/accesskeyConfig",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        AccesskeyType: accesskeyType,
        AccesskeyKey: accesskeyKey,
        AccesskeyKeep: accesskeyKeep,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetAccesskeyConfigRequestWithoutParam() *SetAccesskeyConfigRequest {

    return &SetAccesskeyConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/accesskeyConfig",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetAccesskeyConfigRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param accesskeyType: 鉴权类型，0表示无鉴权，1表示参数鉴权，2表示路径鉴权(Optional) */
func (r *SetAccesskeyConfigRequest) SetAccesskeyType(accesskeyType int) {
    r.AccesskeyType = &accesskeyType
}

/* param accesskeyKey: 密码，长度为8到32(Optional) */
func (r *SetAccesskeyConfigRequest) SetAccesskeyKey(accesskeyKey string) {
    r.AccesskeyKey = &accesskeyKey
}

/* param accesskeyKeep: 是否是回源鉴权 0表示是 1表示否(Optional) */
func (r *SetAccesskeyConfigRequest) SetAccesskeyKeep(accesskeyKeep int) {
    r.AccesskeyKeep = &accesskeyKeep
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetAccesskeyConfigRequest) GetRegionId() string {
    return ""
}

type SetAccesskeyConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetAccesskeyConfigResult `json:"result"`
}

type SetAccesskeyConfigResult struct {
}