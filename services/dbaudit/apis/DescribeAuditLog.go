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
    dbaudit "github.com/lshuining/jdcloud-sdk-go/services/dbaudit/models"
)

type DescribeAuditLogRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 审计实例ID  */
    InsId string `json:"insId"`

    /* 审计日志ID  */
    LogId string `json:"logId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param logId: 审计日志ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAuditLogRequest(
    regionId string,
    insId string,
    logId string,
) *DescribeAuditLogRequest {

	return &DescribeAuditLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{insId}/logs/{logId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InsId: insId,
        LogId: logId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param logId: 审计日志ID (Required)
 */
func NewDescribeAuditLogRequestWithAllParams(
    regionId string,
    insId string,
    logId string,
) *DescribeAuditLogRequest {

    return &DescribeAuditLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/logs/{logId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InsId: insId,
        LogId: logId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAuditLogRequestWithoutParam() *DescribeAuditLogRequest {

    return &DescribeAuditLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/logs/{logId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeAuditLogRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param insId: 审计实例ID(Required) */
func (r *DescribeAuditLogRequest) SetInsId(insId string) {
    r.InsId = insId
}

/* param logId: 审计日志ID(Required) */
func (r *DescribeAuditLogRequest) SetLogId(logId string) {
    r.LogId = logId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAuditLogRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAuditLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAuditLogResult `json:"result"`
}

type DescribeAuditLogResult struct {
    AuditLogDetail dbaudit.AuditLogDetail `json:"auditLogDetail"`
}