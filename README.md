# Circle.so API v1 Go client

Provides access to the [Client.so Platform](https://circle.so/) v1 REST API. This is not an official SDK.

This library does not currently cover all endpoints and will be implemented as and when they are needed by my internal projects.

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
