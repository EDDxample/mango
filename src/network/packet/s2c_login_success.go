package packet

import (
	"fmt"
	"net"
)

func WriteS2CLoginSuccess(connection net.Conn, username string, uuid string) {
	bufferedPacket := BufferedPacket{}

	uuid1 := int64(0xEDD)
	uuid2 := int64(0x1337)

	bufferedPacket.WriteLong(uuid1)
	bufferedPacket.WriteLong(uuid2)
	bufferedPacket.WriteString(username)

	fmt.Printf(`
Sending Pong Packet
- uuid: %d + %d
- username: %s
`, uuid1, uuid2, username)

	connection.Write(bufferedPacket.GetBufferWithHeader(2))
}
