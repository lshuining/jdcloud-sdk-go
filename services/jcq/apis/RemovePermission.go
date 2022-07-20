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

type RemovePermissionRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* 权限类型, [PUB, SUB, PUBSUB]  */
    Permission string `json:"permission"`

    /* 目标用户UserId  */
    TargetUserId string `json:"targetUserId"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param permission: 权限类型, [PUB, SUB, PUBSUB] (Required)
 * param targetUserId: 目标用户UserId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemovePermissionRequest(
    regionId string,
    topicName string,
    permission string,
    targetUserId string,
) *RemovePermissionRequest {

	return &RemovePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics/{topicName}/iam",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicName: topicName,
        Permission: permission,
        TargetUserId: targetUserId,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param permission: 权限类型, [PUB, SUB, PUBSUB] (Required)
 * param targetUserId: 目标用户UserId (Required)
 */
func NewRemovePermissionRequestWithAllParams(
    regionId string,
    topicName string,
    permission string,
    targetUserId string,
) *RemovePermissionRequest {

    return &RemovePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/iam",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicName: topicName,
        Permission: permission,
        TargetUserId: targetUserId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemovePermissionRequestWithoutParam() *RemovePermissionRequest {

    return &RemovePermissionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/iam",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *RemovePermissionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicName: topic 名称(Required) */
func (r *RemovePermissionRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param permission: 权限类型, [PUB, SUB, PUBSUB](Required) */
func (r *RemovePermissionRequest) SetPermission(permission string) {
    r.Permission = permission
}

/* param targetUserId: 目标用户UserId(Required) */
func (r *RemovePermissionRequest) SetTargetUserId(targetUserId string) {
    r.TargetUserId = targetUserId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemovePermissionRequest) GetRegionId() string {
    return r.RegionId
}

type RemovePermissionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemovePermissionResult `json:"result"`
}

type RemovePermissionResult struct {
}