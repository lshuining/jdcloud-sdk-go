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
    cps "github.com/lshuining/jdcloud-sdk-go/services/cps/models"
)

type CreateVpcRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 子网配置  */
    VpcSpec *cps.VpcSpec `json:"vpcSpec"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param vpcSpec: 子网配置 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVpcRequest(
    regionId string,
    vpcSpec *cps.VpcSpec,
) *CreateVpcRequest {

	return &CreateVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcSpec: vpcSpec,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param vpcSpec: 子网配置 (Required)
 */
func NewCreateVpcRequestWithAllParams(
    regionId string,
    clientToken *string,
    vpcSpec *cps.VpcSpec,
) *CreateVpcRequest {

    return &CreateVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        VpcSpec: vpcSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVpcRequestWithoutParam() *CreateVpcRequest {

    return &CreateVpcRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *CreateVpcRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *CreateVpcRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param vpcSpec: 子网配置(Required) */
func (r *CreateVpcRequest) SetVpcSpec(vpcSpec *cps.VpcSpec) {
    r.VpcSpec = vpcSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVpcRequest) GetRegionId() string {
    return r.RegionId
}

type CreateVpcResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVpcResult `json:"result"`
}

type CreateVpcResult struct {
    VpcId string `json:"vpcId"`
}