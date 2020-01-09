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

package models


type DomainLog struct {

    /* 下载链接 (Optional) */
    Url string `json:"url"`

    /* md5值 (Optional) */
    Md5 string `json:"md5"`

    /* 文件名 (Optional) */
    FileName string `json:"fileName"`

    /* 日志格式 (Optional) */
    LogType string `json:"logType"`

    /* 日志粒度 (Optional) */
    Interval string `json:"interval"`

    /* 文件大小 (Optional) */
    Size int64 `json:"size"`

    /* 日志开始时间，UTC时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 日志结束时间，UTC时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 日志修改时间，UTC时间 (Optional) */
    LastModified string `json:"lastModified"`
}
