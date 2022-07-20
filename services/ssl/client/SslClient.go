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
    ssl "github.com/lshuining/jdcloud-sdk-go/services/ssl/apis"
    "encoding/json"
    "errors"
)

type SslClient struct {
    core.JDCloudClient
}

func NewSslClient(credential *core.Credential) *SslClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("ssl.jdcloud-api.com")

    return &SslClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "ssl",
            Revision:    "1.0.2",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *SslClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *SslClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *SslClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 更新证书 [MFA enabled] */
func (c *SslClient) UpdateCert(request *ssl.UpdateCertRequest) (*ssl.UpdateCertResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.UpdateCertResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看证书详情 */
func (c *SslClient) DescribeCert(request *ssl.DescribeCertRequest) (*ssl.DescribeCertResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.DescribeCertResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 上传证书 */
func (c *SslClient) UploadCert(request *ssl.UploadCertRequest) (*ssl.UploadCertResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.UploadCertResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除证书 [MFA enabled] */
func (c *SslClient) DeleteCerts(request *ssl.DeleteCertsRequest) (*ssl.DeleteCertsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.DeleteCertsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 下载证书 [MFA enabled] */
func (c *SslClient) DownloadCert(request *ssl.DownloadCertRequest) (*ssl.DownloadCertResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.DownloadCertResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看证书列表 */
func (c *SslClient) DescribeCerts(request *ssl.DescribeCertsRequest) (*ssl.DescribeCertsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.DescribeCertsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改证书名称 */
func (c *SslClient) UpdateCertName(request *ssl.UpdateCertNameRequest) (*ssl.UpdateCertNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ssl.UpdateCertNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

