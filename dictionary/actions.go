package dictionary

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v2"
	"strings"
	"time"
)

func (d *Dictionary) Add(word, definition string) error {
	// open transaction and set value
	entry := Entry{
		Word:       strings.Title(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}

	// convert struct entry into byte (gob)
	var bf bytes.Buffer
	enc := gob.NewEncoder(&bf)
	err := enc.Encode(entry)
	if err != nil {
		return err
	}

	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), bf.Bytes())
	})
}

func (d *Dictionary) Get(word string) (Entry, error) {
	var entry Entry
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(word))
		if err != nil {
			return err
		}
		entry, err = getEntry(item)

		return err
	})

	return entry, err
}

func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry
	var bf bytes.Buffer
	err := item.Value(func(val []byte) error {
		_, err := bf.Write(val)
		return err
	})

	dec := gob.NewDecoder(&bf)
	err = dec.Decode(&entry)

	return entry, err

}
