package main

import (
	"github.com/carbondai/ch1/lissajous"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//func handler(w http.ResponseWriter, r *http.Request)  {
//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}

