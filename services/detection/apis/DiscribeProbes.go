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
    detection "github.com/lshuining/jdcloud-sdk-go/services/detection/models"
)

type DiscribeProbesRequest struct {

    core.JDCloudRequest

    /* 探测任务的task_id  */
    ProbeTaskID string `json:"probeTaskID"`

    /* 自定义标签 (Optional) */
    Filters []detection.Filter `json:"filters"`
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDiscribeProbesRequest(
    probeTaskID string,
) *DiscribeProbesRequest {

	return &DiscribeProbesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/probeTask/{probeTaskID}/probeList",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        ProbeTaskID: probeTaskID,
	}
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 * param filters: 自定义标签 (Optional)
 */
func NewDiscribeProbesRequestWithAllParams(
    probeTaskID string,
    filters []detection.Filter,
) *DiscribeProbesRequest {

    return &DiscribeProbesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask/{probeTaskID}/probeList",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        ProbeTaskID: probeTaskID,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDiscribeProbesRequestWithoutParam() *DiscribeProbesRequest {

    return &DiscribeProbesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask/{probeTaskID}/probeList",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param probeTaskID: 探测任务的task_id(Required) */
func (r *DiscribeProbesRequest) SetProbeTaskID(probeTaskID string) {
    r.ProbeTaskID = probeTaskID
}

/* param filters: 自定义标签(Optional) */
func (r *DiscribeProbesRequest) SetFilters(filters []detection.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DiscribeProbesRequest) GetRegionId() string {
    return ""
}

type DiscribeProbesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DiscribeProbesResult `json:"result"`
}

type DiscribeProbesResult struct {
    ProbeList []detection.ProbeInfo `json:"probeList"`
}