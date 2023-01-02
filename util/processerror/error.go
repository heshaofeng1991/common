/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    error
	@Date    2022/3/6 8:31 下午
	@Desc
*/

package processerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	Err  error
	Msg  string
	Code int32
}

func (e Error) Error() string {
	return e.Msg
}

func (e Error) ErrorCode() int32 {
	return e.Code
}

func (e Error) Unwrap() error {
	return e.Err
}

// NewError .
func NewError(err error, msg string, code int32) *Error {
	return &Error{
		Err:  err,
		Msg:  msg,
		Code: code,
	}
}

func Exit(info interface{}) error {
	return fmt.Errorf("%w - %s", ErrServer, info)
}

func Stack(format string, args ...interface{}) error {
	exception := fmt.Sprintf(format, args) //nolint:govet

	return fmt.Errorf("%w - %s", ErrStack, exception)
}

func ProductSkuMapped(barcode, reqSku, reqBarcode string) error {
	info := fmt.Sprintf("%v mapped %v, create new sku for %v", reqSku, barcode, reqBarcode)

	err := fmt.Errorf("%w-%s", ErrProduct, info)

	return errors.Wrap(err, "")
}

func ProductNeedNewBarcode(sku, barcode string) error {
	info := fmt.Sprintf("%v existed, create new barcode for %v", barcode, sku)

	err := fmt.Errorf("%w-%s", ErrProduct, info)

	return errors.Wrap(err, "")
}

func BarcodeExisted(barcode string) error {
	info := fmt.Sprintf("the barcode(%v) has existed, please use new barcode", barcode)

	err := fmt.Errorf("%w-%s", ErrProduct, info)

	return errors.Wrap(err, "")
}

func HTTPStatus(url string, code int, postData interface{}) error {
	err := fmt.Errorf("%w : url=%v, postData-%v, statusCode=%v", ErrHTTPStatus, url, postData, code)

	return errors.Wrap(err, "")
}
