package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

var gorillaLambda *gorillamux.GorillaMuxAdapter

func init() {

	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{From: "gorilla", Message: time.Now().Format(time.UnixDate)})
	})

	gorillaLambda = gorillamux.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r, err := gorillaLambda.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&req))
	return *r.Version1(), err
}

func main() {
	lambda.Start(Handler)
}

type Response struct {
	From    string `json:"from"`
	Message string `json:"message"`
}
