package main

import (
	"fmt"
	"my-module/hello"

	checkOddEvenV1 "github.com/bachtiarfr/my-go-library"
	checkOddEvenV2 "github.com/bachtiarfr/my-go-library/v2"
)

func main() {
	hello.SayHello()
	result := checkOddEvenV1.CekGanjilGenap(1)
	result2 := checkOddEvenV2.CekGanjilGenap(1, 2, 3, 4, 5)

	fmt.Println("V1 :", result)
	fmt.Println("V2 :", result2)
}

