package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
)

// TODO: implement Query

func Query(callSign string, queryCase row.QueryCase) {
	rec, _ := Get(row.CallSign(callSign))
	resp := rec.Locators().SortedByTime()
	for _, v := range resp {
		t := v.LogTime.Sprint(true)
		l := v.Locator
		fmt.Println(l, t)
	}
}
