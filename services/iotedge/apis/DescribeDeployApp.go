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
    iotedge "github.com/lshuining/jdcloud-sdk-go/services/iotedge/models"
)

type DescribeDeployAppRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 硬件版本  */
    HardwareId string `json:"hardwareId"`

    /* OSID  */
    OsId string `json:"osId"`

    /* Edge名称  */
    EdgeName string `json:"edgeName"`

    /* Edge对应的产品Key  */
    ProductKey string `json:"productKey"`

    /* App类型(1-设备服务 2-边缘应用)  */
    AppType int `json:"appType"`

    /* 当前的规则位置 (Optional) */
    NowPage *int `json:"nowPage"`

    /* 显示多少个数据 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 排序方式 (Optional) */
    Order *int `json:"order"`

    /* 排序依据的关键词 (Optional) */
    Property *int `json:"property"`

    /* 模糊搜索关键字 (Optional) */
    SearchText *string `json:"searchText"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param hardwareId: 硬件版本 (Required)
 * param osId: OSID (Required)
 * param edgeName: Edge名称 (Required)
 * param productKey: Edge对应的产品Key (Required)
 * param appType: App类型(1-设备服务 2-边缘应用) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeployAppRequest(
    instanceId string,
    regionId string,
    hardwareId string,
    osId string,
    edgeName string,
    productKey string,
    appType int,
) *DescribeDeployAppRequest {

	return &DescribeDeployAppRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:describeDeployApp",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
        HardwareId: hardwareId,
        OsId: osId,
        EdgeName: edgeName,
        ProductKey: productKey,
        AppType: appType,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param hardwareId: 硬件版本 (Required)
 * param osId: OSID (Required)
 * param edgeName: Edge名称 (Required)
 * param productKey: Edge对应的产品Key (Required)
 * param appType: App类型(1-设备服务 2-边缘应用) (Required)
 * param nowPage: 当前的规则位置 (Optional)
 * param pageSize: 显示多少个数据 (Optional)
 * param order: 排序方式 (Optional)
 * param property: 排序依据的关键词 (Optional)
 * param searchText: 模糊搜索关键字 (Optional)
 */
func NewDescribeDeployAppRequestWithAllParams(
    instanceId string,
    regionId string,
    hardwareId string,
    osId string,
    edgeName string,
    productKey string,
    appType int,
    nowPage *int,
    pageSize *int,
    order *int,
    property *int,
    searchText *string,
) *DescribeDeployAppRequest {

    return &DescribeDeployAppRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:describeDeployApp",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        HardwareId: hardwareId,
        OsId: osId,
        EdgeName: edgeName,
        ProductKey: productKey,
        AppType: appType,
        NowPage: nowPage,
        PageSize: pageSize,
        Order: order,
        Property: property,
        SearchText: searchText,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeployAppRequestWithoutParam() *DescribeDeployAppRequest {

    return &DescribeDeployAppRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:describeDeployApp",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *DescribeDeployAppRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *DescribeDeployAppRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param hardwareId: 硬件版本(Required) */
func (r *DescribeDeployAppRequest) SetHardwareId(hardwareId string) {
    r.HardwareId = hardwareId
}

/* param osId: OSID(Required) */
func (r *DescribeDeployAppRequest) SetOsId(osId string) {
    r.OsId = osId
}

/* param edgeName: Edge名称(Required) */
func (r *DescribeDeployAppRequest) SetEdgeName(edgeName string) {
    r.EdgeName = edgeName
}

/* param productKey: Edge对应的产品Key(Required) */
func (r *DescribeDeployAppRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

/* param appType: App类型(1-设备服务 2-边缘应用)(Required) */
func (r *DescribeDeployAppRequest) SetAppType(appType int) {
    r.AppType = appType
}

/* param nowPage: 当前的规则位置(Optional) */
func (r *DescribeDeployAppRequest) SetNowPage(nowPage int) {
    r.NowPage = &nowPage
}

/* param pageSize: 显示多少个数据(Optional) */
func (r *DescribeDeployAppRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param order: 排序方式(Optional) */
func (r *DescribeDeployAppRequest) SetOrder(order int) {
    r.Order = &order
}

/* param property: 排序依据的关键词(Optional) */
func (r *DescribeDeployAppRequest) SetProperty(property int) {
    r.Property = &property
}

/* param searchText: 模糊搜索关键字(Optional) */
func (r *DescribeDeployAppRequest) SetSearchText(searchText string) {
    r.SearchText = &searchText
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeployAppRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeployAppResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeployAppResult `json:"result"`
}

type DescribeDeployAppResult struct {
    AppsResp []iotedge.DescribeAppsRespVO `json:"appsResp"`
    Page interface{} `json:"page"`
}