package packet

import (
	"fmt"
	"net"
)

func WriteS2CPong(connection net.Conn, timestamp int64) {
	bufferedPacket := BufferedPacket{}
	bufferedPacket.WriteLong(timestamp)

	fmt.Printf(`
Sending Pong Packet
- Timestamp: %d
`, timestamp)

	connection.Write(bufferedPacket.GetBufferWithHeader(1))
}
