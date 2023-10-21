# ListCashTransactionsCashTransactionListResponseCashTransactionTransactionTax

Entity representing the transaction tax.


## Fields

| Field                                                                                                                                                                                                    | Type                                                                                                                                                                                                     | Required                                                                                                                                                                                                 | Description                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Amount`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                       | N/A                                                                                                                                                                                                      |
| `Currency`                                                                                                                                                                                               | [*ListCashTransactionsCashTransactionListResponseCashTransactionTransactionTaxCurrency](../../models/operations/listcashtransactionscashtransactionlistresponsecashtransactiontransactiontaxcurrency.md) | :heavy_minus_sign:                                                                                                                                                                                       | Alphabetic three-letter [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.<br/>* EUR - Euro                                                                                               |
| `Type`                                                                                                                                                                                                   | [*ListCashTransactionsCashTransactionListResponseCashTransactionTransactionTaxType](../../models/operations/listcashtransactionscashtransactionlistresponsecashtransactiontransactiontaxtype.md)         | :heavy_minus_sign:                                                                                                                                                                                       | Type of the tax.<br/>* TOTAL - Total taxes                                                                                                                                                               |