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

type DeleteLiveStreamDomainTranslateRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 翻译模板  */
    Template string `json:"template"`
}

/*
 * param publishDomain: 推流域名 (Required)
 * param template: 翻译模板 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveStreamDomainTranslateRequest(
    publishDomain string,
    template string,
) *DeleteLiveStreamDomainTranslateRequest {

	return &DeleteLiveStreamDomainTranslateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/translateDomains/{publishDomain}/templates/{template}:config",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        Template: template,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 * param template: 翻译模板 (Required)
 */
func NewDeleteLiveStreamDomainTranslateRequestWithAllParams(
    publishDomain string,
    template string,
) *DeleteLiveStreamDomainTranslateRequest {

    return &DeleteLiveStreamDomainTranslateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/translateDomains/{publishDomain}/templates/{template}:config",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveStreamDomainTranslateRequestWithoutParam() *DeleteLiveStreamDomainTranslateRequest {

    return &DeleteLiveStreamDomainTranslateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/translateDomains/{publishDomain}/templates/{template}:config",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *DeleteLiveStreamDomainTranslateRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param template: 翻译模板(Required) */
func (r *DeleteLiveStreamDomainTranslateRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveStreamDomainTranslateRequest) GetRegionId() string {
    return ""
}

type DeleteLiveStreamDomainTranslateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveStreamDomainTranslateResult `json:"result"`
}

type DeleteLiveStreamDomainTranslateResult struct {
}