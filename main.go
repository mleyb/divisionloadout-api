package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/mleyb/divisionloadout-api/router"
)

var gmuxLambda *gorillamux.GorillaMuxAdapter

func init() {
	log.Println("Lambda cold start - initialising router")
	r := router.New()
	gmuxLambda = gorillamux.New(r)
}

// Handler is the lambda handler function
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return gmuxLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
