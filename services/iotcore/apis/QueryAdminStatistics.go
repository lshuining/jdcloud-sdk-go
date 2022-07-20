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

type QueryAdminStatisticsRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 过滤条件，产品Key (Optional) */
    ProductKey *string `json:"productKey"`

    /* 针对parentId下的子设备进行统计 (Optional) */
    ParentId *string `json:"parentId"`

    /* 采集器类型 (Optional) */
    DeviceCollectorType *string `json:"deviceCollectorType"`

    /* 查询的用户信息 (Optional) */
    QueryUserPin *string `json:"queryUserPin"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAdminStatisticsRequest(
    instanceId string,
    regionId string,
) *QueryAdminStatisticsRequest {

	return &QueryAdminStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/device:queryAdminStatistics",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param productKey: 过滤条件，产品Key (Optional)
 * param parentId: 针对parentId下的子设备进行统计 (Optional)
 * param deviceCollectorType: 采集器类型 (Optional)
 * param queryUserPin: 查询的用户信息 (Optional)
 */
func NewQueryAdminStatisticsRequestWithAllParams(
    instanceId string,
    regionId string,
    productKey *string,
    parentId *string,
    deviceCollectorType *string,
    queryUserPin *string,
) *QueryAdminStatisticsRequest {

    return &QueryAdminStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/device:queryAdminStatistics",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        ProductKey: productKey,
        ParentId: parentId,
        DeviceCollectorType: deviceCollectorType,
        QueryUserPin: queryUserPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAdminStatisticsRequestWithoutParam() *QueryAdminStatisticsRequest {

    return &QueryAdminStatisticsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/device:queryAdminStatistics",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *QueryAdminStatisticsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *QueryAdminStatisticsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productKey: 过滤条件，产品Key(Optional) */
func (r *QueryAdminStatisticsRequest) SetProductKey(productKey string) {
    r.ProductKey = &productKey
}

/* param parentId: 针对parentId下的子设备进行统计(Optional) */
func (r *QueryAdminStatisticsRequest) SetParentId(parentId string) {
    r.ParentId = &parentId
}

/* param deviceCollectorType: 采集器类型(Optional) */
func (r *QueryAdminStatisticsRequest) SetDeviceCollectorType(deviceCollectorType string) {
    r.DeviceCollectorType = &deviceCollectorType
}

/* param queryUserPin: 查询的用户信息(Optional) */
func (r *QueryAdminStatisticsRequest) SetQueryUserPin(queryUserPin string) {
    r.QueryUserPin = &queryUserPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAdminStatisticsRequest) GetRegionId() string {
    return r.RegionId
}

type QueryAdminStatisticsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAdminStatisticsResult `json:"result"`
}

type QueryAdminStatisticsResult struct {
    Devices int `json:"devices"`
    ActivatedDevices int `json:"activatedDevices"`
    OnlineDevices int `json:"onlineDevices"`
    MonthMessages int64 `json:"monthMessages"`
    MonthDuration int64 `json:"monthDuration"`
    Products int64 `json:"products"`
}