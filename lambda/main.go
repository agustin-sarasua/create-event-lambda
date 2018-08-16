package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var infoLogger = log.New(os.Stdout, "INFO ", log.Llongfile)

func getClaimsSub(req events.APIGatewayProxyRequest) string {
	jc, _ := json.Marshal(req.RequestContext.Authorizer)
	r := make(map[string]map[string]interface{})
	err := json.Unmarshal([]byte(jc), &r)
	if err != nil {
		fmt.Printf("Something went wrong %v", err)
	}
	fmt.Printf("Printing sub: %s ", r["claims"]["sub"].(string))
	return r["claims"]["sub"].(string)
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	e := new(Event)
	e.Sub = getClaimsSub(req)
	if req.Headers["content-type"] != "application/json" {
		fmt.Print(req.Headers)
		fmt.Printf("Unsopported content-type %s", req.Headers["Content-Type"])
		return clientError(http.StatusNotAcceptable)
	}

	err := json.Unmarshal([]byte(req.Body), e)
	if err != nil {
		fmt.Print("Could not unmarshal body")
		return clientError(http.StatusUnprocessableEntity)
	}

	if e.Name == "" {
		fmt.Printf("Name not found")
		return clientError(http.StatusBadRequest)
	}

	err = putItem(e)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/books?isbn=%s", e.Name)},
	}, nil
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//lc, _ := lambdacontext.FromContext(ctx)
	//log.Print(lc.Identity.CognitoIdentityID)

	// ctxJson, _ := json.Marshal(request)
	// fmt.Println(string(ctxJson))
	switch request.HTTPMethod {
	case "GET":
		//return show(req)
		fmt.Print("GET - Method not supported")
		return clientError(http.StatusMethodNotAllowed)
	case "POST":
		return create(request)
	default:
		fmt.Print("Method not supported")
		return clientError(http.StatusMethodNotAllowed)
	}
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
