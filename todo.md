# To do

## Format Parsing

### Accumulative vs. Non-Accumulative Events

The events being parsed should be placed into two different types, ones that are 'accumulative' and ones that are 'non-accumulative'. i.e. "user did 3 pushed into repo/user" vs. "user created repo/user"

### Custom Event Types

Some events require additional information, such as "CreateEvent" and "DeleteEvent" types. These both require parsing of `$.payload` in order to get whether the repo was created or deleted, or a branch or tag was created or deleted.

Can I get away with only one additional modifier?
Or should I create a whole new slew of custom event types that go further: "CreateRepoEvent", "DeleteBranchEvent", and they have their types that contain needed information. I could write an interface of `Event` that they conform to, in which they each put together their own formatted string?

## Performing Requests

### Cache x-ratelimit-reset Reset

Save the time that the rate limit is reset locally in order to respect the rate limit, utilize the `retry-after` header as well.
