# Backend Testing

[Go Back](/README.md)

For the testing, we use standard `testing` library. Also we have several testing utils, that can help to keep the code clean and underdstandable.

## Code organization

Create a file with `_test.go` suffix, e.g. `login_test.go` and it will be automatically recognized as test.

For the clarity, we use nested tests:

```go
// File function_test.go

// Optional setup function for your tests
// That will be executed for each group
func setupGroup() {...}

func TestGroup(t *testing.T) {
  setupGroup()

  t.Run("FunctionTest") {
    // Your tests

    // Also you can nest your tests even more:
    t.Run("NestedTest") {
      // ...
    }
  }
}
```
## Running your tests

Run `go test ./...` in `/backend` directory to run all the tests (please do it before commit) or in your microservice's directory to run only related tests.

Major IDEs (like Visual Studio Code and Idea) should be able to automatically detect and run tests.

All the tests are executed locally.

## Testutils

Testutils is a small library that can help you to write less-repeatable tests.
Just add `testing` to your tests' imports.

### Setting up a Testing DB

We use in-memory SQLite driver with gorm. Use `testutils.SetupDB` function to set up and automigrate your testng DB. It recreates DB after each test (using `t.Cleanup` hook).

**Notes:**
1. Pass an empty struct for the automigration
1. It does not seed any data, so keep it in mind.

### Making requests

Use `rr, req := testing.MakeRequest(...)`  to make a request and get both ResponseRecorder and Request objects.

**Notes:**

1. You can use long or short version of API endpoint, but short one is more preferred.
1. The payload will be marshallized, so you can pass an object

### Asserts

A simple asserting function that takes expected and actual values with some message you can display.
