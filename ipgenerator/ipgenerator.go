package ipgenerator

type IpGenerator struct {
	outChan    chan uint32
	current_ip Ip
}
