package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	log "github.com/sirupsen/logrus"
)

// Delete deletes a build
func Delete(id string) error {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(buildTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	}

	resp, err := db.DeleteItem(params)
	if err != nil {
		return err
	}

	log.Debugln(resp)

	return nil
}
