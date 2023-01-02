/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    jwt_encode_test.go
	@Date    2022/4/20 12:30
	@Desc
*/

package auth_test

// import (
// 	"testing"
//
// 	"github.com/dgrijalva/jwt-go"
// )
//
// func TestGenerateToken(t *testing.T) {
// 	type args struct {
// 		userID int32
// 		secret string
// 		method jwt.SigningMethod
// 	}
//
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		// Add test cases.
// 		{
// 			name: "generate token",
// 			args: args{
// 				userID: 1,
// 				secret: "wms",
// 				method: jwt.SigningMethodHS256,
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, _ = GenerateToken(tt.args.userID, tt.args.secret, tt.args.method)
// 		})
// 	}
// }
