package main

import (
	"fmt"
)

func main() {
	ip := GetLocalIP()

	fmt.Printf("Local IP is %s\n", ip)
}
