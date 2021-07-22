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
    openjrtc "github.com/jdcloud-api/jdcloud-sdk-go/services/openjrtc/apis"
    "encoding/json"
    "errors"
)

type OpenjrtcClient struct {
    core.JDCloudClient
}

func NewOpenjrtcClient(credential *core.Credential) *OpenjrtcClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("openjrtc.jdcloud-api.com")

    return &OpenjrtcClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "openjrtc",
            Revision:    "1.1.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *OpenjrtcClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *OpenjrtcClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *OpenjrtcClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询应用下的房间列表
允许通过条件过滤查询，支持的过滤字段如下：
           - appId[eq] 按应用ID查询
 */
func (c *OpenjrtcClient) DescribeRooms(request *openjrtc.DescribeRoomsRequest) (*openjrtc.DescribeRoomsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询房间内人员列表
允许通过条件过滤查询，支持的过滤字段如下：
           - status[eq] 在线状态 1-在线 2-离线
           - startTime[eq] 用户加入时间段开始时间-UTC时间  startTime,endTime同时有值时生效
           - endTime[eq] 用户加入时间段结束时间-UTC时间    startTime,endTime同时有值时生效
 */
func (c *OpenjrtcClient) DescribeRoomUsers(request *openjrtc.DescribeRoomUsersRequest) (*openjrtc.DescribeRoomUsersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomUsersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 统计房间内人数
 */
func (c *OpenjrtcClient) DescribeRoomUsersNum(request *openjrtc.DescribeRoomUsersNumRequest) (*openjrtc.DescribeRoomUsersNumResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomUsersNumResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建房间
 */
func (c *OpenjrtcClient) CreateRoom(request *openjrtc.CreateRoomRequest) (*openjrtc.CreateRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.CreateRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询注册用户
 */
func (c *OpenjrtcClient) DescribeRegisterUser(request *openjrtc.DescribeRegisterUserRequest) (*openjrtc.DescribeRegisterUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRegisterUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改房间
 */
func (c *OpenjrtcClient) UpdateUserRoom(request *openjrtc.UpdateUserRoomRequest) (*openjrtc.UpdateUserRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.UpdateUserRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询注册房间号列表
允许通过条件过滤查询，支持的过滤字段如下：
           - startTime[eq] 房间注册时间段开始时间-UTC时间 startTime,endTime同时有值时生效
           - endTime[eq] 房间注册时间段结束时间-UTC时间   startTime,endTime同时有值时生效
 */
func (c *OpenjrtcClient) DescribeUserRooms(request *openjrtc.DescribeUserRoomsRequest) (*openjrtc.DescribeUserRoomsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeUserRoomsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询应用appKey
 */
func (c *OpenjrtcClient) DescribeAppKey(request *openjrtc.DescribeAppKeyRequest) (*openjrtc.DescribeAppKeyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeAppKeyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 发送自定义信令给房间内的人员 */
func (c *OpenjrtcClient) SendMessageToUser(request *openjrtc.SendMessageToUserRequest) (*openjrtc.SendMessageToUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.SendMessageToUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改房间
 */
func (c *OpenjrtcClient) UpdateRoom(request *openjrtc.UpdateRoomRequest) (*openjrtc.UpdateRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.UpdateRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 生成token-用户加入房间时携带token校验通过后方能加入
 */
func (c *OpenjrtcClient) CreateToken(request *openjrtc.CreateTokenRequest) (*openjrtc.CreateTokenResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.CreateTokenResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 注册用户房间号-将业务接入方定义的userRoomId注册为jrtc系统内可识别和流转的房间号
 */
func (c *OpenjrtcClient) RegisterUserRoom(request *openjrtc.RegisterUserRoomRequest) (*openjrtc.RegisterUserRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.RegisterUserRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 注册用户-将业务接入方用户体系的userId注册为jrtc系统内可识别和流转的用户id
 */
func (c *OpenjrtcClient) RegisterUser(request *openjrtc.RegisterUserRequest) (*openjrtc.RegisterUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.RegisterUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 发送自定义信令给房间 */
func (c *OpenjrtcClient) SendMessageToRoom(request *openjrtc.SendMessageToRoomRequest) (*openjrtc.SendMessageToRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.SendMessageToRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询peer对应的用户信息
 */
func (c *OpenjrtcClient) DescribeUserByPeer(request *openjrtc.DescribeUserByPeerRequest) (*openjrtc.DescribeUserByPeerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeUserByPeerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询房间实时在线人数
 */
func (c *OpenjrtcClient) DescribeRoomOnlineUserNum(request *openjrtc.DescribeRoomOnlineUserNumRequest) (*openjrtc.DescribeRoomOnlineUserNumResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomOnlineUserNumResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 发送自定义信令给房间内的人员 */
func (c *OpenjrtcClient) PostMessageToUser(request *openjrtc.PostMessageToUserRequest) (*openjrtc.PostMessageToUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.PostMessageToUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户应用列表信息
 */
func (c *OpenjrtcClient) DescribeApps(request *openjrtc.DescribeAppsRequest) (*openjrtc.DescribeAppsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeAppsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除房间
 */
func (c *OpenjrtcClient) DeleteRoom(request *openjrtc.DeleteRoomRequest) (*openjrtc.DeleteRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DeleteRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询注册用户列表
允许通过条件过滤查询，支持的过滤字段如下：
           - startTime[eq] 用户注册时间段开始时间-UTC时间 startTime,endTime同时有值时生效
           - endTime[eq] 用户注册时间段结束时间-UTC时间 startTime,endTime同时有值时生效
 */
func (c *OpenjrtcClient) DescribeRegisterUsers(request *openjrtc.DescribeRegisterUsersRequest) (*openjrtc.DescribeRegisterUsersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRegisterUsersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 发送自定义信令给房间 */
func (c *OpenjrtcClient) PostMessageToUserRoom(request *openjrtc.PostMessageToUserRoomRequest) (*openjrtc.PostMessageToUserRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.PostMessageToUserRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 移除房间内所有人员
 */
func (c *OpenjrtcClient) RemoveAllUsersByUserRoomId(request *openjrtc.RemoveAllUsersByUserRoomIdRequest) (*openjrtc.RemoveAllUsersByUserRoomIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.RemoveAllUsersByUserRoomIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 移除房间内人员
 */
func (c *OpenjrtcClient) RemoveRoomUser(request *openjrtc.RemoveRoomUserRequest) (*openjrtc.RemoveRoomUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.RemoveRoomUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 移除房间内人员
 */
func (c *OpenjrtcClient) RemoveUserByUserRoomId(request *openjrtc.RemoveUserByUserRoomIdRequest) (*openjrtc.RemoveUserByUserRoomIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.RemoveUserByUserRoomIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取房间信息
 */
func (c *OpenjrtcClient) DescribeRoomInfo(request *openjrtc.DescribeRoomInfoRequest) (*openjrtc.DescribeRoomInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeRoomInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询应用信息:
 */
func (c *OpenjrtcClient) DescribeApp(request *openjrtc.DescribeAppRequest) (*openjrtc.DescribeAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 移除房间内所有人员
 */
func (c *OpenjrtcClient) RemoveAllRoomUsers(request *openjrtc.RemoveAllRoomUsersRequest) (*openjrtc.RemoveAllRoomUsersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.RemoveAllRoomUsersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询注册房间号
 */
func (c *OpenjrtcClient) DescribeUserRoom(request *openjrtc.DescribeUserRoomRequest) (*openjrtc.DescribeUserRoomResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.DescribeUserRoomResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建用户
 */
func (c *OpenjrtcClient) CreateUser(request *openjrtc.CreateUserRequest) (*openjrtc.CreateUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &openjrtc.CreateUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

