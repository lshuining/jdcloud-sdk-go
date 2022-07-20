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

type SetMonitorRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 探测周期，取值1和5，单位为分钟 (Optional) */
    Cycle *int `json:"cycle"`

    /* 探测路径 (Optional) */
    MonitorPath *string `json:"monitorPath"`

    /* http请求头 (Optional) */
    HttpRequestHeader *interface{} `json:"httpRequestHeader"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetMonitorRequest(
    domain string,
) *SetMonitorRequest {

	return &SetMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/monitor",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param cycle: 探测周期，取值1和5，单位为分钟 (Optional)
 * param monitorPath: 探测路径 (Optional)
 * param httpRequestHeader: http请求头 (Optional)
 */
func NewSetMonitorRequestWithAllParams(
    domain string,
    cycle *int,
    monitorPath *string,
    httpRequestHeader *interface{},
) *SetMonitorRequest {

    return &SetMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/monitor",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Cycle: cycle,
        MonitorPath: monitorPath,
        HttpRequestHeader: httpRequestHeader,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetMonitorRequestWithoutParam() *SetMonitorRequest {

    return &SetMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/monitor",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetMonitorRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param cycle: 探测周期，取值1和5，单位为分钟(Optional) */
func (r *SetMonitorRequest) SetCycle(cycle int) {
    r.Cycle = &cycle
}

/* param monitorPath: 探测路径(Optional) */
func (r *SetMonitorRequest) SetMonitorPath(monitorPath string) {
    r.MonitorPath = &monitorPath
}

/* param httpRequestHeader: http请求头(Optional) */
func (r *SetMonitorRequest) SetHttpRequestHeader(httpRequestHeader interface{}) {
    r.HttpRequestHeader = &httpRequestHeader
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetMonitorRequest) GetRegionId() string {
    return ""
}

type SetMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetMonitorResult `json:"result"`
}

type SetMonitorResult struct {
    MonitorId int64 `json:"monitorId"`
}