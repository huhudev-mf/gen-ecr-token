package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	region := os.Getenv("AWS_REGION")
	id := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECERT_ACCESS_KEY")

	svc := ecr.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			id,
			secret,
			"",
		),
	})))

	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		panic(err)
	}

	authorizationToken := *result.AuthorizationData[0].AuthorizationToken
	outputBytes, err := base64.StdEncoding.DecodeString(authorizationToken)
	output := string(outputBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
