package main

import (
	"os"
	"fmt"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.9fs\n", time.Since(start).Seconds())	
	//fmt.Println(os.Args[1:])
}

