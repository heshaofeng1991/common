/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    s3
	@Date    2022/4/19 10:19
	@Desc
*/

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/heshaofeng1991/common/util/env"
	"github.com/pkg/errors"
)

type Session struct {
	Session *session.Session
}

// NewS3Session for init s3 session.
func NewS3Session(region string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			env.AwsAccessKeyID,
			env.AwsSecretAccessKey,
			""),
	})

	return sess, errors.Wrap(err, "")
}
