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

type CreateSmsPackageUsingPOSTRequest struct {

    core.JDCloudRequest

    /* 套餐包数量  */
    PackageCount int `json:"packageCount"`

    /* 套餐包类型，1通道短信，2官方短信  */
    PackageType int `json:"packageType"`

    /* 套餐包规格  */
    Specification string `json:"specification"`
}

/*
 * param packageCount: 套餐包数量 (Required)
 * param packageType: 套餐包类型，1通道短信，2官方短信 (Required)
 * param specification: 套餐包规格 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSmsPackageUsingPOSTRequest(
    packageCount int,
    packageType int,
    specification string,
) *CreateSmsPackageUsingPOSTRequest {

	return &CreateSmsPackageUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsPackages",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        PackageCount: packageCount,
        PackageType: packageType,
        Specification: specification,
	}
}

/*
 * param packageCount: 套餐包数量 (Required)
 * param packageType: 套餐包类型，1通道短信，2官方短信 (Required)
 * param specification: 套餐包规格 (Required)
 */
func NewCreateSmsPackageUsingPOSTRequestWithAllParams(
    packageCount int,
    packageType int,
    specification string,
) *CreateSmsPackageUsingPOSTRequest {

    return &CreateSmsPackageUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsPackages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PackageCount: packageCount,
        PackageType: packageType,
        Specification: specification,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSmsPackageUsingPOSTRequestWithoutParam() *CreateSmsPackageUsingPOSTRequest {

    return &CreateSmsPackageUsingPOSTRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsPackages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param packageCount: 套餐包数量(Required) */
func (r *CreateSmsPackageUsingPOSTRequest) SetPackageCount(packageCount int) {
    r.PackageCount = packageCount
}

/* param packageType: 套餐包类型，1通道短信，2官方短信(Required) */
func (r *CreateSmsPackageUsingPOSTRequest) SetPackageType(packageType int) {
    r.PackageType = packageType
}

/* param specification: 套餐包规格(Required) */
func (r *CreateSmsPackageUsingPOSTRequest) SetSpecification(specification string) {
    r.Specification = specification
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSmsPackageUsingPOSTRequest) GetRegionId() string {
    return ""
}

type CreateSmsPackageUsingPOSTResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSmsPackageUsingPOSTResult `json:"result"`
}

type CreateSmsPackageUsingPOSTResult struct {
    BuyId string `json:"buyId"`
}