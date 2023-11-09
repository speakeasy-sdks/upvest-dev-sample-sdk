# AccountValuation


## Fields

| Field                                                                                                                                                                                              | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccountID`                                                                                                                                                                                        | *string*                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                 | Account unique identifier.                                                                                                                                                                         |
| `CreatedAt`                                                                                                                                                                                        | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                 | Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)        |
| `ID`                                                                                                                                                                                               | *string*                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                 | Account valuation unique identifier.                                                                                                                                                               |
| `PriceQuality`                                                                                                                                                                                     | [*operations.ListAccountValuationHistoryPriceQuality](../../../pkg/models/operations/listaccountvaluationhistorypricequality.md)                                                                   | :heavy_minus_sign:                                                                                                                                                                                 | Price quality used for the calculation of the account valuation.<br/>* EOD - end of day price                                                                                                      |
| `SecurityPositions`                                                                                                                                                                                | [][operations.ListAccountValuationHistoryAccountValuationSecurityPosition](../../../pkg/models/operations/listaccountvaluationhistoryaccountvaluationsecurityposition.md)                          | :heavy_minus_sign:                                                                                                                                                                                 | Positions associated with this account valuation.                                                                                                                                                  |
| `TotalSecurityValue`                                                                                                                                                                               | [operations.ListAccountValuationHistoryTotalSecurityValue](../../../pkg/models/operations/listaccountvaluationhistorytotalsecurityvalue.md)                                                        | :heavy_check_mark:                                                                                                                                                                                 | Entity representing the monetary value by amount and currency.                                                                                                                                     |
| `UpdatedAt`                                                                                                                                                                                        | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                 | Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)   |
| `ValuationTime`                                                                                                                                                                                    | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                 | Date and time as of which the value was calculated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html) |