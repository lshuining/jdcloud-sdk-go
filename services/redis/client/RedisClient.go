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
    redis "github.com/jdcloud-api/jdcloud-sdk-go/services/redis/apis"
    "encoding/json"
    "errors"
)

type RedisClient struct {
    core.JDCloudClient
}

func NewRedisClient(credential *core.Credential) *RedisClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("redis.jdcloud-api.com")

    return &RedisClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "redis",
            Revision:    "1.3.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *RedisClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *RedisClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询账户的缓存Redis配额信息 */
func (c *RedisClient) DescribeUserQuota(request *redis.DescribeUserQuotaRequest) (*redis.DescribeUserQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.DescribeUserQuotaResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 变更缓存Redis实例规格（变配），只能变更运行状态的实例规格，变更的规格不能与之前的相同。
预付费用户，从集群版变配到主从版，新规格的内存大小要大于老规格的内存大小，从主从版到集群版，新规格的内存大小要不小于老规格的内存大小。
 */
func (c *RedisClient) ModifyCacheInstanceClass(request *redis.ModifyCacheInstanceClassRequest) (*redis.ModifyCacheInstanceClassResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.ModifyCacheInstanceClassResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询缓存Redis实例列表，可分页、可排序、可搜索、可过滤 */
func (c *RedisClient) DescribeCacheInstances(request *redis.DescribeCacheInstancesRequest) (*redis.DescribeCacheInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.DescribeCacheInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询Redis实例的内部集群信息 */
func (c *RedisClient) DescribeClusterInfo(request *redis.DescribeClusterInfoRequest) (*redis.DescribeClusterInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.DescribeClusterInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改缓存Redis实例的资源名称或描述，二者至少选一 */
func (c *RedisClient) ModifyCacheInstanceAttribute(request *redis.ModifyCacheInstanceAttributeRequest) (*redis.ModifyCacheInstanceAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.ModifyCacheInstanceAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询缓存Redis实例的详细信息 */
func (c *RedisClient) DescribeCacheInstance(request *redis.DescribeCacheInstanceRequest) (*redis.DescribeCacheInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.DescribeCacheInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除按配置计费、或包年包月已到期的缓存Redis实例，包年包月未到期不可删除。
只有处于运行running或者错误error状态才可以删除，其余状态不可以删除。
白名单用户不能删除包年包月已到期的缓存Redis实例。
 */
func (c *RedisClient) DeleteCacheInstance(request *redis.DeleteCacheInstanceRequest) (*redis.DeleteCacheInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.DeleteCacheInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重置缓存Redis实例的密码，可为空 */
func (c *RedisClient) ResetCacheInstancePassword(request *redis.ResetCacheInstancePasswordRequest) (*redis.ResetCacheInstancePasswordResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.ResetCacheInstancePasswordResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询某区域下的缓存Redis实例规格列表 */
func (c *RedisClient) DescribeInstanceClass(request *redis.DescribeInstanceClassRequest) (*redis.DescribeInstanceClassResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.DescribeInstanceClassResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个指定配置的缓存Redis实例：可选择主从版或集群版，每种类型又分为多种规格（按CPU核数、内存容量、磁盘容量、带宽等划分），具体可参考产品规格代码，https://docs.jdcloud.com/cn/jcs-for-redis/specifications
 */
func (c *RedisClient) CreateCacheInstance(request *redis.CreateCacheInstanceRequest) (*redis.CreateCacheInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &redis.CreateCacheInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

