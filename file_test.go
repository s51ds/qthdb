package qthdb

import (
	"fmt"
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/log"
	"testing"
)

func Test_process(t *testing.T) {
	N1mmCallHistoryLines := []string{
		"   ",
		"!!Order!!,Call,Name,Loc1,                  ",
		"# Last Edit,2020-08-24                     ",
		"# Send any correction direct ve2fk@arrl.net",
		"# Thanks Beat HB9THU for updated file      ",
		"# VHFREG1                                  ",
		"# VHFHELV26                                ",
		"#                                          ",
		"",
		"   ",
		"13MEK,,JN55SJ,",
		"2E0BMO;;IO83P0",
		"2E0NEY;;IO81VK",
		"2E0UAC;;IO92FJ",
		"2E0VXX/P;;IO82RJ;IO82QJ;",
		"3Z25ZJP;;JO80VE",
		"5P5T;;JO64GX;JO65CS;",
		"7S7V;;JO65SN",
		"9A/IK0BZY;;JN83LD",
		"9A/OM1II;;JN83DM",
		"9A/S51ML;;JN82LX",
		"9A00C;;JN85AO",
		"9A0C;;JN85AO",
		"9A0V;;JN95NH",
		"9A11P;;JN86MB",
		"9A12AO;;JN75WR",
		"9A1ACD;;JN64VU",
	}

	for _, v := range N1mmCallHistoryLines {
		_, err := process(log.TypeN1mmCallHistory, v)
		if err != nil {
			fmt.Println(err.Error())
		}
		//if !next {
		//	os.Exit(1)
		//}

	}
	fmt.Println(db.String())

	N1mmGenericFileLines := []string{
		"Date     Time    Freq     Mode MyCall        Snt Exchange    Call             Rcvd Exchange   Pts Comment",
		"20200704 1402   144222,08  USB S59ABC         59 001 JN76TO  E74G              59 002 JN94FQ   307",
		"20200704 1402   144222,08  USB S59ABC         59 002 JN76TO  9A9D              59 003 JN85KV   125",
		"20200704 1404   144222,08  USB S59ABC         59 003 JN76TO  S53V              59 006 JN76UH    34",
		"20200704 1405   144222,08  USB S59ABC         59 004 JN76TO  OE6AGD            59 002 JN77RB    53",
		"20200704 1406   144222,08  USB S59ABC         59 005 JN76TO  S57RO             59 001 JN76TO     1",
		"20200704 1407   144222,08  USB S59ABC         59 006 JN76TO  HG1A              59 001 JN86MM   109",
		"20200704 1409   144409,86  USB S59ABC         59 007 JN76TO  YU7AJM            59 008 JN95RE   336",
		"20200704 1410   144409,86  USB S59ABC         59 008 JN76TO  YU1ES             59 007 KN04GG   464",
		"20200704 1411   144409,86  USB S59ABC         59 009 JN76TO  9A3SM             59 008 JN85FW    98",
		"20200704 1413   144409,86  USB S59ABC         59 010 JN76TO  YT1WP             59 001 KN04CV   404",
		"20200704 1413   144409,86  USB S59ABC         59 011 JN76TO  HG1W              59 014 JN87GF    99",
		"20200704 1414   144409,86  USB S59ABC         59 012 JN76TO  E76C              59 006 JN84NS   235",
		"20200704 1414   144409,86  USB S59ABC         59 013 JN76TO  9A1DL             59 005 JN85WF    98",
		"20200704 1415   144409,86  USB S59ABC         59 014 JN76TO  9A2MW             59 010 JN75VW    76",
		"20200704 1416   144409,86  USB S59ABC         59 015 JN76TO  OE6END            59 009 JN77PC    62",
		"20200704 1416   144409,86  USB S59ABC         59 016 JN76TO  9A1E              59 011 JN85QT   161",
		"20200704 1417   144409,86  USB S59ABC         59 017 JN76TO  LZ4BF             59 004 KN23LD   830",
		"20200704 1418   144409,86  USB S59ABC         59 018 JN76TO  YP2DX             59 008 KN05NR   435",
	}
	for _, v := range N1mmGenericFileLines {
		process(log.TypeN1mmGenericFile, v)
	}

	EdiFileLines2021 := []string{
		"210306;1422;9A0BB;1;59;001;59;022;;JN85EI;151;;;;  ",
		"210306;1423;S57O;1;59;002;59;021;;JN86DT;56;;;;    ",
		"210306;1424;9A2AE;1;59;003;59;017;;JN86HF;88;;;;   ",
		"210306;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;    ",
		"210306;1430;OE6V;1;59;005;59;020;;JN76VT;27;;;;    ",
		"210306;1431;OK1FPG;1;59;006;59;016;;JN78DR;257;;;; ",
		"210306;1433;DM5B;1;59;007;59;011;;JO71IW;597;;;;   ",
		"210306;1435;OE6VCG/6;1;59;008;59;011;;JN86AX;53;;;;",
		"210306;1438;9A1I;1;59;009;59;019;;JN85FS;113;;;;   ",
		"210306;1440;OE6KPE;1;59;010;59;007;;JN77QA;51;;;;  ",
		"210306;1442;9A1E;1;59;011;59;025;;JN85QT;161;;;;   ",
	}
	for _, v := range EdiFileLines2021 {
		process(log.TypeEdiFile, v)
	}

	EdiFileLines2020 := []string{
		"200904;1422;9A0BB;1;59;001;59;022;;JN85EI;151;;;;  ",
		"200904;1423;S57O;1;59;002;59;021;;JN86DT;56;;;;    ",
		"200904;1424;9A2AE;1;59;003;59;017;;JN86HF;88;;;;   ",
		"200904;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;    ",
		"200904;1430;OE6V;1;59;005;59;020;;JN76VT;27;;;;    ",
		"200904;1431;OK1FPG;1;59;006;59;016;;JN78DR;257;;;; ",
		"200904;1433;DM5B;1;59;007;59;011;;JO71IW;597;;;;   ",
		"200904;1435;OE6VCG/6;1;59;008;59;011;;JN86AX;53;;;;",
		"200904;1438;9A1I;1;59;009;59;019;;JN85FS;113;;;;   ",
		"200904;1440;OE6KPE;1;59;010;59;007;;JN77QA;51;;;;  ",
		"200904;1442;9A1E;1;59;011;59;025;;JN85QT;161;;;;   ",
	}
	for _, v := range EdiFileLines2020 {
		process(log.TypeEdiFile, v)
	}

	EdiFileLines1999 := []string{
		"990306;1422;9A0BB;1;59;001;59;022;;JN85EI;151;;;;  ",
		"990306;1423;S57O;1;59;002;59;021;;JN86DT;56;;;;    ",
		"990306;1424;9A2AE;1;59;003;59;017;;JN86HF;88;;;;   ",
		"990306;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;    ",
		"990306;1430;OE6V;1;59;005;59;020;;JN76VT;27;;;;    ",
		"990306;1431;OK1FPG;1;59;006;59;016;;JN78DR;257;;;; ",
		"990306;1433;DM5B;1;59;007;59;011;;JO71IW;597;;;;   ",
		"990306;1435;OE6VCG/6;1;59;008;59;011;;JN86AX;53;;;;",
		"990306;1438;9A1I;1;59;009;59;019;;JN85FS;113;;;;   ",
		"990306;1440;OE6KPE;1;59;010;59;007;;JN77QA;51;;;;  ",
		"990306;1442;9A1E;1;59;011;59;025;;JN85QT;161;;;;   ",
	}
	for _, v := range EdiFileLines1999 {
		process(log.TypeEdiFile, v)
	}

	EdiFileLines1999 = []string{
		"990306;1422;9A0BB;1;59;001;59;022;;JN85EI;151;;;;  ",
		"990306;1423;S57O;1;59;002;59;021;;JN86DT;56;;;;    ",
		"990306;1424;9A2AE;1;59;003;59;017;;JN86HF;88;;;;   ",
		"990306;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;    ",
		"990306;1430;OE6V;1;59;005;59;020;;JN76VT;27;;;;    ",
		"990306;1431;OK1FPG;1;59;006;59;016;;JN78DR;257;;;; ",
		"990306;1433;DM5B;1;59;007;59;011;;JO71IW;597;;;;   ",
		"990306;1435;OE6VCG/6;1;59;008;59;011;;JN86AX;53;;;;",
		"990306;1438;9A1I;1;59;009;59;019;;JN85FS;113;;;;   ",
		"990306;1440;OE6KPE;1;59;010;59;007;;JN77QA;51;;;;  ",
		"990306;1442;9A1E;1;59;011;59;025;;JN85QT;161;;;;   ",
	}
	for _, v := range EdiFileLines1999 {
		process(log.TypeEdiFile, v)
	}

	fmt.Println(db.String())
}
