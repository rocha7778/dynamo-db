package user

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

func FindByEmail(email string) (*modelos.User, error) {

	// Get the item from DynamoDB
	result, err := db.DBClient().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(variables.TableUser),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {S: aws.String(email)},
		},
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal the item into a note struct
	var user modelos.User
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func Save(body string) (*modelos.User, error) {

	var user modelos.User
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		return nil, err
	}

	err = user.ValidEmail()

	if err != nil {
		return nil, err
	}

	// Put the item into DynamoDB
	_, err = db.DBClient().PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(variables.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":        {S: aws.String(user.ID)},
			"name":      {S: aws.String(user.Name)},
			"last_name": {S: aws.String(user.LastName)},
			"birthday":  {S: aws.String(user.Birthday.String())},
			"emeail":    {S: aws.String(user.Email)},
			"password":  {S: aws.String(user.Password)},
		},
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
