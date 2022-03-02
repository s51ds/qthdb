package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	if err := Open("../db.gob"); err != nil {
		fmt.Println(err.Error())
		dir, _ := os.Getwd()
		fmt.Println("Working directory:", dir)
	}
	Query("9A5ISS", row.QueryLatestAll)
}
