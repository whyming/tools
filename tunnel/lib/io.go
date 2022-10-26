package lib

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"time"
)

func Decode(from, to net.Conn) {
	var length [2]byte
	for {
		from.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, err := from.Read(length[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			return
		}
		if n < 2 {
			fmt.Println("bad length")
			return
		}

		data, _ := readLength(from, bytesToInt(length))
		result, err := base64.StdEncoding.DecodeString(string(data))
		if err == nil {
			to.Write(result)
		}
	}
}
func readLength(reader io.Reader, length int) ([]byte, error) {
	buf := make([]byte, length)
	data := make([]byte, 0, length)
	left := length
	for {
		n, err := reader.Read(buf[0:left])
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		data = append(data, buf[0:n]...)
		if n >= left { //读够了
			break
		}
		if n < left {
			left -= n
		}
	}
	return data, nil
}

func Encode(from, to net.Conn) {
	buf := make([]byte, 1024)
	for {
		from.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, err := from.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			return
		}
		if n > 0 {
			data := base64.StdEncoding.EncodeToString(buf[0:n])
			length := intToBytes(len([]byte(data)))
			to.Write(length[:])
			to.Write([]byte(data))
		}
	}
}

func intToBytes(n int) [2]byte {
	var buf [2]byte
	buf[0] = uint8(n >> 8)
	buf[1] = uint8(n)
	return buf
}

func bytesToInt(buf [2]byte) int {
	return int(buf[0])<<8 + int(buf[1])
}
