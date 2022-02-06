package main

import (
	"fmt"
	"math"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type CompoundInterestCalcEventBody struct {
	Years             int
	Principal         int
	MonthlyPayment    int
	ExpectedReturn    float64
	ExpectedInflation float64
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request CompoundInterestCalcEventBody) (Response, error) {
	effectiveInterestRate := (request.ExpectedReturn - request.ExpectedInflation) / 100
	principalGrowth := float64(request.Principal) * math.Pow(1+effectiveInterestRate, float64(request.Years))
	monthlyContributions := float64(request.MonthlyPayment) * 12 * ((math.Pow(1+effectiveInterestRate, float64(request.Years)) - 1) / effectiveInterestRate)
	totalAmount := math.Round(100*(monthlyContributions+principalGrowth)) / 100
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            fmt.Sprintf("$%.2f", totalAmount),
		Headers: map[string]string{
			"Content-Type":           "text/plain",
			"X-MyCompany-Func-Reply": "compound-interest-calculator",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
