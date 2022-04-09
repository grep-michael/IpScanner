package ipgenerator

import "sync"

type ip_generator struct {
	IpLock     sync.Mutex
	current_ip Ip
}

func (gen *ip_generator) incrementIp() {
	for i, v := range gen.current_ip.bytes[1:] {
		if v == 255 {
			gen.current_ip.bytes[i]++
			gen.current_ip.bytes[i+1] = 0
		}
		if i == 2 {
			gen.current_ip.bytes[3]++
		}
	}
}
func (gen *ip_generator) NextIp() Ip {
	gen.incrementIp()
	return gen.current_ip
}
