# SetTaxResidenciesResponse


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `TwoHundredApplicationJSONTaxResidencyRecord`                                                                            | [*operations.SetTaxResidenciesTaxResidencyRecord](../../../pkg/models/operations/settaxresidenciestaxresidencyrecord.md) | :heavy_minus_sign:                                                                                                       | User tax residencies                                                                                                     |
| `ContentType`                                                                                                            | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | HTTP response content type for this operation                                                                            |
| `Headers`                                                                                                                | map[string][]*string*                                                                                                    | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `StatusCode`                                                                                                             | *int*                                                                                                                    | :heavy_check_mark:                                                                                                       | HTTP response status code for this operation                                                                             |
| `RawResponse`                                                                                                            | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                   | :heavy_minus_sign:                                                                                                       | Raw HTTP response; suitable for custom response parsing                                                                  |