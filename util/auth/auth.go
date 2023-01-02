/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    auth
	@Date    2022/4/19 11:03
	@Desc
*/

package auth

import (
	"net/http"

	internal "github.com/heshaofeng1991/common"
	"github.com/heshaofeng1991/common/util/env"
	"github.com/pkg/errors"
)

// Authenticate 鉴权 JWT token.
func Authenticate(token string) (id int32, err error) {
	// JWT Secret 存储在环境变量里面的密钥.
	jwtSecret := env.JwtSecret

	if jwtSecret == "" {
		jwtSecret = "wms"
	}

	clm, err := ParseToken(token, jwtSecret)
	if err != nil || clm == nil {
		return 0, internal.NewError("failed to parse token", http.StatusUnauthorized)
	}

	return clm.ID, errors.Wrap(err, "")
}

// OMSAuthenticate 鉴权 JWT token.
func OMSAuthenticate(token string) (id int64, tenantID int64, err error) {
	// JWT Secret 存储在环境变量里面的密钥.
	jwtSecret := env.JwtSecret

	if jwtSecret == "" {
		jwtSecret = "wms"
	}

	clm, err := ParseOMSToken(token, jwtSecret)
	if err != nil || clm == nil {
		return 0, 0, internal.NewError("failed to parse token", http.StatusUnauthorized)
	}

	return clm.ID, clm.TenantID, errors.Wrap(err, "")
}
