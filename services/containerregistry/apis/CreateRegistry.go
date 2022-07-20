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
    containerregistry "github.com/lshuining/jdcloud-sdk-go/services/containerregistry/models"
)

type CreateRegistryRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户定义的registry名称。<br> DNS兼容registry名称规则如下：
 <br> 不可为空，且不能超过32字符 <br> 以小写字母开始和结尾，支持使用小写字母、数字、中划线(-)
  */
    RegistryName string `json:"registryName"`

    /* 注册表描述，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。
 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 用户定义的registry名称。<br> DNS兼容registry名称规则如下：
 <br> 不可为空，且不能超过32字符 <br> 以小写字母开始和结尾，支持使用小写字母、数字、中划线(-)
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRegistryRequest(
    regionId string,
    registryName string,
) *CreateRegistryRequest {

	return &CreateRegistryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryName: registryName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 用户定义的registry名称。<br> DNS兼容registry名称规则如下：
 <br> 不可为空，且不能超过32字符 <br> 以小写字母开始和结尾，支持使用小写字母、数字、中划线(-)
 (Required)
 * param description: 注册表描述，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。
 (Optional)
 */
func NewCreateRegistryRequestWithAllParams(
    regionId string,
    registryName string,
    description *string,
) *CreateRegistryRequest {

    return &CreateRegistryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRegistryRequestWithoutParam() *CreateRegistryRequest {

    return &CreateRegistryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateRegistryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 用户定义的registry名称。<br> DNS兼容registry名称规则如下：
 <br> 不可为空，且不能超过32字符 <br> 以小写字母开始和结尾，支持使用小写字母、数字、中划线(-)
(Required) */
func (r *CreateRegistryRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

/* param description: 注册表描述，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。
(Optional) */
func (r *CreateRegistryRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRegistryRequest) GetRegionId() string {
    return r.RegionId
}

type CreateRegistryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRegistryResult `json:"result"`
}

type CreateRegistryResult struct {
    Registry containerregistry.Registry `json:"registry"`
}