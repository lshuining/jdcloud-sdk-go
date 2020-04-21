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


type CreateOrderResponseObject struct {

    /* 消息 (Optional) */
    Message string `json:"message"`

    /* 订单号 (Optional) */
    OrderNumber string `json:"orderNumber"`

    /* 应付金额（订单总金额-折扣金额） (Optional) */
    ActualFee int `json:"actualFee"`

    /* 是否成功 (Optional) */
    Success bool `json:"success"`

    /* 资源简单列表 (Optional) */
    OrderResourceInfoList []OrderResourceInfo `json:"orderResourceInfoList"`
}
