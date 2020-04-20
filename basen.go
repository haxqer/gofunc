package gofunc

func DecimalToAny(in int, strArr []rune) string {
	baseN := len(strArr)
	if baseN == 0 {
		return ""
	}

	tailInt := in % baseN
	tail := string(strArr[tailInt])
	if tailInt == in {
		return tail
	}
	HeadInt := (in-tailInt)/baseN
	return DecimalToAny(HeadInt, strArr) + tail
}
