<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := upvestdevsamplesdk.New(
		upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
	)

	ctx := context.Background()
	res, err := s.Accounts.CreateAccount(ctx, operations.CreateAccountRequest{
		RequestBody: &operations.CreateAccountRequestBody{
			AccountGroupID: "e9562292-f304-4c6a-8db0-ea541f32fba9",
			Type:           operations.CreateAccountRequestBodyTypeTrading,
			UserID:         "d04cd2d5-ae02-4bb1-9118-75a95a0f2373",
		},
		IdempotencyKey:   "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
		Signature:        "string",
		SignatureInput:   "string",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Account != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->