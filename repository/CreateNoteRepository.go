package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

type CreateNoteRepository struct{}

func (*CreateNoteRepository) PutItem(note *modelos.UserNote) error {
	// Put the item into DynamoDB
	_, err := db.DBClient().PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(variables.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String(note.ID)},
			"text": {S: aws.String(note.Text)},
		},
	})

	return err
}
