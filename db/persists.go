package db

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Persists store DB to disk, file name is db.gob on working directory
func Persists() {
	if file, err := os.Create("./db.gob"); err != nil {
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
