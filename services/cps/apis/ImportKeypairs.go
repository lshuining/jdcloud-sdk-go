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

type ImportKeypairsRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 密钥对名称  */
    Name string `json:"name"`

    /* 公钥  */
    PublicKey string `json:"publicKey"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param name: 密钥对名称 (Required)
 * param publicKey: 公钥 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewImportKeypairsRequest(
    regionId string,
    name string,
    publicKey string,
) *ImportKeypairsRequest {

	return &ImportKeypairsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/keypairs:import",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        PublicKey: publicKey,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param name: 密钥对名称 (Required)
 * param publicKey: 公钥 (Required)
 */
func NewImportKeypairsRequestWithAllParams(
    regionId string,
    clientToken *string,
    name string,
    publicKey string,
) *ImportKeypairsRequest {

    return &ImportKeypairsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs:import",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        Name: name,
        PublicKey: publicKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewImportKeypairsRequestWithoutParam() *ImportKeypairsRequest {

    return &ImportKeypairsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs:import",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *ImportKeypairsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *ImportKeypairsRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param name: 密钥对名称(Required) */
func (r *ImportKeypairsRequest) SetName(name string) {
    r.Name = name
}

/* param publicKey: 公钥(Required) */
func (r *ImportKeypairsRequest) SetPublicKey(publicKey string) {
    r.PublicKey = publicKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ImportKeypairsRequest) GetRegionId() string {
    return r.RegionId
}

type ImportKeypairsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ImportKeypairsResult `json:"result"`
}

type ImportKeypairsResult struct {
    KeypairId string `json:"keypairId"`
    Region string `json:"region"`
    Name string `json:"name"`
    PublicKey string `json:"publicKey"`
    FingerPrint string `json:"fingerPrint"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}