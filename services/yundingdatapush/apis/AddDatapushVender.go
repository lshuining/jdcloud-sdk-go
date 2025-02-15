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
    yundingdatapush "github.com/lshuining/jdcloud-sdk-go/services/yundingdatapush/models"
)

type AddDatapushVenderRequest struct {

    core.JDCloudRequest

    /* 添加数据推送用户对象
  */
    DatapushVender *yundingdatapush.Vender `json:"datapushVender"`
}

/*
 * param datapushVender: 添加数据推送用户对象
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddDatapushVenderRequest(
    datapushVender *yundingdatapush.Vender,
) *AddDatapushVenderRequest {

	return &AddDatapushVenderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/datapushVenders",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DatapushVender: datapushVender,
	}
}

/*
 * param datapushVender: 添加数据推送用户对象
 (Required)
 */
func NewAddDatapushVenderRequestWithAllParams(
    datapushVender *yundingdatapush.Vender,
) *AddDatapushVenderRequest {

    return &AddDatapushVenderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/datapushVenders",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        DatapushVender: datapushVender,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddDatapushVenderRequestWithoutParam() *AddDatapushVenderRequest {

    return &AddDatapushVenderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/datapushVenders",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param datapushVender: 添加数据推送用户对象
(Required) */
func (r *AddDatapushVenderRequest) SetDatapushVender(datapushVender *yundingdatapush.Vender) {
    r.DatapushVender = datapushVender
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddDatapushVenderRequest) GetRegionId() string {
    return ""
}

type AddDatapushVenderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddDatapushVenderResult `json:"result"`
}

type AddDatapushVenderResult struct {
    Success bool `json:"success"`
}