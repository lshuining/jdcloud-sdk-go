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
    jdro "github.com/lshuining/jdcloud-sdk-go/services/jdro/models"
)

type CreateChangeSetRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 资源栈 ID  */
    StackId string `json:"stackId"`

    /*   */
    Environment *jdro.Environment `json:"environment"`

    /* 模板, JSON对象  */
    Template interface{} `json:"template"`
}

/*
 * param regionId: 地域 ID (Required)
 * param stackId: 资源栈 ID (Required)
 * param environment:  (Required)
 * param template: 模板, JSON对象 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateChangeSetRequest(
    regionId string,
    stackId string,
    environment *jdro.Environment,
    template interface{},
) *CreateChangeSetRequest {

	return &CreateChangeSetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/stacks/{stackId}/changesets",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StackId: stackId,
        Environment: environment,
        Template: template,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param stackId: 资源栈 ID (Required)
 * param environment:  (Required)
 * param template: 模板, JSON对象 (Required)
 */
func NewCreateChangeSetRequestWithAllParams(
    regionId string,
    stackId string,
    environment *jdro.Environment,
    template interface{},
) *CreateChangeSetRequest {

    return &CreateChangeSetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stacks/{stackId}/changesets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StackId: stackId,
        Environment: environment,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateChangeSetRequestWithoutParam() *CreateChangeSetRequest {

    return &CreateChangeSetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stacks/{stackId}/changesets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *CreateChangeSetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param stackId: 资源栈 ID(Required) */
func (r *CreateChangeSetRequest) SetStackId(stackId string) {
    r.StackId = stackId
}

/* param environment: (Required) */
func (r *CreateChangeSetRequest) SetEnvironment(environment *jdro.Environment) {
    r.Environment = environment
}

/* param template: 模板, JSON对象(Required) */
func (r *CreateChangeSetRequest) SetTemplate(template interface{}) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateChangeSetRequest) GetRegionId() string {
    return r.RegionId
}

type CreateChangeSetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateChangeSetResult `json:"result"`
}

type CreateChangeSetResult struct {
    Id string `json:"id"`
}