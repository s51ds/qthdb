package row

import (
	"github.com/s51ds/qthdb/timing"
	"sort"
)

type QueryCase int

const (
	QueryLatestOne QueryCase = iota
	QueryLatestTwo
	QueryLatestFife
	QueryLatestAll
	QueryLatestMonthOn
	QueryLatestByMonth
)

func (t LocatorTimes) SortedByTime() []timing.LogTime {
	logTimes := make([]timing.LogTime, len(t), len(t))
	i := 0
	for k := range t {
		logTimes[i] = k
		i++
	}
	sort.Sort(timing.ByTime(logTimes))
	return logTimes
}

type LocatorWithLogTimes struct {
	locator string
	logTime []timing.LogTime
}

type QueryResponse struct {
	Locator string
	LogTime timing.LogTime
}

func (q *QueryResponse) String() string {
	return string(q.Locator) + " " + q.LogTime.GetString()
}

func (l LocatorsMap) SortedByTime() (resp []QueryResponse) {
	mainSlice := make([]LocatorWithLogTimes, 0, 10)
	for k, v := range l {
		lwt := LocatorWithLogTimes{}
		lwt.locator = k
		lwt.logTime = make([]timing.LogTime, 0, 10)
		for k1 := range v {
			lt := timing.LogTime{}
			lt.SetLoggedTime(k1.LoggedTime())
			lwt.logTime = append(lwt.logTime, lt)
		}
		mainSlice = append(mainSlice, lwt)
	}
	resp = make([]QueryResponse, 0, 10)
	for _, v := range mainSlice {
		locator := v.locator
		for _, v1 := range v.logTime {
			qr := QueryResponse{}
			qr.Locator = locator
			lt := timing.LogTime{}
			lt.SetLoggedTime(v1.LoggedTime())
			qr.LogTime = lt
			resp = append(resp, qr)
		}
	}
	sort.Sort(ByTime(resp))

	//for _, v := range resp {
	//	fmt.Println(v.String())
	//}

	return
}

// ByTime implements sort.Interface for QueryResponse based on gotime field
type ByTime []QueryResponse

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].LogTime.GetUnix() > a[j].LogTime.GetUnix() }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
