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
    dbs "github.com/lshuining/jdcloud-sdk-go/services/dbs/models"
)

type DescribeAgentAttributesRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* Agent ID  */
    AgentId string `json:"agentId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param agentId: Agent ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAgentAttributesRequest(
    regionId string,
    agentId string,
) *DescribeAgentAttributesRequest {

	return &DescribeAgentAttributesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/agents/{agentId}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AgentId: agentId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param agentId: Agent ID (Required)
 */
func NewDescribeAgentAttributesRequestWithAllParams(
    regionId string,
    agentId string,
) *DescribeAgentAttributesRequest {

    return &DescribeAgentAttributesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/agents/{agentId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AgentId: agentId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAgentAttributesRequestWithoutParam() *DescribeAgentAttributesRequest {

    return &DescribeAgentAttributesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/agents/{agentId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *DescribeAgentAttributesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agentId: Agent ID(Required) */
func (r *DescribeAgentAttributesRequest) SetAgentId(agentId string) {
    r.AgentId = agentId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAgentAttributesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAgentAttributesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAgentAttributesResult `json:"result"`
}

type DescribeAgentAttributesResult struct {
    AgentAttributes dbs.AgentAttributes `json:"agentAttributes"`
}