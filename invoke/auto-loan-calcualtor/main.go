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
	monthlyInterest := interestRate / 12
	// (Total Amount)(Monthy Interest)(1+monthy interest)^months
	topHalfOfFunction := float64(costToAmort) * (monthlyInterest) * math.Pow(1+monthlyInterest, float64(request.DurationMonths))
	// (1 + monthly interest)^months - 1
	bottomHalfOfFunction := math.Pow(1+monthlyInterest, float64(request.DurationMonths)) - 1
	monthlyCost := math.Floor(100*topHalfOfFunction/bottomHalfOfFunction) / 100
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            fmt.Sprintf("$%.2f", monthlyCost),
		Headers: map[string]string{
			"Content-Type":           "text/plain",
			"X-MyCompany-Func-Reply": "auto-load-calculator",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
