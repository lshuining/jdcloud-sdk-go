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
    cdn "github.com/lshuining/jdcloud-sdk-go/services/cdn/models"
)

type QueryDirBandwidthRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名，该接口仅支持单域名查询 (Optional) */
    Domain *string `json:"domain"`

    /* 需要过滤的目录 (Optional) */
    Dirs *string `json:"dirs"`

    /* 需要过滤的地区 (Optional) */
    Regions *string `json:"regions"`

    /* 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间 (Optional) */
    CacheType *string `json:"cacheType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDirBandwidthRequest(
) *QueryDirBandwidthRequest {

	return &QueryDirBandwidthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/statistics:queryDirBandwidth",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名，该接口仅支持单域名查询 (Optional)
 * param dirs: 需要过滤的目录 (Optional)
 * param regions: 需要过滤的地区 (Optional)
 * param cacheType: 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间 (Optional)
 */
func NewQueryDirBandwidthRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    dirs *string,
    regions *string,
    cacheType *string,
) *QueryDirBandwidthRequest {

    return &QueryDirBandwidthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:queryDirBandwidth",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Dirs: dirs,
        Regions: regions,
        CacheType: cacheType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDirBandwidthRequestWithoutParam() *QueryDirBandwidthRequest {

    return &QueryDirBandwidthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:queryDirBandwidth",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryDirBandwidthRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryDirBandwidthRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名，该接口仅支持单域名查询(Optional) */
func (r *QueryDirBandwidthRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param dirs: 需要过滤的目录(Optional) */
func (r *QueryDirBandwidthRequest) SetDirs(dirs string) {
    r.Dirs = &dirs
}

/* param regions: 需要过滤的地区(Optional) */
func (r *QueryDirBandwidthRequest) SetRegions(regions string) {
    r.Regions = &regions
}

/* param cacheType: 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间(Optional) */
func (r *QueryDirBandwidthRequest) SetCacheType(cacheType string) {
    r.CacheType = &cacheType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDirBandwidthRequest) GetRegionId() string {
    return ""
}

type QueryDirBandwidthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDirBandwidthResult `json:"result"`
}

type QueryDirBandwidthResult struct {
    Domain string `json:"domain"`
    Datas []cdn.DirBandwidthItem `json:"datas"`
}