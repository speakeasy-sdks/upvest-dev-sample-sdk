# Status

The status of the account
* PENDING_APPROVAL - Account approval is pending - the account is visible through our API but cannot be acted on.
* ACTIVE - Account is active - full functionality of the Investment API is accessible.
* CLOSING - Account is closing - only sell orders or the transfer of positions out are permissible before the account is closed.
* CLOSED - Account is closed with zero balance successfully.
* LOCKED - Account is locked for all actions.


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `StatusPendingApproval` | PENDING_APPROVAL        |
| `StatusActive`          | ACTIVE                  |
| `StatusClosing`         | CLOSING                 |
| `StatusClosed`          | CLOSED                  |
| `StatusLocked`          | LOCKED                  |