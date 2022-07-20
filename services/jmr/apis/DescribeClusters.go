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
    jmr "github.com/lshuining/jdcloud-sdk-go/services/jmr/models"
)

type DescribeClustersRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 地域 (Optional) */
    DataCenter *string `json:"dataCenter"`

    /* 集群状态，CREATING，RUNNING，RELEASED，FAILED等 (Optional) */
    Status *string `json:"status"`

    /* 集群名称 (Optional) */
    ClusterName *string `json:"clusterName"`

    /* 排序，比如 id desc (Optional) */
    OrderBy *string `json:"orderBy"`

    /* 页数，默认为1 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 每页数目，默认为10 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeClustersRequest(
    regionId string,
) *DescribeClustersRequest {

	return &DescribeClustersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/clusters",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param dataCenter: 地域 (Optional)
 * param status: 集群状态，CREATING，RUNNING，RELEASED，FAILED等 (Optional)
 * param clusterName: 集群名称 (Optional)
 * param orderBy: 排序，比如 id desc (Optional)
 * param pageNum: 页数，默认为1 (Optional)
 * param pageSize: 每页数目，默认为10 (Optional)
 */
func NewDescribeClustersRequestWithAllParams(
    regionId string,
    dataCenter *string,
    status *string,
    clusterName *string,
    orderBy *string,
    pageNum *int,
    pageSize *int,
) *DescribeClustersRequest {

    return &DescribeClustersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataCenter: dataCenter,
        Status: status,
        ClusterName: clusterName,
        OrderBy: orderBy,
        PageNum: pageNum,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeClustersRequestWithoutParam() *DescribeClustersRequest {

    return &DescribeClustersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeClustersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataCenter: 地域(Optional) */
func (r *DescribeClustersRequest) SetDataCenter(dataCenter string) {
    r.DataCenter = &dataCenter
}

/* param status: 集群状态，CREATING，RUNNING，RELEASED，FAILED等(Optional) */
func (r *DescribeClustersRequest) SetStatus(status string) {
    r.Status = &status
}

/* param clusterName: 集群名称(Optional) */
func (r *DescribeClustersRequest) SetClusterName(clusterName string) {
    r.ClusterName = &clusterName
}

/* param orderBy: 排序，比如 id desc(Optional) */
func (r *DescribeClustersRequest) SetOrderBy(orderBy string) {
    r.OrderBy = &orderBy
}

/* param pageNum: 页数，默认为1(Optional) */
func (r *DescribeClustersRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 每页数目，默认为10(Optional) */
func (r *DescribeClustersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeClustersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeClustersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeClustersResult `json:"result"`
}

type DescribeClustersResult struct {
    TotalNum int `json:"totalNum"`
    Clusters []jmr.ClusterListNode `json:"clusters"`
    Status bool `json:"status"`
}