package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
)

// TODO: implement Query

func Query(callSign string, queryCase row.QueryCase) {
	rec, _ := Get(callSign)
	resp := rec.Locators.SortedByTime()
	if callSign == "SN7L" {
		fmt.Println()
	}
	fmt.Println("\n------------->")
	for _, v := range resp {
		t := v.LogTime.Sprint(true)
		l := v.Locator
		fmt.Println(fmt.Sprintf("%s %s %s", callSign, l, t))
	}
	fmt.Print("******************\n\n")
}
