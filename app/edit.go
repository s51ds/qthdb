package app

//// InsertLog parses log's lines, convert each line into records and puts them into db
//// If file can not be open, non nil error is returned.
//// During line parse process error can be detected (e.g. invalid call or locator or unexpected
//// line fields). In such case, error is logged to stdOut, nothing is put into db
//// but file parsing is continued until the end of file
//func InsertLog(fileName string, logType log.Type) error {
//	file, err := os.Open(fileName)
//	if err != nil {
//		return err
//	}
//	defer func() {
//		_ = file.Close()
//	}()
//
//	foundQSORecords := false
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		switch logType {
//		case log.TypeEdiFile:
//			{
//				if foundQSORecords {
//					line := scanner.Text()
//					if rec, err := log.Parse(logType, line); err != nil {
//						fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
//					} else {
//						if err := db.Put(rec); err != nil {
//							fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
//						}
//					}
//				} else {
//					if strings.HasPrefix(scanner.Text(), "[QSORecords") {
//						foundQSORecords = true
//					}
//				}
//			}
//		case log.TypeN1mmCallHistory:
//			{
//				line := scanner.Text()
//				if rec, err := log.Parse(logType, line); err != nil {
//					fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
//				} else {
//					if err := db.Put(rec); err != nil {
//						fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
//					}
//				}
//			}
//		case log.TypeN1mmGenericFile:
//			{
//				{
//					line := scanner.Text()
//					if !strings.HasPrefix(line, "Date") {
//						if rec, err := log.Parse(logType, line); err != nil {
//							fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
//						} else {
//							if err := db.Put(rec); err != nil {
//								fmt.Printf("%s; file=%s, line=%s\n", err.Error(), fileName, line)
//							}
//						}
//					}
//				}
//			}
//		default:
//			return errors.New(fmt.Sprintf("Unknown file type: %d", logType))
//		}
//
//	}
//
//	if err := scanner.Err(); err != nil {
//		return err
//	}
//
//	if logType == log.TypeEdiFile && !foundQSORecords {
//		return errors.New(fmt.Sprintf("file:%s is not %s", fileName, logType.String()))
//	}
//
//	return nil
//}
