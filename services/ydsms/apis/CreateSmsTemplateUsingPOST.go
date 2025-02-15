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

type CreateSmsTemplateUsingPOSTRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 申请说明 (Optional) */
    ApplyExplanation *string `json:"applyExplanation"`

    /* 模板内容  */
    TemplateContent string `json:"templateContent"`

    /* 模板名称  */
    TemplateName string `json:"templateName"`

    /* 模板类型:0 验证码短信,1 通知短信,2 推广短信  */
    TemplateType int `json:"templateType"`
}

/*
 * param appId: 应用id (Required)
 * param templateContent: 模板内容 (Required)
 * param templateName: 模板名称 (Required)
 * param templateType: 模板类型:0 验证码短信,1 通知短信,2 推广短信 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSmsTemplateUsingPOSTRequest(
    appId string,
    templateContent string,
    templateName string,
    templateType int,
) *CreateSmsTemplateUsingPOSTRequest {

	return &CreateSmsTemplateUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsTemplates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        TemplateContent: templateContent,
        TemplateName: templateName,
        TemplateType: templateType,
	}
}

/*
 * param appId: 应用id (Required)
 * param applyExplanation: 申请说明 (Optional)
 * param templateContent: 模板内容 (Required)
 * param templateName: 模板名称 (Required)
 * param templateType: 模板类型:0 验证码短信,1 通知短信,2 推广短信 (Required)
 */
func NewCreateSmsTemplateUsingPOSTRequestWithAllParams(
    appId string,
    applyExplanation *string,
    templateContent string,
    templateName string,
    templateType int,
) *CreateSmsTemplateUsingPOSTRequest {

    return &CreateSmsTemplateUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsTemplates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        ApplyExplanation: applyExplanation,
        TemplateContent: templateContent,
        TemplateName: templateName,
        TemplateType: templateType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSmsTemplateUsingPOSTRequestWithoutParam() *CreateSmsTemplateUsingPOSTRequest {

    return &CreateSmsTemplateUsingPOSTRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsTemplates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *CreateSmsTemplateUsingPOSTRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param applyExplanation: 申请说明(Optional) */
func (r *CreateSmsTemplateUsingPOSTRequest) SetApplyExplanation(applyExplanation string) {
    r.ApplyExplanation = &applyExplanation
}

/* param templateContent: 模板内容(Required) */
func (r *CreateSmsTemplateUsingPOSTRequest) SetTemplateContent(templateContent string) {
    r.TemplateContent = templateContent
}

/* param templateName: 模板名称(Required) */
func (r *CreateSmsTemplateUsingPOSTRequest) SetTemplateName(templateName string) {
    r.TemplateName = templateName
}

/* param templateType: 模板类型:0 验证码短信,1 通知短信,2 推广短信(Required) */
func (r *CreateSmsTemplateUsingPOSTRequest) SetTemplateType(templateType int) {
    r.TemplateType = templateType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSmsTemplateUsingPOSTRequest) GetRegionId() string {
    return ""
}

type CreateSmsTemplateUsingPOSTResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSmsTemplateUsingPOSTResult `json:"result"`
}

type CreateSmsTemplateUsingPOSTResult struct {
    TemplateId string `json:"templateId"`
}