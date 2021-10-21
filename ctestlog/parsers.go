package ctestlog

import (
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/row"
	"github.com/s51ds/qthdb/timing"
	"regexp"
	"strings"
)

var regex, _ = regexp.Compile("^[a-zA-Z0-9]")

// LineHasData returns false if line is comment, empty ...
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
		return record, err // nil error, just skip the line
	}
	switch logType {
	case TypeN1mmCallHistory:
		sep := DetectSeparator(line)
		return parseN1mmCallHistoryLine(line, sep)
	case TypeN1mmGenericFile:
		return parseN1mmGenericFileLine(line)
	case TypeEdiFile:
		return parseEdiQsoRecord(line)
	default:
		return row.Record{}, errors.New(fmt.Sprintf("Wrong logType:%d", logType))
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

func parseN1mmCallHistoryLine(line string, sep string) (record row.Record, err error) {
	//   0   1  2      3
	// S51IV,,JN76UP,JN76TO
	ss := strings.Split(line, sep)
	inputCase := dataLocatorsInputCase{}
	switch len(ss) {
	case 3:
		{
			if ss[2] != "" {
				inputCase.loc1 = true
			}
		}
	default:
		if len(ss) > 2 && ss[2] != "" {
			inputCase.loc1 = true
		}
		if len(ss) > 3 && ss[3] != "" {
			inputCase.loc2 = true
		}
	}
	switch {
	case inputCase.loc1andLoc2():
		{

			if record, err = row.MakeNewRecord(ss[0], ss[2], "", ""); err != nil {
				return row.Record{}, err
			}
			if err = record.Update(ss[3], "", ""); err != nil {
				return row.Record{}, err
			}
		}
	case inputCase.loc1Only():
		{
			if record, err = row.MakeNewRecord(ss[0], ss[2], "", ""); err != nil {
				return row.Record{}, err
			}
		}
	case inputCase.loc2Only():
		{
			if record, err = row.MakeNewRecord(ss[0], ss[3], "", ""); err != nil {
				return row.Record{}, err
			}
		}
	default:
		if record, err = row.MakeNewRecord(ss[0], "", "", ""); err != nil {
			return row.Record{}, err
		}
	}

	return record, err
}

func parseN1mmGenericFileLine(line string) (record row.Record, err error) {
	// Date     Time    Freq     Mode MyCall        Snt Exchange    Call             Rcvd Exchange   Pts Comment
	// 20200704 1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10
	//    0       1      2         3    4             5  6    7       8               9   10   11      12
	ss := strings.Fields(line)
	if len(ss) < 12 {
		return record, errors.New(fmt.Sprintf("wrong line:%s", line))
	}
	if record, err = row.MakeNewRecord(ss[8], ss[11], ss[0], ss[1]); err != nil {
		return row.Record{}, err
	}
	return
}

func parseEdiQsoRecord(line string) (record row.Record, err error) {
	// 210306;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;
	//    0     1    2   3  4  5  6  7   8  9
	ss := strings.Split(line, ";")
	if len(ss) <= 10 {
		return record, errors.New(fmt.Sprintf("wrong format of EDI QSORecord:%s", line))
	}
	yyyymmdd := ss[0]
	if yyyymmdd, err = timing.FourDigitsYear(yyyymmdd); err != nil {
		return
	}
	if record, err = row.MakeNewRecord(ss[2], ss[9], yyyymmdd, ss[1]); err != nil {
		return row.Record{}, err
	}
	return
}
