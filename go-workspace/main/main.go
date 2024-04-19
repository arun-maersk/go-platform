package main

import (
	"fmt"

	"github.com/arun-maersk/go-platform/go-workspace/module1"
	"github.com/arun-maersk/go-platform/go-workspace/module2"
	"github.com/arun-maersk/go-platform/go-workspace/module2/module3"
)

func main() {
	fmt.Println(module1.Text())
	fmt.Println(module2.Number())
	fmt.Println(module3.Time())

}
