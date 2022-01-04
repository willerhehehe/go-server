package utils

import "fmt"

func StrConvert(split string, a ...interface{}) string {
	str := ""

	for index := 0; index < len(a); index++ {
		str1 := fmt.Sprintf("%v", a[index])
		if index > 0 {
			str += split + str1
		} else {
			str += str1
		}
	}
	return str
}
