package network

type Protocol int

const (
	SHAKE Protocol = iota
	STATUS
	LOGIN
	PLAY
)

func (p Protocol) ToString() string {
	switch p {
	case SHAKE:
		return "HANDSHAKE"
	case STATUS:
		return "STATUS"
	case LOGIN:
		return "LOGIN"
	case PLAY:
		return "PLAY"
	default:
		return "UNKNOWN"
	}
}

func HandlePacket(conn *Connection, packet *[]byte) {
	switch conn.state {
	case SHAKE:
		HandleHandshakePacket(conn, packet)
	case STATUS:
		HandleStatusPacket(conn, packet)
	case LOGIN:
		HandleLoginPacket(conn, packet)
	case PLAY:
		HandlePlayPacket(conn, packet)
	}
}
