/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    sqs
	@Date    2022/4/19 10:19
	@Desc
*/

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/heshaofeng1991/common/util/env"
	"github.com/pkg/errors"
)

// SQSReceiveWaybillDetailMessage .
type SQSReceiveWaybillDetailMessage struct {
	Type        string                  `json:"type"`
	TrackDetail *ReceiveTrackDetailsReq `json:"trackDetail"`
}

// ReceiveTrackDetailsReq 接收运单信息明细.
type ReceiveTrackDetailsReq struct {
	FulfillOrderVO *ReceiveTrackDetailsModel `json:"fulfillOrderVO"` //nolint:tagliatelle
}

// ReceiveTrackDetailsModel .
type ReceiveTrackDetailsModel struct {
	OrderID          int64                      `json:"orderId"`          // 订单ID
	LogisticsChannel string                     `json:"logisticsChannel"` // 物流渠道（NSS）
	FulfillPackages  []*ReceiveFulfillPackageVO `json:"fulfillPackages"`  // 包裹明细，详见ReceiveFulfillPackageVO
}

// ReceiveFulfillPackageVO 包裹明细.
type ReceiveFulfillPackageVO struct {
	PackageID        string  `json:"packageId"`                  // 包裹ID
	ActualLength     float64 `json:"actualLength,omitempty"`     // 实际包裹长
	ActualWidth      float64 `json:"actualWidth,omitempty"`      // 实际包裹宽
	ActualHeight     float64 `json:"actualHeight,omitempty"`     // 实际包裹高
	ActualLengthUnit int     `json:"actualLengthUnit,omitempty"` // 实际包裹长度单位  1厘米，2米，3英尺
	ActualWeight     float64 `json:"actualWeight,omitempty"`     // 实际包裹重量
	CostWeight       float64 `json:"costWeight,omitempty"`       // 计费包裹重量
	ActualWeightUnit int     `json:"actualWeightUnit,omitempty"` // 实际包裹重量单位  1克，2千克，3磅
	BillURL          string  `json:"billUrl,omitempty"`          // 末公里面单URL
	TrackingNo       string  `json:"trackingNo,omitempty"`       // 运单号
	CarrierCode      string  `json:"carrierCode,omitempty"`      // 承运商Code
	CarrierName      string  `json:"carrierName,omitempty"`      // 承运商名称
}

func (s Session) SendTrackInfoToSQSEvent(msg string) error {
	_, err := sqs.New(s.Session).SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(env.AwsWmsBackendQueueURL),
		MessageBody: aws.String(msg),
	})

	return errors.Wrap(err, "")
}

// ReceiveTrackStatusReq 接收轨迹节点回传.
type ReceiveTrackStatusReq struct {
	Type                string                  `json:"type"`
	FulfillOrderTraceVO ReceiveTrackStatusModel `json:"fulfillOrderTraceVO"` //nolint:tagliatelle
}

type ReceiveTrackStatusModel struct {
	OrderID              int64                    `json:"orderId"`              // 订单ID
	LogisticsChannel     string                   `json:"logisticsChannel"`     // 物流渠道（NSS）
	FulfillPackageTraces []*FulfillPackageTraceVO `json:"fulfillPackageTraces"` // 包裹明细，详见FulfillPackageTraceVO
}

// FulfillPackageTraceVO 包裹明细.
type FulfillPackageTraceVO struct {
	PackageID   string `json:"packageId"`   // 包裹ID
	TraceStatus int    `json:"traceStatus"` // 轨迹状态：10-初始，20-待揽件，30-揽收入库成功，
	// 40-发货出库，50-发货国出关，60-到达目的国，70-配送，80-妥 投，90-投递失败，100-发生异常
	TraceOccurTime string `json:"traceOccurTime"` // 状态发生的时间 时间戳，格式为yyyy-MM-dd HH:mm:ss，例如：2011-06-16 13:23:30
	OutTraceURL    string `json:"outTraceUrl"`    // 去NSS查询的轨迹URL
}

// ReceiveTrackStatusRsp 接收轨迹节点回传.
type ReceiveTrackStatusRsp struct {
	Code    string `json:"code"`    // 错误码
	Success bool   `json:"success"` // 是否成功
	Message string `json:"message"` // 错误信息
	Data    int64  `json:"data"`    // 入参的订单ID
}
