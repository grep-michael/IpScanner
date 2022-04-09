package main

import (
	"fmt"

	"michaelknudsen.com/ipscanner/ipgenerator"
)

func main() {
	//ip class
	ip := ipgenerator.Ip{}
	b := [...]int{1, 1, 1, 1}
	ip.SetBytes(b)
	fmt.Println(ip.AsString())

	//ip genorator
	ipgenerator.Main()

}
