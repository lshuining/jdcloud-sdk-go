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
)

type CollectorReadMessageRequest struct {

    core.JDCloudRequest

    /* Hub实例Id  */
    InstanceId string `json:"instanceId"`

    /* 区域Id  */
    RegionId string `json:"regionId"`

    /* 当前的链接码  */
    Identifier string `json:"identifier"`

    /* 当前的协议类型：
语音播报控制器-输入端子,0X00000~X0007：inputTerminal
语音播报控制器-播放信息,0X00024~X0027：playInfo
LR001-516-5B边缘数据采集器-传感器管理：sensor
LR001-516-5B边缘数据采集器-采集器属性：collectorProperty
LR001-516-5B边缘数据采集器-电梯属性：elevatorProperty
LR001-516-5A边缘数据采集器-水质酸碱度(PH)：waterQualityPh
LR001-516-5A水质监测采集器-水质电导率：waterQualityElectroConductivity
  */
    Protocol string `json:"protocol"`
}

/*
 * param instanceId: Hub实例Id (Required)
 * param regionId: 区域Id (Required)
 * param identifier: 当前的链接码 (Required)
 * param protocol: 当前的协议类型：
语音播报控制器-输入端子,0X00000~X0007：inputTerminal
语音播报控制器-播放信息,0X00024~X0027：playInfo
LR001-516-5B边缘数据采集器-传感器管理：sensor
LR001-516-5B边缘数据采集器-采集器属性：collectorProperty
LR001-516-5B边缘数据采集器-电梯属性：elevatorProperty
LR001-516-5A边缘数据采集器-水质酸碱度(PH)：waterQualityPh
LR001-516-5A水质监测采集器-水质电导率：waterQualityElectroConductivity
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCollectorReadMessageRequest(
    instanceId string,
    regionId string,
    identifier string,
    protocol string,
) *CollectorReadMessageRequest {

	return &CollectorReadMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/readCollectorProperty",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
        Identifier: identifier,
        Protocol: protocol,
	}
}

/*
 * param instanceId: Hub实例Id (Required)
 * param regionId: 区域Id (Required)
 * param identifier: 当前的链接码 (Required)
 * param protocol: 当前的协议类型：
语音播报控制器-输入端子,0X00000~X0007：inputTerminal
语音播报控制器-播放信息,0X00024~X0027：playInfo
LR001-516-5B边缘数据采集器-传感器管理：sensor
LR001-516-5B边缘数据采集器-采集器属性：collectorProperty
LR001-516-5B边缘数据采集器-电梯属性：elevatorProperty
LR001-516-5A边缘数据采集器-水质酸碱度(PH)：waterQualityPh
LR001-516-5A水质监测采集器-水质电导率：waterQualityElectroConductivity
 (Required)
 */
func NewCollectorReadMessageRequestWithAllParams(
    instanceId string,
    regionId string,
    identifier string,
    protocol string,
) *CollectorReadMessageRequest {

    return &CollectorReadMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/readCollectorProperty",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        Identifier: identifier,
        Protocol: protocol,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCollectorReadMessageRequestWithoutParam() *CollectorReadMessageRequest {

    return &CollectorReadMessageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/readCollectorProperty",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: Hub实例Id(Required) */
func (r *CollectorReadMessageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 区域Id(Required) */
func (r *CollectorReadMessageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param identifier: 当前的链接码(Required) */
func (r *CollectorReadMessageRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

/* param protocol: 当前的协议类型：
语音播报控制器-输入端子,0X00000~X0007：inputTerminal
语音播报控制器-播放信息,0X00024~X0027：playInfo
LR001-516-5B边缘数据采集器-传感器管理：sensor
LR001-516-5B边缘数据采集器-采集器属性：collectorProperty
LR001-516-5B边缘数据采集器-电梯属性：elevatorProperty
LR001-516-5A边缘数据采集器-水质酸碱度(PH)：waterQualityPh
LR001-516-5A水质监测采集器-水质电导率：waterQualityElectroConductivity
(Required) */
func (r *CollectorReadMessageRequest) SetProtocol(protocol string) {
    r.Protocol = protocol
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CollectorReadMessageRequest) GetRegionId() string {
    return r.RegionId
}

type CollectorReadMessageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CollectorReadMessageResult `json:"result"`
}

type CollectorReadMessageResult struct {
    ReadData interface{} `json:"readData"`
}