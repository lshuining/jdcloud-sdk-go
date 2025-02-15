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

type UpdateCategoryRequest struct {

    core.JDCloudRequest

    /* 分类ID  */
    CategoryId int `json:"categoryId"`

    /* 分类名称 (Optional) */
    Name *string `json:"name"`

    /* 父分类ID，取值为 0 或 null 时，表示该分类为一级分类
 (Optional) */
    ParentId *int64 `json:"parentId"`

    /* 分类描述信息 (Optional) */
    Description *string `json:"description"`
}

/*
 * param categoryId: 分类ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCategoryRequest(
    categoryId int,
) *UpdateCategoryRequest {

	return &UpdateCategoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/categories/{categoryId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        CategoryId: categoryId,
	}
}

/*
 * param categoryId: 分类ID (Required)
 * param name: 分类名称 (Optional)
 * param parentId: 父分类ID，取值为 0 或 null 时，表示该分类为一级分类
 (Optional)
 * param description: 分类描述信息 (Optional)
 */
func NewUpdateCategoryRequestWithAllParams(
    categoryId int,
    name *string,
    parentId *int64,
    description *string,
) *UpdateCategoryRequest {

    return &UpdateCategoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/categories/{categoryId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        CategoryId: categoryId,
        Name: name,
        ParentId: parentId,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCategoryRequestWithoutParam() *UpdateCategoryRequest {

    return &UpdateCategoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/categories/{categoryId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param categoryId: 分类ID(Required) */
func (r *UpdateCategoryRequest) SetCategoryId(categoryId int) {
    r.CategoryId = categoryId
}

/* param name: 分类名称(Optional) */
func (r *UpdateCategoryRequest) SetName(name string) {
    r.Name = &name
}

/* param parentId: 父分类ID，取值为 0 或 null 时，表示该分类为一级分类
(Optional) */
func (r *UpdateCategoryRequest) SetParentId(parentId int64) {
    r.ParentId = &parentId
}

/* param description: 分类描述信息(Optional) */
func (r *UpdateCategoryRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCategoryRequest) GetRegionId() string {
    return ""
}

type UpdateCategoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCategoryResult `json:"result"`
}

type UpdateCategoryResult struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Level int `json:"level"`
    ParentId int64 `json:"parentId"`
    Description string `json:"description"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}