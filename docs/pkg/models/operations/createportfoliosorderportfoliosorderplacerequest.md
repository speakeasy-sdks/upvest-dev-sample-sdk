# CreatePortfoliosOrderPortfoliosOrderPlaceRequest


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `AccountID`                                                                                                  | *string*                                                                                                     | :heavy_check_mark:                                                                                           | Account unique identifier.                                                                                   |
| `CashAmount`                                                                                                 | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `Currency`                                                                                                   | [*operations.CreatePortfoliosOrderCurrency](../../../pkg/models/operations/createportfoliosordercurrency.md) | :heavy_minus_sign:                                                                                           | Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.<br/>* EUR - Euro   |
| `PostTax`                                                                                                    | **bool*                                                                                                      | :heavy_minus_sign:                                                                                           | Cash amount is post-tax value                                                                                |
| `Side`                                                                                                       | [operations.CreatePortfoliosOrderSide](../../../pkg/models/operations/createportfoliosorderside.md)          | :heavy_check_mark:                                                                                           | Side of the portfolio order.<br/>* BUY - <br/>* SELL -                                                       |
| `UserID`                                                                                                     | *string*                                                                                                     | :heavy_check_mark:                                                                                           | User unique identifier.                                                                                      |