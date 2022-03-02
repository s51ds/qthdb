package locators

import (
	"fmt"
	"github.com/s51ds/qthdb/db"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := db.Open("../db.gob"); err != nil {
		fmt.Println(err.Error())
		dir, _ := os.Getwd()
		fmt.Println("Working directory:", dir)
		os.Exit(1)
	}
	fmt.Println(db.NumberOfRows())
	os.Exit(m.Run())

}

func TestLocators(t *testing.T) {
	callSign := "S59ABC"
	resp := Get(callSign)
	if len(resp) == 0 {
		t.Errorf("unexpected, no locators for %s ", callSign)
	}
	for _, v := range resp {
		t := v.LogTime.Sprint(true)
		l := v.Locator
		fmt.Println(fmt.Sprintf("%s %s %s", callSign, l, t))
	}
	callSign = "S57NAW"
	resp = Get("S57NAW")
	if len(resp) == 0 {
		t.Errorf("unexpected, no locators for %s ", callSign)
	}
	for _, v := range resp {
		t := v.LogTime.Sprint(true)
		l := v.Locator
		fmt.Println(fmt.Sprintf("%s %s %s", callSign, l, t))
	}
}
