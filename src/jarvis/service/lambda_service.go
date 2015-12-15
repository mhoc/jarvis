// Contains code to encapsulate running lambda functions on AWS
package service

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"jarvis/config"
)

type Lambda struct{}

func (l Lambda) RunAsync(fName string, payload map[string]interface{}) error {
	access, secret := config.AwsCredentials()
	sess := session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(access, secret, ""),
	})
	svc := lambda.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	b, _ := json.Marshal(payload)
	params := &lambda.InvokeInput{
		FunctionName:   aws.String(fName),
		InvocationType: aws.String(lambda.InvocationTypeEvent),
		Payload:        b,
	}
	_, err := svc.Invoke(params)
	return err
}
