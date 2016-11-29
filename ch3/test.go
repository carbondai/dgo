//package main
//
//import (
//	"log"
//	"time"
//)
//
////func bigSlow() {
////	defer trace("bigSlow")()
////	log.Printf("1")
////
////	time.Sleep(3 * time.Second)
////	log.Printf("3")
////
////}
//
////func trace(msg string) func()  {
////	log.Printf("2")
////	start := time.Now()
////	log.Printf("enter %s", msg)
////	return func() {
////		log.Printf("exit %s (%s)", msg, time.Since(start))
////	}
////}
//
//func double(x int)  int {
//	defer p()
//	log.Printf("2")
//	return x + x
//}
//func p()  {
//	log.Printf("1")
//	time.Sleep(time.Second)
//	log.Printf("3")
//	//log.Printf("double(%d)=%d\n", x, result)
//	log.Printf("4")
//}
//
//func main()  {
//	double(4)
//}

//package main
//
//import (
//	"time"
//	"fmt"
//)
//
//func main()  {
//	go spinner(100*time.Millisecond)
//	const n  = 45
//	fibN := fib(n)
//	fmt.Printf("fibonacci(%d)=%d\n", n, fibN)
//}
//
//func spinner(delay time.Duration)  {
//	for {
//		for _, r := range `-\|/` {
//			fmt.Printf("\r%c", r)
//			time.Sleep(delay)
//		}
//	}
//}
//
//func fib(x int) int  {
//	if x < 2 {
//		return x
//	}
//	return fib(x-1) + fib(x-2)
//}

package main

import (
	"os"
	"fmt"
)

func main()  {
	temp := os.Getenv("127.0.0.1:2375")
	fmt.Println(temp)
}