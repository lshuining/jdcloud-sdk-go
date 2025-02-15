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
    jmr "github.com/lshuining/jdcloud-sdk-go/services/jmr/models"
)

type CreateClusterRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 描述集群配置  */
    ClusterSpec *jmr.ClusterSpec `json:"clusterSpec"`

    /* 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。
 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: 地域ID (Required)
 * param clusterSpec: 描述集群配置 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateClusterRequest(
    regionId string,
    clusterSpec *jmr.ClusterSpec,
) *CreateClusterRequest {

	return &CreateClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cluster:create",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterSpec: clusterSpec,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param clusterSpec: 描述集群配置 (Required)
 * param clientToken: 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。
 (Optional)
 */
func NewCreateClusterRequestWithAllParams(
    regionId string,
    clusterSpec *jmr.ClusterSpec,
    clientToken *string,
) *CreateClusterRequest {

    return &CreateClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cluster:create",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterSpec: clusterSpec,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateClusterRequestWithoutParam() *CreateClusterRequest {

    return &CreateClusterRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cluster:create",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateClusterRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterSpec: 描述集群配置(Required) */
func (r *CreateClusterRequest) SetClusterSpec(clusterSpec *jmr.ClusterSpec) {
    r.ClusterSpec = clusterSpec
}

/* param clientToken: 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。
(Optional) */
func (r *CreateClusterRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateClusterRequest) GetRegionId() string {
    return r.RegionId
}

type CreateClusterResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateClusterResult `json:"result"`
}

type CreateClusterResult struct {
    Status bool `json:"status"`
    ClusterId string `json:"clusterId"`
    Message string `json:"message"`
}