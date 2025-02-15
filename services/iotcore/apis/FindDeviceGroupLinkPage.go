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
    iotcore "github.com/lshuining/jdcloud-sdk-go/services/iotcore/models"
)

type FindDeviceGroupLinkPageRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 产品Key (Optional) */
    ProductKey *string `json:"productKey"`

    /* 设备名称 (Optional) */
    DeviceName *string `json:"deviceName"`

    /* 采集器类型 (Optional) */
    DeviceCollectorType *string `json:"deviceCollectorType"`

    /* 组名称 (Optional) */
    GroupName *string `json:"groupName"`

    /* 查询的用户组 (Optional) */
    UserPin *string `json:"userPin"`

    /* 组标签 (Optional) */
    Tag *string `json:"tag"`

    /* 组ID (Optional) */
    GroupId *string `json:"groupId"`

    /* 厂商名称 (Optional) */
    Manufacturer *string `json:"manufacturer"`

    /* 设备型号 (Optional) */
    Model *string `json:"model"`

    /* 订单号 (Optional) */
    OrderId *int `json:"orderId"`

    /* 设备状态 (Optional) */
    Status *int `json:"status"`

    /* 当前页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 排序字段 (Optional) */
    Order *string `json:"order"`

    /* 排序方式（asc desc） (Optional) */
    Direction *string `json:"direction"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewFindDeviceGroupLinkPageRequest(
    instanceId string,
    regionId string,
) *FindDeviceGroupLinkPageRequest {

	return &FindDeviceGroupLinkPageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/devicegrouplink:get",
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
 * param productKey: 产品Key (Optional)
 * param deviceName: 设备名称 (Optional)
 * param deviceCollectorType: 采集器类型 (Optional)
 * param groupName: 组名称 (Optional)
 * param userPin: 查询的用户组 (Optional)
 * param tag: 组标签 (Optional)
 * param groupId: 组ID (Optional)
 * param manufacturer: 厂商名称 (Optional)
 * param model: 设备型号 (Optional)
 * param orderId: 订单号 (Optional)
 * param status: 设备状态 (Optional)
 * param pageNumber: 当前页码 (Optional)
 * param pageSize: 每页大小 (Optional)
 * param order: 排序字段 (Optional)
 * param direction: 排序方式（asc desc） (Optional)
 */
func NewFindDeviceGroupLinkPageRequestWithAllParams(
    instanceId string,
    regionId string,
    productKey *string,
    deviceName *string,
    deviceCollectorType *string,
    groupName *string,
    userPin *string,
    tag *string,
    groupId *string,
    manufacturer *string,
    model *string,
    orderId *int,
    status *int,
    pageNumber *int,
    pageSize *int,
    order *string,
    direction *string,
) *FindDeviceGroupLinkPageRequest {

    return &FindDeviceGroupLinkPageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/devicegrouplink:get",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        ProductKey: productKey,
        DeviceName: deviceName,
        DeviceCollectorType: deviceCollectorType,
        GroupName: groupName,
        UserPin: userPin,
        Tag: tag,
        GroupId: groupId,
        Manufacturer: manufacturer,
        Model: model,
        OrderId: orderId,
        Status: status,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Order: order,
        Direction: direction,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewFindDeviceGroupLinkPageRequestWithoutParam() *FindDeviceGroupLinkPageRequest {

    return &FindDeviceGroupLinkPageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/devicegrouplink:get",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *FindDeviceGroupLinkPageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *FindDeviceGroupLinkPageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productKey: 产品Key(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetProductKey(productKey string) {
    r.ProductKey = &productKey
}

/* param deviceName: 设备名称(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetDeviceName(deviceName string) {
    r.DeviceName = &deviceName
}

/* param deviceCollectorType: 采集器类型(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetDeviceCollectorType(deviceCollectorType string) {
    r.DeviceCollectorType = &deviceCollectorType
}

/* param groupName: 组名称(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetGroupName(groupName string) {
    r.GroupName = &groupName
}

/* param userPin: 查询的用户组(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetUserPin(userPin string) {
    r.UserPin = &userPin
}

/* param tag: 组标签(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetTag(tag string) {
    r.Tag = &tag
}

/* param groupId: 组ID(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetGroupId(groupId string) {
    r.GroupId = &groupId
}

/* param manufacturer: 厂商名称(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetManufacturer(manufacturer string) {
    r.Manufacturer = &manufacturer
}

/* param model: 设备型号(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetModel(model string) {
    r.Model = &model
}

/* param orderId: 订单号(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetOrderId(orderId int) {
    r.OrderId = &orderId
}

/* param status: 设备状态(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetStatus(status int) {
    r.Status = &status
}

/* param pageNumber: 当前页码(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页大小(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param order: 排序字段(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetOrder(order string) {
    r.Order = &order
}

/* param direction: 排序方式（asc desc）(Optional) */
func (r *FindDeviceGroupLinkPageRequest) SetDirection(direction string) {
    r.Direction = &direction
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r FindDeviceGroupLinkPageRequest) GetRegionId() string {
    return r.RegionId
}

type FindDeviceGroupLinkPageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result FindDeviceGroupLinkPageResult `json:"result"`
}

type FindDeviceGroupLinkPageResult struct {
    PageSize int `json:"pageSize"`
    NowPage int `json:"nowPage"`
    TotalSize int `json:"totalSize"`
    TotalPage int `json:"totalPage"`
    Data []iotcore.DeviceGroupInfo `json:"data"`
}