package ch2o

import (
	"fmt"
	"strconv"
)

type Ch2oInfo struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

func Append(ch2o float64) {
	appendCh2o(ch2o)
}

func GetCh2oInfo(startDate string, endTime string) []Ch2oInfo {
	ch2oInfo := []Ch2oInfo{}

	ch2os := findCh2os(startDate, endTime)

	for _, val := range ch2os {
		c, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", val.Ch2o), 3)
		ch2oInfo = append(ch2oInfo, Ch2oInfo{
			Time:  val.CreatedAt.Format("2006-01-02 15:04:05"),
			Value: c,
		})
	}

	return ch2oInfo
}
