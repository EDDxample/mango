package packet

import "fmt"

func ReadC2SHandshake(bufferedPacket *BufferedPacket) int32 {
	// header
	packetLength, packetID := bufferedPacket.ReadPacketHeader()

	// payload
	protocolVersion := bufferedPacket.ReadVarInt()
	address := bufferedPacket.ReadString(255)
	port := bufferedPacket.ReadUShort()
	nextState := bufferedPacket.ReadVarInt()
	fmt.Printf(`
New Handshake Packet (length: %d, id: %d)
- Protocol Version: %d
- Address: %s
- Port: %d
- Next State: %d
`, packetLength, packetID, protocolVersion, address, port, nextState)
	return nextState
}
