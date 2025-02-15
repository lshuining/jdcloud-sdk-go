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
    vm "github.com/lshuining/jdcloud-sdk-go/services/vm/models"
)

type DescribeInstanceTemplateRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 实例模板ID。  */
    InstanceTemplateId string `json:"instanceTemplateId"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceTemplateId: 实例模板ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceTemplateRequest(
    regionId string,
    instanceTemplateId string,
) *DescribeInstanceTemplateRequest {

	return &DescribeInstanceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceTemplateId: instanceTemplateId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceTemplateId: 实例模板ID。 (Required)
 */
func NewDescribeInstanceTemplateRequestWithAllParams(
    regionId string,
    instanceTemplateId string,
) *DescribeInstanceTemplateRequest {

    return &DescribeInstanceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceTemplateId: instanceTemplateId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceTemplateRequestWithoutParam() *DescribeInstanceTemplateRequest {

    return &DescribeInstanceTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DescribeInstanceTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceTemplateId: 实例模板ID。(Required) */
func (r *DescribeInstanceTemplateRequest) SetInstanceTemplateId(instanceTemplateId string) {
    r.InstanceTemplateId = instanceTemplateId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceTemplateResult `json:"result"`
}

type DescribeInstanceTemplateResult struct {
    InstanceTemplate vm.InstanceTemplate `json:"instanceTemplate"`
}