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
    mongodb "github.com/lshuining/jdcloud-sdk-go/services/mongodb/models"
    charge "github.com/lshuining/jdcloud-sdk-go/services/charge/models"
)

type CreateShardingInstanceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例规格  */
    ShardingInstanceSpec *mongodb.ShardingDBInstanceSpec `json:"shardingInstanceSpec"`

    /* 付费方式 (Optional) */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param shardingInstanceSpec: 实例规格 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateShardingInstanceRequest(
    regionId string,
    shardingInstanceSpec *mongodb.ShardingDBInstanceSpec,
) *CreateShardingInstanceRequest {

	return &CreateShardingInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/shardingInstances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ShardingInstanceSpec: shardingInstanceSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param shardingInstanceSpec: 实例规格 (Required)
 * param chargeSpec: 付费方式 (Optional)
 */
func NewCreateShardingInstanceRequestWithAllParams(
    regionId string,
    shardingInstanceSpec *mongodb.ShardingDBInstanceSpec,
    chargeSpec *charge.ChargeSpec,
) *CreateShardingInstanceRequest {

    return &CreateShardingInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/shardingInstances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ShardingInstanceSpec: shardingInstanceSpec,
        ChargeSpec: chargeSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateShardingInstanceRequestWithoutParam() *CreateShardingInstanceRequest {

    return &CreateShardingInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/shardingInstances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateShardingInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param shardingInstanceSpec: 实例规格(Required) */
func (r *CreateShardingInstanceRequest) SetShardingInstanceSpec(shardingInstanceSpec *mongodb.ShardingDBInstanceSpec) {
    r.ShardingInstanceSpec = shardingInstanceSpec
}

/* param chargeSpec: 付费方式(Optional) */
func (r *CreateShardingInstanceRequest) SetChargeSpec(chargeSpec *charge.ChargeSpec) {
    r.ChargeSpec = chargeSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateShardingInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateShardingInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateShardingInstanceResult `json:"result"`
}

type CreateShardingInstanceResult struct {
    InstanceId string `json:"instanceId"`
    OrderId string `json:"orderId"`
}