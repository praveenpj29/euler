package main

import (
	"fmt"
	euler "euler/goeuler"
	"time"
)

func main() {
	start_time := time.Now()
	fmt.Println(euler.P1())
	fmt.Println("Time taken: ", time.Since(start_time).Abs())
}