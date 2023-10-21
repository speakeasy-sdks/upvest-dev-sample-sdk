# CreateAccountGroupAccountGroup

Account group created.


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CreatedAt`                                                                                                                                                                                                                                                                                                                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)                                                                                                                                                                                          |
| `ID`                                                                                                                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | Account group unique identifier.                                                                                                                                                                                                                                                                                                                                                     |
| `SecuritiesAccountNumber`                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | Securities account number                                                                                                                                                                                                                                                                                                                                                            |
| `Status`                                                                                                                                                                                                                                                                                                                                                                             | [CreateAccountGroupAccountGroupStatus](../../models/operations/createaccountgroupaccountgroupstatus.md)                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | Status of the account group<br/>* PENDING_APPROVAL - Account group approval is pending - the account group is visible through our API but cannot be acted on.<br/>* ACTIVE - Account group is active - full functionality of the Investment API is accessible.<br/>* CLOSING - Account group is closing.<br/>* CLOSED - Account group is closed.<br/>* LOCKED - Account group is locked for all actions. |
| `Type`                                                                                                                                                                                                                                                                                                                                                                               | [CreateAccountGroupAccountGroupType](../../models/operations/createaccountgroupaccountgrouptype.md)                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | Account group type.<br/>* PERSONAL - Account group of a person holding assets on their own behalf.<br/>* LEGAL_ENTITY - Account group of a legal entity holding assets on behalf of their users.                                                                                                                                                                                     |
| `UpdatedAt`                                                                                                                                                                                                                                                                                                                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)                                                                                                                                                                                     |
| `Users`                                                                                                                                                                                                                                                                                                                                                                              | [][CreateAccountGroupAccountGroupUsers](../../models/operations/createaccountgroupaccountgroupusers.md)                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                   | N/A                                                                                                                                                                                                                                                                                                                                                                                  |