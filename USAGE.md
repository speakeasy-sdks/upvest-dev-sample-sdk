<!-- Start SDK Example Usage [usage] -->
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
	s := upvestdevsamplesdk.New()

	operationSecurity := operations.CreateAccountSecurity{
		OauthClientCredentials: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
	}

	ctx := context.Background()
	res, err := s.Accounts.CreateAccount(ctx, operations.CreateAccountRequest{
		IdempotencyKey:   "ccb07f42-4104-44ad-8e1f-c660bb7b269c",
		Signature:        "<value>",
		SignatureInput:   "<value>",
		UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
		UpvestClientID:   "ebabcf4d-61c3-4942-875c-e265a7c2d062",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->