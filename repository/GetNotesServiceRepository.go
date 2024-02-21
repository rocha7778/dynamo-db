package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/variables"
)

type GetNotesServiceRepository struct{}

func (*GetNotesServiceRepository) Scam() (*dynamodb.ScanOutput, error) {

	result, err := db.DBClient().Scan(&dynamodb.ScanInput{
		TableName: aws.String(variables.TableName),
	})
	return result, err
}
