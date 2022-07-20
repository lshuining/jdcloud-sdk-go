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

type UpdateThingShadowRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 设备唯一标识  */
    Identifier string `json:"identifier"`

    /* 产品Key  */
    ProductKey string `json:"productKey"`

    /* 运行状态 (Optional) */
    State *interface{} `json:"state"`

    /* 设备影子版本,当前版本加1，当前版本默认其实版本为-1
用户主动更新版本号时，设备影子会检查请求中的主动更新版本号是否大于当前版本号。
如果大于当前版本号，则更新设备影子，并将影子版本值更新到请求的版本中，反之则会拒绝更新设备影子。
影子版本参数为Integer型
取值范围：0到2147483647(2的31次方-1)
当取值达到最大值2147483647(2的31次方-1)时，请求中的主动更新版本号应为-1
 (Optional) */
    Version *int `json:"version"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param identifier: 设备唯一标识 (Required)
 * param productKey: 产品Key (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateThingShadowRequest(
    instanceId string,
    regionId string,
    identifier string,
    productKey string,
) *UpdateThingShadowRequest {

	return &UpdateThingShadowRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/devices/{identifier}/shadow",
			Method:  "PATCH",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
        Identifier: identifier,
        ProductKey: productKey,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param identifier: 设备唯一标识 (Required)
 * param productKey: 产品Key (Required)
 * param state: 运行状态 (Optional)
 * param version: 设备影子版本,当前版本加1，当前版本默认其实版本为-1
用户主动更新版本号时，设备影子会检查请求中的主动更新版本号是否大于当前版本号。
如果大于当前版本号，则更新设备影子，并将影子版本值更新到请求的版本中，反之则会拒绝更新设备影子。
影子版本参数为Integer型
取值范围：0到2147483647(2的31次方-1)
当取值达到最大值2147483647(2的31次方-1)时，请求中的主动更新版本号应为-1
 (Optional)
 */
func NewUpdateThingShadowRequestWithAllParams(
    instanceId string,
    regionId string,
    identifier string,
    productKey string,
    state *interface{},
    version *int,
) *UpdateThingShadowRequest {

    return &UpdateThingShadowRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/devices/{identifier}/shadow",
            Method:  "PATCH",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        Identifier: identifier,
        ProductKey: productKey,
        State: state,
        Version: version,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateThingShadowRequestWithoutParam() *UpdateThingShadowRequest {

    return &UpdateThingShadowRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/devices/{identifier}/shadow",
            Method:  "PATCH",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *UpdateThingShadowRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *UpdateThingShadowRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param identifier: 设备唯一标识(Required) */
func (r *UpdateThingShadowRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

/* param productKey: 产品Key(Required) */
func (r *UpdateThingShadowRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

/* param state: 运行状态(Optional) */
func (r *UpdateThingShadowRequest) SetState(state interface{}) {
    r.State = &state
}

/* param version: 设备影子版本,当前版本加1，当前版本默认其实版本为-1
用户主动更新版本号时，设备影子会检查请求中的主动更新版本号是否大于当前版本号。
如果大于当前版本号，则更新设备影子，并将影子版本值更新到请求的版本中，反之则会拒绝更新设备影子。
影子版本参数为Integer型
取值范围：0到2147483647(2的31次方-1)
当取值达到最大值2147483647(2的31次方-1)时，请求中的主动更新版本号应为-1
(Optional) */
func (r *UpdateThingShadowRequest) SetVersion(version int) {
    r.Version = &version
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateThingShadowRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateThingShadowResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateThingShadowResult `json:"result"`
}

type UpdateThingShadowResult struct {
}