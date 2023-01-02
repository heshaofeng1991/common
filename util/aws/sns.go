/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    sns
	@Date    2022/5/4 11:49
	@Desc
*/

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/pkg/errors"
)

func (s Session) ListTopics() (*sns.ListTopicsOutput, error) {
	result, err := sns.New(s.Session).ListTopics(&sns.ListTopicsInput{
		// NextToken: aws.String(token),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

func (s Session) ListSubscriptions() (*sns.ListSubscriptionsOutput, error) {
	result, err := sns.New(s.Session).ListSubscriptions(&sns.ListSubscriptionsInput{
		// NextToken: aws.String(token),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

func (s Session) CreateSnsTopic(topic string) (*sns.CreateTopicOutput, error) {
	result, err := sns.New(s.Session).CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(topic),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

func (s Session) DeleteSnsTopic(topic string) (*sns.DeleteTopicOutput, error) {
	result, err := sns.New(s.Session).DeleteTopic(&sns.DeleteTopicInput{
		TopicArn: aws.String(topic),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

func (s Session) SnsPublish(message, topic string) (*sns.PublishOutput, error) {
	result, err := sns.New(s.Session).Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topic),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

func (s Session) SnsSubscribe(endpoint, protocol, topic string) (*sns.SubscribeOutput, error) {
	result, err := sns.New(s.Session).Subscribe(&sns.SubscribeInput{
		Endpoint:              aws.String(endpoint),
		Protocol:              aws.String(protocol),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              aws.String(topic),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

func (s Session) SnsUnsubscribe(arn string) (*sns.UnsubscribeOutput, error) {
	result, err := sns.New(s.Session).Unsubscribe(&sns.UnsubscribeInput{
		SubscriptionArn: aws.String(arn),
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
