package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/modelos"
)

const tableName = "UserNote"

var dynamoDBClient *dynamodb.DynamoDB

func init() {
	// Create a new DynamoDB client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoDBClient = dynamodb.New(sess)
}

func createNoteHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var note modelos.UserNote
	err := json.Unmarshal([]byte(request.Body), &note)
	fmt.Println(request)
	log.SetPrefix("XXXXXXX: " + request.Body)

	if err != nil {
		errMsg := fmt.Sprintf("THE BODY %s, Error unmarshaling request body: %s", request.Body, err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errMsg}, nil
	}

	/*if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}*/

	if note.ID == "" || note.Text == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "ID and Text fields are required"}, nil
	}

	// Marshal the note struct into JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error marshaling JSON"}, nil
	}

	// Put the item into DynamoDB
	_, err = dynamoDBClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String(note.ID)},
			"text": {S: aws.String(note.Text)},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error saving note to DynamoDB"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(noteJSON)}, nil

}

func main() {
	lambda.Start(createNoteHandler)
}
