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
    vod "github.com/lshuining/jdcloud-sdk-go/services/vod/models"
)

type CreateLiveToVodTaskRequest struct {

    core.JDCloudRequest

    /* 视频标题  */
    Title string `json:"title"`

    /* 文件名称  */
    FileName string `json:"fileName"`

    /* 文件大小 (Optional) */
    FileSize *int64 `json:"fileSize"`

    /* 封面地址 (Optional) */
    CoverUrl *string `json:"coverUrl"`

    /* 视频描述 (Optional) */
    Description *string `json:"description"`

    /* 分类ID (Optional) */
    CategoryId *int64 `json:"categoryId"`

    /* 视频标签集合 (Optional) */
    Tags []string `json:"tags"`

    /* 转码模板组ID。若此字段不为空，则将以模板组方式提交转码作业，transcodeTemplateIds字段将被忽略。 (Optional) */
    TranscodeTemplateGroupId *string `json:"transcodeTemplateGroupId"`

    /* 转码模板ID集合 (Optional) */
    TranscodeTemplateIds []int64 `json:"transcodeTemplateIds"`

    /* 水印ID集合 (Optional) */
    WatermarkIds []int64 `json:"watermarkIds"`

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 流名称  */
    StreamName string `json:"streamName"`

    /* 录制时间段集合
- 支持自定义1-10个时间段,拼接成一个文件
- 每个时间段不小于10s
- 总跨度不超过12小时
- 时间段按升序排列且无重叠
  */
    RecordTimes []vod.RecordTime `json:"recordTimes"`

    /* 录制文件类型:
- 取值: ts, flv, mp4
- 不区分大小写
  */
    RecordFileType string `json:"recordFileType"`

    /* 直播录制任务外键 (Optional) */
    TaskExternalId *string `json:"taskExternalId"`

    /* 任务优先级:
- 取值: low, medium, high
- 不区分大小写
 (Optional) */
    Priority *string `json:"priority"`
}

/*
 * param title: 视频标题 (Required)
 * param fileName: 文件名称 (Required)
 * param publishDomain: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param recordTimes: 录制时间段集合
- 支持自定义1-10个时间段,拼接成一个文件
- 每个时间段不小于10s
- 总跨度不超过12小时
- 时间段按升序排列且无重叠
 (Required)
 * param recordFileType: 录制文件类型:
- 取值: ts, flv, mp4
- 不区分大小写
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLiveToVodTaskRequest(
    title string,
    fileName string,
    publishDomain string,
    appName string,
    streamName string,
    recordTimes []vod.RecordTime,
    recordFileType string,
) *CreateLiveToVodTaskRequest {

	return &CreateLiveToVodTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/createLiveToVodTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Title: title,
        FileName: fileName,
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        RecordTimes: recordTimes,
        RecordFileType: recordFileType,
	}
}

/*
 * param title: 视频标题 (Required)
 * param fileName: 文件名称 (Required)
 * param fileSize: 文件大小 (Optional)
 * param coverUrl: 封面地址 (Optional)
 * param description: 视频描述 (Optional)
 * param categoryId: 分类ID (Optional)
 * param tags: 视频标签集合 (Optional)
 * param transcodeTemplateGroupId: 转码模板组ID。若此字段不为空，则将以模板组方式提交转码作业，transcodeTemplateIds字段将被忽略。 (Optional)
 * param transcodeTemplateIds: 转码模板ID集合 (Optional)
 * param watermarkIds: 水印ID集合 (Optional)
 * param publishDomain: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param recordTimes: 录制时间段集合
- 支持自定义1-10个时间段,拼接成一个文件
- 每个时间段不小于10s
- 总跨度不超过12小时
- 时间段按升序排列且无重叠
 (Required)
 * param recordFileType: 录制文件类型:
- 取值: ts, flv, mp4
- 不区分大小写
 (Required)
 * param taskExternalId: 直播录制任务外键 (Optional)
 * param priority: 任务优先级:
- 取值: low, medium, high
- 不区分大小写
 (Optional)
 */
