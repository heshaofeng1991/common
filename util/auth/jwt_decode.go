/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    jwt_decode
	@Date    2022/4/19 11:02
	@Desc
*/

package auth

import (
	"net/http"

	internal "github.com/heshaofeng1991/common"
	"github.com/pkg/errors"
)

// ParseToken 解析token.
func ParseToken(tokenString, secret string) (*WMSClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &WMSClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, internal.NewError("invalid signature method", http.StatusUnauthorized)
		}

		return []byte(secret), nil
	})

	if token == nil {
		return nil, internal.NewError("parse token failed", http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(*WMSClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.Wrap(err, "")
}

// ParseOMSToken 解析oms token.
func ParseOMSToken(tokenString, secret string) (*OMSClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &OMSClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, internal.NewError("invalid signature method", http.StatusUnauthorized)
		}

		return []byte(secret), nil
	})

	if token == nil {
		return nil, internal.NewError("parse token failed", http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(*OMSClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.Wrap(err, "")
}
