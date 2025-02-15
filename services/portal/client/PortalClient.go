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
    portal "github.com/lshuining/jdcloud-sdk-go/services/portal/apis"
    "encoding/json"
    "errors"
)

type PortalClient struct {
    core.JDCloudClient
}

func NewPortalClient(credential *core.Credential) *PortalClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("portal.jdcloud-api.com")

    return &PortalClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "portal",
            Revision:    "0.1.4",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *PortalClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *PortalClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询产品的详细信息
 */
func (c *PortalClient) DescribeProduct(request *portal.DescribeProductRequest) (*portal.DescribeProductResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &portal.DescribeProductResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 产品页列表查询接口
 */
func (c *PortalClient) DescribeProductsById(request *portal.DescribeProductsByIdRequest) (*portal.DescribeProductsByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &portal.DescribeProductsByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

