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
    jdccs "github.com/lshuining/jdcloud-sdk-go/services/jdccs/models"
)

type LastDownsampleRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 监控项英文标识(id)  */
    Metric string `json:"metric"`

    /* 资源ID，支持多个resourceId批量查询，每个id用英文竖线分隔  */
    ResourceId string `json:"resourceId"`
}

/*
 * param idc: IDC机房ID (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param resourceId: 资源ID，支持多个resourceId批量查询，每个id用英文竖线分隔 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewLastDownsampleRequest(
    idc string,
    metric string,
    resourceId string,
) *LastDownsampleRequest {

	return &LastDownsampleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/metrics/{metric}/lastDownsample",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
        Metric: metric,
        ResourceId: resourceId,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param resourceId: 资源ID，支持多个resourceId批量查询，每个id用英文竖线分隔 (Required)
 */
func NewLastDownsampleRequestWithAllParams(
    idc string,
    metric string,
    resourceId string,
) *LastDownsampleRequest {

    return &LastDownsampleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/metrics/{metric}/lastDownsample",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        Metric: metric,
        ResourceId: resourceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewLastDownsampleRequestWithoutParam() *LastDownsampleRequest {

    return &LastDownsampleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/metrics/{metric}/lastDownsample",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *LastDownsampleRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param metric: 监控项英文标识(id)(Required) */
func (r *LastDownsampleRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param resourceId: 资源ID，支持多个resourceId批量查询，每个id用英文竖线分隔(Required) */
func (r *LastDownsampleRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r LastDownsampleRequest) GetRegionId() string {
    return ""
}

type LastDownsampleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result LastDownsampleResult `json:"result"`
}

type LastDownsampleResult struct {
    Items []jdccs.LastDownsampleRespItem `json:"items"`
}