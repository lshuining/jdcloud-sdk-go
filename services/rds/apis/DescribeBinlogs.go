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
    rds "github.com/lshuining/jdcloud-sdk-go/services/rds/models"
)

type DescribeBinlogsRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：10、20、30、50、100 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天 (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBinlogsRequest(
    regionId string,
    instanceId string,
) *DescribeBinlogsRequest {

	return &DescribeBinlogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/binlogs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。 (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：10、20、30、50、100 (Optional)
 * param startTime: 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天 (Optional)
 * param endTime: 查询结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天 (Optional)
 */
func NewDescribeBinlogsRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
    startTime *string,
    endTime *string,
) *DescribeBinlogsRequest {

    return &DescribeBinlogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/binlogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBinlogsRequestWithoutParam() *DescribeBinlogsRequest {

    return &DescribeBinlogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/binlogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeBinlogsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeBinlogsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。(Optional) */
func (r *DescribeBinlogsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：10、20、30、50、100(Optional) */
func (r *DescribeBinlogsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param startTime: 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天(Optional) */
func (r *DescribeBinlogsRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天(Optional) */
func (r *DescribeBinlogsRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBinlogsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBinlogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBinlogsResult `json:"result"`
}

type DescribeBinlogsResult struct {
    TotalCount int `json:"totalCount"`
    Binlogs []rds.Binlog `json:"binlogs"`
}