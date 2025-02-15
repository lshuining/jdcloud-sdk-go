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
    elite "github.com/lshuining/jdcloud-sdk-go/services/elite/models"
)

type GetStoreServiceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 购买用户pin  */
    BuyerPin string `json:"buyerPin"`

    /* 业务数据，与下单时的业务数据一致  */
    BusinessData string `json:"businessData"`

    /* 是否查询全部，如果传入false，则只查询当前时间有效的，否则查询所有的 (Optional) */
    QueryAll *bool `json:"queryAll"`
}

/*
 * param regionId: 地域ID (Required)
 * param buyerPin: 购买用户pin (Required)
 * param businessData: 业务数据，与下单时的业务数据一致 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetStoreServiceRequest(
    regionId string,
    buyerPin string,
    businessData string,
) *GetStoreServiceRequest {

	return &GetStoreServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/getStoreService",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BuyerPin: buyerPin,
        BusinessData: businessData,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param buyerPin: 购买用户pin (Required)
 * param businessData: 业务数据，与下单时的业务数据一致 (Required)
 * param queryAll: 是否查询全部，如果传入false，则只查询当前时间有效的，否则查询所有的 (Optional)
 */
func NewGetStoreServiceRequestWithAllParams(
    regionId string,
    buyerPin string,
    businessData string,
    queryAll *bool,
) *GetStoreServiceRequest {

    return &GetStoreServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/getStoreService",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BuyerPin: buyerPin,
        BusinessData: businessData,
        QueryAll: queryAll,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetStoreServiceRequestWithoutParam() *GetStoreServiceRequest {

    return &GetStoreServiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/getStoreService",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetStoreServiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param buyerPin: 购买用户pin(Required) */
func (r *GetStoreServiceRequest) SetBuyerPin(buyerPin string) {
    r.BuyerPin = buyerPin
}

/* param businessData: 业务数据，与下单时的业务数据一致(Required) */
func (r *GetStoreServiceRequest) SetBusinessData(businessData string) {
    r.BusinessData = businessData
}

/* param queryAll: 是否查询全部，如果传入false，则只查询当前时间有效的，否则查询所有的(Optional) */
func (r *GetStoreServiceRequest) SetQueryAll(queryAll bool) {
    r.QueryAll = &queryAll
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetStoreServiceRequest) GetRegionId() string {
    return r.RegionId
}

type GetStoreServiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetStoreServiceResult `json:"result"`
}

type GetStoreServiceResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data elite.StoreServiceVo `json:"data"`
}