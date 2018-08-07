package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	var answer int
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Scanln(&answer)
	if answer == 0 {
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			if b == nil {
				b, err := tx.CreateBucket([]byte("MyBucket"))
				err = b.Put([]byte("answer"), []byte("42"))
				if err != nil {
					log.Fatal(err)
				}
			} else {
				err := b.Put([]byte("answer"), []byte("42"))
				if err != nil {
					log.Fatal(err)
				}

			}

			return err
		})
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			v := b.Get([]byte("answer"))
			fmt.Printf("The answer is: %s\n", v)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
