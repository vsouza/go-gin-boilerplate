package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/vsouza/go-gin-boilerplate/config"
)

var db *dynamodb.DynamoDB

func Init() {
	c := config.GetConfig()
	db = dynamodb.New(session.New(&aws.Config{
		Region:      aws.String(c.GetString("db.region")),
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    aws.String(c.GetString("db.endpoint")),
		DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),
	}))
}

func GetDB() *dynamodb.DynamoDB {
	return db
}
