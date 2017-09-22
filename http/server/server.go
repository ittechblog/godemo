package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
)
type Counter struct {
	n int
}
func (ctr *Counter) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(c, "counter = %d\n", ctr.n)
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	err := req.ParseForm()
	if(err != nil){
		return
	}
	id := req.Form["id"]
	println(string(id))
}

func main() {
	http.HandleFunc("/",HelloServer)
	http.Handle("/counter", new(Counter))
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":12345", nil))
}