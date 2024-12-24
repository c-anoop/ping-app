package ping

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

type ICMPHeader struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	ID       uint16
	Sequence uint16
}

func checksum(data []byte) uint16 {
	var sum uint32
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(binary.BigEndian.Uint16(data[i:]))
	}
	if len(data)%2 == 1 {
		sum += uint32(data[len(data)-1]) << 8
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum = sum + (sum >> 16)
	return uint16(^sum)
}

func Ping(ip string) {
	conn, err := net.Dial("ip4:icmp", ip)
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		return
	}
	defer conn.Close()

	pid := uint16(os.Getpid() & 0xFFFF)
	payload := []byte("ping-test")

	for i := 0; i < 5; i++ {
		bytes := make([]byte, 8+len(payload))
		header := ICMPHeader{
			Type:     8,
			Code:     0,
			Checksum: 0,
			ID:       pid,
			Sequence: uint16(i),
		}

		binary.BigEndian.PutUint16(bytes[0:2], uint16(header.Type)<<8|uint16(header.Code))
		binary.BigEndian.PutUint16(bytes[2:4], header.Checksum)
		binary.BigEndian.PutUint16(bytes[4:6], header.ID)
		binary.BigEndian.PutUint16(bytes[6:8], header.Sequence)
		copy(bytes[8:], payload)

		header.Checksum = checksum(bytes)
		binary.BigEndian.PutUint16(bytes[2:4], header.Checksum)

		start := time.Now()
		_, err := conn.Write(bytes)
		if err != nil {
			fmt.Printf("Error sending packet: %v\n", err)
			return
		}

		reply := make([]byte, 1500)
		err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			fmt.Printf("Error setting deadline: %v\n", err)
			return
		}

		n, err := conn.Read(reply)
		if err != nil {
			fmt.Printf("Error reading: %v\n", err)
			continue
		}

		duration := time.Since(start)
		fmt.Printf("Reply from %s: bytes=%d time=%v\n", ip, n, duration)
		time.Sleep(time.Second)
	}
}
