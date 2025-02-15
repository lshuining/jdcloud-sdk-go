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

type CertificateTemplateRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 模板ID  */
    TemplateId int `json:"templateId"`

    /* 所有人证件号码  */
    IdentityNo string `json:"identityNo"`

    /* 注册人证件类型
1.个人
  (1)身份证 SFZ
2.企业
  (1)组织机构代码证 ORG
  (2)工商营业执照 YYZZ
  (3)统一社会信用代码证书 TYDMZ
  (4)部队代号 BDDH
  (5)军队单位对外有偿服务许可证 JDXKZ
  (6)事业单位法人证书 SYZS
  (7)社会团体法人登记证书 STDJZ
  (8)宗教活动场所登记证 ZJDJZ
  (9)民办非企业单位登记证书 MBDJZ
  (10)基金会法人登记证书 JJDJZ
  (11)律师事务所执业许可证 LSXKZ
  (12)登记证 GWLYDJZ
  (13)司法鉴定许可证 SFXKZ
  (14)社会服务机构登记证书 SHFWJGZ
  (15)民办学校办学许可证 MBXXXKZ
  (16)医疗机构执业许可证 YLJGXKZ
  */
    IdentityType string `json:"identityType"`

    /* 所有人证件，jpg 图片的 base64 编码，必填（大小 55KB~1MB）  */
    File string `json:"file"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param templateId: 模板ID (Required)
 * param identityNo: 所有人证件号码 (Required)
 * param identityType: 注册人证件类型
1.个人
  (1)身份证 SFZ
2.企业
  (1)组织机构代码证 ORG
  (2)工商营业执照 YYZZ
  (3)统一社会信用代码证书 TYDMZ
  (4)部队代号 BDDH
  (5)军队单位对外有偿服务许可证 JDXKZ
  (6)事业单位法人证书 SYZS
  (7)社会团体法人登记证书 STDJZ
  (8)宗教活动场所登记证 ZJDJZ
  (9)民办非企业单位登记证书 MBDJZ
  (10)基金会法人登记证书 JJDJZ
  (11)律师事务所执业许可证 LSXKZ
  (12)登记证 GWLYDJZ
  (13)司法鉴定许可证 SFXKZ
  (14)社会服务机构登记证书 SHFWJGZ
  (15)民办学校办学许可证 MBXXXKZ
  (16)医疗机构执业许可证 YLJGXKZ
 (Required)
 * param file: 所有人证件，jpg 图片的 base64 编码，必填（大小 55KB~1MB） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCertificateTemplateRequest(
    regionId string,
    templateId int,
    identityNo string,
    identityType string,
    file string,
) *CertificateTemplateRequest {

	return &CertificateTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/template/{templateId}/certificate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TemplateId: templateId,
        IdentityNo: identityNo,
        IdentityType: identityType,
        File: file,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param templateId: 模板ID (Required)
 * param identityNo: 所有人证件号码 (Required)
 * param identityType: 注册人证件类型
1.个人
  (1)身份证 SFZ
2.企业
  (1)组织机构代码证 ORG
  (2)工商营业执照 YYZZ
  (3)统一社会信用代码证书 TYDMZ
  (4)部队代号 BDDH
  (5)军队单位对外有偿服务许可证 JDXKZ
  (6)事业单位法人证书 SYZS
  (7)社会团体法人登记证书 STDJZ
  (8)宗教活动场所登记证 ZJDJZ
  (9)民办非企业单位登记证书 MBDJZ
  (10)基金会法人登记证书 JJDJZ
  (11)律师事务所执业许可证 LSXKZ
  (12)登记证 GWLYDJZ
  (13)司法鉴定许可证 SFXKZ
  (14)社会服务机构登记证书 SHFWJGZ
  (15)民办学校办学许可证 MBXXXKZ
  (16)医疗机构执业许可证 YLJGXKZ
 (Required)
 * param file: 所有人证件，jpg 图片的 base64 编码，必填（大小 55KB~1MB） (Required)
 */
func NewCertificateTemplateRequestWithAllParams(
    regionId string,
    templateId int,
    identityNo string,
    identityType string,
    file string,
) *CertificateTemplateRequest {

    return &CertificateTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/template/{templateId}/certificate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TemplateId: templateId,
        IdentityNo: identityNo,
        IdentityType: identityType,
        File: file,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCertificateTemplateRequestWithoutParam() *CertificateTemplateRequest {

    return &CertificateTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/template/{templateId}/certificate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *CertificateTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param templateId: 模板ID(Required) */
func (r *CertificateTemplateRequest) SetTemplateId(templateId int) {
    r.TemplateId = templateId
}

/* param identityNo: 所有人证件号码(Required) */
func (r *CertificateTemplateRequest) SetIdentityNo(identityNo string) {
    r.IdentityNo = identityNo
}

/* param identityType: 注册人证件类型
1.个人
  (1)身份证 SFZ
2.企业
  (1)组织机构代码证 ORG
  (2)工商营业执照 YYZZ
  (3)统一社会信用代码证书 TYDMZ
  (4)部队代号 BDDH
  (5)军队单位对外有偿服务许可证 JDXKZ
  (6)事业单位法人证书 SYZS
  (7)社会团体法人登记证书 STDJZ
  (8)宗教活动场所登记证 ZJDJZ
  (9)民办非企业单位登记证书 MBDJZ
  (10)基金会法人登记证书 JJDJZ
  (11)律师事务所执业许可证 LSXKZ
  (12)登记证 GWLYDJZ
  (13)司法鉴定许可证 SFXKZ
  (14)社会服务机构登记证书 SHFWJGZ
  (15)民办学校办学许可证 MBXXXKZ
  (16)医疗机构执业许可证 YLJGXKZ
(Required) */
func (r *CertificateTemplateRequest) SetIdentityType(identityType string) {
    r.IdentityType = identityType
}

/* param file: 所有人证件，jpg 图片的 base64 编码，必填（大小 55KB~1MB）(Required) */
func (r *CertificateTemplateRequest) SetFile(file string) {
    r.File = file
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CertificateTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type CertificateTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CertificateTemplateResult `json:"result"`
}

type CertificateTemplateResult struct {
    TemplateId int64 `json:"templateId"`
}