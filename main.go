package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mleyb/divisionloadout-api/router"

	gmux "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	log "github.com/sirupsen/logrus"
)

var gmuxLambda *gmux.GorillaMuxAdapter

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	log.Infoln("Lambda cold start - initialising router")
	r := router.New()
	gmuxLambda = gmux.New(r)
}

// Handler is the lambda handler function
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return gmuxLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
