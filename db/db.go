package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamoDBClient *dynamodb.DynamoDB

func init() {
	// Initialize the DynamoDB client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoDBClient = dynamodb.New(sess)
}

// NewDynamoDBClient returns the initialized DynamoDB client
func DBClient() *dynamodb.DynamoDB {
	return dynamoDBClient
}
