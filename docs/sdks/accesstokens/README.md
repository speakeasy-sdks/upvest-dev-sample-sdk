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
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/shared"
	upvestdevsamplesdk "github.com/speakeasy-sdks/upvest-dev-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/upvest-dev-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := upvestdevsamplesdk.New(
        upvestdevsamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.AccessTokens.IssueToken(ctx, operations.IssueTokenRequest{
        RequestBody: &operations.IssueTokenRequestAuthRequestAccessToken{
            ClientID: "66f33cc6-ccf4-4562-8f8d-7c9213d11eda",
            ClientSecret: "string",
            Scope: "string",
        },
        Signature: "string",
        SignatureInput: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.IssueTokenRequest](../../pkg/models/operations/issuetokenrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.IssueTokenResponse](../../pkg/models/operations/issuetokenresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.IssueTokenError                        | 400                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensError            | 401                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensResponseError    | 403                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensResponse406Error | 406                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensResponse429Error | 429                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensResponse500Error | 500                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensResponse503Error | 503                                              | application/problem+json                         |
| sdkerrors.IssueTokenAccessTokensResponse504Error | 504                                              | application/problem+json                         |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |
