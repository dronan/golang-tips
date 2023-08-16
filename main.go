package main

import (
	"fmt"
)

func main() {

	// print json response on terminal
	res := debugJsonResponse()
	fmt.Printf("%+v\n", res)

}
