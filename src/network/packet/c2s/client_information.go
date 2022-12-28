package c2s

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type ClientInformation struct {
	Header              packet.PacketHeader
	Locale              dt.String
	ViewDistance        dt.Byte
	ChatMode            dt.VarInt
	ChatColors          dt.Boolean
	DisplayedSkinsPart  dt.UByte
	MainHand            dt.VarInt
	EnableTextFiltering dt.Boolean
	AllowServerListing  dt.Boolean
}

func (ci *ClientInformation) ReadPacket(reader io.Reader) {
	ci.Header.ReadHeader(reader)
	ci.Locale.ReadFrom(reader)
	ci.ViewDistance.ReadFrom(reader)
	ci.ChatMode.ReadFrom(reader)
	ci.ChatColors.ReadFrom(reader)
	ci.DisplayedSkinsPart.ReadFrom(reader)
	ci.MainHand.ReadFrom(reader)
	ci.EnableTextFiltering.ReadFrom(reader)
	ci.AllowServerListing.ReadFrom(reader)
}
