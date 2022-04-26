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

type WMSClaims struct {
	ID int32 `json:"id"`
}

func (m WMSClaims) Valid() error {
	return nil
}

func GenerateToken(userID int32,  secret string, method jwt.SigningMethod) (string, error) {
	claims := &WMSClaims{
		ID: userID,
	}

	token, err := jwt.NewWithClaims(method, claims).SignedString([]byte(secret))

	return token, errors.Wrap(err, "")
}

type OMSClaims struct {
  ID int64 `json:"id"`
  TenantID int64 `json:"tenant_id"`
}


func (m OMSClaims) Valid() error {
  return nil
}

func GenerateOMSToken(userID int64, tenantID int64, secret string, method jwt.SigningMethod) (string, error) {
  claims := &OMSClaims{
    ID: userID,
    TenantID: tenantID,
  }

  token, err := jwt.NewWithClaims(method, claims).SignedString([]byte(secret))

  return token, errors.Wrap(err, "")
}
