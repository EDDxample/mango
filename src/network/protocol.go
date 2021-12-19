package network

const (
	PLAY int32 = iota - 1
	HANDSHAKE
	STATUS
	LOGIN
)

func Handshake(conn net.Conn) {
	var handshake c2s.Handshake
	handshake.ReadFrom(conn)
	fmt.Println(handshake)

	var request c2s.Request
	request.ReadFrom(conn)
	fmt.Println(request)

	// S2C_response

	var ping c2s.Ping
	ping.ReadFrom(conn)
	fmt.Println(ping)
}
