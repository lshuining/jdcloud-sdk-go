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
    iam "github.com/lshuining/jdcloud-sdk-go/services/iam/models"
)

type DescribeAttachedGroupPoliciesRequest struct {

    core.JDCloudRequest

    /* 用户组名称  */
    GroupName string `json:"groupName"`
}

/*
 * param groupName: 用户组名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAttachedGroupPoliciesRequest(
    groupName string,
) *DescribeAttachedGroupPoliciesRequest {

	return &DescribeAttachedGroupPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/group/{groupName}/policies",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        GroupName: groupName,
	}
}

/*
 * param groupName: 用户组名称 (Required)
 */
func NewDescribeAttachedGroupPoliciesRequestWithAllParams(
    groupName string,
) *DescribeAttachedGroupPoliciesRequest {

    return &DescribeAttachedGroupPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/group/{groupName}/policies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        GroupName: groupName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAttachedGroupPoliciesRequestWithoutParam() *DescribeAttachedGroupPoliciesRequest {

    return &DescribeAttachedGroupPoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/group/{groupName}/policies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param groupName: 用户组名称(Required) */
func (r *DescribeAttachedGroupPoliciesRequest) SetGroupName(groupName string) {
    r.GroupName = groupName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAttachedGroupPoliciesRequest) GetRegionId() string {
    return ""
}

type DescribeAttachedGroupPoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAttachedGroupPoliciesResult `json:"result"`
}

type DescribeAttachedGroupPoliciesResult struct {
    Total int `json:"total"`
    Policies []iam.Policy `json:"policies"`
}