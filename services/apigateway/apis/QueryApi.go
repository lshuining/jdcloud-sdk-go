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
    apigateway "github.com/lshuining/jdcloud-sdk-go/services/apigateway/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type QueryApiRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本号  */
    Revision string `json:"revision"`

    /* 接口ID  */
    ApiId string `json:"apiId"`

    /* isApiProduct - 是否API产品，精确匹配，1为是
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param apiId: 接口ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryApiRequest(
    regionId string,
    apiGroupId string,
    revision string,
    apiId string,
) *QueryApiRequest {

	return &QueryApiRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis/{apiId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiId: apiId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param apiId: 接口ID (Required)
 * param filters: isApiProduct - 是否API产品，精确匹配，1为是
 (Optional)
 */
func NewQueryApiRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    apiId string,
    filters []common.Filter,
) *QueryApiRequest {

    return &QueryApiRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis/{apiId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiId: apiId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryApiRequestWithoutParam() *QueryApiRequest {

    return &QueryApiRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis/{apiId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QueryApiRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *QueryApiRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 版本号(Required) */
func (r *QueryApiRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param apiId: 接口ID(Required) */
func (r *QueryApiRequest) SetApiId(apiId string) {
    r.ApiId = apiId
}

/* param filters: isApiProduct - 是否API产品，精确匹配，1为是
(Optional) */
func (r *QueryApiRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryApiRequest) GetRegionId() string {
    return r.RegionId
}

type QueryApiResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryApiResult `json:"result"`
}

type QueryApiResult struct {
    Api apigateway.Api `json:"api"`
}