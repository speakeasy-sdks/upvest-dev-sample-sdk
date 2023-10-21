# RetrieveUserUserGetResponseUserBYOL


## Fields

| Field                                                                                                                                                                                            | Type                                                                                                                                                                                             | Required                                                                                                                                                                                         | Description                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Address`                                                                                                                                                                                        | [RetrieveUserUserGetResponseUserBYOLAddress](../../models/operations/retrieveuserusergetresponseuserbyoladdress.md)                                                                              | :heavy_check_mark:                                                                                                                                                                               | Address. Must not be a P.O. box or c/o address.                                                                                                                                                  |
| `BirthCity`                                                                                                                                                                                      | **string*                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                               | N/A                                                                                                                                                                                              |
| `BirthCountry`                                                                                                                                                                                   | **string*                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                               | Country code. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).                                                                                                        |
| `BirthDate`                                                                                                                                                                                      | [types.Date](../../types/date.md)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                               | Birth date of the user in YYYY-MM-DD format. [RFC 3339, section 5.6](https://json-schema.org/draft/2020-12/json-schema-validation.html#RFC3339)                                                  |
| `BirthName`                                                                                                                                                                                      | **string*                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                               | If applicable, birth name of the user.                                                                                                                                                           |
| `CreatedAt`                                                                                                                                                                                      | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                               | Date and time when the resource was created. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html)      |
| `FirstName`                                                                                                                                                                                      | *string*                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                               | First name of the user.                                                                                                                                                                          |
| `ID`                                                                                                                                                                                             | *string*                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                               | User unique identifier.                                                                                                                                                                          |
| `LastName`                                                                                                                                                                                       | *string*                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                               | Last name of the user.                                                                                                                                                                           |
| `Nationalities`                                                                                                                                                                                  | []*string*                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                               | Nationalities of the user. [ISO 3166 alpha-2 Codes](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).                                                                                           |
| `PostalAddress`                                                                                                                                                                                  | [*RetrieveUserUserGetResponseUserBYOLAddress](../../models/operations/retrieveuserusergetresponseuserbyoladdress.md)                                                                             | :heavy_minus_sign:                                                                                                                                                                               | User postal address. Needs to be specified if different to the residential address, otherwise it is automatically populated.                                                                     |
| `Salutation`                                                                                                                                                                                     | [*RetrieveUserUserGetResponseUserBYOLSalutation](../../models/operations/retrieveuserusergetresponseuserbyolsalutation.md)                                                                       | :heavy_minus_sign:                                                                                                                                                                               | Salutation of the user used in reports and statements.<br/>* (empty string) - <br/>* SALUTATION_MALE - <br/>* SALUTATION_FEMALE - <br/>* SALUTATION_FEMALE_MARRIED - <br/>* SALUTATION_DIVERSE -  |
| `Status`                                                                                                                                                                                         | [RetrieveUserUserGetResponseUserBYOLStatus](../../models/operations/retrieveuserusergetresponseuserbyolstatus.md)                                                                                | :heavy_check_mark:                                                                                                                                                                               | Status of the user.<br/>* ACTIVE - <br/>* INACTIVE - <br/>* OFFBOARDING - <br/>* OFFBOARDED -                                                                                                    |
| `Title`                                                                                                                                                                                          | [*RetrieveUserUserGetResponseUserBYOLTitle](../../models/operations/retrieveuserusergetresponseuserbyoltitle.md)                                                                                 | :heavy_minus_sign:                                                                                                                                                                               | Title of the user used in reports and statements.<br/>* (empty string) - <br/>* DR - Doctor<br/>* PROF - Professor<br/>* PROF_DR - <br/>* DIPL_ING - Graduate engineer (Diplom-Ingenieur)<br/>* MAGISTER -  |
| `UpdatedAt`                                                                                                                                                                                      | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                               | Date and time when the resource was last updated. [RFC 3339-5](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6), [ISO8601 UTC](https://www.iso.org/iso-8601-date-and-time-format.html) |