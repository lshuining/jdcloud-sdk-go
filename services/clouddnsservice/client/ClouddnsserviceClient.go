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
    clouddnsservice "github.com/lshuining/jdcloud-sdk-go/services/clouddnsservice/apis"
    "encoding/json"
    "errors"
)

type ClouddnsserviceClient struct {
    core.JDCloudClient
}

func NewClouddnsserviceClient(credential *core.Credential) *ClouddnsserviceClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("clouddnsservice.jdcloud-api.com")

    return &ClouddnsserviceClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "clouddnsservice",
            Revision:    "1.0.12",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *ClouddnsserviceClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *ClouddnsserviceClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 删除主域名 */
func (c *ClouddnsserviceClient) DelDomain(request *clouddnsservice.DelDomainRequest) (*clouddnsservice.DelDomainResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.DelDomainResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看主域名的解析次数 */
func (c *ClouddnsserviceClient) GetDomainQueryCount(request *clouddnsservice.GetDomainQueryCountRequest) (*clouddnsservice.GetDomainQueryCountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetDomainQueryCountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 同一个主域名下，批量新增或者批量更新导入解析记录。
如果解析记录的ID为0，是新增解析记录，如果不为0，则是更新解析记录。
 */
func (c *ClouddnsserviceClient) BatchSetDnsResolve(request *clouddnsservice.BatchSetDnsResolveRequest) (*clouddnsservice.BatchSetDnsResolveResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.BatchSetDnsResolveResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看用户在云解析服务下的操作记录 */
func (c *ClouddnsserviceClient) GetActionLog(request *clouddnsservice.GetActionLogRequest) (*clouddnsservice.GetActionLogResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetActionLogResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看域名的查询流量 */
func (c *ClouddnsserviceClient) GetDomainQueryTraffic(request *clouddnsservice.GetDomainQueryTrafficRequest) (*clouddnsservice.GetDomainQueryTrafficResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetDomainQueryTrafficResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 监控项的操作集合，包括：删除，暂停，启动, 手动恢复, 手动切换 */
func (c *ClouddnsserviceClient) OperateMonitor(request *clouddnsservice.OperateMonitorRequest) (*clouddnsservice.OperateMonitorResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.OperateMonitorResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加主域名的自定义解析线路的IP段 */
func (c *ClouddnsserviceClient) AddUserViewIP(request *clouddnsservice.AddUserViewIPRequest) (*clouddnsservice.AddUserViewIPResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.AddUserViewIPResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加子域名的某些特定监控对象为监控项 */
func (c *ClouddnsserviceClient) AddMonitorTarget(request *clouddnsservice.AddMonitorTargetRequest) (*clouddnsservice.AddMonitorTargetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.AddMonitorTargetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加主域名的解析记录 */
func (c *ClouddnsserviceClient) AddRR(request *clouddnsservice.AddRRRequest) (*clouddnsservice.AddRRResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.AddRRResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除主域名的自定义解析线路 */
func (c *ClouddnsserviceClient) DelUserView(request *clouddnsservice.DelUserViewRequest) (*clouddnsservice.DelUserViewResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.DelUserViewResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用、停用、删除主域名下的解析记录 */
func (c *ClouddnsserviceClient) OperateRR(request *clouddnsservice.OperateRRRequest) (*clouddnsservice.OperateRRResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.OperateRRResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除主域名的自定义解析线路的IP段 */
func (c *ClouddnsserviceClient) DelUserViewIP(request *clouddnsservice.DelUserViewIPRequest) (*clouddnsservice.DelUserViewIPResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.DelUserViewIPResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看主域名的监控项的配置以及状态 */
func (c *ClouddnsserviceClient) GetMonitor(request *clouddnsservice.GetMonitorRequest) (*clouddnsservice.GetMonitorResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetMonitorResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 域名的监控项修改 */
func (c *ClouddnsserviceClient) UpdateMonitor(request *clouddnsservice.UpdateMonitorRequest) (*clouddnsservice.UpdateMonitorResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.UpdateMonitorResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询主域名的自定义解析线路的IP段 */
func (c *ClouddnsserviceClient) GetUserViewIP(request *clouddnsservice.GetUserViewIPRequest) (*clouddnsservice.GetUserViewIPResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetUserViewIPResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询子域名的可用监控对象 */
func (c *ClouddnsserviceClient) GetTargets(request *clouddnsservice.GetTargetsRequest) (*clouddnsservice.GetTargetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetTargetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 主域名的监控项的报警信息 */
func (c *ClouddnsserviceClient) GetMonitorAlarmInfo(request *clouddnsservice.GetMonitorAlarmInfoRequest) (*clouddnsservice.GetMonitorAlarmInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetMonitorAlarmInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加子域名的监控项，默认把子域名的所有监控项都添加上监控 */
func (c *ClouddnsserviceClient) AddMonitor(request *clouddnsservice.AddMonitorRequest) (*clouddnsservice.AddMonitorResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.AddMonitorResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改主域名 */
func (c *ClouddnsserviceClient) UpdateDomain(request *clouddnsservice.UpdateDomainRequest) (*clouddnsservice.UpdateDomainResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.UpdateDomainResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改主域名的某个解析记录 */
func (c *ClouddnsserviceClient) UpdateRR(request *clouddnsservice.UpdateRRRequest) (*clouddnsservice.UpdateRRResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.UpdateRRResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加主域名的自定义解析线路 */
func (c *ClouddnsserviceClient) AddUserView(request *clouddnsservice.AddUserViewRequest) (*clouddnsservice.AddUserViewResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.AddUserViewResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询主域名的解析记录。  
在使用解析记录相关的接口之前，请调用此接口获取解析记录的列表。
 */
func (c *ClouddnsserviceClient) SearchRR(request *clouddnsservice.SearchRRRequest) (*clouddnsservice.SearchRRResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.SearchRRResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询主域名的自定义解析线路 */
func (c *ClouddnsserviceClient) GetUserView(request *clouddnsservice.GetUserViewRequest) (*clouddnsservice.GetUserViewResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetUserViewResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户所属的主域名列表。   
请在调用域名相关的接口之前，调用此接口获取相关的domainId和domainName。  
主域名的相关概念，请查阅<a href="https://docs.jdcloud.com/cn/jd-cloud-dns/product-overview">云解析文档</a>
 */
func (c *ClouddnsserviceClient) GetDomains(request *clouddnsservice.GetDomainsRequest) (*clouddnsservice.GetDomainsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetDomainsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加主域名  
如何添加免费域名，详细情况请查阅<a href="https://docs.jdcloud.com/cn/jd-cloud-dns/domainadd">文档</a>
添加收费域名，请查阅<a href="https://docs.jdcloud.com/cn/jd-cloud-dns/purchase-process">文档</a>，
添加收费域名前，请确保用户的京东云账户有足够的资金支付，Openapi接口回返回订单号，可以用此订单号向计费系统查阅详情。
 */
func (c *ClouddnsserviceClient) AddDomain(request *clouddnsservice.AddDomainRequest) (*clouddnsservice.AddDomainResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.AddDomainResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询云解析所有的基础解析线路。  
在使用解析线路的参数之前，请调用此接口获取解析线路的ID。
 */
func (c *ClouddnsserviceClient) GetViewTree(request *clouddnsservice.GetViewTreeRequest) (*clouddnsservice.GetViewTreeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &clouddnsservice.GetViewTreeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

