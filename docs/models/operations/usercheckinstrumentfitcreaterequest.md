# UserCheckInstrumentFitCreateRequest

Instrument fit check is completed by the client providing the user's answers to the instrument appropriateness or suitability questionnaire.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `CheckConfirmedAt`                                                                   | [time.Time](https://pkg.go.dev/time#Time)                                            | :heavy_check_mark:                                                                   | Completion date and time of the instrument fit check.                                |
| `InstrumentSuitability`                                                              | [operations.InstrumentSuitability](../../models/operations/instrumentsuitability.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `Type`                                                                               | **string*                                                                            | :heavy_minus_sign:                                                                   | The type of check must be INSTRUMENT_FIT.                                            |