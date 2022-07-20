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
    jcq "github.com/lshuining/jdcloud-sdk-go/services/jcq/models"
)

type DescribeMessagesRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* 开始时间  */
    StartTime string `json:"startTime"`

    /* 结束时间  */
    EndTime string `json:"endTime"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param startTime: 开始时间 (Required)
 * param endTime: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMessagesRequest(
    regionId string,
    topicName string,
    startTime string,
    endTime string,
) *DescribeMessagesRequest {

	return &DescribeMessagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics/{topicName}/messages",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicName: topicName,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param startTime: 开始时间 (Required)
 * param endTime: 结束时间 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param pageNumber: 页码 (Optional)
 */
func NewDescribeMessagesRequestWithAllParams(
    regionId string,
    topicName string,
    startTime string,
    endTime string,
    pageSize *int,
    pageNumber *int,
) *DescribeMessagesRequest {

    return &DescribeMessagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/messages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicName: topicName,
        StartTime: startTime,
        EndTime: endTime,
        PageSize: pageSize,
        PageNumber: pageNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMessagesRequestWithoutParam() *DescribeMessagesRequest {

    return &DescribeMessagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/messages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *DescribeMessagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicName: topic 名称(Required) */
func (r *DescribeMessagesRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param startTime: 开始时间(Required) */
func (r *DescribeMessagesRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间(Required) */
func (r *DescribeMessagesRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeMessagesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeMessagesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMessagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMessagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMessagesResult `json:"result"`
}

type DescribeMessagesResult struct {
    Messages []jcq.Message `json:"messages"`
}