package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a directory.")
		os.Exit(1)
	}
	directory := arguments[1]

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Mode(), " ", file.ModTime(), " ", file.Name())
	}

}
