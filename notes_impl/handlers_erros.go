package notes_impl

import (
	"github.com/aws/aws-lambda-go/events"
)

func handleError(msg string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{StatusCode: statusCode, Body: msg}
}
