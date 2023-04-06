package server

import (
	"mango/src/config"
	"mango/src/logger"
	"mango/src/network"
	net "mango/src/network"
	"time"
)

type IPlayerList interface {
	BroadcastPacket()
	Tick()
}

type IWorld interface {
	Tick()
	GetDimension() string
}

// =====================================

type MinecraftServer struct {
	connListener network.IConnectionListener
	playerList   IPlayerList
	running      bool
	tickCount    int
	worlds       []IWorld
}

func NewServer() *MinecraftServer {
	connListener := net.NewConnectionListener()
	return &MinecraftServer{
		running:      true,
		tickCount:    0,
		connListener: connListener,
	}
}

// =====================================

// Start @ MinecraftServer.runServer
func (server *MinecraftServer) Start() {
	if err := server.bootstrap(); err != nil {
		logger.Fatal("Failed to initialize server %s", err.Error())
	}

	// Set metadata (description, protocol, motd...)

	nextTickTime := time.Now().UTC().UnixMilli()
	lastOverloadWarning := time.Now().UTC().UnixMilli()

	for server.running {
		loopMillis := time.Now().UTC().UnixMilli() - nextTickTime

		// Time duration check (overload)
		if loopMillis > 2000 && nextTickTime-lastOverloadWarning >= 15000 {
			loopTicks := loopMillis / 50
			logger.Warn("Can't keep up! Is the server overloaded? Running %dms or %d ticks behind", loopMillis, loopTicks)

			nextTickTime += loopTicks * 50
			lastOverloadWarning = nextTickTime
		}

		nextTickTime += 50
		server.Tick()
		/// handleDelayedTasks (AKA tileticks)

		// Wait until next tick
		diff := nextTickTime - time.Now().UTC().UnixMilli()
		time.Sleep(time.Duration(diff) * time.Millisecond)
	}

}

// See .notes/01_init_server.png
func (server *MinecraftServer) bootstrap() error {
	host := config.Host()
	port := config.Port()

	// Start TCP listener
	if err := server.connListener.Start(host, port); err != nil {
		return err
	}

	logger.Info("Server running on %s:%d...", host, port)

	return nil
}

// Tick @ MinecraftServer.tickServer
func (server *MinecraftServer) Tick() error {
	// timestampNanos := time.Now().UTC()
	server.tickCount++

	server.TickChildren()

	/// TODO: update status every 5s (ping, players, etc)

	/// TODO: autosave on tickCount % 6000

	/// TODO: log frame duration and so on...

	return nil
}

// TickChildren @ MinecraftServer.tickChildren
func (server *MinecraftServer) TickChildren() {
	// TODO: tick cmd functions

	// tick worlds
	for _, _ = range server.worlds {

		/*
			// sync world time every second
			if server.tickCount%20 == 0 {
				server.playerList.BroadcastPacket()
			}

			world.Tick()
		*/
	}

	server.connListener.Tick()
	// server.playerList.Tick()
}
