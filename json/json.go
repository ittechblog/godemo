package main

import (
	"fmt"
	"encoding/json"
)

type Response struct {
	Page   int
	Fruits []string
}

func main() {
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	mapC := map[string]int{}
	json.Unmarshal(mapB,&mapC)
	for k,v:= range mapC {
		fmt.Printf("%s----%d\n",k,v)
	}

	res1D := &Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	resp := &Response{}
	json.Unmarshal(res1B,&resp)
	fmt.Println(resp.Page)
}
