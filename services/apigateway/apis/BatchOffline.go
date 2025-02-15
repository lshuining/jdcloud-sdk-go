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

type BatchOfflineRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 要删除的部署ID集合，以,分隔 (Optional) */
    DeploymentIds *string `json:"deploymentIds"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchOfflineRequest(
    regionId string,
    apiGroupId string,
) *BatchOfflineRequest {

	return &BatchOfflineRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments:offline",
			Method:  "POST",
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
 * param deploymentIds: 要删除的部署ID集合，以,分隔 (Optional)
 */
func NewBatchOfflineRequestWithAllParams(
    regionId string,
    apiGroupId string,
    deploymentIds *string,
) *BatchOfflineRequest {

    return &BatchOfflineRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments:offline",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        DeploymentIds: deploymentIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchOfflineRequestWithoutParam() *BatchOfflineRequest {

    return &BatchOfflineRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments:offline",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *BatchOfflineRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *BatchOfflineRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param deploymentIds: 要删除的部署ID集合，以,分隔(Optional) */
func (r *BatchOfflineRequest) SetDeploymentIds(deploymentIds string) {
    r.DeploymentIds = &deploymentIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchOfflineRequest) GetRegionId() string {
    return r.RegionId
}

type BatchOfflineResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchOfflineResult `json:"result"`
}

type BatchOfflineResult struct {
    SuccessCount int `json:"successCount"`
    Failed []common.ErrorItem `json:"failed"`
}