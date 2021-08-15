package packet

import "fmt"

func ReadC2SLoginStart(bufferedPacket *BufferedPacket) string {
	// header
	packetLength, packetID := bufferedPacket.ReadPacketHeader()

	// payload
	username := bufferedPacket.ReadString(16)

	fmt.Printf(`
New LoginStart Packet (length: %d, id: %d)
- User Name: %s
`, packetLength, packetID, username)

	return username
}
