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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type UpdateInstanceTemplateRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 实例模板ID。  */
    InstanceTemplateId string `json:"instanceTemplateId"`

    /* 实例模板的名称，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。  */
    Name string `json:"name"`

    /* 实例模板的描述，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceTemplateId: 实例模板ID。 (Required)
 * param name: 实例模板的名称，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateInstanceTemplateRequest(
    regionId string,
    instanceTemplateId string,
    name string,
) *UpdateInstanceTemplateRequest {

	return &UpdateInstanceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceTemplateId: instanceTemplateId,
        Name: name,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceTemplateId: 实例模板ID。 (Required)
 * param name: 实例模板的名称，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。 (Required)
 * param description: 实例模板的描述，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。 (Optional)
 */
func NewUpdateInstanceTemplateRequestWithAllParams(
    regionId string,
    instanceTemplateId string,
    name string,
    description *string,
) *UpdateInstanceTemplateRequest {

    return &UpdateInstanceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceTemplateId: instanceTemplateId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateInstanceTemplateRequestWithoutParam() *UpdateInstanceTemplateRequest {

    return &UpdateInstanceTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *UpdateInstanceTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceTemplateId: 实例模板ID。(Required) */
func (r *UpdateInstanceTemplateRequest) SetInstanceTemplateId(instanceTemplateId string) {
    r.InstanceTemplateId = instanceTemplateId
}

/* param name: 实例模板的名称，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。(Required) */
func (r *UpdateInstanceTemplateRequest) SetName(name string) {
    r.Name = name
}

/* param description: 实例模板的描述，参考 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。(Optional) */
func (r *UpdateInstanceTemplateRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateInstanceTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateInstanceTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateInstanceTemplateResult `json:"result"`
}

type UpdateInstanceTemplateResult struct {
}