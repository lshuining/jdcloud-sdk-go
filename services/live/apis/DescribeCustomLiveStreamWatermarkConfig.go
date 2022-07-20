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
    live "github.com/lshuining/jdcloud-sdk-go/services/live/models"
)

type DescribeCustomLiveStreamWatermarkConfigRequest struct {

    core.JDCloudRequest

    /* 页码
- 取值范围 [1, 100000]
 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小
- 取值范围 [10, 100]
 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 水印配置查询过滤条件:
  - name:   publishDomain，必填(推流域名)
  - value:  参数
  - name:   appName，必填(应用名称)
  - value:  参数
  - name:   streamName，非必填(流名称)
  - value:  参数
 (Optional) */
    Filters []live.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomLiveStreamWatermarkConfigRequest(
) *DescribeCustomLiveStreamWatermarkConfigRequest {

	return &DescribeCustomLiveStreamWatermarkConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/watermarks:config",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNum: 页码
- 取值范围 [1, 100000]
 (Optional)
 * param pageSize: 分页大小
- 取值范围 [10, 100]
 (Optional)
 * param filters: 水印配置查询过滤条件:
  - name:   publishDomain，必填(推流域名)
  - value:  参数
  - name:   appName，必填(应用名称)
  - value:  参数
  - name:   streamName，非必填(流名称)
  - value:  参数
 (Optional)
 */
func NewDescribeCustomLiveStreamWatermarkConfigRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    filters []live.Filter,
) *DescribeCustomLiveStreamWatermarkConfigRequest {

    return &DescribeCustomLiveStreamWatermarkConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarks:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCustomLiveStreamWatermarkConfigRequestWithoutParam() *DescribeCustomLiveStreamWatermarkConfigRequest {

    return &DescribeCustomLiveStreamWatermarkConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarks:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码
- 取值范围 [1, 100000]
(Optional) */
func (r *DescribeCustomLiveStreamWatermarkConfigRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小
- 取值范围 [10, 100]
(Optional) */
func (r *DescribeCustomLiveStreamWatermarkConfigRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 水印配置查询过滤条件:
  - name:   publishDomain，必填(推流域名)
  - value:  参数
  - name:   appName，必填(应用名称)
  - value:  参数
  - name:   streamName，非必填(流名称)
  - value:  参数
(Optional) */
func (r *DescribeCustomLiveStreamWatermarkConfigRequest) SetFilters(filters []live.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomLiveStreamWatermarkConfigRequest) GetRegionId() string {
    return ""
}

type DescribeCustomLiveStreamWatermarkConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomLiveStreamWatermarkConfigResult `json:"result"`
}

type DescribeCustomLiveStreamWatermarkConfigResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    WatermarkConfigs []live.LiveStreamRecordConfig `json:"watermarkConfigs"`
}