func NewCreateLiveToVodTaskRequestWithAllParams(
    title string,
    fileName string,
    fileSize *int64,
    coverUrl *string,
    description *string,
    categoryId *int64,
    tags []string,
    transcodeTemplateGroupId *string,
    transcodeTemplateIds []int64,
    watermarkIds []int64,
    publishDomain string,
    appName string,
    streamName string,
    recordTimes []vod.RecordTime,
    recordFileType string,
    taskExternalId *string,
    priority *string,
) *CreateLiveToVodTaskRequest {

    return &CreateLiveToVodTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/createLiveToVodTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Title: title,
        FileName: fileName,
        FileSize: fileSize,
        CoverUrl: coverUrl,
        Description: description,
        CategoryId: categoryId,
        Tags: tags,
        TranscodeTemplateGroupId: transcodeTemplateGroupId,
        TranscodeTemplateIds: transcodeTemplateIds,
        WatermarkIds: watermarkIds,
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        RecordTimes: recordTimes,
        RecordFileType: recordFileType,
        TaskExternalId: taskExternalId,
        Priority: priority,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLiveToVodTaskRequestWithoutParam() *CreateLiveToVodTaskRequest {

    return &CreateLiveToVodTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/createLiveToVodTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param title: 视频标题(Required) */
func (r *CreateLiveToVodTaskRequest) SetTitle(title string) {
    r.Title = title
}

/* param fileName: 文件名称(Required) */
func (r *CreateLiveToVodTaskRequest) SetFileName(fileName string) {
    r.FileName = fileName
}

/* param fileSize: 文件大小(Optional) */
func (r *CreateLiveToVodTaskRequest) SetFileSize(fileSize int64) {
    r.FileSize = &fileSize
}

/* param coverUrl: 封面地址(Optional) */
func (r *CreateLiveToVodTaskRequest) SetCoverUrl(coverUrl string) {
    r.CoverUrl = &coverUrl
}

/* param description: 视频描述(Optional) */
func (r *CreateLiveToVodTaskRequest) SetDescription(description string) {
    r.Description = &description
}

/* param categoryId: 分类ID(Optional) */
func (r *CreateLiveToVodTaskRequest) SetCategoryId(categoryId int64) {
    r.CategoryId = &categoryId
}

/* param tags: 视频标签集合(Optional) */
func (r *CreateLiveToVodTaskRequest) SetTags(tags []string) {
    r.Tags = tags
}

/* param transcodeTemplateGroupId: 转码模板组ID。若此字段不为空，则将以模板组方式提交转码作业，transcodeTemplateIds字段将被忽略。(Optional) */
func (r *CreateLiveToVodTaskRequest) SetTranscodeTemplateGroupId(transcodeTemplateGroupId string) {
    r.TranscodeTemplateGroupId = &transcodeTemplateGroupId
}

/* param transcodeTemplateIds: 转码模板ID集合(Optional) */
func (r *CreateLiveToVodTaskRequest) SetTranscodeTemplateIds(transcodeTemplateIds []int64) {
    r.TranscodeTemplateIds = transcodeTemplateIds
}

/* param watermarkIds: 水印ID集合(Optional) */
func (r *CreateLiveToVodTaskRequest) SetWatermarkIds(watermarkIds []int64) {
    r.WatermarkIds = watermarkIds
}

/* param publishDomain: 推流域名(Required) */
func (r *CreateLiveToVodTaskRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 应用名称(Required) */
func (r *CreateLiveToVodTaskRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param streamName: 流名称(Required) */
func (r *CreateLiveToVodTaskRequest) SetStreamName(streamName string) {
    r.StreamName = streamName
}

/* param recordTimes: 录制时间段集合
- 支持自定义1-10个时间段,拼接成一个文件
- 每个时间段不小于10s
- 总跨度不超过12小时
- 时间段按升序排列且无重叠
(Required) */
func (r *CreateLiveToVodTaskRequest) SetRecordTimes(recordTimes []vod.RecordTime) {
    r.RecordTimes = recordTimes
}

/* param recordFileType: 录制文件类型:
- 取值: ts, flv, mp4
- 不区分大小写
(Required) */
func (r *CreateLiveToVodTaskRequest) SetRecordFileType(recordFileType string) {
    r.RecordFileType = recordFileType
}

/* param taskExternalId: 直播录制任务外键(Optional) */
func (r *CreateLiveToVodTaskRequest) SetTaskExternalId(taskExternalId string) {
    r.TaskExternalId = &taskExternalId
}

/* param priority: 任务优先级:
- 取值: low, medium, high
- 不区分大小写
(Optional) */
func (r *CreateLiveToVodTaskRequest) SetPriority(priority string) {
    r.Priority = &priority
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLiveToVodTaskRequest) GetRegionId() string {
    return ""
}

type CreateLiveToVodTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLiveToVodTaskResult `json:"result"`
}

type CreateLiveToVodTaskResult struct {
    FlowId string `json:"flowId"`
}