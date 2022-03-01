package file

import (
	"bufio"
	"errors"
	"fmt"
	hamLog "github.com/s51ds/qthdb/ctestlog"
	"github.com/s51ds/qthdb/db"
	"os"
	"strings"
)

// Open is just a wrapper around os.Open. It returns err if
// file does not exist included info about working directory
func Open(fileName string) (file *os.File, err error) {
	file, err = os.Open(fileName)
	if err != nil {
		wd, _ := os.Getwd()
		if wd != "" {
			wd = "Working directory=" + wd
		}
		return nil, fmt.Errorf("%v %s", err, wd)
	}
	return file, nil
}

// InsertLog parses log's lines, convert each line into records and puts them into db
// If file can not be open, non nil error is returned.
// During line parse error can be detected (e.g. invalid call or locator or unexpected
// line fields). In such case, error is logged to stdOut, nothing is put into db
// but file parsing is continued until the end of file
func InsertLog(fileName string, logType hamLog.Type) error {
	file, err := Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	foundQSORecords := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch logType {
		case hamLog.TypeEdiFile:
			{
				if foundQSORecords {
					line := scanner.Text()
					if rec, err := hamLog.Parse(logType, line); err != nil {
						fmt.Printf("-----> WARNING -----> %s; file=%s, line=%s\n", err.Error(), fileName, line)
					} else {
						if err := db.Put(rec); err != nil {
							fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
						}
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
					fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
				} else {
					if err := db.Put(rec); err != nil {
						fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
					}
				}
			}
		case hamLog.TypeN1mmGenericFile:
			{
				{
					line := scanner.Text()
					if !strings.HasPrefix(line, "Date") {
						if rec, err := hamLog.Parse(logType, line); err != nil {
							fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
						} else {
							if err := db.Put(rec); err != nil {
								fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
							}
						}
					}
				}
			}
		default:
			return errors.New(fmt.Sprintf("Unknown file type: %d", logType))
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if logType == hamLog.TypeEdiFile && !foundQSORecords {
		return errors.New(fmt.Sprintf("file:%s is not %s", fileName, logType.String()))
	}

	return nil
}
