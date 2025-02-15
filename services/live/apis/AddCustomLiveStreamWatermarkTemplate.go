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

type AddCustomLiveStreamWatermarkTemplateRequest struct {

    core.JDCloudRequest

    /* 水印位置
- 取值范围：左上：1，右上：3， 左下：7，右下：9，默认：1
 (Optional) */
    Position *int `json:"position"`

    /* 偏移量单位
- 取值: percent,pixel
- percent:按百分比; pixel:像素 默认:pixel
 (Optional) */
    OffsetUnit *string `json:"offsetUnit"`

    /* x轴偏移量
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
  */
    OffsetX int `json:"offsetX"`

    /* y轴偏移量:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
  */
    OffsetY int `json:"offsetY"`

    /* 水印大小单位
- 取值: percent,pixel
- percent:按百分比; pixel:像素 默认:pixel
 (Optional) */
    SizeUnit *string `json:"sizeUnit"`

    /* 水印宽度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
  */
    Width int `json:"width"`

    /* 水印高度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
  */
    Height int `json:"height"`

    /* 自定义水印模板名称
-&ensp;取值要求: 数字、大小写字母、短横线("-")、下划线("_"),
&ensp;&ensp;首尾不能有特殊字符("-"),
&ensp;&ensp;不超过50字符,utf-8格式
-&ensp;<b>注意: 不能与已定义命名重复</b>
  */
    Template string `json:"template"`

    /* 创建上传任务时返回的uploadId参数，当通过接口上传水印图片时，uploadId必填
 (Optional) */
    UploadId *string `json:"uploadId"`

    /* 水印地址<br>-&ensp;以&ensp;http:// 开头,可公开访问地址<br>  */
    Url string `json:"url"`
}

/*
 * param offsetX: x轴偏移量
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param offsetY: y轴偏移量:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param width: 水印宽度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param height: 水印高度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param template: 自定义水印模板名称
-&ensp;取值要求: 数字、大小写字母、短横线("-")、下划线("_"),
&ensp;&ensp;首尾不能有特殊字符("-"),
&ensp;&ensp;不超过50字符,utf-8格式
-&ensp;<b>注意: 不能与已定义命名重复</b>
 (Required)
 * param url: 水印地址<br>-&ensp;以&ensp;http:// 开头,可公开访问地址<br> (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCustomLiveStreamWatermarkTemplateRequest(
    offsetX int,
    offsetY int,
    width int,
    height int,
    template string,
    url string,
) *AddCustomLiveStreamWatermarkTemplateRequest {

	return &AddCustomLiveStreamWatermarkTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/watermarkCustoms:template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        OffsetX: offsetX,
        OffsetY: offsetY,
        Width: width,
        Height: height,
        Template: template,
        Url: url,
	}
}

/*
 * param position: 水印位置
- 取值范围：左上：1，右上：3， 左下：7，右下：9，默认：1
 (Optional)
 * param offsetUnit: 偏移量单位
- 取值: percent,pixel
- percent:按百分比; pixel:像素 默认:pixel
 (Optional)
 * param offsetX: x轴偏移量
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param offsetY: y轴偏移量:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param sizeUnit: 水印大小单位
- 取值: percent,pixel
- percent:按百分比; pixel:像素 默认:pixel
 (Optional)
 * param width: 水印宽度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param height: 水印高度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
 (Required)
 * param template: 自定义水印模板名称
-&ensp;取值要求: 数字、大小写字母、短横线("-")、下划线("_"),
&ensp;&ensp;首尾不能有特殊字符("-"),
&ensp;&ensp;不超过50字符,utf-8格式
-&ensp;<b>注意: 不能与已定义命名重复</b>
 (Required)
 * param uploadId: 创建上传任务时返回的uploadId参数，当通过接口上传水印图片时，uploadId必填
 (Optional)
 * param url: 水印地址<br>-&ensp;以&ensp;http:// 开头,可公开访问地址<br> (Required)
 */
func NewAddCustomLiveStreamWatermarkTemplateRequestWithAllParams(
    position *int,
    offsetUnit *string,
    offsetX int,
    offsetY int,
    sizeUnit *string,
    width int,
    height int,
    template string,
    uploadId *string,
    url string,
) *AddCustomLiveStreamWatermarkTemplateRequest {

    return &AddCustomLiveStreamWatermarkTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarkCustoms:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Position: position,
        OffsetUnit: offsetUnit,
        OffsetX: offsetX,
        OffsetY: offsetY,
        SizeUnit: sizeUnit,
        Width: width,
        Height: height,
        Template: template,
        UploadId: uploadId,
        Url: url,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCustomLiveStreamWatermarkTemplateRequestWithoutParam() *AddCustomLiveStreamWatermarkTemplateRequest {

    return &AddCustomLiveStreamWatermarkTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarkCustoms:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param position: 水印位置
- 取值范围：左上：1，右上：3， 左下：7，右下：9，默认：1
(Optional) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetPosition(position int) {
    r.Position = &position
}

/* param offsetUnit: 偏移量单位
- 取值: percent,pixel
- percent:按百分比; pixel:像素 默认:pixel
(Optional) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetOffsetUnit(offsetUnit string) {
    r.OffsetUnit = &offsetUnit
}

/* param offsetX: x轴偏移量
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetOffsetX(offsetX int) {
    r.OffsetX = offsetX
}

/* param offsetY: y轴偏移量:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetOffsetY(offsetY int) {
    r.OffsetY = offsetY
}

/* param sizeUnit: 水印大小单位
- 取值: percent,pixel
- percent:按百分比; pixel:像素 默认:pixel
(Optional) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetSizeUnit(sizeUnit string) {
    r.SizeUnit = &sizeUnit
}

/* param width: 水印宽度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetWidth(width int) {
    r.Width = width
}

/* param height: 水印高度:
- 取值范围
  percent: (0,100]
  pixel: (0,1920]
(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetHeight(height int) {
    r.Height = height
}

/* param template: 自定义水印模板名称
-&ensp;取值要求: 数字、大小写字母、短横线("-")、下划线("_"),
&ensp;&ensp;首尾不能有特殊字符("-"),
&ensp;&ensp;不超过50字符,utf-8格式
-&ensp;<b>注意: 不能与已定义命名重复</b>
(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetTemplate(template string) {
    r.Template = template
}

/* param uploadId: 创建上传任务时返回的uploadId参数，当通过接口上传水印图片时，uploadId必填
(Optional) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetUploadId(uploadId string) {
    r.UploadId = &uploadId
}

/* param url: 水印地址<br>-&ensp;以&ensp;http:// 开头,可公开访问地址<br>(Required) */
func (r *AddCustomLiveStreamWatermarkTemplateRequest) SetUrl(url string) {
    r.Url = url
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCustomLiveStreamWatermarkTemplateRequest) GetRegionId() string {
    return ""
}

type AddCustomLiveStreamWatermarkTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCustomLiveStreamWatermarkTemplateResult `json:"result"`
}

type AddCustomLiveStreamWatermarkTemplateResult struct {
}