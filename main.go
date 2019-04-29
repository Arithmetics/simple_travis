package main

import (
	"fmt"
)

func goodFunc() string {
	return "looks good"
}

func badFunc() string {
	return "this one bad"
}

func main() {
	fmt.Printf("%+v\n", goodFunc())
	fmt.Printf("%+v\n", badFunc())
}
