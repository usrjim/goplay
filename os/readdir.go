package main

import (
	"fmt"
	"os"
	"time"
)

var deep int

func main() {
	target := "/tmp"
	dig(target)
}

func dig(p string) {
	dh, _ := os.Open(p)
	defer dh.Close()

	f, _ := dh.Readdir(-1)
	for _, fi := range f {
		if fi.IsDir() {
			child := p + string(os.PathSeparator) + fi.Name()
			dig(child)
		} else {
			fmt.Println(p + fi.Name())
			time.Sleep(time.Millisecond * 10)
		}
	}
}
