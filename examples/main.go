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
