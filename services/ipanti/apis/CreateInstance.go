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
    ipanti "github.com/lshuining/jdcloud-sdk-go/services/ipanti/models"
)

type CreateInstanceRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 新购或升级实例请求参数  */
    CreateInstanceSpec *ipanti.CreateInstanceSpec `json:"createInstanceSpec"`

    /* 自动续费配置, 默认不开通, 仅新购实例时可设置 (Optional) */
    AutoRenewalSpec *ipanti.AutoRenewalSpec `json:"autoRenewalSpec"`

    /* 自动支付标识 (Optional) */
    AutoPay *bool `json:"autoPay"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param createInstanceSpec: 新购或升级实例请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstanceRequest(
    regionId string,
    createInstanceSpec *ipanti.CreateInstanceSpec,
) *CreateInstanceRequest {

	return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CreateInstanceSpec: createInstanceSpec,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param createInstanceSpec: 新购或升级实例请求参数 (Required)
 * param autoRenewalSpec: 自动续费配置, 默认不开通, 仅新购实例时可设置 (Optional)
 * param autoPay: 自动支付标识 (Optional)
 */
func NewCreateInstanceRequestWithAllParams(
    regionId string,
    createInstanceSpec *ipanti.CreateInstanceSpec,
    autoRenewalSpec *ipanti.AutoRenewalSpec,
    autoPay *bool,
) *CreateInstanceRequest {

    return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CreateInstanceSpec: createInstanceSpec,
        AutoRenewalSpec: autoRenewalSpec,
        AutoPay: autoPay,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateInstanceRequestWithoutParam() *CreateInstanceRequest {

    return &CreateInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *CreateInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param createInstanceSpec: 新购或升级实例请求参数(Required) */
func (r *CreateInstanceRequest) SetCreateInstanceSpec(createInstanceSpec *ipanti.CreateInstanceSpec) {
    r.CreateInstanceSpec = createInstanceSpec
}

/* param autoRenewalSpec: 自动续费配置, 默认不开通, 仅新购实例时可设置(Optional) */
func (r *CreateInstanceRequest) SetAutoRenewalSpec(autoRenewalSpec *ipanti.AutoRenewalSpec) {
    r.AutoRenewalSpec = autoRenewalSpec
}

/* param autoPay: 自动支付标识(Optional) */
func (r *CreateInstanceRequest) SetAutoPay(autoPay bool) {
    r.AutoPay = &autoPay
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstanceResult `json:"result"`
}

type CreateInstanceResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
    OrderNumber string `json:"orderNumber"`
}