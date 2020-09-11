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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    cloudauth "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudauth/apis"
    "encoding/json"
    "errors"
)

type CloudauthClient struct {
    core.JDCloudClient
}

func NewCloudauthClient(credential *core.Credential) *CloudauthClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("cloudauth.jdcloud-api.com")

    return &CloudauthClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "cloudauth",
            Revision:    "1.0.3",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *CloudauthClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *CloudauthClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *CloudauthClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询城市下银行分行列表 */
func (c *CloudauthClient) QueryBankBranchList(request *cloudauth.QueryBankBranchListRequest) (*cloudauth.QueryBankBranchListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.QueryBankBranchListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询服务开通状态 */
func (c *CloudauthClient) DescribeApplyStatus(request *cloudauth.DescribeApplyStatusRequest) (*cloudauth.DescribeApplyStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.DescribeApplyStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 法人信息校验 */
func (c *CloudauthClient) CheckLegalPerson(request *cloudauth.CheckLegalPersonRequest) (*cloudauth.CheckLegalPersonResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.CheckLegalPersonResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 个人实名认证 */
func (c *CloudauthClient) PersonalAuth(request *cloudauth.PersonalAuthRequest) (*cloudauth.PersonalAuthResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.PersonalAuthResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询所有支持的银行 */
func (c *CloudauthClient) QueryBankList(request *cloudauth.QueryBankListRequest) (*cloudauth.QueryBankListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.QueryBankListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 对公银行账户打款(随机小额) */
func (c *CloudauthClient) CompanyTransfer(request *cloudauth.CompanyTransferRequest) (*cloudauth.CompanyTransferResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.CompanyTransferResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 代理人信息核验 */
func (c *CloudauthClient) CheckAgent(request *cloudauth.CheckAgentRequest) (*cloudauth.CheckAgentResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.CheckAgentResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询省份下城市编码 */
func (c *CloudauthClient) QueryCityList(request *cloudauth.QueryCityListRequest) (*cloudauth.QueryCityListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.QueryCityListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 对公打款查询 */
func (c *CloudauthClient) CheckCompanyTransfer(request *cloudauth.CheckCompanyTransferRequest) (*cloudauth.CheckCompanyTransferResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.CheckCompanyTransferResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 企业基础信息核验 */
func (c *CloudauthClient) CheckCompanyInfo(request *cloudauth.CheckCompanyInfoRequest) (*cloudauth.CheckCompanyInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.CheckCompanyInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 企业所有信息核验 */
func (c *CloudauthClient) CheckLegalPersonAndAgent(request *cloudauth.CheckLegalPersonAndAgentRequest) (*cloudauth.CheckLegalPersonAndAgentResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.CheckLegalPersonAndAgentResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询所有省份列表 */
func (c *CloudauthClient) QueryProvinceList(request *cloudauth.QueryProvinceListRequest) (*cloudauth.QueryProvinceListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &cloudauth.QueryProvinceListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

