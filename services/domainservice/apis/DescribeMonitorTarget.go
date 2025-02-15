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

type DescribeMonitorTargetRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 子域名  */
    SubDomainName string `json:"subDomainName"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param subDomainName: 子域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMonitorTargetRequest(
    regionId string,
    domainId string,
    subDomainName string,
) *DescribeMonitorTargetRequest {

	return &DescribeMonitorTargetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/monitorTarget",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        SubDomainName: subDomainName,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param subDomainName: 子域名 (Required)
 */
func NewDescribeMonitorTargetRequestWithAllParams(
    regionId string,
    domainId string,
    subDomainName string,
) *DescribeMonitorTargetRequest {

    return &DescribeMonitorTargetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/monitorTarget",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        SubDomainName: subDomainName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMonitorTargetRequestWithoutParam() *DescribeMonitorTargetRequest {

    return &DescribeMonitorTargetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/monitorTarget",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DescribeMonitorTargetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *DescribeMonitorTargetRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param subDomainName: 子域名(Required) */
func (r *DescribeMonitorTargetRequest) SetSubDomainName(subDomainName string) {
    r.SubDomainName = subDomainName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMonitorTargetRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMonitorTargetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMonitorTargetResult `json:"result"`
}

type DescribeMonitorTargetResult struct {
    Data []string `json:"data"`
}