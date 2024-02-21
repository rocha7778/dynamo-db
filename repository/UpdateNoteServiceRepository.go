package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

type UpdateNoteServiceRepository struct{}

func (*UpdateNoteServiceRepository) UpdateItem(note *modelos.UserNote) error {

	_, err := db.DBClient().UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(variables.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(note.ID)},
		},
		UpdateExpression: aws.String("SET #text = :text"),
		ExpressionAttributeNames: map[string]*string{
			"#text": aws.String("text"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":text": {S: aws.String(note.Text)},
		},
		ReturnValues: aws.String("ALL_NEW"),
	})

	return err
}
