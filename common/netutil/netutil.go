package netutil

import (
	"io"
	"net"

	"github.com/net-byte/gexpose/common/cipher"
)

func Copy(src, dst net.Conn) {
	defer dst.Close()
	defer src.Close()
	buf := make([]byte, 16*1024)
	for {
		n, err := src.Read(buf)
		if err != nil || err == io.EOF {
			break
		}
		b := buf[:n]
		b = cipher.XOR(b)
		_, err = dst.Write(b)
		if err != nil {
			break
		}
	}
}
