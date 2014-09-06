package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	components := []string{"a", "path", "..", "with", "relative", "elements"}

	path := path.Join(components...)
	fmt.Printf("Path: %s\n", path)

	decomposed := filepath.SplitList(path)
	fmt.Printf("%c\n", filepath.Separator)
	for _, dir := range decomposed {
		fmt.Printf("%s%c", dir, filepath.Separator)
	}

}
