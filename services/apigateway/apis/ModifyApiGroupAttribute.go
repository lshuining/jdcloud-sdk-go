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

type ModifyApiGroupAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 名称 (Optional) */
    GroupName *string `json:"groupName"`

    /* 描述 (Optional) */
    Description *string `json:"description"`

    /* 分组路径前缀 (Optional) */
    Prefix *string `json:"prefix"`

    /* 密钥验证方式：check_exist（密钥必须在访问授权中已配置）、no_check_exist（无需事先配置） (Optional) */
    KeyCheck *string `json:"keyCheck"`

    /* 访问授权方式：None（免鉴权）、jd_cloud（开启访问授权，且必须使用京东云的AK、SK验签）、hufu（虎符用户） (Optional) */
    AuthType *string `json:"authType"`

    /* 是否转发分组路径到后端服务：0（不转发）、1（转发）默认为1 (Optional) */
    PrefixStrip *int `json:"prefixStrip"`

    /* 分组类型：api_group（api分组）、jdsf_group（微服务分组）默认为 api_group (Optional) */
    GroupType *string `json:"groupType"`

    /* 微服务网关名称 (Optional) */
    JdsfName *string `json:"jdsfName"`

    /* 微服务注册中心ID (Optional) */
    JdsfRegistryName *string `json:"jdsfRegistryName"`

    /* 微服务网关ID (Optional) */
    JdsfId *string `json:"jdsfId"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyApiGroupAttributeRequest(
    regionId string,
    apiGroupId string,
) *ModifyApiGroupAttributeRequest {

	return &ModifyApiGroupAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param groupName: 名称 (Optional)
 * param description: 描述 (Optional)
 * param prefix: 分组路径前缀 (Optional)
 * param keyCheck: 密钥验证方式：check_exist（密钥必须在访问授权中已配置）、no_check_exist（无需事先配置） (Optional)
 * param authType: 访问授权方式：None（免鉴权）、jd_cloud（开启访问授权，且必须使用京东云的AK、SK验签）、hufu（虎符用户） (Optional)
 * param prefixStrip: 是否转发分组路径到后端服务：0（不转发）、1（转发）默认为1 (Optional)
 * param groupType: 分组类型：api_group（api分组）、jdsf_group（微服务分组）默认为 api_group (Optional)
 * param jdsfName: 微服务网关名称 (Optional)
 * param jdsfRegistryName: 微服务注册中心ID (Optional)
 * param jdsfId: 微服务网关ID (Optional)
 */
func NewModifyApiGroupAttributeRequestWithAllParams(
    regionId string,
    apiGroupId string,
    groupName *string,
    description *string,
    prefix *string,
    keyCheck *string,
    authType *string,
    prefixStrip *int,
    groupType *string,
    jdsfName *string,
    jdsfRegistryName *string,
    jdsfId *string,
) *ModifyApiGroupAttributeRequest {

    return &ModifyApiGroupAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        GroupName: groupName,
        Description: description,
        Prefix: prefix,
        KeyCheck: keyCheck,
        AuthType: authType,
        PrefixStrip: prefixStrip,
        GroupType: groupType,
        JdsfName: jdsfName,
        JdsfRegistryName: jdsfRegistryName,
        JdsfId: jdsfId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyApiGroupAttributeRequestWithoutParam() *ModifyApiGroupAttributeRequest {

    return &ModifyApiGroupAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyApiGroupAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *ModifyApiGroupAttributeRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param groupName: 名称(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetGroupName(groupName string) {
    r.GroupName = &groupName
}

/* param description: 描述(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetDescription(description string) {
    r.Description = &description
}

/* param prefix: 分组路径前缀(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetPrefix(prefix string) {
    r.Prefix = &prefix
}

/* param keyCheck: 密钥验证方式：check_exist（密钥必须在访问授权中已配置）、no_check_exist（无需事先配置）(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetKeyCheck(keyCheck string) {
    r.KeyCheck = &keyCheck
}

/* param authType: 访问授权方式：None（免鉴权）、jd_cloud（开启访问授权，且必须使用京东云的AK、SK验签）、hufu（虎符用户）(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetAuthType(authType string) {
    r.AuthType = &authType
}

/* param prefixStrip: 是否转发分组路径到后端服务：0（不转发）、1（转发）默认为1(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetPrefixStrip(prefixStrip int) {
    r.PrefixStrip = &prefixStrip
}

/* param groupType: 分组类型：api_group（api分组）、jdsf_group（微服务分组）默认为 api_group(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetGroupType(groupType string) {
    r.GroupType = &groupType
}

/* param jdsfName: 微服务网关名称(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetJdsfName(jdsfName string) {
    r.JdsfName = &jdsfName
}

/* param jdsfRegistryName: 微服务注册中心ID(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetJdsfRegistryName(jdsfRegistryName string) {
    r.JdsfRegistryName = &jdsfRegistryName
}

/* param jdsfId: 微服务网关ID(Optional) */
func (r *ModifyApiGroupAttributeRequest) SetJdsfId(jdsfId string) {
    r.JdsfId = &jdsfId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyApiGroupAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyApiGroupAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyApiGroupAttributeResult `json:"result"`
}

type ModifyApiGroupAttributeResult struct {
}