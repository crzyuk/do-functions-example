package main

import "fmt"

func Main(args map[string]interface{}) map[string]interface{} {
	fmt.Println("running function")

	msg := make(map[string]interface{})
	msg["body"] = "This is a test"
	return msg
}
