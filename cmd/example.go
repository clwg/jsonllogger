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
