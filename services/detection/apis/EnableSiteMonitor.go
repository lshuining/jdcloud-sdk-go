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

type EnableSiteMonitorRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    List []detection.EnableSiteMonitorReqItem `json:"list"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableSiteMonitorRequest(
) *EnableSiteMonitorRequest {

	return &EnableSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/siteMonitor:switch",
			Method:  "PUT",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param list:  (Optional)
 */
func NewEnableSiteMonitorRequestWithAllParams(
    list []detection.EnableSiteMonitorReqItem,
) *EnableSiteMonitorRequest {

    return &EnableSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor:switch",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
        List: list,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableSiteMonitorRequestWithoutParam() *EnableSiteMonitorRequest {

    return &EnableSiteMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor:switch",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param list: (Optional) */
func (r *EnableSiteMonitorRequest) SetList(list []detection.EnableSiteMonitorReqItem) {
    r.List = list
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableSiteMonitorRequest) GetRegionId() string {
    return ""
}

type EnableSiteMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableSiteMonitorResult `json:"result"`
}

type EnableSiteMonitorResult struct {
    Suc bool `json:"suc"`
}