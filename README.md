# Circle.so API v1 Go client

[![Go](https://github.com/miguelbemartin/circleso-go/actions/workflows/go.yml/badge.svg)](https://github.com/miguelbemartin/circleso-go/actions/workflows/go.yml)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/miguelbemartin/circleso-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/miguelbemartin/circleso-go)](https://goreportcard.com/report/github.com/miguelbemartin/circleso-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

Provides access to the [Client.so](https://circle.so/) Platform v1 REST API. This is not an official SDK. Golang wrapper around the Circle.so API v1.

This library does not currently cover all endpoints and will be implemented as and when they are needed by my internal projects. Feel free to add contribute to the project.

## Installation

```
go get github.com/miguelbemartin/circleso-go
```

## Usage

```golang
package main

import (
	"context"
	"fmt"
	circle "github.com/miguelbemartin/circleso-go"
)

func main() {

	c := circle.NewClient("your-api-token")

	members, err := c.GetMembers(context.Background(), nil)
	if err != nil {
		// do something with the error
	}

	// iterate the members
	for _, member := range members {
		// do something
	}

}
```
