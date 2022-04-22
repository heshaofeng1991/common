/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    jwt_encode
	@Date    2022/4/19 11:02
	@Desc
*/

package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type MyClaims struct {
	ID int32 `json:"id"`
}

func (m MyClaims) Valid() error {
	return nil
}

func GenerateToken(userID int32, secret string, method jwt.SigningMethod) (string, error) {
	claims := &MyClaims{
		ID: userID,
	}

	token, err := jwt.NewWithClaims(method, claims).SignedString([]byte(secret))

	return token, errors.Wrap(err, "")
}
