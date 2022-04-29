package ipgenerator

import (
	"strconv"
	"strings"
)

type Ip uint32

func (ip *Ip) ToString() string {
	var sb strings.Builder
	//printBytes(strconv.FormatInt(int64(*ipg), 2))

	for i := 1; i <= 4; i++ {
		//bit shift
		shifted := int(*ip) >> (8 * (4 - i))

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
