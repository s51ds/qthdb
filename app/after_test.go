package app

import (
	"fmt"
	"testing"
)

func TestAfterTheContest(t *testing.T) {
	if err := AfterTheContest("../testdata/scp/maj/S59ABC-MAY-2021.edi", "../testdata/scp/maj/vhf.txt"); err != nil {
		fmt.Println(err.Error())
	}

}
