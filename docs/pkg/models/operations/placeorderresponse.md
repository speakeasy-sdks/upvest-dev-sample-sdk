# PlaceOrderResponse


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ContentType`                                                                    | *string*                                                                         | :heavy_check_mark:                                                               | HTTP response content type for this operation                                    |
| `Headers`                                                                        | map[string][]*string*                                                            | :heavy_check_mark:                                                               | N/A                                                                              |
| `Order`                                                                          | [*operations.PlaceOrderOrder](../../../pkg/models/operations/placeorderorder.md) | :heavy_minus_sign:                                                               | The request for the order creation has been accepted for processing.             |
| `StatusCode`                                                                     | *int*                                                                            | :heavy_check_mark:                                                               | HTTP response status code for this operation                                     |
| `RawResponse`                                                                    | [*http.Response](https://pkg.go.dev/net/http#Response)                           | :heavy_check_mark:                                                               | Raw HTTP response; suitable for custom response parsing                          |