package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/variables"
)

type DeleteServiceRepository struct {
}

func (*DeleteServiceRepository) DeleteItem(noteId string) error {
	_, err := db.DBClient().DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(variables.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: &noteId},
		},
	})
	return err
}
