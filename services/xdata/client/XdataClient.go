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
    xdata "github.com/lshuining/jdcloud-sdk-go/services/xdata/apis"
    "encoding/json"
    "errors"
)

type XdataClient struct {
    core.JDCloudClient
}

func NewXdataClient(credential *core.Credential) *XdataClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("xdata.jdcloud-api.com")

    return &XdataClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "xdata",
            Revision:    "1.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *XdataClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *XdataClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建属于用户实例的数据库 */
func (c *XdataClient) CreateDatabase(request *xdata.CreateDatabaseRequest) (*xdata.CreateDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.CreateDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 终止用户Spark SQL脚本查询 */
func (c *XdataClient) CancelRasQuery(request *xdata.CancelRasQueryRequest) (*xdata.CancelRasQueryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.CancelRasQueryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户实例的所有数据库信息 */
func (c *XdataClient) ListDatabaseInfo(request *xdata.ListDatabaseInfoRequest) (*xdata.ListDatabaseInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ListDatabaseInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户PySpark脚本的执行状态 */
func (c *XdataClient) GetPySparkExecuteState(request *xdata.GetPySparkExecuteStateRequest) (*xdata.GetPySparkExecuteStateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetPySparkExecuteStateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户实例的指定数据库信息 */
func (c *XdataClient) GetDatabaseInfo(request *xdata.GetDatabaseInfoRequest) (*xdata.GetDatabaseInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetDatabaseInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户实例指定数据库下的所有数据表信息 */
func (c *XdataClient) ListTableInfo(request *xdata.ListTableInfoRequest) (*xdata.ListTableInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ListTableInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户Spark SQL脚本的查询状态 */
func (c *XdataClient) GetRasQueryState(request *xdata.GetRasQueryStateRequest) (*xdata.GetRasQueryStateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetRasQueryStateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 执行用户编写的Spark SQL脚本 */
func (c *XdataClient) ExecuteRasQuery(request *xdata.ExecuteRasQueryRequest) (*xdata.ExecuteRasQueryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ExecuteRasQueryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 终止用户PySpark脚本任务 */
func (c *XdataClient) CancelPySparkJob(request *xdata.CancelPySparkJobRequest) (*xdata.CancelPySparkJobResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.CancelPySparkJobResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户PySpark脚本的执行结果 */
func (c *XdataClient) GetPySparkExecuteResult(request *xdata.GetPySparkExecuteResultRequest) (*xdata.GetPySparkExecuteResultResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetPySparkExecuteResultResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户所属的实例信息 */
func (c *XdataClient) ListInstanceInfo(request *xdata.ListInstanceInfoRequest) (*xdata.ListInstanceInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ListInstanceInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除用户实例的指定数据库 */
func (c *XdataClient) DeleteDatabase(request *xdata.DeleteDatabaseRequest) (*xdata.DeleteDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.DeleteDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户实例的指定数据表信息 */
func (c *XdataClient) GetTableInfo(request *xdata.GetTableInfoRequest) (*xdata.GetTableInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetTableInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 执行用户编写的PySpark脚本 */
func (c *XdataClient) ExecutePySparkQuery(request *xdata.ExecutePySparkQueryRequest) (*xdata.ExecutePySparkQueryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ExecutePySparkQueryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户Spark SQL脚本的查询结果 */
func (c *XdataClient) GetRasQueryResult(request *xdata.GetRasQueryResultRequest) (*xdata.GetRasQueryResultResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetRasQueryResultResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建用户实例的数据表 */
func (c *XdataClient) CreateTable(request *xdata.CreateTableRequest) (*xdata.CreateTableResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.CreateTableResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户Spark SQL脚本的查询日志 */
func (c *XdataClient) GetRasQueryLog(request *xdata.GetRasQueryLogRequest) (*xdata.GetRasQueryLogResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetRasQueryLogResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除用户实例的指定数据表 */
func (c *XdataClient) DeleteTable(request *xdata.DeleteTableRequest) (*xdata.DeleteTableResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.DeleteTableResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

