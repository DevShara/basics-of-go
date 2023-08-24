package main

import (
	"fmt"
	"os"

	"frontendmasters.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	data, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")

	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(data)
	}
}
