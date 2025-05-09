package wal

import (
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func InitDB() error {
	var err error
	db, err = bbolt.Open("chat.db", 0600, nil)
	if err != nil {
		return err
	}

	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("messages"))
		return err
	})
}

func SaveMessage(key, message string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		return b.Put([]byte(key), []byte(message))
	})
}

func GetMessage(key string) (string, error) {
	var msg string
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("messages"))
		msg = string(b.Get([]byte(key)))
		return nil
	})
	return msg, err
}
