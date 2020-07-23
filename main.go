package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"log"
)

func main() {
	fmt.Println("Kickstart ...")

	db, err := BadgerConnect()
	if err != nil {
		panic(err)
	}


}

func BadgerConnect() (*badger.DB, error) {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db, err
}