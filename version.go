package gofunc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	versionRe = regexp.MustCompile(`^(\d+.){0,3}\d+$`)
	numberRe  = regexp.MustCompile(`(\d+)`)
)

func VersionToInt64(version string) (int64, error) {
	version = strings.TrimSpace(version)
	if len(version) > 19 {
		return 0, errors.New(`version string err`)
	}
	if !versionRe.MatchString(version) {
		return 0, errors.New(`version string err`)
	}
	sp := strings.Split(version, ".")
	res := int64(0)
	for _, v := range sp {
		versionInt, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		res *= 10000
		res += int64(versionInt)

	}

	return res, nil
}

func VersionInt64ToStr(versionInt int64) (string, error) {
	if versionInt <= 0 {
		return "", errors.New(`version int64 err`)
	}
	resArr := make([]string, 0)
	for versionInt > 0 {
		subVersion := versionInt % 10000
		versionInt /= 10000
		resArr = append(resArr, strconv.Itoa(int(subVersion)))
	}
	for i, j := 0, len(resArr)-1; i < j; i, j = i+1, j-1 {
		resArr[i], resArr[j] = resArr[j], resArr[i]
	}

	return strings.Join(resArr, "."), nil
}
