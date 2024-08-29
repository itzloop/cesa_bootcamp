## Features to Implement:

### 1. Displaying the Correct `userId`
Currently, all three simulated work functions (`Work1`, `Work2`, and `Work3`) are printing `userId = 0` regardless of the actual user associated with the request. The task is to ensure that the correct `userId` is printed in each of these work functions without modifying the `main` function.

### 2. Timeouts for Requests
You need to implement a mechanism to ensure that the entire request (handled by `HandleRequest`) does not take longer than **10 seconds**. Use Go’s `context` package to enforce this timeout.

### 3. Error Propagation and Cancellation
If an error occurs in `Work1`, the entire request should be canceled, meaning that all ongoing tasks (`Work2`, `Work3`, etc.) should be stopped. Again, use the `context` package to handle this cancellation logic.