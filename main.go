package main

import (
	"dictionnary/dictionary"
	"flag"
	"fmt"
	"os"
	"os/user"
)

func main() {
	action := flag.String("action", "list", "What action do you perform ?")

	u, err := user.Current()
	if err != nil {
		fmt.Println("Unable to find user home directory")
		os.Exit(1)
	}

	homeDir := u.HomeDir

	dir := fmt.Sprintf("%v/%s", homeDir, ".badger/")
	d, err := dictionary.New(dir)
	handleErr(err)
	defer d.Close()

	flag.Parse()

	// de referenced value of Action
	switch *action {
	case "list":
		actionList(d)
	case "get":
		actionGet(d, flag.Args())
	case "add":
		actionAdd(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Todolist content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionGet(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Printf("%v\t%v added to the todolist\n", entry.Word, entry.Definition)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the todolist\n", word)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' remove from the todolist\n", word)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
	}
}
