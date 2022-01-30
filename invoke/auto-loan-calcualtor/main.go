package main

import (
	"fmt"
	"math"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type AutoLoanCalcEventBody struct {
	DurationMonths    int
	TotalLoanAmount   int
	InterestInPercent float64
	DownPayment       int
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request AutoLoanCalcEventBody) (events.APIGatewayProxyResponse, error) {
	costToAmort := request.TotalLoanAmount - request.DownPayment
	interestRate := request.InterestInPercent / 100
	interestCost := math.Pow(1+interestRate, float64(request.DurationMonths)/12)
	totalLoanCost := float64(costToAmort) * interestCost
	monthlyCost := totalLoanCost / 12
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            fmt.Sprintf("%f", monthlyCost),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
