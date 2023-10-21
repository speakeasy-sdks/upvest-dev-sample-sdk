# CreateAccountLiquidationAccountLiquidationAccountLiquidation


## Fields

| Field                                                                                                                                                               | Type                                                                                                                                                                | Required                                                                                                                                                            | Description                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                                                                | *string*                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                  | N/A                                                                                                                                                                 |
| `Side`                                                                                                                                                              | **string*                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                  | Side of the order.<br/>* SELL -                                                                                                                                     |
| `Status`                                                                                                                                                            | [CreateAccountLiquidationAccountLiquidationAccountLiquidationStatus](../../models/operations/createaccountliquidationaccountliquidationaccountliquidationstatus.md) | :heavy_check_mark:                                                                                                                                                  | Execution status of the Account liquidation order.<br/>* NEW - <br/>* PROCESSING - <br/>* FILLED - <br/>* CANCELLED -                                               |