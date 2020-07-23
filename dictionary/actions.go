package dictionary

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v2"
	"sort"
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

// Get dictionary content sorted
func (d *Dictionary) List() ([]string, map[string]Entry, error) {
	entries := make(map[string]Entry)
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			entry, err := getEntry(item)
			if err != nil {
				return err
			}
			entries[entry.Word] = entry
		}
		return nil
	})

	return sortedKeys(entries), entries, err
}

func sortedKeys(entries map[string]Entry) []string {
	// make a map of size entries slice, range, sort and return
	keys := make([]string, len(entries))
	for key, _ := range entries {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry
	var bf bytes.Buffer

	// write all value into the buffer
	err := item.Value(func(val []byte) error {
		_, err := bf.Write(val)
		return err
	})

	// Decode and return
	dec := gob.NewDecoder(&bf)
	err = dec.Decode(&entry)

	return entry, err
}

func (d *Dictionary) Remove(word string) error {
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(word))
	})
}
