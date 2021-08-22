package packet

type PacketHeader struct {
	PacketLength int32 `packet_dt:"varint"`
	ID           int32 `packet_dt:"varint"`
}

type C2SHandshake struct {
	PacketHeader    `packet_dt:"header"`
	ProtocolVersion int32  `packet_dt:"varint"`
	Address         string `packet_dt:"string" max_length:"255"`
	Port            int32  `packet_dt:"short"`
	NextState       int32  `packet_dt:"varint"`
}

type C2SRequest struct {
	PacketHeader `packet_dt:"header"`
}

type C2SPing struct {
	PacketHeader `packet_dt:"header"`
	Timestamp    int64 `packet_dt:"long"`
}

type C2SLoginStart struct {
	PacketHeader `packet_dt:"header"`
	Username     string `packet_dt:"string" max_length:"16"`
}
