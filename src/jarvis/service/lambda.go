
// Contains code to encapsulate running lambda functions on AWS
package service

import (
  "encoding/json"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/lambda"
)

type Lambda struct {}

func (l Lambda) RunAsync(fName string, payload map[string]interface{}) error {
  b, _ := json.Marshal(payload)
  svc := lambda.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
  params := &lambda.InvokeInput{
    FunctionName: aws.String(fName),
    InvocationType: aws.String(lambda.InvocationTypeEvent),
    Payload: b,
  }
  _, err := svc.Invoke(params)
  return err
}
