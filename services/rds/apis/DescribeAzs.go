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

type DescribeAzsRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS引擎类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)  */
    Engine string `json:"engine"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param engine: RDS引擎类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAzsRequest(
    regionId string,
    engine string,
) *DescribeAzsRequest {

	return &DescribeAzsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/azs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Engine: engine,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param engine: RDS引擎类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Required)
 */
func NewDescribeAzsRequestWithAllParams(
    regionId string,
    engine string,
) *DescribeAzsRequest {

    return &DescribeAzsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/azs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Engine: engine,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAzsRequestWithoutParam() *DescribeAzsRequest {

    return &DescribeAzsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/azs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeAzsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param engine: RDS引擎类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)(Required) */
func (r *DescribeAzsRequest) SetEngine(engine string) {
    r.Engine = engine
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAzsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAzsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAzsResult `json:"result"`
}

type DescribeAzsResult struct {
    Azs []string `json:"azs"`
}