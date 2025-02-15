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

package apis

import (
    "github.com/lshuining/jdcloud-sdk-go/core"
    detection "github.com/lshuining/jdcloud-sdk-go/services/detection/models"
)

type CreateSiteMonitorRequest struct {

    core.JDCloudRequest

    /* 地址  */
    Address string `json:"address"`

    /*  (Optional) */
    AdvanceChecked *string `json:"advanceChecked"`

    /*  (Optional) */
    CreatedTime *int64 `json:"createdTime"`

    /* 探测频率  */
    Cycle int64 `json:"cycle"`

    /*  (Optional) */
    DefaultSource *string `json:"defaultSource"`

    /*  (Optional) */
    DnsOption *detection.SiteMonitorDnsOption `json:"dnsOption"`

    /*  (Optional) */
    Enabled *string `json:"enabled"`

    /*  (Optional) */
    FtpOption *detection.SiteMonitorFtpOption `json:"ftpOption"`

    /*  (Optional) */
    HawkeyeId *int64 `json:"hawkeyeId"`

    /*  (Optional) */
    HttpOption *detection.SiteMonitorHttpOption `json:"httpOption"`

    /*  (Optional) */
    Id *string `json:"id"`

    /*  (Optional) */
    IsDeleted *string `json:"isDeleted"`

    /* 任务名称  */
    Name string `json:"name"`

    /*  (Optional) */
    Pin *string `json:"pin"`

    /*  (Optional) */
    PingOption *detection.SiteMonitorPingOption `json:"pingOption"`

    /*  (Optional) */
    Pop3Option *detection.SiteMonitorPop3Option `json:"pop3Option"`

    /* 端口 (Optional) */
    Port *string `json:"port"`

    /*  (Optional) */
    SmtpOption *detection.SiteMonitorSmtpOption `json:"smtpOption"`

    /* 探测源  */
    Source []detection.SiteMonitorSource `json:"source"`

    /*  (Optional) */
    Stats *interface{} `json:"stats"`

    /* 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP  */
    TaskType string `json:"taskType"`

    /*  (Optional) */
    TcpOption *detection.SiteMonitorTcpOption `json:"tcpOption"`

    /*  (Optional) */
    UdpOption *detection.SiteMonitorUdpOption `json:"udpOption"`

    /*  (Optional) */
    UpdatedTime *int64 `json:"updatedTime"`
}

/*
 * param address: 地址 (Required)
 * param cycle: 探测频率 (Required)
 * param name: 任务名称 (Required)
 * param source: 探测源 (Required)
 * param taskType: 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSiteMonitorRequest(
    address string,
    cycle int64,
    name string,
    source []detection.SiteMonitorSource,
    taskType string,
) *CreateSiteMonitorRequest {

	return &CreateSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/siteMonitor",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        Address: address,
        Cycle: cycle,
        Name: name,
        Source: source,
        TaskType: taskType,
	}
}

/*
 * param address: 地址 (Required)
 * param advanceChecked:  (Optional)
 * param createdTime:  (Optional)
 * param cycle: 探测频率 (Required)
 * param defaultSource:  (Optional)
 * param dnsOption:  (Optional)
 * param enabled:  (Optional)
 * param ftpOption:  (Optional)
 * param hawkeyeId:  (Optional)
 * param httpOption:  (Optional)
 * param id:  (Optional)
 * param isDeleted:  (Optional)
 * param name: 任务名称 (Required)
 * param pin:  (Optional)
 * param pingOption:  (Optional)
 * param pop3Option:  (Optional)
 * param port: 端口 (Optional)
 * param smtpOption:  (Optional)
 * param source: 探测源 (Required)
 * param stats:  (Optional)
 * param taskType: 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP (Required)
 * param tcpOption:  (Optional)
 * param udpOption:  (Optional)
 * param updatedTime:  (Optional)
 */
