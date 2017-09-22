package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"flag"
)

func main() {
	//flag简单使用方法
	backup_dir := flag.String("b", "/home/default_dir", "backup path")
	debug_mode := flag.Bool("d", false, "debug mode")
	flag.Parse()
	fmt.Println("backup_dir: ",*backup_dir)
	fmt.Println("debug_mode: ",*debug_mode)

	username := flag.String("name", "111", "Input your username")
	flag.Parse()
	fmt.Println("Hello, ", *username)

	resp, err := http.Get("http://localhost:12345/?id=1")
	if err != nil {
		// handle error
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
