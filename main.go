package main

import (
	"fmt"

	"github.com/georgehao/gomodtesta"
	"github.com/georgehao/gomodtestb"
)

func main() {

	fmt.Println("print test a:", gomodtesta.LookUp("hello"))
	fmt.Println("print test b:", gomodtestb.Find("hello", 1234))

}
