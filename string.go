package gofunc

import "math/rand"

func RandStringRunes(n int, strArr []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = strArr[rand.Intn(len(strArr))]
	}
	return string(b)
}
