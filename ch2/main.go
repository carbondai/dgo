package main

import (
	"github.com/carbondai/ch2/popcount"
	"fmt"

	"time"
)

func main()  {
	//var i uint64
	//i = 7
	//fmt.Println(byte(i))
	t := time.Now()
	fmt.Println(popcount.PopCount(65536))
	fmt.Printf("%v\n", time.Since(t).Seconds())
}
