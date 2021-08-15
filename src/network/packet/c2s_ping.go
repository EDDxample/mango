package packet

import "fmt"

func ReadC2SPing(bufferedPacket *BufferedPacket) int64 {
	// header
	packetLength, packetID := bufferedPacket.ReadPacketHeader()
	// payload
	timestamp := bufferedPacket.ReadLong()

	fmt.Printf(`
New Ping Packet (length: %d, id: %d)
- Timestamp: %d
`, packetLength, packetID, timestamp)

	return timestamp
}
