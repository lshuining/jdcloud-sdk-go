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
    clouddnsservice "github.com/lshuining/jdcloud-sdk-go/services/clouddnsservice/models"
)

type DelUserViewRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 删除自定义线路的参数  */
    Req *clouddnsservice.DelView `json:"req"`
}

/*
 * param regionId: 地域ID (Required)
 * param req: 删除自定义线路的参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDelUserViewRequest(
    regionId string,
    req *clouddnsservice.DelView,
) *DelUserViewRequest {

	return &DelUserViewRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/userview/delUserView",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Req: req,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param req: 删除自定义线路的参数 (Required)
 */
func NewDelUserViewRequestWithAllParams(
    regionId string,
    req *clouddnsservice.DelView,
) *DelUserViewRequest {

    return &DelUserViewRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userview/delUserView",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDelUserViewRequestWithoutParam() *DelUserViewRequest {

    return &DelUserViewRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userview/delUserView",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DelUserViewRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param req: 删除自定义线路的参数(Required) */
func (r *DelUserViewRequest) SetReq(req *clouddnsservice.DelView) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DelUserViewRequest) GetRegionId() string {
    return r.RegionId
}

type DelUserViewResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DelUserViewResult `json:"result"`
}

type DelUserViewResult struct {
}