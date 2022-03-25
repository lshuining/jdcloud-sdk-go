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
    openjrtc "github.com/jdcloud-api/jdcloud-sdk-go/services/openjrtc/models"
)

type DescribeUserRecordByRoomRequest struct {

    core.JDCloudRequest

    /* 页码；默认值为 1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认值为 10；取值范围 [10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 传参字段描述:
  appId:   应用ID (必填)
  startTime: 房间使用起始时间 UTC格式 (必填)
  endTime：房间使用截止时间 UTC格式 (必填)
  userRoomId：业务接入方定义的且在JRTC系统内注册过的房间号(必填)
  userId：业务接入方定义的且在JRTC系统内注册过的用户id
  */
    Filters []openjrtc.Filter `json:"filters"`
}

/*
 * param filters: 传参字段描述:
  appId:   应用ID (必填)
  startTime: 房间使用起始时间 UTC格式 (必填)
  endTime：房间使用截止时间 UTC格式 (必填)
  userRoomId：业务接入方定义的且在JRTC系统内注册过的房间号(必填)
  userId：业务接入方定义的且在JRTC系统内注册过的用户id
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUserRecordByRoomRequest(
    filters []openjrtc.Filter,
) *DescribeUserRecordByRoomRequest {

	return &DescribeUserRecordByRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeUserRecordByRoom",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Filters: filters,
	}
}

/*
 * param pageNumber: 页码；默认值为 1 (Optional)
 * param pageSize: 分页大小；默认值为 10；取值范围 [10, 100] (Optional)
 * param filters: 传参字段描述:
  appId:   应用ID (必填)
  startTime: 房间使用起始时间 UTC格式 (必填)
  endTime：房间使用截止时间 UTC格式 (必填)
  userRoomId：业务接入方定义的且在JRTC系统内注册过的房间号(必填)
  userId：业务接入方定义的且在JRTC系统内注册过的用户id
 (Required)
 */
func NewDescribeUserRecordByRoomRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    filters []openjrtc.Filter,
) *DescribeUserRecordByRoomRequest {

    return &DescribeUserRecordByRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeUserRecordByRoom",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUserRecordByRoomRequestWithoutParam() *DescribeUserRecordByRoomRequest {

    return &DescribeUserRecordByRoomRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeUserRecordByRoom",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认值为 1(Optional) */
func (r *DescribeUserRecordByRoomRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认值为 10；取值范围 [10, 100](Optional) */
func (r *DescribeUserRecordByRoomRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 传参字段描述:
  appId:   应用ID (必填)
  startTime: 房间使用起始时间 UTC格式 (必填)
  endTime：房间使用截止时间 UTC格式 (必填)
  userRoomId：业务接入方定义的且在JRTC系统内注册过的房间号(必填)
  userId：业务接入方定义的且在JRTC系统内注册过的用户id
(Required) */
func (r *DescribeUserRecordByRoomRequest) SetFilters(filters []openjrtc.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUserRecordByRoomRequest) GetRegionId() string {
    return ""
}

type DescribeUserRecordByRoomResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUserRecordByRoomResult `json:"result"`
}

type DescribeUserRecordByRoomResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalElements int `json:"totalElements"`
    TotalPages int `json:"totalPages"`
    Content []openjrtc.RoomUserRecord `json:"content"`
}