package app

//func TestInsertLog(t *testing.T) {
//	db.Clear()
//	if err := file.InsertLog("./testdata/S59ABC-NOV2020.edi", log.TypeEdiFile); err != nil {
//		fmt.Println(err.Error())
//	}
//	if db.NumberOfRows() != 205 {
//		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 205, db.NumberOfRows()))
//	}
//
//	if err := file.InsertLog("./testdata/S59ABC-MAR2021.edi", log.TypeEdiFile); err != nil {
//		fmt.Println(err.Error())
//	}
//	if db.NumberOfRows() != 413 {
//		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 413, db.NumberOfRows()))
//	}
//
//	if err := file.InsertLog("./testdata/vhf.txt", log.TypeN1mmCallHistory); err != nil {
//		fmt.Println(err.Error())
//	}
//	if db.NumberOfRows() != 43654 {
//		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 43654, db.NumberOfRows()))
//	}
//
//	if err := file.InsertLog("./testdata/S59ABC-NOV2020.edi", log.TypeEdiFile); err != nil {
//		fmt.Println(err.Error())
//	}
//	if db.NumberOfRows() != 43654 {
//		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 43654, db.NumberOfRows()))
//	}
//
//	if err := file.InsertLog("./testdata/S59ABC-VHF-SEP2019.txt", log.TypeEdiFile); err == nil {
//		t.Errorf("Expected error file:S59ABC-VHF-SEP2019.txt is not TypeEdiFile")
//	}
//	if db.NumberOfRows() != 43654 {
//		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 43654, db.NumberOfRows()))
//	}
//
//	if err := file.InsertLog("./testdata/S59ABC-VHF-SEP2019.txt", log.TypeN1mmGenericFile); err != nil {
//		t.Error(err.Error())
//	}
//	if db.NumberOfRows() != 43656 {
//		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 43656, db.NumberOfRows()))
//	}
//
//	fmt.Println(db.NumberOfRows())
//
//	//fmt.Println(db.String())
//	db.Clear()
//}
