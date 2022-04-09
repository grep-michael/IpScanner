package ipgenerator

import "fmt"

type Ipgenerator struct {
	current_ip Ip
}

func (gen *Ipgenerator) incrementIp() {
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
func (gen *Ipgenerator) NextIp() Ip {
	gen.incrementIp()
	return gen.current_ip
}

func Main() {
	gen := Ipgenerator{}
	gen.current_ip.SetBytes([4]int{1, 1, 1, 255})
	fmt.Println(gen.current_ip.AsString())

	for i := 0; i < 256; i++ {
		gen.incrementIp()
		fmt.Println(gen.current_ip.AsString())
	}
}

/*
func (gen *Ipgenerator) NextIp() Ip {
	curip = gen.current_ip

}*/
