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
    jcq "github.com/lshuining/jdcloud-sdk-go/services/jcq/models"
)

type DescribeSubscriptionRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* consumerGroupId  */
    ConsumerGroupId string `json:"consumerGroupId"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param consumerGroupId: consumerGroupId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSubscriptionRequest(
    regionId string,
    topicName string,
    consumerGroupId string,
) *DescribeSubscriptionRequest {

	return &DescribeSubscriptionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicName: topicName,
        ConsumerGroupId: consumerGroupId,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param consumerGroupId: consumerGroupId (Required)
 */
func NewDescribeSubscriptionRequestWithAllParams(
    regionId string,
    topicName string,
    consumerGroupId string,
) *DescribeSubscriptionRequest {

    return &DescribeSubscriptionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicName: topicName,
        ConsumerGroupId: consumerGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSubscriptionRequestWithoutParam() *DescribeSubscriptionRequest {

    return &DescribeSubscriptionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *DescribeSubscriptionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicName: topic 名称(Required) */
func (r *DescribeSubscriptionRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param consumerGroupId: consumerGroupId(Required) */
func (r *DescribeSubscriptionRequest) SetConsumerGroupId(consumerGroupId string) {
    r.ConsumerGroupId = consumerGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubscriptionRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSubscriptionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubscriptionResult `json:"result"`
}

type DescribeSubscriptionResult struct {
    Subscription jcq.Subscription `json:"subscription"`
}