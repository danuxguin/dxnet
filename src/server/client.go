package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"lanstonetech.com/common"
	// "lanstonetech.com/network"
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

		_, err := conn.Write(data[0:pos])
		if err != nil {
			fmt.Printf("net.Write err=%v\n", err)
			return
		}

		c, err := io.ReadAtLeast(conn, data, 16)
		if err != nil {
			fmt.Printf("err=%v\n", err)
			return
		}

		fmt.Printf("header=%v\n", data[0:c])

		// header := network.ParseHeader(data[0:common.PACKET_HEAD_LEN])
		//
		// _, err = conn.Read(data[0:header.MsgLen])
		// if err != nil {
		// 	fmt.Printf("err=%v\n", err)
		// 	return
		// }
		result, err := ioutil.ReadAll(conn)
		if err != nil {
			fmt.Printf("err=%v\n", err)
			return
		}
		fmt.Printf("result=%v\n", string(result))

		// fmt.Printf("header=%#v\n", header)
		// fmt.Printf("data=%v len=%v\n", data[0:c], c)

		conn.Close()
	}
}
