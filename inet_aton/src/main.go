package main

import (
	"ip_helper"
	"fmt"
)

func main() {
	ip := "172.168.5.1"

	result, _ := ip_helper.IpConvert(ip)
	fmt.Println(result)
}