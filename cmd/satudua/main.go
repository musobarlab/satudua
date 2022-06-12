package main

import (
	"fmt"
	"github.com/telkomdev/satudua"
)

func main() {

	name := "gopher"

	res := satudua.Hello(name)

	fmt.Println(res)
}