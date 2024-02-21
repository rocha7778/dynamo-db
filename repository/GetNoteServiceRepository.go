package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/variables"
)

type GetNoteServiceRepository struct{}

func (*GetNoteServiceRepository) GetItem(noteID string) (*dynamodb.GetItemOutput, error) {
	result, err := db.DBClient().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(variables.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(noteID)},
		},
	})

	return result, err
}
