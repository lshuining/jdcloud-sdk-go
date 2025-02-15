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
    yundingdatapush "github.com/lshuining/jdcloud-sdk-go/services/yundingdatapush/apis"
    "encoding/json"
    "errors"
)

type YundingdatapushClient struct {
    core.JDCloudClient
}

func NewYundingdatapushClient(credential *core.Credential) *YundingdatapushClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("yundingdatapush.jdcloud-api.com")

    return &YundingdatapushClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "yundingdatapush",
            Revision:    "1.0.4",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *YundingdatapushClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *YundingdatapushClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *YundingdatapushClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 添加数据推送用户 */
func (c *YundingdatapushClient) AddDatapushVender(request *yundingdatapush.AddDatapushVenderRequest) (*yundingdatapush.AddDatapushVenderResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yundingdatapush.AddDatapushVenderResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建历史订单同步 */
func (c *YundingdatapushClient) CreateOrderSync(request *yundingdatapush.CreateOrderSyncRequest) (*yundingdatapush.CreateOrderSyncResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yundingdatapush.CreateOrderSyncResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询已经开通的用户 */
func (c *YundingdatapushClient) DescribeDatapushVenders(request *yundingdatapush.DescribeDatapushVendersRequest) (*yundingdatapush.DescribeDatapushVendersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yundingdatapush.DescribeDatapushVendersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询已绑定数据推送的数据库实例 */
func (c *YundingdatapushClient) DescribeRdsInstances(request *yundingdatapush.DescribeRdsInstancesRequest) (*yundingdatapush.DescribeRdsInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yundingdatapush.DescribeRdsInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除数据推送用户 */
func (c *YundingdatapushClient) DeleteDatapushVender(request *yundingdatapush.DeleteDatapushVenderRequest) (*yundingdatapush.DeleteDatapushVenderResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &yundingdatapush.DeleteDatapushVenderResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

