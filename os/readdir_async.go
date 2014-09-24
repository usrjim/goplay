package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var deep int
var wg sync.WaitGroup

func main() {
	target := "/tmp"
	wg.Add(1)
	go dig(target)
	wg.Wait()
}

func dig(p string) {
	dh, _ := os.Open(p)
	defer dh.Close()
	defer wg.Done()

	f, _ := dh.Readdir(-1)
	for _, fi := range f {
		if fi.IsDir() {
			child := p + string(os.PathSeparator) + fi.Name()
			wg.Add(1)
			go dig(child)
		} else {
			fmt.Println(p + fi.Name())
			time.Sleep(time.Millisecond * 10)
		}
	}
}
