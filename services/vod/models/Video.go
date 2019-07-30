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

package models


type Video struct {

    /* 视频编码。取值范围：h265、h264 (Optional) */
    Codec *string `json:"codec"`

    /* 视频码率。取值范围 [128、10000]，单位为 Kbps (Optional) */
    Bitrate *int `json:"bitrate"`

    /* 视频帧率。取值范围为 [1、60]，单位为 fps (Optional) */
    Fps *int `json:"fps"`

    /* 视频输出宽度。取值范围 [128，4096]，取值需为2的倍数 (Optional) */
    Width *int `json:"width"`

    /* 视频输出高度。取值范围 [128，4096]，取值需为2的倍数 (Optional) */
    Height *int `json:"height"`
}
