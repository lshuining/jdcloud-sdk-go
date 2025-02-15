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
    elite "github.com/lshuining/jdcloud-sdk-go/services/elite/models"
)

type JdxQueryPriceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 查询价格参数  */
    QueryPriceParam *elite.QueryPriceParam `json:"queryPriceParam"`
}

/*
 * param regionId: 地域ID (Required)
 * param queryPriceParam: 查询价格参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewJdxQueryPriceRequest(
    regionId string,
    queryPriceParam *elite.QueryPriceParam,
) *JdxQueryPriceRequest {

	return &JdxQueryPriceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/jdxQueryPrice",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        QueryPriceParam: queryPriceParam,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param queryPriceParam: 查询价格参数 (Required)
 */
func NewJdxQueryPriceRequestWithAllParams(
    regionId string,
    queryPriceParam *elite.QueryPriceParam,
) *JdxQueryPriceRequest {

    return &JdxQueryPriceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jdxQueryPrice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        QueryPriceParam: queryPriceParam,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewJdxQueryPriceRequestWithoutParam() *JdxQueryPriceRequest {

    return &JdxQueryPriceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jdxQueryPrice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *JdxQueryPriceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queryPriceParam: 查询价格参数(Required) */
func (r *JdxQueryPriceRequest) SetQueryPriceParam(queryPriceParam *elite.QueryPriceParam) {
    r.QueryPriceParam = queryPriceParam
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r JdxQueryPriceRequest) GetRegionId() string {
    return r.RegionId
}

type JdxQueryPriceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result JdxQueryPriceResult `json:"result"`
}

type JdxQueryPriceResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data elite.QueryPriceResultVo `json:"data"`
}