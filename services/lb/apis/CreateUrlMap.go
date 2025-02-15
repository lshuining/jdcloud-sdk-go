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

type CreateUrlMapRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    UrlMapName string `json:"urlMapName"`

    /* 转发规则组所属负载均衡的Id  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param urlMapName: 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param loadBalancerId: 转发规则组所属负载均衡的Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateUrlMapRequest(
    regionId string,
    urlMapName string,
    loadBalancerId string,
) *CreateUrlMapRequest {

	return &CreateUrlMapRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/urlMaps/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UrlMapName: urlMapName,
        LoadBalancerId: loadBalancerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param urlMapName: 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param loadBalancerId: 转发规则组所属负载均衡的Id (Required)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 */
func NewCreateUrlMapRequestWithAllParams(
    regionId string,
    urlMapName string,
    loadBalancerId string,
    description *string,
) *CreateUrlMapRequest {

    return &CreateUrlMapRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UrlMapName: urlMapName,
        LoadBalancerId: loadBalancerId,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateUrlMapRequestWithoutParam() *CreateUrlMapRequest {

    return &CreateUrlMapRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateUrlMapRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param urlMapName: 转发规则组名称，同一个负载均衡下，名称不能重复,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateUrlMapRequest) SetUrlMapName(urlMapName string) {
    r.UrlMapName = urlMapName
}

/* param loadBalancerId: 转发规则组所属负载均衡的Id(Required) */
func (r *CreateUrlMapRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateUrlMapRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateUrlMapRequest) GetRegionId() string {
    return r.RegionId
}

type CreateUrlMapResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateUrlMapResult `json:"result"`
}

type CreateUrlMapResult struct {
    UrlMapId string `json:"urlMapId"`
}