func NewCreateSiteMonitorRequestWithAllParams(
    address string,
    advanceChecked *string,
    createdTime *int64,
    cycle int64,
    defaultSource *string,
    dnsOption *detection.SiteMonitorDnsOption,
    enabled *string,
    ftpOption *detection.SiteMonitorFtpOption,
    hawkeyeId *int64,
    httpOption *detection.SiteMonitorHttpOption,
    id *string,
    isDeleted *string,
    name string,
    pin *string,
    pingOption *detection.SiteMonitorPingOption,
    pop3Option *detection.SiteMonitorPop3Option,
    port *string,
    smtpOption *detection.SiteMonitorSmtpOption,
    source []detection.SiteMonitorSource,
    stats *interface{},
    taskType string,
    tcpOption *detection.SiteMonitorTcpOption,
    udpOption *detection.SiteMonitorUdpOption,
    updatedTime *int64,
) *CreateSiteMonitorRequest {

    return &CreateSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        Address: address,
        AdvanceChecked: advanceChecked,
        CreatedTime: createdTime,
        Cycle: cycle,
        DefaultSource: defaultSource,
        DnsOption: dnsOption,
        Enabled: enabled,
        FtpOption: ftpOption,
        HawkeyeId: hawkeyeId,
        HttpOption: httpOption,
        Id: id,
        IsDeleted: isDeleted,
        Name: name,
        Pin: pin,
        PingOption: pingOption,
        Pop3Option: pop3Option,
        Port: port,
        SmtpOption: smtpOption,
        Source: source,
        Stats: stats,
        TaskType: taskType,
        TcpOption: tcpOption,
        UdpOption: udpOption,
        UpdatedTime: updatedTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSiteMonitorRequestWithoutParam() *CreateSiteMonitorRequest {

    return &CreateSiteMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param address: 地址(Required) */
func (r *CreateSiteMonitorRequest) SetAddress(address string) {
    r.Address = address
}

/* param advanceChecked: (Optional) */
func (r *CreateSiteMonitorRequest) SetAdvanceChecked(advanceChecked string) {
    r.AdvanceChecked = &advanceChecked
}

/* param createdTime: (Optional) */
func (r *CreateSiteMonitorRequest) SetCreatedTime(createdTime int64) {
    r.CreatedTime = &createdTime
}

/* param cycle: 探测频率(Required) */
func (r *CreateSiteMonitorRequest) SetCycle(cycle int64) {
    r.Cycle = cycle
}

/* param defaultSource: (Optional) */
func (r *CreateSiteMonitorRequest) SetDefaultSource(defaultSource string) {
    r.DefaultSource = &defaultSource
}

/* param dnsOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetDnsOption(dnsOption *detection.SiteMonitorDnsOption) {
    r.DnsOption = dnsOption
}

/* param enabled: (Optional) */
func (r *CreateSiteMonitorRequest) SetEnabled(enabled string) {
    r.Enabled = &enabled
}

/* param ftpOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetFtpOption(ftpOption *detection.SiteMonitorFtpOption) {
    r.FtpOption = ftpOption
}

/* param hawkeyeId: (Optional) */
func (r *CreateSiteMonitorRequest) SetHawkeyeId(hawkeyeId int64) {
    r.HawkeyeId = &hawkeyeId
}

/* param httpOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetHttpOption(httpOption *detection.SiteMonitorHttpOption) {
    r.HttpOption = httpOption
}

/* param id: (Optional) */
func (r *CreateSiteMonitorRequest) SetId(id string) {
    r.Id = &id
}

/* param isDeleted: (Optional) */
func (r *CreateSiteMonitorRequest) SetIsDeleted(isDeleted string) {
    r.IsDeleted = &isDeleted
}

/* param name: 任务名称(Required) */
func (r *CreateSiteMonitorRequest) SetName(name string) {
    r.Name = name
}

/* param pin: (Optional) */
func (r *CreateSiteMonitorRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param pingOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetPingOption(pingOption *detection.SiteMonitorPingOption) {
    r.PingOption = pingOption
}

/* param pop3Option: (Optional) */
func (r *CreateSiteMonitorRequest) SetPop3Option(pop3Option *detection.SiteMonitorPop3Option) {
    r.Pop3Option = pop3Option
}

/* param port: 端口(Optional) */
func (r *CreateSiteMonitorRequest) SetPort(port string) {
    r.Port = &port
}

/* param smtpOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetSmtpOption(smtpOption *detection.SiteMonitorSmtpOption) {
    r.SmtpOption = smtpOption
}

/* param source: 探测源(Required) */
func (r *CreateSiteMonitorRequest) SetSource(source []detection.SiteMonitorSource) {
    r.Source = source
}

/* param stats: (Optional) */
func (r *CreateSiteMonitorRequest) SetStats(stats interface{}) {
    r.Stats = &stats
}

/* param taskType: 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP(Required) */
func (r *CreateSiteMonitorRequest) SetTaskType(taskType string) {
    r.TaskType = taskType
}

/* param tcpOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetTcpOption(tcpOption *detection.SiteMonitorTcpOption) {
    r.TcpOption = tcpOption
}

/* param udpOption: (Optional) */
func (r *CreateSiteMonitorRequest) SetUdpOption(udpOption *detection.SiteMonitorUdpOption) {
    r.UdpOption = udpOption
}

/* param updatedTime: (Optional) */
func (r *CreateSiteMonitorRequest) SetUpdatedTime(updatedTime int64) {
    r.UpdatedTime = &updatedTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSiteMonitorRequest) GetRegionId() string {
    return ""
}

type CreateSiteMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSiteMonitorResult `json:"result"`
}

type CreateSiteMonitorResult struct {
    Id string `json:"id"`
}