# CashTransactionReference

Entity representing cash transaction reference.


## Fields

| Field                                                                                                                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                        | Unique identifier for a resource of given type.                                                                                                                                                                                                                                                                           |
| `Type`                                                                                                                                                                                                                                                                                                                    | [operations.ListCashTransactionsType](../../../pkg/models/operations/listcashtransactionstype.md)                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                        | Type of the reference.<br/>* ORDER - Order<br/>* ORDER_EXECUTION - Order execution<br/>* WITHDRAWAL - Cash withdrawal<br/>* DIRECT_DEBIT - Direct debit funding request<br/>* CORPORATE_ACTION - Corporate action<br/>* CORPORATE_ACTION_TRANSACTION_ID - Corporate action transaction ID<br/>* TOPUP - Cash top up<br/>* FEE_COLLECTION - Fee collection |