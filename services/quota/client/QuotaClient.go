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
    "github.com/lshuining/jdcloud-sdk-go/core"
    quota "github.com/lshuining/jdcloud-sdk-go/services/quota/apis"
    "encoding/json"
    "errors"
)

type QuotaClient struct {
    core.JDCloudClient
}

func NewQuotaClient(credential *core.Credential) *QuotaClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("quota.jdcloud-api.com")

    return &QuotaClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "quota",
            Revision:    "0.0.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *QuotaClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *QuotaClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *QuotaClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 校验配额 */
func (c *QuotaClient) VerifyUserQuota(request *quota.VerifyUserQuotaRequest) (*quota.VerifyUserQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &quota.VerifyUserQuotaResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 校验配额  */
func (c *QuotaClient) GetUserQuotaDetail(request *quota.GetUserQuotaDetailRequest) (*quota.GetUserQuotaDetailResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &quota.GetUserQuotaDetailResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

