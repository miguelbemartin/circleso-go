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
	fmt.Println("hello world")

	c := circle.NewClient("your-api-token")

	members, err := c.GetMembers(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	// iterate and print members
	for _, member := range members {
		fmt.Println(member)
	}

}
```
