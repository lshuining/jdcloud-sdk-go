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
    vod "github.com/lshuining/jdcloud-sdk-go/services/vod/models"
)

type ListSnapshotTasksRequest struct {

    core.JDCloudRequest

    /* 页码；默认值为 1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认值为 10；取值范围 [10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Filters []vod.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListSnapshotTasksRequest(
) *ListSnapshotTasksRequest {

	return &ListSnapshotTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshotTasks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认值为 1 (Optional)
 * param pageSize: 分页大小；默认值为 10；取值范围 [10, 100] (Optional)
 * param filters:  (Optional)
 */
func NewListSnapshotTasksRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    filters []vod.Filter,
) *ListSnapshotTasksRequest {

    return &ListSnapshotTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotTasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListSnapshotTasksRequestWithoutParam() *ListSnapshotTasksRequest {

    return &ListSnapshotTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotTasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认值为 1(Optional) */
func (r *ListSnapshotTasksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认值为 10；取值范围 [10, 100](Optional) */
func (r *ListSnapshotTasksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: (Optional) */
func (r *ListSnapshotTasksRequest) SetFilters(filters []vod.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListSnapshotTasksRequest) GetRegionId() string {
    return ""
}

type ListSnapshotTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListSnapshotTasksResult `json:"result"`
}

type ListSnapshotTasksResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalElements int `json:"totalElements"`
    TotalPages int `json:"totalPages"`
    Content []vod.SnapshotTaskSummary `json:"content"`
}