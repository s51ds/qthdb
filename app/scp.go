package app

import "fmt"

func sPrintN1mmScpFormat(callSign, loc1, loc2 string) string {
	// S58M,,JN76PL,JN76JC
	return fmt.Sprintf("%s,,%s,%s", callSign, loc1, loc2)
}
