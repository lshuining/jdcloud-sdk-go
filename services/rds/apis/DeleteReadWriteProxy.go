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

type DeleteReadWriteProxyRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 读写分离代理服务ID  */
    ReadWriteProxyId string `json:"readWriteProxyId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param readWriteProxyId: 读写分离代理服务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteReadWriteProxyRequest(
    regionId string,
    readWriteProxyId string,
) *DeleteReadWriteProxyRequest {

	return &DeleteReadWriteProxyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/readWriteProxy/{readWriteProxyId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ReadWriteProxyId: readWriteProxyId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param readWriteProxyId: 读写分离代理服务ID (Required)
 */
func NewDeleteReadWriteProxyRequestWithAllParams(
    regionId string,
    readWriteProxyId string,
) *DeleteReadWriteProxyRequest {

    return &DeleteReadWriteProxyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/readWriteProxy/{readWriteProxyId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ReadWriteProxyId: readWriteProxyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteReadWriteProxyRequestWithoutParam() *DeleteReadWriteProxyRequest {

    return &DeleteReadWriteProxyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/readWriteProxy/{readWriteProxyId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DeleteReadWriteProxyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param readWriteProxyId: 读写分离代理服务ID(Required) */
func (r *DeleteReadWriteProxyRequest) SetReadWriteProxyId(readWriteProxyId string) {
    r.ReadWriteProxyId = readWriteProxyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteReadWriteProxyRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteReadWriteProxyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteReadWriteProxyResult `json:"result"`
}

type DeleteReadWriteProxyResult struct {
}