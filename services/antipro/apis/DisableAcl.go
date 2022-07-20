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
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DisableAclRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 访问控制规则 Id  */
    AclId string `json:"aclId"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param aclId: 访问控制规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableAclRequest(
    regionId string,
    instanceId string,
    aclId string,
) *DisableAclRequest {

	return &DisableAclRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/acl/{aclId}:disableAcl",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AclId: aclId,
	}
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param aclId: 访问控制规则 Id (Required)
 */
func NewDisableAclRequestWithAllParams(
    regionId string,
    instanceId string,
    aclId string,
) *DisableAclRequest {

    return &DisableAclRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/acl/{aclId}:disableAcl",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        AclId: aclId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableAclRequestWithoutParam() *DisableAclRequest {

    return &DisableAclRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/acl/{aclId}:disableAcl",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *DisableAclRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *DisableAclRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param aclId: 访问控制规则 Id(Required) */
func (r *DisableAclRequest) SetAclId(aclId string) {
    r.AclId = aclId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableAclRequest) GetRegionId() string {
    return r.RegionId
}

type DisableAclResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableAclResult `json:"result"`
}

type DisableAclResult struct {
    Data common.BatchResultDetail `json:"data"`
}