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

type LoongrayQueryPageRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 设备名称，模糊匹配 (Optional) */
    DeviceName *string `json:"deviceName"`

    /* 设备厂商，模糊匹配 (Optional) */
    Manufacturer *string `json:"manufacturer"`

    /* 设备型号，模糊匹配 (Optional) */
    Model *string `json:"model"`

    /* 设备状态 0-未激活，1-激活离线，2-激活在线 (Optional) */
    Status *int `json:"status"`

    /* 设备所归属的产品Key (Optional) */
    ProductKey *string `json:"productKey"`

    /* 设备类型，同产品类型，0-设备，1-网关 (Optional) */
    DeviceType *int `json:"deviceType"`

    /* 当前页数 (Optional) */
    NowPage *int `json:"nowPage"`

    /* 每页的数据条数 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 排序关键字--name,type,productKey,status--最多支持一个字段 (Optional) */
    Order *string `json:"order"`

    /* 顺序，升序降序--asc,desc (Optional) */
    Direction *string `json:"direction"`

    /* 父设备Id (Optional) */
    ParentId *string `json:"parentId"`

    /* 订单号 (Optional) */
    OrderId *int `json:"orderId"`

    /* 设备采集器类型 (Optional) */
    DeviceCollectorType *string `json:"deviceCollectorType"`

    /* 查询的userPin (Optional) */
    QueryUserPin *string `json:"queryUserPin"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewLoongrayQueryPageRequest(
    instanceId string,
    regionId string,
) *LoongrayQueryPageRequest {

	return &LoongrayQueryPageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/devices:loongrayQueryPage",
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
 * param deviceName: 设备名称，模糊匹配 (Optional)
 * param manufacturer: 设备厂商，模糊匹配 (Optional)
 * param model: 设备型号，模糊匹配 (Optional)
 * param status: 设备状态 0-未激活，1-激活离线，2-激活在线 (Optional)
 * param productKey: 设备所归属的产品Key (Optional)
 * param deviceType: 设备类型，同产品类型，0-设备，1-网关 (Optional)
 * param nowPage: 当前页数 (Optional)
 * param pageSize: 每页的数据条数 (Optional)
 * param order: 排序关键字--name,type,productKey,status--最多支持一个字段 (Optional)
 * param direction: 顺序，升序降序--asc,desc (Optional)
 * param parentId: 父设备Id (Optional)
 * param orderId: 订单号 (Optional)
 * param deviceCollectorType: 设备采集器类型 (Optional)
 * param queryUserPin: 查询的userPin (Optional)
 */
func NewLoongrayQueryPageRequestWithAllParams(
    instanceId string,
    regionId string,
    deviceName *string,
    manufacturer *string,
    model *string,
    status *int,
    productKey *string,
    deviceType *int,
    nowPage *int,
    pageSize *int,
    order *string,
    direction *string,
    parentId *string,
    orderId *int,
    deviceCollectorType *string,
    queryUserPin *string,
) *LoongrayQueryPageRequest {

    return &LoongrayQueryPageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/devices:loongrayQueryPage",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        DeviceName: deviceName,
        Manufacturer: manufacturer,
        Model: model,
        Status: status,
        ProductKey: productKey,
        DeviceType: deviceType,
        NowPage: nowPage,
        PageSize: pageSize,
        Order: order,
        Direction: direction,
        ParentId: parentId,
        OrderId: orderId,
        DeviceCollectorType: deviceCollectorType,
        QueryUserPin: queryUserPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewLoongrayQueryPageRequestWithoutParam() *LoongrayQueryPageRequest {

    return &LoongrayQueryPageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/devices:loongrayQueryPage",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *LoongrayQueryPageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *LoongrayQueryPageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param deviceName: 设备名称，模糊匹配(Optional) */
func (r *LoongrayQueryPageRequest) SetDeviceName(deviceName string) {
    r.DeviceName = &deviceName
}

/* param manufacturer: 设备厂商，模糊匹配(Optional) */
func (r *LoongrayQueryPageRequest) SetManufacturer(manufacturer string) {
    r.Manufacturer = &manufacturer
}

/* param model: 设备型号，模糊匹配(Optional) */
func (r *LoongrayQueryPageRequest) SetModel(model string) {
    r.Model = &model
}

/* param status: 设备状态 0-未激活，1-激活离线，2-激活在线(Optional) */
func (r *LoongrayQueryPageRequest) SetStatus(status int) {
    r.Status = &status
}

/* param productKey: 设备所归属的产品Key(Optional) */
func (r *LoongrayQueryPageRequest) SetProductKey(productKey string) {
    r.ProductKey = &productKey
}

/* param deviceType: 设备类型，同产品类型，0-设备，1-网关(Optional) */
func (r *LoongrayQueryPageRequest) SetDeviceType(deviceType int) {
    r.DeviceType = &deviceType
}

/* param nowPage: 当前页数(Optional) */
func (r *LoongrayQueryPageRequest) SetNowPage(nowPage int) {
    r.NowPage = &nowPage
}

/* param pageSize: 每页的数据条数(Optional) */
func (r *LoongrayQueryPageRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param order: 排序关键字--name,type,productKey,status--最多支持一个字段(Optional) */
func (r *LoongrayQueryPageRequest) SetOrder(order string) {
    r.Order = &order
}

/* param direction: 顺序，升序降序--asc,desc(Optional) */
func (r *LoongrayQueryPageRequest) SetDirection(direction string) {
    r.Direction = &direction
}

/* param parentId: 父设备Id(Optional) */
func (r *LoongrayQueryPageRequest) SetParentId(parentId string) {
    r.ParentId = &parentId
}

/* param orderId: 订单号(Optional) */
func (r *LoongrayQueryPageRequest) SetOrderId(orderId int) {
    r.OrderId = &orderId
}

/* param deviceCollectorType: 设备采集器类型(Optional) */
func (r *LoongrayQueryPageRequest) SetDeviceCollectorType(deviceCollectorType string) {
    r.DeviceCollectorType = &deviceCollectorType
}

/* param queryUserPin: 查询的userPin(Optional) */
func (r *LoongrayQueryPageRequest) SetQueryUserPin(queryUserPin string) {
    r.QueryUserPin = &queryUserPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r LoongrayQueryPageRequest) GetRegionId() string {
    return r.RegionId
}

type LoongrayQueryPageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result LoongrayQueryPageResult `json:"result"`
}

type LoongrayQueryPageResult struct {
    PageSize int `json:"pageSize"`
    NowPage int `json:"nowPage"`
    TotalSize int `json:"totalSize"`
    TotalPage int `json:"totalPage"`
    Data []iotcore.DeviceVO `json:"data"`
}