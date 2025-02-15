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
    sts "github.com/lshuining/jdcloud-sdk-go/services/sts/apis"
    "encoding/json"
    "errors"
)

type StsClient struct {
    core.JDCloudClient
}

func NewStsClient(credential *core.Credential) *StsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("sts.jdcloud-api.com")

    return &StsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "sts",
            Revision:    "0.2.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *StsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *StsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *StsClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 扮演用户角色，获取临时凭证 */
func (c *StsClient) AssumeRole(request *sts.AssumeRoleRequest) (*sts.AssumeRoleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &sts.AssumeRoleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

