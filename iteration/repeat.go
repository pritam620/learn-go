package iteration

import "strings"

func Repeat(character string, times int) string {
	// repeated := ""
	// for i := 0; i < times; i++ {
	// 	repeated = repeated + character
	// }
	// return repeated
	return strings.Repeat(character, times)
}
