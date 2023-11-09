# Allocation


## Fields

| Field                                                                                                                                            | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `InstrumentID`                                                                                                                                   | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `InstrumentIDType`                                                                                                                               | [*operations.CreatePortfoliosAllocationInstrumentIDType](../../../pkg/models/operations/createportfoliosallocationinstrumentidtype.md)           | :heavy_minus_sign:                                                                                                                               | The type of the ID used in the request.<br/>* ISIN - International Securities Identification Number<br/>* UPVEST - UPVEST's unique instrument identifier |
| `Weight`                                                                                                                                         | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | Instrument allocation weight                                                                                                                     |