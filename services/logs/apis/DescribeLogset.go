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

type DescribeLogsetRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集 UID  */
    LogsetUID string `json:"logsetUID"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLogsetRequest(
    regionId string,
    logsetUID string,
) *DescribeLogsetRequest {

	return &DescribeLogsetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 */
func NewDescribeLogsetRequestWithAllParams(
    regionId string,
    logsetUID string,
) *DescribeLogsetRequest {

    return &DescribeLogsetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLogsetRequestWithoutParam() *DescribeLogsetRequest {

    return &DescribeLogsetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeLogsetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUID: 日志集 UID(Required) */
func (r *DescribeLogsetRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLogsetRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLogsetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLogsetResult `json:"result"`
}

type DescribeLogsetResult struct {
    UID string `json:"uID"`
    CreateTime string `json:"createTime"`
    Description string `json:"description"`
    HasTopic bool `json:"hasTopic"`
    LifeCycle int64 `json:"lifeCycle"`
    Name string `json:"name"`
    Region string `json:"region"`
}