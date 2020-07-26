// Copyright 2019 Andrew Merenbach
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Portions may be copyright Amazon Web Services.
// Adapted from https://aws.amazon.com/blogs/compute/implementing-safe-aws-lambda-deployments-with-aws-codedeploy/

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	this "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Request is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Request events.APIGatewayProxyRequest

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event codedeploy.PutLifecycleEventHookExecutionStatusInput) error {
	log.Println("Entering PreTraffic Hook!")

	// Read the DeploymentId & LifecycleEventHookExecutionId from the event payload
	deploymentID := event.DeploymentId
	lifecycleEventHookExecutionID := event.LifecycleEventHookExecutionId

	functionToTest := os.Getenv("NewVersion")
	log.Println("Testing new function version:", functionToTest)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	lambdaParams := &lambda.InvokeInput{
		FunctionName:   aws.String(functionToTest),
		InvocationType: aws.String(lambda.InvocationTypeRequestResponse),
	}

	lambdaResult := "Failed"

	client := lambda.New(sess)
	// result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("MyGetItemsFunction"), Payload: payload})
	output, err := client.Invoke(lambdaParams)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check the response for valid results
	// The response will be a JSON payload with statusCode and body properties. ie:
	// {
	//		"statusCode": 200,
	//		"body": 51
	// }

	var result Response
	if err := json.Unmarshal(output.Payload, &result); err != nil {
		return err
	}
	log.Println("Result:", result)

	lambdaResult = "Succeeded"

	// re := regexp.MustCompile(`^output$`)
	// if re.MatchString(result.Body) {
	// 	lambdaResult = "Succeeded"
	// 	log.Println("Validation testing succeeded!")
	// } else {
	// 	lambdaResult = "Failed"
	// 	log.Println("Validation testing failed!")
	// }

	// Complete the PreTraffic Hook by sending CodeDeploy the validation status
	var params = &codedeploy.PutLifecycleEventHookExecutionStatusInput{
		DeploymentId:                  deploymentID,
		LifecycleEventHookExecutionId: lifecycleEventHookExecutionID,
		Status:                        aws.String(lambdaResult), // status can be 'Succeeded' or 'Failed'
	}

	// Pass AWS CodeDeploy the prepared validation test results.
	codedeployClient := codedeploy.New(sess)
	req, err := codedeployClient.PutLifecycleEventHookExecutionStatus(params)
	if err != nil {
		return err
	}
	log.Println("Execution ID:", req.LifecycleEventHookExecutionId)

	return nil
}

func main() {
	this.Start(Handler)
}
