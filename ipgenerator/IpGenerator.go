package ipgenerator

import (
	"bytes"
	"strconv"
)

type Ip struct {
	bytes [4]int
}

func (ip *Ip) AsString() string {
	var buf bytes.Buffer
	for i, v := range ip.bytes {
		buf.WriteString(strconv.FormatInt(int64(v), 10))
		if i >= 0 && i < len(ip.bytes)-1 {
			buf.WriteString(".")
		}
	}
	return buf.String()
}

func (ip *Ip) SetBytes(b [4]int) {
	for i, v := range b {
		ip.bytes[i] = v
	}
}
