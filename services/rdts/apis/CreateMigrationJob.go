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

type CreateMigrationJobRequest struct {

    core.JDCloudRequest

    /* 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1)  */
    RegionId string `json:"regionId"`

    /* 迁移任务的唯一标识  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1) (Required)
 * param instanceId: 迁移任务的唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateMigrationJobRequest(
    regionId string,
    instanceId string,
) *CreateMigrationJobRequest {

	return &CreateMigrationJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceId}:start",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1) (Required)
 * param instanceId: 迁移任务的唯一标识 (Required)
 */
func NewCreateMigrationJobRequestWithAllParams(
    regionId string,
    instanceId string,
) *CreateMigrationJobRequest {

    return &CreateMigrationJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceId}:start",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateMigrationJobRequestWithoutParam() *CreateMigrationJobRequest {

    return &CreateMigrationJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceId}:start",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1)(Required) */
func (r *CreateMigrationJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 迁移任务的唯一标识(Required) */
func (r *CreateMigrationJobRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateMigrationJobRequest) GetRegionId() string {
    return r.RegionId
}

type CreateMigrationJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateMigrationJobResult `json:"result"`
}

type CreateMigrationJobResult struct {
}