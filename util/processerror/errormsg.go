/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    error
	@Date    2022/3/3 10:17 下午
	@Desc
*/

package processerror

import (
	"errors"
)

// ErrInner 内部错误信息.
var (
	ErrInner      = errors.New("inner error")
	ErrServer     = errors.New("server error")
	ErrStack      = errors.New("stack error")
	ErrHTTPStatus = errors.New("http request error")
)

// ErrEmptyToken 鉴权错误信息.
var (
	ErrEmptyToken        = errors.New("token is empty")
	ErrInvalidSignMethod = errors.New("invalid signature method")
	ErrInvalidToken      = errors.New("token is invalid")
)

// ErrOrderNotSupported 入库单错误信息.
var (
	ErrInboundOrderNotSupported      = errors.New("the order is not supported update")
	ErrInboundOrderLimitQuery        = errors.New("only supported query 50 order is allowed")
	ErrInboundTrackNumberExisted     = errors.New("track number already existed")
	ErrInboundCustomerOrderIDExisted = errors.New("customer order id already existed")
	ErrInboundMode                   = errors.New("orders of different modes do not support updating, " +
		"please change the mode or create a new order")
)

// ErrOutboundOrderNotSupported 出库单错误信息.
var (
	ErrOutboundOrderNotFound     = errors.New("the order not found")
	ErrOutboundOrderNotSupported = errors.New("the order is not supported update")
	ErrOutboundOrderLimitQuery   = errors.New("only supported query 50 order is allowed")
	ErrOutboundOrderExisted      = errors.New("the order already existed")
	ErrOutboundMode              = errors.New("orders of different modes do not support updating, " +
		"please change the mode or create a new order")
	ErrOutboundQty           = errors.New("invalid product qty")
	ErrOutboundDeleteItem    = errors.New("the order is not supported delete item")
	ErrOutboundUpdateItem    = errors.New("the order is not supported update item")
	ErrOutboundAddItem       = errors.New("the order is not supported add item")
	ErrOutboundUpdateItemQty = errors.New("the order is not supported update qty")
	ErrOutboundShip          = errors.New("the order is not supported ship")
	ErrOutboundResubmit      = errors.New("the order is not supported re-submit")
	ErrOutboundBackToReady   = errors.New("the order is not supported back to ready")
	ErrOutboundItemExist     = errors.New("the order item is exists")
	ErrOutboundShipValidate  = errors.New("Ship order verification failed")
	ErrOutboundStockOut      = errors.New("Insufficient Stock")
	ErrOutboundHasHold       = errors.New("the order still have some hold reasons")
	ErrOutboundQuote         = errors.New("Get quote failed")
	ErrInvalidTaxType        = errors.New("invalid tax type")
)

var (
	ErrInventoryNotFound = errors.New("Inventory Not Found")
	ErrInventoryRelease  = errors.New("Release Qty ganter then Prepare Ship Qty")
)

// ErrDB 数据库错误信息.
var (
	ErrDB = errors.New("dbr error")
)

// ErrPickupOrderAddress 取件单错误信息.
var (
	ErrPickupOrderAddress = errors.New("invalid pickup order address")
	ErrParsePickupOrder   = errors.New("parse pickup order address error")
)

// ErrFile 文件错误信息.
var (
	ErrFile     = errors.New("file error")
	ErrFileSize = errors.New("no supported image size")
)

// ErrProductBarcode 产品错误信息.
var (
	ErrProductBarcode = errors.New("don't need NSS Barcode required params of Barcode can't empty string")
)

// ErrInvalidTrackingNumber 轨迹错误信息.
var (
	ErrInvalidTrackingNumber = errors.New("invalid tracking number")
	ErrQueryTrackLimit       = errors.New("only supported query 50 records with tracking number")
	ErrOrderReceived         = errors.New("order received")
	ErrInvalidQueryOrderID   = errors.New("invalid query order id")
)

// ErrTenantCheckFailed 租户错误信息.
var (
	ErrTenantCheckFailed   = errors.New("check tenant failed")
	ErrTenantInvalidParams = errors.New("invalid params")
)

// ErrChannelNotSupported 渠道错误信息.
var (
	ErrChannelNotSupported = errors.New("not supported channel")
	ErrNoExistedChannels   = errors.New("no existed channels")
)

// // ErrNoAvailableChannels 运费试算错误信息.
// var (
// 	ErrNoAvailableChannels           = errors.New("no available channels")
// 	ErrNoAvailableChannelCostBatches = errors.New("no available channels cost batches")
// )

// ErrSnowFlakeNodeID 雪花算法错误码.
var (
	ErrSnowFlakeNodeID     = errors.New("invalid snowflake node id")
	ErrSnowFlakeDistrictID = errors.New("invalid snowflake district id")
	ErrSnowFlakeNodeNum    = errors.New("invalid snowflake node number")
	ErrSnowFlakeReject     = errors.New("reject generate id")
)

// ErrInvalidID .
var (
	ErrInvalidID = errors.New("invalid id")
)

// ErrBarcodeExisted .
var (
	ErrBarcodeExisted = errors.New("barcode already existed")
)

