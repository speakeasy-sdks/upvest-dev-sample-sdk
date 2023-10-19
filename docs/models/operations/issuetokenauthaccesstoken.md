# IssueTokenAuthAccessToken

Schema for an access token response.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AccessToken`                                                   | *string*                                                        | :heavy_check_mark:                                              | The generated access token.                                     |
| `ExpiresIn`                                                     | *int64*                                                         | :heavy_check_mark:                                              | How many seconds the access token is valid for.                 |
| `Scope`                                                         | *string*                                                        | :heavy_check_mark:                                              | List of space delimited scopes requested for this access token. |
| `TokenType`                                                     | **string*                                                       | :heavy_minus_sign:                                              | This is always 'bearer'.                                        |