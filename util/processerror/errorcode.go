/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    errorcode
	@Date    2022/3/6 8:30 下午
	@Desc
*/

package processerror

// 服务器内部错误码（服务本身错误，系统错误，数据库错误等等）[1-10000).
const (
	ErrParseTokenCode   = 100
	ErrInvalidTokenCode = 101
)

/*
	错误码           错误码说明
	0 (default)     正常返回码
	其余区间         正常http返回码，其它预留
	[1, 100)        系统内部错误
	[1000-10000)    组件错误码
	[10000,20000)   订单服务错误码
	[20000,30000)   产品服务错误码
	[30000,40000)   物流服务错误码
	[40000,50000)   文件服务错误码
	[50000,60000)   运费服务错误码
*/

// SystemError 系统内部错误码.
const (
	SystemError = iota + 1
)

// PluginError 组件错误码.
const (
	PluginError                = 1000
	IntegrationExistStoreError = 1001
	// 数据库错误码.
	// redis错误码.
	// mem cache错误码.
	// 其它组件库错误码.
	// ...
)

// http 状态码 goa.design/goa/v3/dsl.

// TrackLengthErrorCode 轨迹服务错误码.
const (
	TrackLengthErrorCode = 30000
)

// 订单服务错误码(入库单，出库单，取件单).
const (
	CodeNoShippingAddress  int32 = 10101
	CodeNoShippingCountry  int32 = 10102
	CodeNoShippingProvince int32 = 10103
	CodeNoShippingCity     int32 = 10104
	CodeEmptyProduct       int32 = 10105
	CodeNoWarehouse        int32 = 10107
	CodeZeroDeclareValue   int32 = 10112

	CodeNoStockIn         int32 = 10200
	CodeStockOut          int32 = 10202
	CodeQuickTronStockOut int32 = 10203
	CodeInternalStockOut  int32 = 10204
	CodeProductNoMapping  int32 = 10205

	CodeTenantNotFound            int32 = 10301
	CodeTenantInsufficientBalance int32 = 10302
)

const (
	HoldTypeBaseInfo     = "baseInfo"
	HoldTypeShippingInfo = "shippingInfo"
	HoldTypeInventory    = "inventory"
	HoldTypeNoBalance    = "noBalance"
	HoldTypeOther        = "other"
)

var HoldTypeMap = map[string][]int32{
	HoldTypeBaseInfo:     {CodeNoWarehouse, CodeZeroDeclareValue, CodeEmptyProduct, CodeTenantNotFound},
	HoldTypeShippingInfo: {CodeNoShippingAddress, CodeNoShippingCountry, CodeNoShippingProvince, CodeNoShippingCity},
	HoldTypeInventory:    {CodeStockOut, CodeQuickTronStockOut, CodeInternalStockOut},
	HoldTypeNoBalance:    {CodeTenantInsufficientBalance},
}

// 文件服务错误码.
// 产品服务错误码(product+barcode).
// 渠道服务错误码.
// 运费服务错误码.

// JD 服务错误码.
const (
	JDServerInnerCode                     = 500
	JDServerBusyCode                      = 503
	JDInterfaceExceptionCode              = 504
	JDNULLAppKeyCode                      = 100001
	JDInvalidAppKeyCode                   = 100002
	JDForbidAppKeyCode                    = 100003
	JDUnAuthenticationApplicationCode     = 100004
	JDExpiredAuthTokenCode                = 100005
	JDEmptyTokenCode                      = 100006
	JDInvalidTokenCode                    = 100007
	JDEmptyMethodCode                     = 100008
	JDEmptyMethodVersionCode              = 100009
	JDInvalidMethodCode                   = 100010
	JDUnAuthenticationMethodCode          = 100011
	JDForbidMethodCode                    = 100012
	JDEmptyFormatCode                     = 100013
	JDInvalidFormatCode                   = 100014
	JDInconsistentDataFormatCode          = 100015
	JDInvalidSignCode                     = 100016
	JDEmptyTimestampCode                  = 100017
	JDErrorFormatTimestampCode            = 100018
	JDOverTenMinutesTimestampCode         = 100019
	JDCallApplicationOverLimitCode        = 100020
	JDCallMethodOverLimitCode             = 100021
	JDCallMethodFrequentlyCode            = 100022
	JDErrorAuthenticationCode             = 200001
	JDEmptyAppKeyCode                     = 200002
	JDEmptyAppSecretCode                  = 200003
	JDNoExistedAppKeyCode                 = 200004
	JDErrorAppSecretCode                  = 200005
	JDForbidApplicationAuthenticationCode = 200006
	JDEmptyRefreshTokenCode               = 200007
	JDErrorRefreshTokenCode               = 200008
	JDExpiredRefreshTokenCode             = 200009
	JDExpiredBusinessTokenCode            = 2000010
	JDRefreshTokenFrequentlyCode          = 200011
)

const (
	UploadFileToAwsError       = 400001
	ReadFileFromExcelError     = 400002
	ReadFileFromExcelRowsError = 400003
	FileRowsCountError         = 400004
	WriteFileToExcelError      = 400005
)
