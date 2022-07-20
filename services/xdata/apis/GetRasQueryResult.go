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

type GetRasQueryResultRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 用户名称  */
    UserName string `json:"userName"`

    /* 查询id  */
    QueryId string `json:"queryId"`
}

/*
 * param regionId: 地域ID (Required)
 * param userName: 用户名称 (Required)
 * param queryId: 查询id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetRasQueryResultRequest(
    regionId string,
    userName string,
    queryId string,
) *GetRasQueryResultRequest {

	return &GetRasQueryResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwQuery:getRasQueryResult",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UserName: userName,
        QueryId: queryId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param userName: 用户名称 (Required)
 * param queryId: 查询id (Required)
 */
func NewGetRasQueryResultRequestWithAllParams(
    regionId string,
    userName string,
    queryId string,
) *GetRasQueryResultRequest {

    return &GetRasQueryResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:getRasQueryResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UserName: userName,
        QueryId: queryId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetRasQueryResultRequestWithoutParam() *GetRasQueryResultRequest {

    return &GetRasQueryResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:getRasQueryResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetRasQueryResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param userName: 用户名称(Required) */
func (r *GetRasQueryResultRequest) SetUserName(userName string) {
    r.UserName = userName
}

/* param queryId: 查询id(Required) */
func (r *GetRasQueryResultRequest) SetQueryId(queryId string) {
    r.QueryId = queryId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetRasQueryResultRequest) GetRegionId() string {
    return r.RegionId
}

type GetRasQueryResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetRasQueryResultResult `json:"result"`
}

type GetRasQueryResultResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}