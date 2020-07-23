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

	d.Add("php", "another one")
	d.Add("java", "no way")
	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Println(entries[word])
	}

	entry, _ := d.Get("go")
	fmt.Println(entry)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
	}
}
