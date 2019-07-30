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
)

type ExportLiveSnapshotDataRequest struct {

    core.JDCloudRequest

    /* 推流域名 (Optional) */
    PublishDomain *string `json:"publishDomain"`

    /* 应用名称 (Optional) */
    AppName *string `json:"appName"`

    /* 流名称 (Optional) */
    StreamName *string `json:"streamName"`

    /* 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
  */
    StartTime string `json:"startTime"`

    /* 结束时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 为空,默认当前时间
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param startTime: 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportLiveSnapshotDataRequest(
    startTime string,
) *ExportLiveSnapshotDataRequest {

	return &ExportLiveSnapshotDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/exportLiveSnapshotData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
	}
}

/*
 * param publishDomain: 推流域名 (Optional)
 * param appName: 应用名称 (Optional)
 * param streamName: 流名称 (Optional)
 * param startTime: 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
 (Required)
 * param endTime: 结束时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 为空,默认当前时间
 (Optional)
 */
func NewExportLiveSnapshotDataRequestWithAllParams(
    publishDomain *string,
    appName *string,
    streamName *string,
    startTime string,
    endTime *string,
) *ExportLiveSnapshotDataRequest {

    return &ExportLiveSnapshotDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/exportLiveSnapshotData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportLiveSnapshotDataRequestWithoutParam() *ExportLiveSnapshotDataRequest {

    return &ExportLiveSnapshotDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/exportLiveSnapshotData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Optional) */
func (r *ExportLiveSnapshotDataRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = &publishDomain
}

/* param appName: 应用名称(Optional) */
func (r *ExportLiveSnapshotDataRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param streamName: 流名称(Optional) */
func (r *ExportLiveSnapshotDataRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}

/* param startTime: 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
(Required) */
func (r *ExportLiveSnapshotDataRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 为空,默认当前时间
(Optional) */
func (r *ExportLiveSnapshotDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportLiveSnapshotDataRequest) GetRegionId() string {
    return ""
}

type ExportLiveSnapshotDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportLiveSnapshotDataResult `json:"result"`
}

type ExportLiveSnapshotDataResult struct {
    FilePath string `json:"filePath"`
}