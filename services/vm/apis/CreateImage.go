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
    vm "github.com/lshuining/jdcloud-sdk-go/services/vm/models"
)

type CreateImageRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 云主机ID。  */
    InstanceId string `json:"instanceId"`

    /* 镜像名称，长度为2\~32个字符，只允许中文、数字、大小写字母、英文下划线（\_）、连字符（-）及点（.）。
  */
    Name string `json:"name"`

    /* 镜像描述。256字符以内。
 (Optional) */
    Description *string `json:"description"`

    /* 数据盘列表。
在不指定该参数的情况下，制作镜像的过程中会将该实例中的所有云盘数据盘制作快照，并与系统盘一起，制作成打包镜像。
如果不希望将实例中的某个云盘数据盘制作快照，可使用 `noDevice` 的方式排除，例如：`deviceName=vdb`、`noDevice=true` 就不会将 `vdb` 制作快照。
如果希望在打包镜像中插入一块新盘，该盘不在实例中，可通过指定新的 `deviceName` 的方式实现，例如：`deviceName=vdx` 将会在打包镜像中插入一块盘符为 `vdx` 的新盘，支持新盘使用或不使用快照都可以。
如果使用 `deviceName` 指定了与实例中相同的盘符，那么实例中对应的云盘数据盘也不会制作快照，并使用新指定的参数进行替换。
 (Optional) */
    DataDisks []vm.InstanceDiskAttachmentSpec `json:"dataDisks"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param name: 镜像名称，长度为2\~32个字符，只允许中文、数字、大小写字母、英文下划线（\_）、连字符（-）及点（.）。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateImageRequest(
    regionId string,
    instanceId string,
    name string,
) *CreateImageRequest {

	return &CreateImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:createImage",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Name: name,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param name: 镜像名称，长度为2\~32个字符，只允许中文、数字、大小写字母、英文下划线（\_）、连字符（-）及点（.）。
 (Required)
 * param description: 镜像描述。256字符以内。
 (Optional)
 * param dataDisks: 数据盘列表。
在不指定该参数的情况下，制作镜像的过程中会将该实例中的所有云盘数据盘制作快照，并与系统盘一起，制作成打包镜像。
如果不希望将实例中的某个云盘数据盘制作快照，可使用 `noDevice` 的方式排除，例如：`deviceName=vdb`、`noDevice=true` 就不会将 `vdb` 制作快照。
如果希望在打包镜像中插入一块新盘，该盘不在实例中，可通过指定新的 `deviceName` 的方式实现，例如：`deviceName=vdx` 将会在打包镜像中插入一块盘符为 `vdx` 的新盘，支持新盘使用或不使用快照都可以。
如果使用 `deviceName` 指定了与实例中相同的盘符，那么实例中对应的云盘数据盘也不会制作快照，并使用新指定的参数进行替换。
 (Optional)
 */
func NewCreateImageRequestWithAllParams(
    regionId string,
    instanceId string,
    name string,
    description *string,
    dataDisks []vm.InstanceDiskAttachmentSpec,
) *CreateImageRequest {

    return &CreateImageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:createImage",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Name: name,
        Description: description,
        DataDisks: dataDisks,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateImageRequestWithoutParam() *CreateImageRequest {

    return &CreateImageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:createImage",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *CreateImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID。(Required) */
func (r *CreateImageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param name: 镜像名称，长度为2\~32个字符，只允许中文、数字、大小写字母、英文下划线（\_）、连字符（-）及点（.）。
(Required) */
func (r *CreateImageRequest) SetName(name string) {
    r.Name = name
}

/* param description: 镜像描述。256字符以内。
(Optional) */
func (r *CreateImageRequest) SetDescription(description string) {
    r.Description = &description
}

/* param dataDisks: 数据盘列表。
在不指定该参数的情况下，制作镜像的过程中会将该实例中的所有云盘数据盘制作快照，并与系统盘一起，制作成打包镜像。
如果不希望将实例中的某个云盘数据盘制作快照，可使用 `noDevice` 的方式排除，例如：`deviceName=vdb`、`noDevice=true` 就不会将 `vdb` 制作快照。
如果希望在打包镜像中插入一块新盘，该盘不在实例中，可通过指定新的 `deviceName` 的方式实现，例如：`deviceName=vdx` 将会在打包镜像中插入一块盘符为 `vdx` 的新盘，支持新盘使用或不使用快照都可以。
如果使用 `deviceName` 指定了与实例中相同的盘符，那么实例中对应的云盘数据盘也不会制作快照，并使用新指定的参数进行替换。
(Optional) */
func (r *CreateImageRequest) SetDataDisks(dataDisks []vm.InstanceDiskAttachmentSpec) {
    r.DataDisks = dataDisks
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateImageRequest) GetRegionId() string {
    return r.RegionId
}

type CreateImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateImageResult `json:"result"`
}

type CreateImageResult struct {
    ImageId string `json:"imageId"`
}