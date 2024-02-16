package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/modelos"
)

func CreateNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	var note modelos.UserNote
	err := json.Unmarshal([]byte(request.Body), &note)

	if err != nil {
		errMsg := fmt.Sprintf("THE BODY %s, Error unmarshaling request body: %s", request.Body, err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errMsg}, nil
	}

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
func GetNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	fmt.Println(ctx)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(request.Body)}, nil

}
func DeleteNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	fmt.Println(ctx)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(request.Body)}, nil

}

func UnhandledMethod(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	fmt.Println(ctx)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(request.Body)}, nil

}
