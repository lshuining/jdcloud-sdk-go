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

type DescribeSlowLogDimensionsStatisticRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceGid string `json:"instanceGid"`

    /* 查询开始时间，格式为：2021-11-12T15:04:05Z  */
    StartTime string `json:"startTime"`

    /* 查询截止时间，格式为：2021-11-11T15:04:05Z  */
    EndTime string `json:"endTime"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: 实例ID (Required)
 * param startTime: 查询开始时间，格式为：2021-11-12T15:04:05Z (Required)
 * param endTime: 查询截止时间，格式为：2021-11-11T15:04:05Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSlowLogDimensionsStatisticRequest(
    regionId string,
    instanceGid string,
    startTime string,
    endTime string,
) *DescribeSlowLogDimensionsStatisticRequest {

	return &DescribeSlowLogDimensionsStatisticRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/SlowLogDimensionsStatistic",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: 实例ID (Required)
 * param startTime: 查询开始时间，格式为：2021-11-12T15:04:05Z (Required)
 * param endTime: 查询截止时间，格式为：2021-11-11T15:04:05Z (Required)
 */
func NewDescribeSlowLogDimensionsStatisticRequestWithAllParams(
    regionId string,
    instanceGid string,
    startTime string,
    endTime string,
) *DescribeSlowLogDimensionsStatisticRequest {

    return &DescribeSlowLogDimensionsStatisticRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/SlowLogDimensionsStatistic",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSlowLogDimensionsStatisticRequestWithoutParam() *DescribeSlowLogDimensionsStatisticRequest {

    return &DescribeSlowLogDimensionsStatisticRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/SlowLogDimensionsStatistic",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeSlowLogDimensionsStatisticRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: 实例ID(Required) */
func (r *DescribeSlowLogDimensionsStatisticRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

/* param startTime: 查询开始时间，格式为：2021-11-12T15:04:05Z(Required) */
func (r *DescribeSlowLogDimensionsStatisticRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询截止时间，格式为：2021-11-11T15:04:05Z(Required) */
func (r *DescribeSlowLogDimensionsStatisticRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSlowLogDimensionsStatisticRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSlowLogDimensionsStatisticResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSlowLogDimensionsStatisticResult `json:"result"`
}

type DescribeSlowLogDimensionsStatisticResult struct {
    DbName interface{} `json:"dbName"`
    UserName interface{} `json:"userName"`
    ClientIP interface{} `json:"clientIP"`
}