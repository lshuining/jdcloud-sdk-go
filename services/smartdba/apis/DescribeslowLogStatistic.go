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

type DescribeslowLogStatisticRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 查询时间，格式为：2021-11-11T15:04:05Z  */
    EndTime string `json:"endTime"`

    /* 数据库类型，默认MySQL (Optional) */
    DbType *string `json:"dbType"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param endTime: 查询时间，格式为：2021-11-11T15:04:05Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeslowLogStatisticRequest(
    regionId string,
    endTime string,
) *DescribeslowLogStatisticRequest {

	return &DescribeslowLogStatisticRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/slowLogStatistic",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param endTime: 查询时间，格式为：2021-11-11T15:04:05Z (Required)
 * param dbType: 数据库类型，默认MySQL (Optional)
 */
func NewDescribeslowLogStatisticRequestWithAllParams(
    regionId string,
    endTime string,
    dbType *string,
) *DescribeslowLogStatisticRequest {

    return &DescribeslowLogStatisticRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slowLogStatistic",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        EndTime: endTime,
        DbType: dbType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeslowLogStatisticRequestWithoutParam() *DescribeslowLogStatisticRequest {

    return &DescribeslowLogStatisticRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slowLogStatistic",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeslowLogStatisticRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param endTime: 查询时间，格式为：2021-11-11T15:04:05Z(Required) */
func (r *DescribeslowLogStatisticRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param dbType: 数据库类型，默认MySQL(Optional) */
func (r *DescribeslowLogStatisticRequest) SetDbType(dbType string) {
    r.DbType = &dbType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeslowLogStatisticRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeslowLogStatisticResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeslowLogStatisticResult `json:"result"`
}

type DescribeslowLogStatisticResult struct {
    Count3h int `json:"count3h"`
    Count24h int `json:"count24h"`
    Count72h int `json:"count72h"`
}