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
    jdro "github.com/lshuining/jdcloud-sdk-go/services/jdro/apis"
    "encoding/json"
    "errors"
)

type JdroClient struct {
    core.JDCloudClient
}

func NewJdroClient(credential *core.Credential) *JdroClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("jdro.jdcloud-api.com")

    return &JdroClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "jdro",
            Revision:    "0.0.4",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *JdroClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *JdroClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询资源栈列表 */
func (c *JdroClient) DescribeStacks(request *jdro.DescribeStacksRequest) (*jdro.DescribeStacksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeStacksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询资源栈事件列表 */
func (c *JdroClient) DescribeStackEvents(request *jdro.DescribeStackEventsRequest) (*jdro.DescribeStackEventsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeStackEventsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询资源栈中资源列表 */
func (c *JdroClient) DescribeStackResources(request *jdro.DescribeStackResourcesRequest) (*jdro.DescribeStackResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeStackResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建更改集 */
func (c *JdroClient) CreateChangeSet(request *jdro.CreateChangeSetRequest) (*jdro.CreateChangeSetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.CreateChangeSetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建资源栈 */
func (c *JdroClient) CreateStack(request *jdro.CreateStackRequest) (*jdro.CreateStackResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.CreateStackResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询资源栈详情 */
func (c *JdroClient) DescribeStack(request *jdro.DescribeStackRequest) (*jdro.DescribeStackResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeStackResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询资源栈使用的模板 */
func (c *JdroClient) DescribeStackTemplate(request *jdro.DescribeStackTemplateRequest) (*jdro.DescribeStackTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeStackTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 执行更改集 */
func (c *JdroClient) ExecuteChangeSet(request *jdro.ExecuteChangeSetRequest) (*jdro.ExecuteChangeSetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.ExecuteChangeSetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除资源栈 */
func (c *JdroClient) DeleteStack(request *jdro.DeleteStackRequest) (*jdro.DeleteStackResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DeleteStackResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询支持的资源结构详情 */
func (c *JdroClient) DescribeResourceTypeSpecification(request *jdro.DescribeResourceTypeSpecificationRequest) (*jdro.DescribeResourceTypeSpecificationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeResourceTypeSpecificationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 模板校验 */
func (c *JdroClient) ValidateTemplate(request *jdro.ValidateTemplateRequest) (*jdro.ValidateTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.ValidateTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询支持的资源列表 */
func (c *JdroClient) DescribeResourceTypeList(request *jdro.DescribeResourceTypeListRequest) (*jdro.DescribeResourceTypeListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdro.DescribeResourceTypeListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

