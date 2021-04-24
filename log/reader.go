package log

import (
	"github.com/s51ds/qthdb/row"
	"log"
	"regexp"
	"strings"
)

type Type int

const (
	TypeN1mmCallHistory Type = iota
	TypeN1mmGenericFile
)

var regex, _ = regexp.Compile("^[a-zA-Z0-9]")

func LineHasData(line string) bool {
	return regex.MatchString(line)
}

// DetectSeparator checks if separator is comma or semicolon. If no separators is detected
// comma is returned
func DetectSeparator(line string) string {
	if strings.Contains(line, ";") {
		return ";"
	} else {
		return ","
	}

}

func Parse(logType Type, line string) (record row.Record, err error) {
	if !LineHasData(line) {
		return record, err // no error, just skip the line
	}
	sep := DetectSeparator(line)
	switch logType {
	case TypeN1mmCallHistory:
		return parseN1mmCallHistoryLine(line, sep)
	case TypeN1mmGenericFile:
		return parseN1mmGenericFileLine(line, sep)
	default:
		log.Fatalln("WTF")
		return
	}
}

type dataLocatorsInputCase struct {
	loc1 bool
	loc2 bool
}

func (d *dataLocatorsInputCase) loc1andLoc2() bool {
	return d.loc2 && d.loc1
}

func (d *dataLocatorsInputCase) loc1Only() bool {
	return d.loc1 && !d.loc2
}

func (d *dataLocatorsInputCase) loc2Only() bool {
	return !d.loc1 && d.loc2
}

func (d *dataLocatorsInputCase) bothLocators() bool {
	return d.loc2 && d.loc1
}

func parseN1mmCallHistoryLine(line string, sep string) (record row.Record, err error) {
	//   0   1  2      3
	// S51IV,,JN76UP,JN76TO
	ss := strings.Split(line, sep)
	switch len(ss) {
	case 1:
		{
			return row.MakeNewRecord(row.CallSign(ss[0]), "", "", "")
		}
	case 3:
		{
			if ss[2] != "" {
				return row.MakeNewRecord(row.CallSign(ss[0]), row.Locator(ss[2]), "", "")
			} else {
				return row.MakeNewRecord(row.CallSign(ss[0]), "", "", "")
			}
		}
	default:
		locators := dataLocatorsInputCase{}
		if len(ss) > 2 {
			locators.loc1 = ss[2] != ""
		}
		if len(ss) > 3 {
			locators.loc1 = ss[2] != ""
			locators.loc2 = ss[3] != ""
		}

		if record, err = row.MakeNewRecord(row.CallSign(ss[0]), row.Locator(ss[2]), "", ""); err != nil {
			return
		} else {
			if err = record.Update(row.Locator(ss[2]), "", ""); err != nil {
				return row.Record{}, err
			} else {
				return record, nil
			}
		}
	}
}

//func parseN1mmCallHistoryLine(line string, sep string) (record row.Record, err error) {
//	//   0   1  2      3
//	// S51IV,,JN76UP,JN76TO
//	ss := strings.Split(line, sep)
//	switch len(ss) {
//	case 1:
//		{
//			return row.MakeNewRecord(row.CallSign(ss[0]), "", "", "")
//		}
//	case 3:
//		{
//			if ss[2] != "" {
//				return row.MakeNewRecord(row.CallSign(ss[0]), row.Locator(ss[2]), "", "")
//			} else {
//				return row.MakeNewRecord(row.CallSign(ss[0]), "", "", "")
//			}
//		}
//	default:
//		type data struct {
//			loc1     bool
//			loc2     bool
//		}
//		locators := data{}
//		if len(ss) > 2 {
//			locators.loc1 = ss[2] != ""
//		}
//		if len(ss) > 3 {
//			locators.loc1 = ss[2] != ""
//			locators.loc2 = ss[3] != ""
//		}
//
//		switch locators {
//		case locators.loc2 && locators.loc2: {}
//
//		}
//
//
//		switch {
//		case (!loc1 && loc2):
//			{
//
//			}
//
//		case loc1 && loc2:
//			{
//
//			}
//		case loc1 && !loc2:
//			{
//
//			}
//		case !loc1 && loc2:
//			{
//
//			}
//
//		}
//
//		if record, err = row.MakeNewRecord(row.CallSign(ss[0]), row.Locator(ss[2]), "", ""); err != nil {
//			return
//		} else {
//			if err = record.Update(row.Locator(ss[2]), "", ""); err != nil {
//				return row.Record{}, err
//			} else {
//				return record, nil
//			}
//		}
//	}
//}

func parseN1mmGenericFileLine(line string, sep string) (record row.Record, err error) {

	return
}
