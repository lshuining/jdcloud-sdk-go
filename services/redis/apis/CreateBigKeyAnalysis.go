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

type CreateBigKeyAnalysisRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* String类型阈值 (Optional) */
    StringSize *int `json:"stringSize"`

    /* List类型阈值 (Optional) */
    ListSize *int `json:"listSize"`

    /* Hash类型阈值 (Optional) */
    HashSize *int `json:"hashSize"`

    /* Set类型阈值 (Optional) */
    SetSize *int `json:"setSize"`

    /* Zset类型阈值 (Optional) */
    ZsetSize *int `json:"zsetSize"`

    /* top值，范围10~1000 (Optional) */
    Top *int `json:"top"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBigKeyAnalysisRequest(
    regionId string,
    cacheInstanceId string,
) *CreateBigKeyAnalysisRequest {

	return &CreateBigKeyAnalysisRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKey",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param stringSize: String类型阈值 (Optional)
 * param listSize: List类型阈值 (Optional)
 * param hashSize: Hash类型阈值 (Optional)
 * param setSize: Set类型阈值 (Optional)
 * param zsetSize: Zset类型阈值 (Optional)
 * param top: top值，范围10~1000 (Optional)
 */
func NewCreateBigKeyAnalysisRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    stringSize *int,
    listSize *int,
    hashSize *int,
    setSize *int,
    zsetSize *int,
    top *int,
) *CreateBigKeyAnalysisRequest {

    return &CreateBigKeyAnalysisRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKey",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        StringSize: stringSize,
        ListSize: listSize,
        HashSize: hashSize,
        SetSize: setSize,
        ZsetSize: zsetSize,
        Top: top,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBigKeyAnalysisRequestWithoutParam() *CreateBigKeyAnalysisRequest {

    return &CreateBigKeyAnalysisRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKey",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *CreateBigKeyAnalysisRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *CreateBigKeyAnalysisRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param stringSize: String类型阈值(Optional) */
func (r *CreateBigKeyAnalysisRequest) SetStringSize(stringSize int) {
    r.StringSize = &stringSize
}

/* param listSize: List类型阈值(Optional) */
func (r *CreateBigKeyAnalysisRequest) SetListSize(listSize int) {
    r.ListSize = &listSize
}

/* param hashSize: Hash类型阈值(Optional) */
func (r *CreateBigKeyAnalysisRequest) SetHashSize(hashSize int) {
    r.HashSize = &hashSize
}

/* param setSize: Set类型阈值(Optional) */
func (r *CreateBigKeyAnalysisRequest) SetSetSize(setSize int) {
    r.SetSize = &setSize
}

/* param zsetSize: Zset类型阈值(Optional) */
func (r *CreateBigKeyAnalysisRequest) SetZsetSize(zsetSize int) {
    r.ZsetSize = &zsetSize
}

/* param top: top值，范围10~1000(Optional) */
func (r *CreateBigKeyAnalysisRequest) SetTop(top int) {
    r.Top = &top
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBigKeyAnalysisRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBigKeyAnalysisResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBigKeyAnalysisResult `json:"result"`
}

type CreateBigKeyAnalysisResult struct {
}