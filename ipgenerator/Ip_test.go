package ipgenerator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generator_default(t *testing.T) {
	assert := assert.New(t)
	ip := new(Ip)
	assert.Equal(fmt.Sprintf("%x", *ip), "0", "Initial ip should be zero")
	*ip = 3232235521 //192.168.0.1 = 0a . 80 . 00 . 01
}

func Test_generator_actual_ip(t *testing.T) {
	assert := assert.New(t)
	ip := new(Ip)
	*ip = 3232235521 //192.168.0.1 = c0 . a8 . 00 . 01
	assert.Equal(fmt.Sprintf("%x", *ip), "c0a80001", "Test ip 192.168.0.1 as hex")
}
func Test_ip_to_string(t *testing.T) {
	assert := assert.New(t)
	ip := new(Ip)
	*ip = 3232235521

	assert.Equal("192.168.0.1", ip.ToString(), "print ip as string")
}
func Test_ip_to_string2(t *testing.T) {
	assert := assert.New(t)
	ip := new(Ip)
	*ip = 3232298295
	assert.Equal("192.168.245.55", ip.ToString(), "print ip as string")
}
func Test_increment(t *testing.T) {
	assert := assert.New(t)
	ip := new(Ip)
	*ip = 3232235521
	assert.Equal("192.168.0.1", ip.ToString(), "print ip as string")
	*ip++
	assert.Equal("192.168.0.2", ip.ToString(), "print ip as string")
}
