package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	gobExample()
}

func gobExample() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J", "N2": "J2"}
	if err := enc.Encode(data); err != nil {
		log.Println(err)
		return
	}
	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("%x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		log.Println(err)
	}
	fmt.Println(restored)
}
