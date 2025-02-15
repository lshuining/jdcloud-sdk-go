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
    ydsms "github.com/lshuining/jdcloud-sdk-go/services/ydsms/models"
)

type SendMessagesByEncPinUsingPOSTRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 签名id  */
    SignId string `json:"signId"`

    /* 模板id  */
    TemplateId string `json:"templateId"`

    /* 加密pin集合，不能超过100个  */
    EncPins []string `json:"encPins"`

    /* 短信模板变量对应的数据值 (Optional) */
    Params []string `json:"params"`
}

/*
 * param appId: 应用id (Required)
 * param signId: 签名id (Required)
 * param templateId: 模板id (Required)
 * param encPins: 加密pin集合，不能超过100个 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendMessagesByEncPinUsingPOSTRequest(
    appId string,
    signId string,
    templateId string,
    encPins []string,
) *SendMessagesByEncPinUsingPOSTRequest {

	return &SendMessagesByEncPinUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sendMessagesByEncPin",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        SignId: signId,
        TemplateId: templateId,
        EncPins: encPins,
	}
}

/*
 * param appId: 应用id (Required)
 * param signId: 签名id (Required)
 * param templateId: 模板id (Required)
 * param encPins: 加密pin集合，不能超过100个 (Required)
 * param params: 短信模板变量对应的数据值 (Optional)
 */
func NewSendMessagesByEncPinUsingPOSTRequestWithAllParams(
    appId string,
    signId string,
    templateId string,
    encPins []string,
    params []string,
) *SendMessagesByEncPinUsingPOSTRequest {

    return &SendMessagesByEncPinUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sendMessagesByEncPin",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        SignId: signId,
        TemplateId: templateId,
        EncPins: encPins,
        Params: params,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendMessagesByEncPinUsingPOSTRequestWithoutParam() *SendMessagesByEncPinUsingPOSTRequest {

    return &SendMessagesByEncPinUsingPOSTRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sendMessagesByEncPin",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *SendMessagesByEncPinUsingPOSTRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param signId: 签名id(Required) */
func (r *SendMessagesByEncPinUsingPOSTRequest) SetSignId(signId string) {
    r.SignId = signId
}

/* param templateId: 模板id(Required) */
func (r *SendMessagesByEncPinUsingPOSTRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

/* param encPins: 加密pin集合，不能超过100个(Required) */
func (r *SendMessagesByEncPinUsingPOSTRequest) SetEncPins(encPins []string) {
    r.EncPins = encPins
}

/* param params: 短信模板变量对应的数据值(Optional) */
func (r *SendMessagesByEncPinUsingPOSTRequest) SetParams(params []string) {
    r.Params = params
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendMessagesByEncPinUsingPOSTRequest) GetRegionId() string {
    return ""
}

type SendMessagesByEncPinUsingPOSTResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendMessagesByEncPinUsingPOSTResult `json:"result"`
}

type SendMessagesByEncPinUsingPOSTResult struct {
    MtResVO ydsms.MtResVO `json:"mtResVO"`
}