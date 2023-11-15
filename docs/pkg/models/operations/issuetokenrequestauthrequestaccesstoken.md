# IssueTokenRequestAuthRequestAccessToken

Schema for access token request.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ClientID`                                                       | *string*                                                         | :heavy_check_mark:                                               | Client ID given during onboarding.                               |
| `ClientSecret`                                                   | *string*                                                         | :heavy_check_mark:                                               | Client Secret given during onboarding.                           |
| `GrantType`                                                      | **string*                                                        | :heavy_minus_sign:                                               | This must always be `client_credentials`.                        |
| `Scope`                                                          | *string*                                                         | :heavy_check_mark:                                               | List of space delimited scopes to request for this access token. |