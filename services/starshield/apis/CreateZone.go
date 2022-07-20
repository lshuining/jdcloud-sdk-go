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
    starshield "github.com/lshuining/jdcloud-sdk-go/services/starshield/models"
)

type CreateZoneRequest struct {

    core.JDCloudRequest

    /* 域名  */
    Name string `json:"name"`

    /*   */
    Account *starshield.Account `json:"account"`

    /* 自动尝试获取现有DNS记录 (Optional) */
    Jump_start *bool `json:"jump_start"`

    /* 全接入域意味着DNS由星盾托管。半接入域通常是合作伙伴托管的域或CNAME设置。 (Optional) */
    Ty_pe *string `json:"ty_pe"`
}

/*
 * param name: 域名 (Required)
 * param account:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateZoneRequest(
    name string,
    account *starshield.Account,
) *CreateZoneRequest {

	return &CreateZoneRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Name: name,
        Account: account,
	}
}

/*
 * param name: 域名 (Required)
 * param account:  (Required)
 * param jump_start: 自动尝试获取现有DNS记录 (Optional)
 * param ty_pe: 全接入域意味着DNS由星盾托管。半接入域通常是合作伙伴托管的域或CNAME设置。 (Optional)
 */
func NewCreateZoneRequestWithAllParams(
    name string,
    account *starshield.Account,
    jump_start *bool,
    ty_pe *string,
) *CreateZoneRequest {

    return &CreateZoneRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        Account: account,
        Jump_start: jump_start,
        Ty_pe: ty_pe,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateZoneRequestWithoutParam() *CreateZoneRequest {

    return &CreateZoneRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: 域名(Required) */
func (r *CreateZoneRequest) SetName(name string) {
    r.Name = name
}

/* param account: (Required) */
func (r *CreateZoneRequest) SetAccount(account *starshield.Account) {
    r.Account = account
}

/* param jump_start: 自动尝试获取现有DNS记录(Optional) */
func (r *CreateZoneRequest) SetJump_start(jump_start bool) {
    r.Jump_start = &jump_start
}

/* param ty_pe: 全接入域意味着DNS由星盾托管。半接入域通常是合作伙伴托管的域或CNAME设置。(Optional) */
func (r *CreateZoneRequest) SetTy_pe(ty_pe string) {
    r.Ty_pe = &ty_pe
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateZoneRequest) GetRegionId() string {
    return ""
}

type CreateZoneResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateZoneResult `json:"result"`
}

type CreateZoneResult struct {
    Data starshield.Zone `json:"data"`
}