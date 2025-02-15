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
    dbs "github.com/lshuining/jdcloud-sdk-go/services/dbs/models"
)

type DescribePhysicalBackupsRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：[10,100]，且为10的整数倍 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePhysicalBackupsRequest(
    regionId string,
    backupPlanId string,
) *DescribePhysicalBackupsRequest {

	return &DescribePhysicalBackupsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:describePhysicalBackups",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BackupPlanId: backupPlanId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页 (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，且为10的整数倍 (Optional)
 */
func NewDescribePhysicalBackupsRequestWithAllParams(
    regionId string,
    backupPlanId string,
    pageNumber *int,
    pageSize *int,
) *DescribePhysicalBackupsRequest {

    return &DescribePhysicalBackupsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:describePhysicalBackups",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePhysicalBackupsRequestWithoutParam() *DescribePhysicalBackupsRequest {

    return &DescribePhysicalBackupsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:describePhysicalBackups",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *DescribePhysicalBackupsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *DescribePhysicalBackupsRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页(Optional) */
func (r *DescribePhysicalBackupsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，且为10的整数倍(Optional) */
func (r *DescribePhysicalBackupsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePhysicalBackupsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePhysicalBackupsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePhysicalBackupsResult `json:"result"`
}

type DescribePhysicalBackupsResult struct {
    TotalCount int `json:"totalCount"`
    PhysicalBackups []dbs.PhysicalBackups `json:"physicalBackups"`
}