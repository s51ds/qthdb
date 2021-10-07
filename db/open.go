package db

import (
	"encoding/gob"
	"fmt"
	"github.com/s51ds/qthdb/row"
	"os"
)

func Open(fileName string) {
	table = Table{}
	table.Rows = make(map[string]row.Record)

	// try to load from disk
	if file, err := os.Open(fileName); err != nil {
		fmt.Println("init()", err.Error())
	} else {
		decoder := gob.NewDecoder(file)
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("init()", err.Error())
			}
		}()
		if err = decoder.Decode(&table); err != nil {
			fmt.Println("init()", err.Error())
		} else {
			fmt.Println("db load from disk, file=" + file.Name())
			fmt.Println(fmt.Sprintf("number of rows:%d", NumberOfRows()))

		}
	}

}
