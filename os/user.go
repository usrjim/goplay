package main

import (
	"fmt"
	"os/user"
)

func main() {
	me, _ := user.Current()

	fmt.Println(me.Username)
	fmt.Println(me.Name)
	fmt.Println(me.HomeDir)
}
