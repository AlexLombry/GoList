package main

import (
	"dictionnary/dictionary"
	"fmt"
)

func main() {
	fmt.Println("Kickstart ...")
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
	}
}
