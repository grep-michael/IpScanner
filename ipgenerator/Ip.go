package ipgenerator

import (
	"fmt"
	"strconv"
	"strings"
)

type IpGenerator uint32

func (ipg *IpGenerator) ToString() string {
	var sb strings.Builder
	//printBytes(strconv.FormatInt(int64(*ipg), 2))

	for i := 1; i <= 4; i++ {
		//bit shift
		shifted := int(*ipg) >> (8 * (4 - i))

		//apply mask
		mask := 255
		masked := shifted & mask
		sb.WriteString(strconv.FormatInt(int64(masked), 10))
		if i <= 3 {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func printBytes(s string) {
	//fmt.Println(s)
	oct := 0
	for i := 0; i < len(s); i++ {
		fmt.Printf(string(s[i]))
		if oct == 7 {
			fmt.Printf(",")
			oct = 0
		} else {
			oct++
		}
	}
	fmt.Printf("\n")
}
