package ipgenerator

func NewIpGenerator() ip_generator {
	gen := ip_generator{}
	gen.current_ip.bytes = [4]int{1, 1, 1, 0}
	return gen
}
