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
    bri "github.com/lshuining/jdcloud-sdk-go/services/bri/apis"
    "encoding/json"
    "errors"
)

type BriClient struct {
    core.JDCloudClient
}

func NewBriClient(credential *core.Credential) *BriClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("bri.jdcloud-api.com")

    return &BriClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "bri",
            Revision:    "1.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *BriClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *BriClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *BriClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 设置黑白名单 */
func (c *BriClient) SetBWList(request *bri.SetBWListRequest) (*bri.SetBWListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &bri.SetBWListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取数据信用评分 */
func (c *BriClient) CreditScore(request *bri.CreditScoreRequest) (*bri.CreditScoreResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &bri.CreditScoreResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除黑白名单 */
func (c *BriClient) DelBWList(request *bri.DelBWListRequest) (*bri.DelBWListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &bri.DelBWListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取黑白名单列表 */
func (c *BriClient) DescribeBWList(request *bri.DescribeBWListRequest) (*bri.DescribeBWListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &bri.DescribeBWListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

