package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/log"
	"os"
	"strings"
)

func InsertLog(fileName string, logType log.Type) error {
	file, err := os.Open(fileName)
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
		case log.TypeEdiFile:
			{
				if foundQSORecords {
					line := scanner.Text()
					if rec, err := log.Parse(logType, line); err != nil {
						fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
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
		case log.TypeN1mmCallHistory:
			{
				line := scanner.Text()
				if rec, err := log.Parse(logType, line); err != nil {
					fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
				} else {
					if err := db.Put(rec); err != nil {
						fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
					}
				}
			}
		case log.TypeN1mmGenericFile:
			{
				{
					line := scanner.Text()
					if !strings.HasPrefix(line, "Date") {
						if rec, err := log.Parse(logType, line); err != nil {
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

	if logType == log.TypeEdiFile && !foundQSORecords {
		return errors.New(fmt.Sprintf("file:%s is not %s", fileName, logType.String()))
	}

	return nil
}
