# jsonllogger

The jsonllogger package provides a structured logging solution in Go, specifically designed for managing JSON log files. It supports automatic log file rotation based on either the number of log lines or a specified time duration. The package also includes features for file compression and organized storage.

## Features
- Configurable Logging: Set parameters such as log directory, filename prefix, maximum lines per file, and log rotation time.
- Automatic Log Rotation: Log files are rotated either after reaching a maximum line count or after a specified time duration.
- Compression and Archiving: Old log files are automatically compressed using gzip and stored in an archive directory within the specified log directory.
- Concurrency Safe: Thread-safe operations ensure that logging can be used in concurrent applications without data races.


## Installation

To install the package, use the `go get` command:

```sh
go get github.com/clwg/jsonllogger
```


## Usage

Define the logger configuration using the LoggerConfig struct:

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

## Creating a New Logger

Instantiate a new logger with the configuration:

```go
logger, err := jsonllogger.NewLogger(config)
if err != nil {
    // Handle error
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

## Integration in Your Application

Integrate the logger into your application logic. For example, if you have a function that needs to log detailed information, you can pass the logger instance to it:

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
