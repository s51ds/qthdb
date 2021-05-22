package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/file"
	hamLog "github.com/s51ds/qthdb/log"
	"github.com/s51ds/qthdb/row"
	"strings"
)

func AfterTheContest(ediLogFileName, scpFileName string) error {
	logData, err := readData(ediLogFileName, hamLog.TypeEdiFile)
	if err != nil {
		return err
	}
	scpData, err := readData(scpFileName, hamLog.TypeN1mmCallHistory)
	if err != nil {
		return err
	}
	if err := doReport(logData, scpData); err != nil {
		return errors.New(fmt.Sprintf("%s in %s", err.Error(), scpFileName))
	}
	return nil
}

func doReport(logData, scpData []data) error {
	// convert scpData slice into scpMap
	scpMap := make(map[string]string, len(scpData))
	for _, v := range scpData {
		if _, has := scpMap[v.callSign]; has { // only possible if scp has "manual" intervention
			return errors.New("exited with error: duplicated " + v.callSign)
		}
		scpMap[v.callSign] = v.locators
	}

	for i, qso := range logData {
		if locs, has := scpMap[qso.callSign]; has {
			if !strings.Contains(locs, qso.locators) {
				fmt.Println(fmt.Sprintf("qso %03d %s %s -> qso locator differs from locators in scp:%s", i+1, qso.callSign, qso.locators, locs))
			}
		} else {
			fmt.Println(fmt.Sprintf("qso %03d %s %s -> locator not in scp", i+1, qso.callSign, qso.locators))
		}
	}
	return nil

}

type data struct {
	callSign string
	locators string //space separated locators, two of them are available in scp file
}

func (d *data) String() string {
	return d.callSign + " " + d.locators
}

func readData(fileName string, logType hamLog.Type) (resp []data, err error) {
	f, err := file.Open(fileName)
	if err != nil {
		return resp, err
	}
	defer func() {
		_ = f.Close()
	}()

	resp = make([]data, 0, 1000)

	foundQSORecords := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		switch logType {
		case hamLog.TypeEdiFile:
			{
				if foundQSORecords {
					line := scanner.Text()
					if rec, err := hamLog.Parse(logType, line); err != nil {
						fmt.Printf("%s; f=%s, line=%s\n", err.Error(), fileName, line)
					} else {
						data := recData(rec)
						resp = append(resp, data)
					}
				} else {
					if strings.HasPrefix(scanner.Text(), "[QSORecords") {
						foundQSORecords = true
					}
				}
			}
		case hamLog.TypeN1mmCallHistory:
			{
				line := scanner.Text()
				if rec, err := hamLog.Parse(logType, line); err != nil {
					fmt.Printf("%s; f=%s, line=%s\n", err.Error(), fileName, line)
				} else {
					data := recData(rec)
					resp = append(resp, data)
				}
			}
		case hamLog.TypeN1mmGenericFile:
			{
				{
					line := scanner.Text()
					if !strings.HasPrefix(line, "Date") {
						if rec, err := hamLog.Parse(logType, line); err != nil {
							fmt.Printf("%s; f=%s, line=%s\n", err.Error(), fileName, line)
						} else {
							data := recData(rec)
							resp = append(resp, data)
						}
					}
				}
			}
		default:
			return resp, errors.New(fmt.Sprintf("Unknown f type: %d", logType))
		}
	}

	if err := scanner.Err(); err != nil {
		return resp, err
	}

	if logType == hamLog.TypeEdiFile && !foundQSORecords {
		return resp, errors.New(fmt.Sprintf("f:%s is not %s", fileName, logType.String()))
	}

	return

}

func recData(rec row.Record) data {
	d := data{
		callSign: string(rec.CallSign()),
		locators: rec.Locators().StringLocators(),
	}
	return d
}
