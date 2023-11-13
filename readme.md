# jsonllogger

jsonllogger is a Go package for logging with features like log rotation, gzip compression, and JSON Lines (JSONL) format logging. It's ideal for applications that require efficient logging with automatic log file management.

## Installation

To install the package, use the `go get` command:

```sh
go get github.com/clwg/jsonllogger
```


## Configuration

First, import the package in your Go file:

```go
import "github.com/clwg/jsonllogger"
```

Configure the logger by creating an instance of `LoggerConfig`:

```go
config := jsonllogger.LoggerConfig{
    FilenamePrefix: "yourPrefix",
    LogDir:         "./path/to/logs",
    MaxLines:       1000, // Max lines before rotating
    RotationTime:   1 * time.Hour, // Duration for log rotation
}
```

- `FilenamePrefix`: A string prefix for log filenames.
- `LogDir`: Directory path where logs will be stored.
- `MaxLines`: Maximum number of lines in a log file before it rotates.
- `RotationTime`: Duration after which the log file will rotate.

## Creating a Logger Instance

To create a logger instance, use `NewLogger`:

```go
logger, err := jsonllogger.NewLogger(config)
if err != nil {
    panic(err) // Handle error appropriately
}
```

## Logging

To log a message, call the `Log` method with a struct or map that can be marshaled to JSON:

```go
err = logger.Log(yourDataStructure)
if err != nil {
    // Handle logging error
}
```

Replace `yourDataStructure` with the data you want to log.

## Integration in Your Application

Integrate the logger into your application logic. For example, if you have a function that needs to log detailed information, you can pass the logger instance to it:

```go
go yourFunction(logger) // Replace 'yourFunction' with your actual function
```

Ensure that `yourFunction` accepts a `*jsonllogger.Logger` parameter and uses its `Log` method to log messages.

## Example

Here's a complete example of setting up and using the logger:

```go
package main

import (
	"time"

	"github.com/clwg/jsonllogger"
)

type ExampleStruct struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	config := jsonllogger.LoggerConfig{
		FilenamePrefix: "example",
		LogDir:         "./logs",
		MaxLines:       100,
		RotationTime:   30 * time.Minute,
	}

	logger, err := jsonllogger.NewLogger(config)
	if err != nil {
		panic(err)
	}

	// Example of logging a message
	logger.Log(map[string]string{"message": "Hello, Logger!"})

	// Invoke yourFunction in a goroutine
	go yourFunction(logger)

	// Wait for the goroutine to finish
	// Implement a proper synchronization mechanism here if necessary
	time.Sleep(time.Second * 5)
}

func yourFunction(logger *jsonllogger.Logger) {
	// Function logic that uses logger
	logger.Log(map[string]string{"message": "Hello, yourFunction!"})
}

```

## License

This package is licensed under BSD 3-Clause License.
