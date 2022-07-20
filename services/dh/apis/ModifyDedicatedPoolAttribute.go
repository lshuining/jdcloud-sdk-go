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

type ModifyDedicatedPoolAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 专有宿主机ID  */
    DedicatedPoolId string `json:"dedicatedPoolId"`

    /* 名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
    Name *string `json:"name"`

    /* 描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
    Description *string `json:"description"`

    /* 指定宿主机池申请专有宿主机时默认继承的可用区。<br>
修改可用区时旧可用区必须是新可用区的子集，即可用区只能添加，不能减少。
 (Optional) */
    Az []string `json:"az"`
}

/*
 * param regionId: 地域ID (Required)
 * param dedicatedPoolId: 专有宿主机ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyDedicatedPoolAttributeRequest(
    regionId string,
    dedicatedPoolId string,
) *ModifyDedicatedPoolAttributeRequest {

	return &ModifyDedicatedPoolAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dedicatedPool/{dedicatedPoolId}:modifyAttribute",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DedicatedPoolId: dedicatedPoolId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param dedicatedPoolId: 专有宿主机ID (Required)
 * param name: 名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional)
 * param description: 描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional)
 * param az: 指定宿主机池申请专有宿主机时默认继承的可用区。<br>
修改可用区时旧可用区必须是新可用区的子集，即可用区只能添加，不能减少。
 (Optional)
 */
func NewModifyDedicatedPoolAttributeRequestWithAllParams(
    regionId string,
    dedicatedPoolId string,
    name *string,
    description *string,
    az []string,
) *ModifyDedicatedPoolAttributeRequest {

    return &ModifyDedicatedPoolAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedPool/{dedicatedPoolId}:modifyAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DedicatedPoolId: dedicatedPoolId,
        Name: name,
        Description: description,
        Az: az,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyDedicatedPoolAttributeRequestWithoutParam() *ModifyDedicatedPoolAttributeRequest {

    return &ModifyDedicatedPoolAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedPool/{dedicatedPoolId}:modifyAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyDedicatedPoolAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dedicatedPoolId: 专有宿主机ID(Required) */
func (r *ModifyDedicatedPoolAttributeRequest) SetDedicatedPoolId(dedicatedPoolId string) {
    r.DedicatedPoolId = dedicatedPoolId
}

/* param name: 名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。(Optional) */
func (r *ModifyDedicatedPoolAttributeRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。(Optional) */
func (r *ModifyDedicatedPoolAttributeRequest) SetDescription(description string) {
    r.Description = &description
}

/* param az: 指定宿主机池申请专有宿主机时默认继承的可用区。<br>
修改可用区时旧可用区必须是新可用区的子集，即可用区只能添加，不能减少。
(Optional) */
func (r *ModifyDedicatedPoolAttributeRequest) SetAz(az []string) {
    r.Az = az
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyDedicatedPoolAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyDedicatedPoolAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyDedicatedPoolAttributeResult `json:"result"`
}

type ModifyDedicatedPoolAttributeResult struct {
}