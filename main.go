package main

import (
	euler "euler/goeuler"
	"flag"
	"fmt"
	"reflect"
	"time"
	"os"
)

func main() {
	e := euler.EulerFuncs{}
	problemFlag := flag.String("p", "1", "problem number")
	flag.Parse()

	funcName := fmt.Sprintf("P%s", *problemFlag)

	funcValue := reflect.ValueOf(&e).MethodByName(funcName)
	start_time := time.Now()
	if !funcValue.IsValid() {
		fmt.Printf("Function %s does not exist\n", funcName)
		os.Exit(1)
	}
	result := funcValue.Call(nil)
	fmt.Println(result[0])
	fmt.Println("Time taken: ", time.Since(start_time).Abs())
}
