package packet

import "fmt"

func ReadC2SRequest(bufferedPacket *BufferedPacket) {
	// header
	packetLength, packetID := bufferedPacket.ReadPacketHeader()
	fmt.Printf(`
New Request Packet (length: %d, id: %d)
`, packetLength, packetID)
}
