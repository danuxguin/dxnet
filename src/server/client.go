package main

import (
	"fmt"
	"lanstonetech.com/common"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "123.57.81.38:10001")
	if err != nil {
		fmt.Printf("net.Dial err=%v\n", err)
		return
	}

	for {
		data := make([]byte, 1024)

		pos := 0
		common.WriteUint32(data[pos:pos+4], uint32(1001))
		pos += 4
		common.WriteUint32(data[pos:pos+4], uint32(4))
		pos += 4
		common.WriteUint32(data[pos:pos+4], uint32(4))
		pos += 4
		common.WriteUint32(data[pos:pos+4], uint32(4))
		pos += 4
		common.WriteUint32(data[pos:pos+4], uint32(1))
		pos += 4

		c, err := conn.Write(data[0:pos])
		if err != nil {
			fmt.Printf("net.Write err=%v\n", err)
			return
		}

		c, err = conn.Read(data)
		if err != nil {
			fmt.Printf("net.Read err=%v\n", err)
			return
		}

		fmt.Printf("data=%v len=%v\n", data[0:c], c)
	}
}
