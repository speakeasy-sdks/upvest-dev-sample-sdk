# ListUserAccountsRequest


## Fields

| Field                                                                                                                            | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      | Example                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `Limit`                                                                                                                          | **int*                                                                                                                           | :heavy_minus_sign:                                                                                                               | Use the `limit` argument to specify the maximum number of items returned.                                                        |                                                                                                                                  |
| `Offset`                                                                                                                         | **int*                                                                                                                           | :heavy_minus_sign:                                                                                                               | Use the `offset` argument to specify where in the list of results to start when returning items for a particular query.          |                                                                                                                                  |
| `Order`                                                                                                                          | [*shared.Order](../../models/shared/order.md)                                                                                    | :heavy_minus_sign:                                                                                                               | Sort order of the result list if the `sort` parameter is specified. Use `ASC` for ascending or `DESC` for descending sort order. |                                                                                                                                  |
| `Signature`                                                                                                                      | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-http-header                           |                                                                                                                                  |
| `SignatureInput`                                                                                                                 | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | https://tools.ietf.org/id/draft-ietf-httpbis-message-signatures-01.html#name-the-signature-input-http-he                         |                                                                                                                                  |
| `Sort`                                                                                                                           | [*ListUserAccountsSort](../../models/operations/listuseraccountssort.md)                                                         | :heavy_minus_sign:                                                                                                               | Sort the result by `created_at`.                                                                                                 |                                                                                                                                  |
| `UpvestAPIVersion`                                                                                                               | [*shared.APIVersion](../../models/shared/apiversion.md)                                                                          | :heavy_minus_sign:                                                                                                               | Upvest API version (Note: Do not include quotation marks)                                                                        | 1                                                                                                                                |
| `UpvestClientID`                                                                                                                 | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | Tenant Client ID                                                                                                                 | ebabcf4d-61c3-4942-875c-e265a7c2d062                                                                                             |
| `UserID`                                                                                                                         | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |                                                                                                                                  |