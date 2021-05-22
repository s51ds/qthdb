package app

import (
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/db"
	"os"
	"sort"
	"time"
)

func sPrintN1mmScpFormat(callSign, loc1, loc2 string) string {
	// S58M,,JN76PL,JN76JC
	return fmt.Sprintf("%s,,%s,%s", callSign, loc1, loc2)
}

func MakeN1mmScpFile(scpFileName string, month time.Month) error {
	if month < 0 || month > 12 {
		return errors.New(fmt.Sprintf("invalid Month:%d", month))
	}

	scpFile, err := os.Create(scpFileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = scpFile.Close()
	}()

	scpLines := make([]string, 0, 10000)
	rows := db.GetAll()
	for _, r := range rows {
		callSign := string(r.CallSign())
		var loc1, loc2 string
		resp := r.Locators().SortedByTime()
		switch len(resp) {
		case 0:
			{
				// do nothing
			}
		case 1:
			{
				loc1 = string(resp[0].Locator)
			}
		case 2:
			{
				loc1 = string(resp[0].Locator)
				loc2 = string(resp[1].Locator)
				if loc1 == loc2 { // same locator
					loc2 = ""
				}
			}
		default:
			{ // more than two entries

				loc1 = string(resp[0].Locator) // loc1 is always the latest
				///////////
				// what is the loc2?
				// looking for latest locator which is different than loc1
				//
				// first look the Month of contest
				monthFound := false
				if month != 0 {
					for i, v := range resp {
						if i == 0 {
							// skip, loc1 is already set
						}
						if loc1 != string(v.Locator) && v.LogTime.LoggedTime().Month() == month {
							// found
							monthFound = true
							loc2 = string(v.Locator)
							break
						}
					}

				}
				if !monthFound {
					for i, v := range resp {
						if i == 0 {
							// skip, loc1 is already set
						}
						if loc1 != string(v.Locator) {
							// found
							loc2 = string(v.Locator)
							break
						}
					}

				}

			}
		}
		if loc1 == loc2 {
			fmt.Println("BUUUG")
			os.Exit(2)
		}
		s := sPrintN1mmScpFormat(callSign, loc1, loc2)
		scpLines = append(scpLines, s)
	}
	sort.Strings(scpLines)
	for _, v := range scpLines {
		if _, err := scpFile.WriteString(v + "\n"); err != nil {
		}
	}
	return nil
}
