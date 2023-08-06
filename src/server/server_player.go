package server

import "mango/src/network"

type IServer interface{}
type IBlockPos interface{}

type GameMode int

const (
	Survival GameMode = iota
	Creative
	Adventure
	Spectator
)

type ServerPlayer struct {

	// server related
	connection   *network.Connection
	server       IServer
	disconnected bool

	// respawn
	respawnPosition IBlockPos
	respawnAngle    float32

	// stats
	gamemode       GameMode
	lastHealth     int
	lastFoodLevel  int
	lastArmor      int
	lastLevel      int
	lastXP         int
	lastActionTime int
}
