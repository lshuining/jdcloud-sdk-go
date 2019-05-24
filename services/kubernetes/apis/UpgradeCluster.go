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

type UpgradeClusterRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 集群 ID  */
    ClusterId string `json:"clusterId"`

    /* 升级范围  */
    Scope string `json:"scope"`

    /* 节点组 id (Optional) */
    NodeGroupIds []string `json:"nodeGroupIds"`

    /* 指定升级到的版本  */
    Verison string `json:"verison"`
}

/*
 * param regionId: 地域 ID (Required)
 * param clusterId: 集群 ID (Required)
 * param scope: 升级范围 (Required)
 * param verison: 指定升级到的版本 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpgradeClusterRequest(
    regionId string,
    clusterId string,
    scope string,
    verison string,
) *UpgradeClusterRequest {

	return &UpgradeClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/clusters/{clusterId}:upgradeCluster",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterId: clusterId,
        Scope: scope,
        Verison: verison,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param clusterId: 集群 ID (Required)
 * param scope: 升级范围 (Required)
 * param nodeGroupIds: 节点组 id (Optional)
 * param verison: 指定升级到的版本 (Required)
 */
func NewUpgradeClusterRequestWithAllParams(
    regionId string,
    clusterId string,
    scope string,
    nodeGroupIds []string,
    verison string,
) *UpgradeClusterRequest {

    return &UpgradeClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}:upgradeCluster",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterId: clusterId,
        Scope: scope,
        NodeGroupIds: nodeGroupIds,
        Verison: verison,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpgradeClusterRequestWithoutParam() *UpgradeClusterRequest {

    return &UpgradeClusterRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}:upgradeCluster",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *UpgradeClusterRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterId: 集群 ID(Required) */
func (r *UpgradeClusterRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

/* param scope: 升级范围(Required) */
func (r *UpgradeClusterRequest) SetScope(scope string) {
    r.Scope = scope
}

/* param nodeGroupIds: 节点组 id(Optional) */
func (r *UpgradeClusterRequest) SetNodeGroupIds(nodeGroupIds []string) {
    r.NodeGroupIds = nodeGroupIds
}

/* param verison: 指定升级到的版本(Required) */
func (r *UpgradeClusterRequest) SetVerison(verison string) {
    r.Verison = verison
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpgradeClusterRequest) GetRegionId() string {
    return r.RegionId
}

type UpgradeClusterResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpgradeClusterResult `json:"result"`
}

type UpgradeClusterResult struct {
}