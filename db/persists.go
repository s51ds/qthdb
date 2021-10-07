package db

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Persists store DB to disk, gobFileName is set during db/open() process
func Persists() {
	if file, err := os.Create(gobFileName); err != nil {
		fmt.Println("Persists", err.Error())
	} else {
		encoder := gob.NewEncoder(file)
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("Persists()", err.Error())
			}
		}()
		if err = encoder.Encode(&table); err != nil {
			fmt.Println("Persists()", err.Error())
		}
	}
}
