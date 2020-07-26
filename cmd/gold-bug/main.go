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

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/merenbach/gold-bug/internal/api"
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
type Response = events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req Request) (events.APIGatewayProxyResponse, error) {
	var (
		out string
		err error
	)

	if c, ok := req.PathParameters["cipher"]; ok {
		switch c {
		case "affine":
			out, err = api.Affine(req.Body)
		case "atbash":
			out, err = api.Atbash(req.Body)
		case "caesar":
			out, err = api.Caesar(req.Body)
		case "decimation":
			out, err = api.Decimation(req.Body)
		case "keyword":
			out, err = api.Keyword(req.Body)
		case "rot13":
			out, err = api.Rot13(req.Body)
		case "vigenere":
			out, err = api.Vigenere(req.Body)
		case "beaufort":
			out, err = api.Beaufort(req.Body)
		case "dellaporta":
			out, err = api.DellaPorta(req.Body)
		case "gronsfeld":
			out, err = api.Gronsfeld(req.Body)
		case "trithemius":
			out, err = api.Trithemius(req.Body)
		case "variantbeaufort":
			out, err = api.VariantBeaufort(req.Body)
		default:
			return Response{StatusCode: http.StatusNotFound}, nil
		}
	}

	log.Printf("finished processing, any error is: %+v", err)
	mr := struct {
		Message string `json:"message"`
		Error   error  `json:"error"`
	}{out, err}

	bb, err := json.Marshal(mr)
	if err != nil {
		// TODO: don't return internal server error if we can avoid it
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	var buf bytes.Buffer
	if _, err := buf.Write(bb); err != nil {
		// TODO: don't return internal server error if we can avoid it
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	resp := Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Cache-Control":             "no-cache",
			"Content-Security-Policy":   "default-src 'none'",
			"Content-Type":              "application/json",
			"Referrer-Policy":           "no-referrer",
			"Strict-Transport-Security": "max-age=63072000; includeSubDomains; preload",
			"X-Content-Type-Options":    "nosniff",
			"X-Frame-Options":           "deny",
			"X-XSS-Protection":          "1; mode=block",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
