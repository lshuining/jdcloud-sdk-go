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
    disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/apis"
    "encoding/json"
    "errors"
)

type DiskClient struct {
    core.JDCloudClient
}

func NewDiskClient(credential *core.Credential) *DiskClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("disk.jdcloud-api.com")

    return &DiskClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "disk",
            Revision:    "0.12.3",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *DiskClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *DiskClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *DiskClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 修改快照的名字或描述信息 */
func (c *DiskClient) ModifySnapshotAttribute(request *disk.ModifySnapshotAttributeRequest) (*disk.ModifySnapshotAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.ModifySnapshotAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   创建一块或多块按配置或者按使用时长付费的云硬盘。
-   云硬盘类型包括高效云盘(premium-hdd)、SSD云盘(ssd)、通用型SSD(ssd.gp1)、性能型SSD(ssd.io1)、容量型HDD(hdd.std1)。
-   计费方式默认为按配置付费。
-   创建完成后，云硬盘状态为 available。
-   可选参数快照 ID用于从快照创建新盘。
-   批量创建时，云硬盘的命名为 硬盘名称-数字，例如 myDisk-1，myDisk-2。
-   maxCount为最大努力，不保证一定能达到maxCount。
 */
func (c *DiskClient) CreateDisks(request *disk.CreateDisksRequest) (*disk.CreateDisksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.CreateDisksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   删除单个云硬盘快照:快照状态必须为 available 或 error 状态。
-   快照独立于云硬盘生命周期，删除快照不会对创建快照的云硬盘有任何影响。
-   快照删除后不可恢复，请谨慎操作。
 */
func (c *DiskClient) DeleteSnapshot(request *disk.DeleteSnapshotRequest) (*disk.DeleteSnapshotResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.DeleteSnapshotResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询云硬盘快照列表，filters多个过滤条件之间是逻辑与(AND)，每个条件内部的多个取值是逻辑或(OR) */
func (c *DiskClient) DescribeSnapshots(request *disk.DescribeSnapshotsRequest) (*disk.DescribeSnapshotsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.DescribeSnapshotsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   仅可对制作快照的源硬盘进行数据恢复操作。
-   仅源硬盘处于可用状态时才能使用快照进行数据恢复操作。
-   云硬盘恢复后，当前数据将被清除，请您谨慎操作。
 */
func (c *DiskClient) RestoreDisk(request *disk.RestoreDiskRequest) (*disk.RestoreDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.RestoreDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询某一块云硬盘的信息详情 */
func (c *DiskClient) DescribeDisk(request *disk.DescribeDiskRequest) (*disk.DescribeDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.DescribeDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   扩容云硬盘到指定大小，云硬盘状态必须为 available。
-   当云硬盘正在创建快照时，不允许扩容。
 */
func (c *DiskClient) ExtendDisk(request *disk.ExtendDiskRequest) (*disk.ExtendDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.ExtendDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改云硬盘的名字或描述信息，名字或描述信息至少要指定一个。 */
func (c *DiskClient) ModifyDiskAttribute(request *disk.ModifyDiskAttributeRequest) (*disk.ModifyDiskAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.ModifyDiskAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   查询您已经创建的云硬盘。
-   filters多个过滤条件之间是逻辑与(AND)，每个条件内部的多个取值是逻辑或(OR)
 */
func (c *DiskClient) DescribeDisks(request *disk.DescribeDisksRequest) (*disk.DescribeDisksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.DescribeDisksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   删除一块按配置计费的云硬盘，云盘类型包括高效云盘、SSD云盘、通用型SSD、性能型SSD和容量型HDD。
-   删除云盘时，云盘的状态必须为 待挂载（Available）。
-   云盘被删除后，云硬盘快照可以被保留。
 */
func (c *DiskClient) DeleteDisk(request *disk.DeleteDiskRequest) (*disk.DeleteDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.DeleteDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* -   为指定云硬盘创建快照，新生成的快照的状态为creating。
-   同一地域下单用户快照的配额为15块。
-   为保证数据完整性，请您在创建快照之前，停止对云硬盘进行写入操作，以保证快照数据的完整性。
-   在执行创建快照前，建议您对云硬盘进行卸载操作，创建快照后再重新挂载到云主机上。
-   手动快照的生命周期独立于云硬盘，请您及时删除不需要的快照。
-   创建快照所需时间取决于云硬盘容量的大小，云硬盘容量越大耗时越长。
 */
func (c *DiskClient) CreateSnapshot(request *disk.CreateSnapshotRequest) (*disk.CreateSnapshotResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.CreateSnapshotResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询云硬盘快照信息详情 */
func (c *DiskClient) DescribeSnapshot(request *disk.DescribeSnapshotRequest) (*disk.DescribeSnapshotResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &disk.DescribeSnapshotResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

