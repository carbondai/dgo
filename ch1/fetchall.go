package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"io"
	"net/http"
	"time"
	//"bufio"
	"strconv"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		//fo, err := os.Create(i+".txt")
		go fetch(i, url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(i int, url string, ch chan<- string) {
	fo, err := os.Create(strconv.Itoa(i+1)+".txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	//w := bufio.NewWriter(fo)
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(fo, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}