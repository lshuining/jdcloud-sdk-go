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

type DescribeTicketRequest struct {

    core.JDCloudRequest

    /* 工单编号  */
    TicketNo string `json:"ticketNo"`
}

/*
 * param ticketNo: 工单编号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTicketRequest(
    ticketNo string,
) *DescribeTicketRequest {

	return &DescribeTicketRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/tickets/{ticketNo}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TicketNo: ticketNo,
	}
}

/*
 * param ticketNo: 工单编号 (Required)
 */
func NewDescribeTicketRequestWithAllParams(
    ticketNo string,
) *DescribeTicketRequest {

    return &DescribeTicketRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/tickets/{ticketNo}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TicketNo: ticketNo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTicketRequestWithoutParam() *DescribeTicketRequest {

    return &DescribeTicketRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/tickets/{ticketNo}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param ticketNo: 工单编号(Required) */
func (r *DescribeTicketRequest) SetTicketNo(ticketNo string) {
    r.TicketNo = ticketNo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTicketRequest) GetRegionId() string {
    return ""
}

type DescribeTicketResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTicketResult `json:"result"`
}

type DescribeTicketResult struct {
    Ticket jdccs.Ticket `json:"ticket"`
}