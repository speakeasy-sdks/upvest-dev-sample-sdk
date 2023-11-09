# PlaceOrderOrder

The request for the order creation has been accepted for processing.


## Fields

| Field                                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccountID`                                                                                                                                                                                                                                                    | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | The ID of the account that owns the order                                                                                                                                                                                                                      |
| `CancellationReason`                                                                                                                                                                                                                                           | [*operations.CancellationReason](../../../pkg/models/operations/cancellationreason.md)                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                             | Reason for Order cancellation. The field is present in case the Order has a status of CANCELLED.<br/>* CANCELLED_BY_CLIENT - <br/>* CANCELLED_BY_UPVEST_OPERATIONS - <br/>* CANCELLED_BY_TRADING_PARTNER - <br/>* CANCELLED_BY_UPVEST_PLATFORM -               |
| `CashAmount`                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `ClientReference`                                                                                                                                                                                                                                              | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | An ID provided by the client                                                                                                                                                                                                                                   |
| `CreatedAt`                                                                                                                                                                                                                                                    | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                             | Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)                                                                    |
| `Currency`                                                                                                                                                                                                                                                     | [*operations.PlaceOrderOrdersCurrency](../../../pkg/models/operations/placeorderorderscurrency.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `ExecutionFlow`                                                                                                                                                                                                                                                | [*operations.PlaceOrderExecutionFlow](../../../pkg/models/operations/placeorderexecutionflow.md)                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                             | Execution flow that the order processing goes through. If no value is specified, the default value is assumed - `STRAIGHT_THROUGH`.<br/>* STRAIGHT_THROUGH - <br/>* BLOCK -                                                                                    |
| `Executions`                                                                                                                                                                                                                                                   | [][operations.OrderExecution](../../../pkg/models/operations/orderexecution.md)                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                             | Order executions associated with this order                                                                                                                                                                                                                    |
| `ExpiryDate`                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `Fee`                                                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `ID`                                                                                                                                                                                                                                                           | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `InitiationFlow`                                                                                                                                                                                                                                               | [operations.InitiationFlow](../../../pkg/models/operations/initiationflow.md)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                             | Initiation flow used during order creation, i.e. what triggered the order.<br/>* API - <br/>* PORTFOLIO - <br/>* CASH_DIVIDEND_REINVESTMENT - <br/>* PORTFOLIO_REBALANCING - <br/>* SELL_TO_COVER_FEES - <br/>* SELL_TO_COVER_TAXES - <br/>* ACCOUNT_LIQUIDATION - <br/>* UPVEST_OPERATIONS -  |
| `InstrumentID`                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | International securities identification number defined by [ISO 6166](https://en.wikipedia.org/wiki/International_Securities_Identification_Number).                                                                                                            |
| `InstrumentIDType`                                                                                                                                                                                                                                             | [*operations.PlaceOrderInstrumentIDType](../../../pkg/models/operations/placeorderinstrumentidtype.md)                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                             | The type of the ID used in the request.<br/>* ISIN -                                                                                                                                                                                                           |
| `LimitPrice`                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `OrderType`                                                                                                                                                                                                                                                    | [operations.PlaceOrderOrderType](../../../pkg/models/operations/placeorderordertype.md)                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                             | Type of the order.<br/>* MARKET - <br/>* LIMIT - <br/>* STOP -                                                                                                                                                                                                 |
| `Quantity`                                                                                                                                                                                                                                                     | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `Side`                                                                                                                                                                                                                                                         | [operations.PlaceOrderSide](../../../pkg/models/operations/placeorderside.md)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                             | Side of the order.<br/>* BUY - <br/>* SELL -                                                                                                                                                                                                                   |
| `Status`                                                                                                                                                                                                                                                       | [operations.PlaceOrderStatus](../../../pkg/models/operations/placeorderstatus.md)                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                             | The execution status of the order.<br/>* NEW - <br/>* PROCESSING - <br/>* FILLED - <br/>* CANCELLED -                                                                                                                                                          |
| `StopPrice`                                                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                             | N/A                                                                                                                                                                                                                                                            |
| `UpdatedAt`                                                                                                                                                                                                                                                    | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                             | Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)                                                               |
| `UserID`                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                             | The ID of the user                                                                                                                                                                                                                                             |
| `UserInstrumentFitAcknowledgement`                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                             | Only applicable if the user has failed the instrument fit check for the instrument type being ordered. True if the user has acknowledged their willingness to trade.                                                                                           |