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

type AttachKeypairRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 密钥名称。  */
    KeyName string `json:"keyName"`

    /* 要绑定的云主机Id列表。  */
    InstanceIds []string `json:"instanceIds"`

    /* 绑定密钥后，根据此参数决定是否允许使用密码登录。可选范围：
`yes`：允许SSH密码登录。
`no`：禁止SSH密码登录。
  */
    PassWordAuth string `json:"passWordAuth"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param keyName: 密钥名称。 (Required)
 * param instanceIds: 要绑定的云主机Id列表。 (Required)
 * param passWordAuth: 绑定密钥后，根据此参数决定是否允许使用密码登录。可选范围：
`yes`：允许SSH密码登录。
`no`：禁止SSH密码登录。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAttachKeypairRequest(
    regionId string,
    keyName string,
    instanceIds []string,
    passWordAuth string,
) *AttachKeypairRequest {

	return &AttachKeypairRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/keypairs/{keyName}:attach",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        KeyName: keyName,
        InstanceIds: instanceIds,
        PassWordAuth: passWordAuth,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param keyName: 密钥名称。 (Required)
 * param instanceIds: 要绑定的云主机Id列表。 (Required)
 * param passWordAuth: 绑定密钥后，根据此参数决定是否允许使用密码登录。可选范围：
`yes`：允许SSH密码登录。
`no`：禁止SSH密码登录。
 (Required)
 */
func NewAttachKeypairRequestWithAllParams(
    regionId string,
    keyName string,
    instanceIds []string,
    passWordAuth string,
) *AttachKeypairRequest {

    return &AttachKeypairRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs/{keyName}:attach",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        KeyName: keyName,
        InstanceIds: instanceIds,
        PassWordAuth: passWordAuth,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAttachKeypairRequestWithoutParam() *AttachKeypairRequest {

    return &AttachKeypairRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs/{keyName}:attach",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *AttachKeypairRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param keyName: 密钥名称。(Required) */
func (r *AttachKeypairRequest) SetKeyName(keyName string) {
    r.KeyName = keyName
}

/* param instanceIds: 要绑定的云主机Id列表。(Required) */
func (r *AttachKeypairRequest) SetInstanceIds(instanceIds []string) {
    r.InstanceIds = instanceIds
}

/* param passWordAuth: 绑定密钥后，根据此参数决定是否允许使用密码登录。可选范围：
`yes`：允许SSH密码登录。
`no`：禁止SSH密码登录。
(Required) */
func (r *AttachKeypairRequest) SetPassWordAuth(passWordAuth string) {
    r.PassWordAuth = passWordAuth
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AttachKeypairRequest) GetRegionId() string {
    return r.RegionId
}

type AttachKeypairResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AttachKeypairResult `json:"result"`
}

type AttachKeypairResult struct {
    SuccessInstanceId []string `json:"successInstanceId"`
    FailInstanceId []string `json:"failInstanceId"`
}