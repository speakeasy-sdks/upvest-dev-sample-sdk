# Keys


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Crv`                                                     | [*operations.Crv](../../models/operations/crv.md)         | :heavy_minus_sign:                                        | Elliptic curve family.<br/>* P-521 -                      |
| `Kid`                                                     | **string*                                                 | :heavy_minus_sign:                                        | Key ID                                                    |
| `Kty`                                                     | [*operations.Kty](../../models/operations/kty.md)         | :heavy_minus_sign:                                        | Cryptographic algorithm family used with the key.<br/>* EC -  |
| `X`                                                       | **string*                                                 | :heavy_minus_sign:                                        | Curve parameter                                           |
| `Y`                                                       | **string*                                                 | :heavy_minus_sign:                                        | Curve parameter                                           |