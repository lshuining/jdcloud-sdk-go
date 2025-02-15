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

type DeleteRepositoryRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 注册表名称  */
    RegistryName string `json:"registryName"`

    /* 镜像仓库名称  */
    RepositoryName string `json:"repositoryName"`
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 注册表名称 (Required)
 * param repositoryName: 镜像仓库名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteRepositoryRequest(
    regionId string,
    registryName string,
    repositoryName string,
) *DeleteRepositoryRequest {

	return &DeleteRepositoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries/{registryName}/repositories/{repositoryName}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryName: registryName,
        RepositoryName: repositoryName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 注册表名称 (Required)
 * param repositoryName: 镜像仓库名称 (Required)
 */
func NewDeleteRepositoryRequestWithAllParams(
    regionId string,
    registryName string,
    repositoryName string,
) *DeleteRepositoryRequest {

    return &DeleteRepositoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}/repositories/{repositoryName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
        RepositoryName: repositoryName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteRepositoryRequestWithoutParam() *DeleteRepositoryRequest {

    return &DeleteRepositoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}/repositories/{repositoryName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteRepositoryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 注册表名称(Required) */
func (r *DeleteRepositoryRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

/* param repositoryName: 镜像仓库名称(Required) */
func (r *DeleteRepositoryRequest) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteRepositoryRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteRepositoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteRepositoryResult `json:"result"`
}

type DeleteRepositoryResult struct {
}