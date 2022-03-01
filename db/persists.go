package db

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Persists store DB to disk, gobFileName is set during db/open() process
func Persists() error {
	if NumberOfRows() < 1 {
		return errors.New("db is empty")
	}

	if gobFileName == "" {
		fmt.Println("-----> WARNING -----> THAT SHOULD BE GO TEST, IF NOT IT IS A BUG")
		gobFileName = "test.gob"
		wd, _ := os.Getwd()
		fmt.Println("Persists:", wd+string(filepath.Separator)+gobFileName)

	}

	if file, err := os.Create(gobFileName); err != nil {
		fmt.Println("Persists", err.Error())
		return err
	} else {
		encoder := gob.NewEncoder(file)
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("Persists()", err.Error())
			}
		}()
		if err = encoder.Encode(&table); err != nil {
			fmt.Println("Persists()", err.Error())
			return err
		}
	}
	return nil
}
