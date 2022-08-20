package gofunc

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var (
	factor         = []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	parity         = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	validIdCardReg = regexp.MustCompile(`^[1-9]\d{5}(19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	// http://www.gov.cn/xinwen/2016-05/09/content_5071481.htm
	chineseNameRegexp = regexp.MustCompile("^[\u4e00-\u9fa5\u00b7]{2,25}$")
	chinaRegionCode   = map[string]string{
		"11": "北京",
		"12": "天津",
		"13": "河北",
		"14": "山西",
		"15": "内蒙古",
		"21": "辽宁",
		"22": "吉林",
		"23": "黑龙江",
		"31": "上海",
		"32": "江苏",
		"33": "浙江",
		"34": "安徽",
		"35": "福建",
		"36": "江西",
		"37": "山东",
		"41": "河南",
		"42": "湖北",
		"43": "湖南",
		"44": "广东",
		"45": "广西",
		"46": "海南",
		"50": "重庆",
		"51": "四川",
		"52": "贵州",
		"53": "云南",
		"54": "西藏",
		"61": "陕西",
		"62": "甘肃",
		"63": "青海",
		"64": "宁夏",
		"65": "新疆",
		"71": "台湾",
		"81": "香港",
		"82": "澳门",
	}
)

func GenRandomIdCardByBirthDay() string {

	return ""
}

// ChineseIdCardVerify 校验身份证号码
// 已校验方式: 校验位校验、地区码校验
func ChineseIdCardVerify(idCard string) bool {
	if len(idCard) != 18 || !validIdCardReg.MatchString(idCard) {
		return false
	}

	idByte := []byte(strings.ToUpper(idCard))

	regionCode := string([]byte{idByte[0], idByte[1]})
	if _, found := chinaRegionCode[regionCode]; !found {
		return false
	}

	sum := int32(0)
	for i := 0; i < 17; i++ {
		sum += int32(idByte[i]-byte('0')) * factor[i]
	}
	return parity[sum%11] == idByte[17]
}

// ParseChineIdCardAge 根据身份证解析年龄
func ParseChineIdCardAge(idCard string) (int, error) {
	if !ChineseIdCardVerify(idCard) {
		return 0, errors.New("verify idCard error")
	}
	idByte := []byte(strings.ToUpper(idCard))
	str := string([]byte{idByte[6], idByte[7], idByte[8], idByte[9], idByte[10], idByte[11], idByte[12], idByte[13]})

	layout := "20060102"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, errors.New("parse time error")
	}
	now := time.Now()
	age := AgeAt(t, now)
	return age, nil
}

// ChineseNameVerify 中文名称校验
// refer:  http://www.gov.cn/xinwen/2016-05/09/content_5071481.htm
// todo 百家姓
func ChineseNameVerify(name string) bool {
	return chineseNameRegexp.MatchString(name)
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// AgeAt 计算年龄
func AgeAt(birthDate time.Time, now time.Time) int {
	years := now.Year() - birthDate.Year()
	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		years -= 1
	}

	return years
}

func Age(birthDate time.Time) int {
	return AgeAt(birthDate, time.Now())
}

func getAdjustedBirthDay(birthDate time.Time, now time.Time) int {
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func GenBirthDayByAge(age int) time.Time {
	days := rand.Intn(300)
	return time.Now().AddDate(-age, 0, -days)
}
