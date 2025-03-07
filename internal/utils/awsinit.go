package utils

import (
	"code-execution-sandbox/internal/config"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var awsSession *session.Session
var err error

func init() {
	config.LoadEnv()

	awsSession, err = session.NewSession(&aws.Config{
		Region: aws.String(config.GetEnv("AWS_REGION", "")),
	})

	if err != nil {
		log.Println("Error while loading aws session manager " + err.Error())
	}
}

func GetS3Client() (*s3.S3, error) {
	if err != nil {
		return nil, err
	}

	return s3.New(awsSession), nil
}
