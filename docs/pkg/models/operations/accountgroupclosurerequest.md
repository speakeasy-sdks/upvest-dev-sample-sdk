# AccountGroupClosureRequest


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `AccountGroupID`                                                                                         | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `Signature`                                                                                              | *string*                                                                                                 | :heavy_check_mark:                                                                                       | https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header   |                                                                                                          |
| `SignatureInput`                                                                                         | *string*                                                                                                 | :heavy_check_mark:                                                                                       | https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he |                                                                                                          |
| `UpvestAPIVersion`                                                                                       | [*shared.APIVersion](../../../pkg/models/shared/apiversion.md)                                           | :heavy_minus_sign:                                                                                       | Upvest API version (Note: Do not include quotation marks)                                                | 1                                                                                                        |
| `UpvestClientID`                                                                                         | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Tenant Client ID                                                                                         | ebabcf4d-61c3-4942-875c-e265a7c2d062                                                                     |