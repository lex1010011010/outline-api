# Outline API Go Client

This is a Go client library for interacting with the [Outline](https://github.com/Jigsaw-Code/outline-server) API. The Outline API allows you to manage your Outline server programmatically.

## Installation

To use this package, you can simply add it to your Go module:

```bash
go get github.com/lex1010011010/outline-api
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/lex1010011010/outline-api/outline"
	"time"
)

func main() {
	// Initialize a new Outline API manager
	apiURL := "https://your-outline-server-url.com"
	apiCrt := "your-api-certificate"
	timeout := 30 * time.Second // Optional, defaults to 30 seconds
	manager := outline.NewManager(apiURL, apiCrt, timeout)

	// Get server information
	serverInfo, err := manager.GetServerInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Name:", serverInfo.Name)
}

```

## Documentation

For more details on how to use this package and available methods, please refer to the GoDoc documentation.

The API documentation is based on the Outline API specification.
