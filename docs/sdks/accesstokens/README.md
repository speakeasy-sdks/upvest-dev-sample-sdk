# AccessTokens
(*AccessTokens*)

## Overview

All authentication related paths.

### Available Operations

* [IssueToken](#issuetoken) - Get an access token for requested scopes

## IssueToken

Get an access token to use with the API with specified scopes.

You should _always_ scope your access tokens. You get one for read-access and separate ones for updating, creating or deleting resources.

### Example Usage

```go
package main

import(
	"context"
	"log"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity("YOUR_TOKEN"),
    )

    ctx := context.Background()
    res, err := s.AccessTokens.IssueToken(ctx, operations.IssueTokenRequest{
        RequestBody: &operations.IssueTokenRequestAuthRequestAccessToken{
            ClientID: "66f33cc6-ccf4-4562-8f8d-7c9213d11eda",
            ClientSecret: "Montana",
            Scope: "Northeast",
        },
        Signature: "olive",
        SignatureInput: "Minivan",
        UpvestAPIVersion: shared.APIVersionOne.ToPointer(),
        UpvestClientID: "ebabcf4d-61c3-4942-875c-e265a7c2d062",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthAccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.IssueTokenRequest](../../models/operations/issuetokenrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.IssueTokenResponse](../../models/operations/issuetokenresponse.md), error**

