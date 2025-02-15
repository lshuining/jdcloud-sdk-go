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
    smartdba "github.com/lshuining/jdcloud-sdk-go/services/smartdba/models"
)

type GetPanelRelateClustersRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 监控大盘id  */
    PanelGid string `json:"panelGid"`

    /* 数据库类型,默认MySQL (Optional) */
    DbType *string `json:"dbType"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param panelGid: 监控大盘id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetPanelRelateClustersRequest(
    regionId string,
    panelGid string,
) *GetPanelRelateClustersRequest {

	return &GetPanelRelateClustersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/panels/{panelGid}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        PanelGid: panelGid,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param panelGid: 监控大盘id (Required)
 * param dbType: 数据库类型,默认MySQL (Optional)
 */
func NewGetPanelRelateClustersRequestWithAllParams(
    regionId string,
    panelGid string,
    dbType *string,
) *GetPanelRelateClustersRequest {

    return &GetPanelRelateClustersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/panels/{panelGid}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        PanelGid: panelGid,
        DbType: dbType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetPanelRelateClustersRequestWithoutParam() *GetPanelRelateClustersRequest {

    return &GetPanelRelateClustersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/panels/{panelGid}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *GetPanelRelateClustersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param panelGid: 监控大盘id(Required) */
func (r *GetPanelRelateClustersRequest) SetPanelGid(panelGid string) {
    r.PanelGid = panelGid
}

/* param dbType: 数据库类型,默认MySQL(Optional) */
func (r *GetPanelRelateClustersRequest) SetDbType(dbType string) {
    r.DbType = &dbType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetPanelRelateClustersRequest) GetRegionId() string {
    return r.RegionId
}

type GetPanelRelateClustersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetPanelRelateClustersResult `json:"result"`
}

type GetPanelRelateClustersResult struct {
    Data []smartdba.PanelClusterItem `json:"data"`
}