# ListCashWithdrawalsWithdrawalsListResponseMeta


## Fields

| Field                                                                                                                                  | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `Count`                                                                                                                                | *int64*                                                                                                                                | :heavy_check_mark:                                                                                                                     | Count of the resources returned in the response.                                                                                       |
| `Limit`                                                                                                                                | *int64*                                                                                                                                | :heavy_check_mark:                                                                                                                     | Total limit of the response.                                                                                                           |
| `Offset`                                                                                                                               | *int64*                                                                                                                                | :heavy_check_mark:                                                                                                                     | Amount of resource to offset in the response.                                                                                          |
| `Order`                                                                                                                                | [*ListCashWithdrawalsWithdrawalsListResponseMetaOrder](../../models/operations/listcashwithdrawalswithdrawalslistresponsemetaorder.md) | :heavy_minus_sign:                                                                                                                     | The ordering of the response.<br/>* ASC - Ascending order<br/>* DESC - Descending order                                                |
| `Sort`                                                                                                                                 | **string*                                                                                                                              | :heavy_minus_sign:                                                                                                                     | The field that the list is sorted by.                                                                                                  |
| `TotalCount`                                                                                                                           | *int64*                                                                                                                                | :heavy_check_mark:                                                                                                                     | Total count of all the resources.                                                                                                      |