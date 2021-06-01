package gofunc

import (
	"time"
)

var (
	chineseFestivalCompensatory = map[string]string{
		"20210207" : "春节",
		"20210220" : "春节",
		"20210425" : "劳动",
		"20210508" : "劳动",
		"20210918" : "中秋",
		"20210926" : "国庆",
		"20211009" : "国庆",
	}
	chineseFestival = map[string]string{
		"20210101" : "元旦",
		"20210102" : "元旦",
		"20210103" : "元旦",
		"20210211" : "春节",
		"20210212" : "春节",
		"20210213" : "春节",
		"20210214" : "春节",
		"20210215" : "春节",
		"20210216" : "春节",
		"20210217" : "春节",
		"20210403" : "清明",
		"20210404" : "清明",
		"20210405" : "清明",
		"20210501" : "劳动",
		"20210502" : "劳动",
		"20210503" : "劳动",
		"20210504" : "劳动",
		"20210505" : "劳动",
		"20210612" : "端午",
		"20210613" : "端午",
		"20210614" : "端午",
		"20210919" : "中秋",
		"20210920" : "中秋",
		"20210921" : "中秋",
		"20211001" : "国庆",
		"20211002" : "国庆",
		"20211003" : "国庆",
		"20211004" : "国庆",
		"20211005" : "国庆",
		"20211006" : "国庆",
		"20211007" : "国庆",
	}
)

// IsChineseHoliday 是否是法定节假日
// 只支持 2021 年
func IsChineseHoliday(t time.Time) bool {
	layout := "20060102"
	d := t.Format(layout)
	if _, found := chineseFestival[d]; found {
		// 法定节假日，一定是假期
		return true
	}

	if _, found := chineseFestivalCompensatory[d]; found {
		// 法定调休，一定不是假期
		return false
	}

	return IsCommonWeekend(t)
}

// IsCommonWeekend 是否是通常的周末
func IsCommonWeekend(t time.Time) bool {
	weekday := t.Weekday()
	if time.Sunday == weekday || time.Saturday == weekday {
		return true
	}

	return false
}
