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
		Word: strings.Title(word),
		Definition: definition,
		CreatedAt: time.Now(),
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