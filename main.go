package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("/tmp/captain"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(db)
}
