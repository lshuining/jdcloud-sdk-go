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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryCustomizedDirBandWidthRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2020-12-07T16:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2020-12-07T16:20:00Z，开始时间和结束时间跨度 不能超过4个小时 (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名，该接口仅支持单域名查询 (Optional) */
    Domain *string `json:"domain"`

    /* 需要过滤的目录，以正斜线(/)开头，不填表示查询所有目录。查询目录同时需要以正斜线(/)结尾。 如:/path1/path2/path3/ (Optional) */
    Dir *string `json:"dir"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryCustomizedDirBandWidthRequest(
) *QueryCustomizedDirBandWidthRequest {

	return &QueryCustomizedDirBandWidthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/statistics:queryCustomizedDirBandWidth",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2020-12-07T16:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2020-12-07T16:20:00Z，开始时间和结束时间跨度 不能超过4个小时 (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名，该接口仅支持单域名查询 (Optional)
 * param dir: 需要过滤的目录，以正斜线(/)开头，不填表示查询所有目录。查询目录同时需要以正斜线(/)结尾。 如:/path1/path2/path3/ (Optional)
 */
func NewQueryCustomizedDirBandWidthRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    dir *string,
) *QueryCustomizedDirBandWidthRequest {

    return &QueryCustomizedDirBandWidthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:queryCustomizedDirBandWidth",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Dir: dir,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCustomizedDirBandWidthRequestWithoutParam() *QueryCustomizedDirBandWidthRequest {

    return &QueryCustomizedDirBandWidthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:queryCustomizedDirBandWidth",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2020-12-07T16:00:00Z(Optional) */
func (r *QueryCustomizedDirBandWidthRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2020-12-07T16:20:00Z，开始时间和结束时间跨度 不能超过4个小时(Optional) */
func (r *QueryCustomizedDirBandWidthRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名，该接口仅支持单域名查询(Optional) */
func (r *QueryCustomizedDirBandWidthRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param dir: 需要过滤的目录，以正斜线(/)开头，不填表示查询所有目录。查询目录同时需要以正斜线(/)结尾。 如:/path1/path2/path3/(Optional) */
func (r *QueryCustomizedDirBandWidthRequest) SetDir(dir string) {
    r.Dir = &dir
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCustomizedDirBandWidthRequest) GetRegionId() string {
    return ""
}

type QueryCustomizedDirBandWidthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCustomizedDirBandWidthResult `json:"result"`
}

type QueryCustomizedDirBandWidthResult struct {
    Domain string `json:"domain"`
    Data []cdn.FlowItem `json:"data"`
}