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

type CheckPodNameRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户定义的 pod 名称，符合 DNS-1123 subdomain 规范。  */
    PodName string `json:"podName"`

    /* 需要创建的 pod 总数，默认创建一个，不同的总数会对校验结果产生影响。 (Optional) */
    MaxCount *int `json:"maxCount"`
}

/*
 * param regionId: Region ID (Required)
 * param podName: 用户定义的 pod 名称，符合 DNS-1123 subdomain 规范。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckPodNameRequest(
    regionId string,
    podName string,
) *CheckPodNameRequest {

	return &CheckPodNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods:checkPodName",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PodName: podName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podName: 用户定义的 pod 名称，符合 DNS-1123 subdomain 规范。 (Required)
 * param maxCount: 需要创建的 pod 总数，默认创建一个，不同的总数会对校验结果产生影响。 (Optional)
 */
func NewCheckPodNameRequestWithAllParams(
    regionId string,
    podName string,
    maxCount *int,
) *CheckPodNameRequest {

    return &CheckPodNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods:checkPodName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodName: podName,
        MaxCount: maxCount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckPodNameRequestWithoutParam() *CheckPodNameRequest {

    return &CheckPodNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods:checkPodName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CheckPodNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param podName: 用户定义的 pod 名称，符合 DNS-1123 subdomain 规范。(Required) */
func (r *CheckPodNameRequest) SetPodName(podName string) {
    r.PodName = podName
}

/* param maxCount: 需要创建的 pod 总数，默认创建一个，不同的总数会对校验结果产生影响。(Optional) */
func (r *CheckPodNameRequest) SetMaxCount(maxCount int) {
    r.MaxCount = &maxCount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckPodNameRequest) GetRegionId() string {
    return r.RegionId
}

type CheckPodNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckPodNameResult `json:"result"`
}

type CheckPodNameResult struct {
    Code int `json:"code"`
    Reason string `json:"reason"`
}