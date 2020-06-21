package data

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const buildTableName = "Build"

var db *dynamodb.DynamoDB

func init() {
	sess := session.Must(session.NewSession())

	db = dynamodb.New(sess)
}
