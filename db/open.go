package db

import (
	"encoding/gob"
	"fmt"
	"github.com/s51ds/qthdb/row"
	"os"
)

var gobFileName string

func Open(fileName string) error {
	gobFileName = fileName
	table = Table{}
	table.Rows = make(map[string]row.Record)

	// try to load from disk
	if file, err := os.Open(gobFileName); err != nil {
		return err
	} else {
		decoder := gob.NewDecoder(file)
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("init()", err.Error())
			}
		}()
		if err = decoder.Decode(&table); err != nil {
			return err
		} else {
			fmt.Println("db load from disk, file=" + file.Name())
			fmt.Println(fmt.Sprintf("number of rows:%d", NumberOfRows()))

		}
	}
	return nil
}