// ErrInvalidSku 产品错误信息.
var (
	ErrProductNotFound       = errors.New("product not found")
	ErrInvalidSku            = errors.New("invalid sku")
	ErrInvalidProductName    = errors.New("invalid product name")
	ErrInvalidDeclaredCnName = errors.New("invalid declared cn name")
	ErrInvalidDeclaredEnName = errors.New("invalid declared en name")
	ErrInvalidDeclaredInUsd  = errors.New("invalid declared value in usd")
	ErrInvalidHsCode         = errors.New("invalid HsCode")
	ErrInvalidProductBarcode = errors.New("invalid barcode")
	ErrInvalidProductWeight  = errors.New("invalid weight")
	ErrInvalidProductLength  = errors.New("invalid length")
	ErrInvalidProductWidth   = errors.New("invalid width")
	ErrInvalidProductHeight  = errors.New("invalid height")
	ErrProduct               = errors.New("product error")
)

// JD 错误信息.
const (
	JDServerInnerMsg                     = "服务器内部错误，请联系管理员"
	JDServerBusyMsg                      = "服务忙，请稍后再试"
	JDInterfaceExceptionMsg              = "访问 JSF 接口异常"
	JDNULLAppKeyMsg                      = "APP_KEY 为空"
	JDInvalidAppKeyMsg                   = "无效的 APP_KEY"
	JDForbidAppKeyMsg                    = "APP_KEY 禁用"
	JDUnAuthenticationApplicationMsg     = "应用未授权"
	JDExpiredAuthTokenMsg                = "AuthToken 已过期" //nolint:gosec
	JDEmptyTokenMsg                      = "Token为空"       //nolint:gosec
	JDInvalidTokenMsg                    = "无效的Token"      //nolint:gosec
	JDEmptyMethodMsg                     = "方法为空"
	JDEmptyMethodVersionMsg              = "方法版本号为空"
	JDInvalidMethodMsg                   = "无效的 Method"
	JDUnAuthenticationMethodMsg          = "方法未授权"
	JDForbidMethodMsg                    = "方法被禁用"
	JDEmptyFormatMsg                     = "format 为空"
	JDInvalidFormatMsg                   = "无效的 format"
	JDInconsistentDataFormatMsg          = "数据格式与 format 不一致"
	JDInvalidSignMsg                     = "无效的签名"
	JDEmptyTimestampMsg                  = "Timestamp 为空"
	JDErrorFormatTimestampMsg            = "Timestamp 格式错误"
	JDOverTenMinutesTimestampMsg         = "Timestamp 时差超过 10 分钟"
	JDCallApplicationOverLimitMsg        = "应用调用次数已达上限(${0}次/天)，请 24 小时后再试"
	JDCallMethodOverLimitMsg             = "方法调用次数已达上限(${0}次/天)，请 24 小时后再试"
	JDCallMethodFrequentlyMsg            = "方法调用太过频繁，请稍后再试"
	JDErrorAuthenticationMsg             = "授权类型错误"
	JDEmptyAppKeyMsg                     = "授权 AppKey 不能为空"
	JDEmptyAppSecretMsg                  = "授权应用秘钥不能为空" //nolint:gosec
	JDNoExistedAppKeyMsg                 = "授权 AppKey 不存在"
	JDErrorAppSecretMsg                  = "授权应用秘钥错误" //nolint:gosec
	JDForbidApplicationAuthenticationMsg = "授权应用被禁用"
	JDEmptyRefreshTokenMsg               = "刷新 Token 为空"  //nolint:gosec
	JDErrorRefreshTokenMsg               = "刷新 Token 错误"  //nolint:gosec
	JDExpiredRefreshTokenMsg             = "刷新 Token 已过期" //nolint:gosec
	JDExpiredBusinessTokenMsg            = "Business Token 已过期"
	JDRefreshTokenFrequentlyMsg          = "Token 刷新过于频繁" //nolint:gosec
)

// ErrJDQueryOrderLimit .
var (
	ErrJDQueryOrderLimit  = errors.New("query order length is not correct")
	ErrJDQueryOrderFailed = errors.New("query order failed")
	ErrRequestFailed      = errors.New("request failed")
	ErrEmptyPackage       = errors.New("package is empty")
	ErrEmptyWarehouse     = errors.New("ware is empty")
	ErrJDLimit            = errors.New("请求体超限")
)

// ErrSNSEmptyOrderShipped .
var (
	ErrSNSEmptyOrderShipped = errors.New("received empty order shipped message")
)

// ErrCovertType .
var (
	ErrCovertType   = errors.New("covert type failed")
	ErrCovertLength = errors.New("covert length is not a number")
)

// ErrSignupExistedEmail .
var (
	ErrSignupExistedEmail = errors.New("the email has been already registered")
	ErrLoginEmail         = errors.New("invalid login email")
	ErrLoginPassword      = errors.New("invalid login password")
)

// ErrStoreTenant .
var (
	ErrStoreTenant = errors.New("invalid store tenant id")
)

// ErrUserInvalidParams 用户错误信息.
var (
	ErrUserInvalidParams = errors.New("invalid params")
	ErrUserCheckFailed   = errors.New("check user failed")
)
