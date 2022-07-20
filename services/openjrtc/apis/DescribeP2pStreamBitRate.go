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
    openjrtc "github.com/lshuining/jdcloud-sdk-go/services/openjrtc/models"
)

type DescribeP2pStreamBitRateRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 业务接入方定义的且在JRTC系统内注册过的用户房间号  */
    UserRoomId string `json:"userRoomId"`

    /* 业务接入方定义的且在JRTC系统内注册过的用户id  */
    UserId string `json:"userId"`

    /* audio/video  */
    Kind string `json:"kind"`

    /* producer 发布流 consumer 订阅流  */
    Type string `json:"type"`

    /* 加入时间 UTC格式  */
    JoinTime string `json:"joinTime"`

    /* 离开时间 UTC格式 (Optional) */
    LeaveTime *string `json:"leaveTime"`

    /* 业务接入方定义的且在JRTC系统内注册过的用户id type=consumer时选择发送端用户id切换码率 (Optional) */
    FromUserId *string `json:"fromUserId"`

    /* 粒度 支持 1m 1h 1d (Optional) */
    Period *string `json:"period"`
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的用户房间号 (Required)
 * param userId: 业务接入方定义的且在JRTC系统内注册过的用户id (Required)
 * param kind: audio/video (Required)
 * param type_: producer 发布流 consumer 订阅流 (Required)
 * param joinTime: 加入时间 UTC格式 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeP2pStreamBitRateRequest(
    appId string,
    userRoomId string,
    userId string,
    kind string,
    type_ string,
    joinTime string,
) *DescribeP2pStreamBitRateRequest {

	return &DescribeP2pStreamBitRateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeP2pStreamBitRate",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        UserRoomId: userRoomId,
        UserId: userId,
        Kind: kind,
        Type: type_,
        JoinTime: joinTime,
	}
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的用户房间号 (Required)
 * param userId: 业务接入方定义的且在JRTC系统内注册过的用户id (Required)
 * param kind: audio/video (Required)
 * param type_: producer 发布流 consumer 订阅流 (Required)
 * param joinTime: 加入时间 UTC格式 (Required)
 * param leaveTime: 离开时间 UTC格式 (Optional)
 * param fromUserId: 业务接入方定义的且在JRTC系统内注册过的用户id type=consumer时选择发送端用户id切换码率 (Optional)
 * param period: 粒度 支持 1m 1h 1d (Optional)
 */
func NewDescribeP2pStreamBitRateRequestWithAllParams(
    appId string,
    userRoomId string,
    userId string,
    kind string,
    type_ string,
    joinTime string,
    leaveTime *string,
    fromUserId *string,
    period *string,
) *DescribeP2pStreamBitRateRequest {

    return &DescribeP2pStreamBitRateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeP2pStreamBitRate",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserRoomId: userRoomId,
        UserId: userId,
        Kind: kind,
        Type: type_,
        JoinTime: joinTime,
        LeaveTime: leaveTime,
        FromUserId: fromUserId,
        Period: period,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeP2pStreamBitRateRequestWithoutParam() *DescribeP2pStreamBitRateRequest {

    return &DescribeP2pStreamBitRateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeP2pStreamBitRate",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *DescribeP2pStreamBitRateRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param userRoomId: 业务接入方定义的且在JRTC系统内注册过的用户房间号(Required) */
func (r *DescribeP2pStreamBitRateRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = userRoomId
}

/* param userId: 业务接入方定义的且在JRTC系统内注册过的用户id(Required) */
func (r *DescribeP2pStreamBitRateRequest) SetUserId(userId string) {
    r.UserId = userId
}

/* param kind: audio/video(Required) */
func (r *DescribeP2pStreamBitRateRequest) SetKind(kind string) {
    r.Kind = kind
}

/* param type_: producer 发布流 consumer 订阅流(Required) */
func (r *DescribeP2pStreamBitRateRequest) SetType(type_ string) {
    r.Type = type_
}

/* param joinTime: 加入时间 UTC格式(Required) */
func (r *DescribeP2pStreamBitRateRequest) SetJoinTime(joinTime string) {
    r.JoinTime = joinTime
}

/* param leaveTime: 离开时间 UTC格式(Optional) */
func (r *DescribeP2pStreamBitRateRequest) SetLeaveTime(leaveTime string) {
    r.LeaveTime = &leaveTime
}

/* param fromUserId: 业务接入方定义的且在JRTC系统内注册过的用户id type=consumer时选择发送端用户id切换码率(Optional) */
func (r *DescribeP2pStreamBitRateRequest) SetFromUserId(fromUserId string) {
    r.FromUserId = &fromUserId
}

/* param period: 粒度 支持 1m 1h 1d(Optional) */
func (r *DescribeP2pStreamBitRateRequest) SetPeriod(period string) {
    r.Period = &period
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeP2pStreamBitRateRequest) GetRegionId() string {
    return ""
}

type DescribeP2pStreamBitRateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeP2pStreamBitRateResult `json:"result"`
}

type DescribeP2pStreamBitRateResult struct {
    Content []openjrtc.StreamBitRate `json:"content"`
}