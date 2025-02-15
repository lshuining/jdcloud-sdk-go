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

type CreateRoleRequest struct {

    core.JDCloudRequest

    /* 角色信息  */
    CreateRoleInfo *iam.CreateRoleInfo `json:"createRoleInfo"`
}

/*
 * param createRoleInfo: 角色信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRoleRequest(
    createRoleInfo *iam.CreateRoleInfo,
) *CreateRoleRequest {

	return &CreateRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/role",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CreateRoleInfo: createRoleInfo,
	}
}

/*
 * param createRoleInfo: 角色信息 (Required)
 */
func NewCreateRoleRequestWithAllParams(
    createRoleInfo *iam.CreateRoleInfo,
) *CreateRoleRequest {

    return &CreateRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/role",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CreateRoleInfo: createRoleInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRoleRequestWithoutParam() *CreateRoleRequest {

    return &CreateRoleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/role",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param createRoleInfo: 角色信息(Required) */
func (r *CreateRoleRequest) SetCreateRoleInfo(createRoleInfo *iam.CreateRoleInfo) {
    r.CreateRoleInfo = createRoleInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRoleRequest) GetRegionId() string {
    return ""
}

type CreateRoleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRoleResult `json:"result"`
}

type CreateRoleResult struct {
    RoleInfo iam.RoleInfo `json:"roleInfo"`
}