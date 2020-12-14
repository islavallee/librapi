package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/boltdb/bolt"
	"time"
)

var (
	// ErrNotFound is returned when the key supplied to a Get or Delete
	// method does not exist in the database.
	ErrNotFound = errors.New("librapi: key not found")
	bucketName  = []byte("librapi")
)

type BoltDB struct {
	path string
}

// Create a BoltDB storage
func NewBoltDB(path string) *BoltDB {
	return &BoltDB{path: path}
}

// Delete entry with the given key
func (bdb *BoltDB) Delete(key string) error {
	db, err := bdb.open()
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucketName).Cursor()
		if k, _ := c.Seek([]byte(key)); k == nil || string(k) != key {
			return ErrNotFound
		} else {
			return c.Delete()
		}
	})

	// If an error occurs on close we won't see it
	db.Close()

	return err
}

// Get entry from BoltDB writes down result in value
func (bdb *BoltDB) Get(key string, value interface{}) error {

	db, err := bdb.open()
	if err != nil {
		return err
	}

	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucketName).Cursor()
		if k, v := c.Seek([]byte(key)); k == nil || string(k) != key {
			return ErrNotFound
		} else if value == nil {
			return nil
		} else {
			d := gob.NewDecoder(bytes.NewReader(v))
			return d.Decode(value)
		}
	})

	// If an error occurs on close we won't see it
	db.Close()

	return err
}

// Save value in BoltDB on the given key
func (bdb *BoltDB) Put(key string, value interface{}) error {

	db, err := bdb.open()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Put([]byte(key), buf.Bytes())
	})

	// If an error occurs on close we won't see it
	db.Close()

	return err
}

// Open BoltDB from file
func (bdb *BoltDB) open() (*bolt.DB, error) {

	opts := &bolt.Options{
		Timeout: 50 * time.Millisecond,
	}

	db, err := bolt.Open(bdb.path, 0640, opts)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
