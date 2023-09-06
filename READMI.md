# Outline Manager API Client in Go

This is a Go client for interacting with the Outline Manager API.

## Installation

`go get github.com/lex1010011010/outline-api`

## Usage

```go
package main

import (
	"fmt"
	"github.com/lex1010011010/outline-api/outline"
)

func main() {
	manager := outlineManager.NewManager("apiURL", "apiCrt", 30)
	info, err := manager.ServerInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Server Info:", info)
}
