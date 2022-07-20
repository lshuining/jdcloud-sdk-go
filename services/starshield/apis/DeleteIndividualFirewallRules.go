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

type DeleteIndividualFirewallRulesRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /*   */
    Id string `json:"id"`

    /*  (Optional) */
    Delete_filter_if_unused *bool `json:"delete_filter_if_unused"`
}

/*
 * param zone_identifier:  (Required)
 * param id:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteIndividualFirewallRulesRequest(
    zone_identifier string,
    id string,
) *DeleteIndividualFirewallRulesRequest {

	return &DeleteIndividualFirewallRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/firewall$$rules/{id}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
        Id: id,
	}
}

/*
 * param zone_identifier:  (Required)
 * param id:  (Required)
 * param delete_filter_if_unused:  (Optional)
 */
func NewDeleteIndividualFirewallRulesRequestWithAllParams(
    zone_identifier string,
    id string,
    delete_filter_if_unused *bool,
) *DeleteIndividualFirewallRulesRequest {

    return &DeleteIndividualFirewallRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/firewall$$rules/{id}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Id: id,
        Delete_filter_if_unused: delete_filter_if_unused,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteIndividualFirewallRulesRequestWithoutParam() *DeleteIndividualFirewallRulesRequest {

    return &DeleteIndividualFirewallRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/firewall$$rules/{id}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *DeleteIndividualFirewallRulesRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param id: (Required) */
func (r *DeleteIndividualFirewallRulesRequest) SetId(id string) {
    r.Id = id
}

/* param delete_filter_if_unused: (Optional) */
func (r *DeleteIndividualFirewallRulesRequest) SetDelete_filter_if_unused(delete_filter_if_unused bool) {
    r.Delete_filter_if_unused = &delete_filter_if_unused
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteIndividualFirewallRulesRequest) GetRegionId() string {
    return ""
}

type DeleteIndividualFirewallRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteIndividualFirewallRulesResult `json:"result"`
}

type DeleteIndividualFirewallRulesResult struct {
    Data starshield.FirewallRule `json:"data"`
}