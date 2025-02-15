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
    smartdba "github.com/lshuining/jdcloud-sdk-go/services/smartdba/models"
)

type DescribeLockAnalyseRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceGid string `json:"instanceGid"`

    /* 锁诊断Id  */
    AnalyseId string `json:"analyseId"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param analyseId: 锁诊断Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLockAnalyseRequest(
    regionId string,
    instanceGid string,
    analyseId string,
) *DescribeLockAnalyseRequest {

	return &DescribeLockAnalyseRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/lockAnalyse",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
        AnalyseId: analyseId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param analyseId: 锁诊断Id (Required)
 */
func NewDescribeLockAnalyseRequestWithAllParams(
    regionId string,
    instanceGid string,
    analyseId string,
) *DescribeLockAnalyseRequest {

    return &DescribeLockAnalyseRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/lockAnalyse",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
        AnalyseId: analyseId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLockAnalyseRequestWithoutParam() *DescribeLockAnalyseRequest {

    return &DescribeLockAnalyseRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/lockAnalyse",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeLockAnalyseRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeLockAnalyseRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

/* param analyseId: 锁诊断Id(Required) */
func (r *DescribeLockAnalyseRequest) SetAnalyseId(analyseId string) {
    r.AnalyseId = analyseId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLockAnalyseRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLockAnalyseResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLockAnalyseResult `json:"result"`
}

type DescribeLockAnalyseResult struct {
    AnalyseId string `json:"analyseId"`
    OccurTime string `json:"occurTime"`
    DeadLockLog string `json:"deadLockLog"`
    Data []smartdba.DeadLockDetail `json:"data"`
}