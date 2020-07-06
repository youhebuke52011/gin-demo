package main

import (
	"fmt"
	"sync"
)

func main() {
	t := []map[string]string{}
	t = append(t, map[string]string{"a": "1"})
	t = append(t, map[string]string{"b": "2"})
	t = append(t, map[string]string{"c": "3"})
	var wg sync.WaitGroup
	flag := true
	for _, tmp := range t {
		//data := tmp
		wg.Add(1)
		go func() {
			defer wg.Done()
			if flag {
				tmp["a"] = "11"
				tmp["b"] = "22"
				tmp["c"] = "33"
				flag = false
			}
			//data["a"] = "11"
			//data["b"] = "22"
			//data["c"] = "33"
		}()
	}
	wg.Wait()
	for _, tmp := range t {
		fmt.Println(tmp)
	}
}
