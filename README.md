# mango
Minecraft Advanced Network in go (Minecraft Server) 

![banner](assets/img.png)

`$ go run man.go`

## TODO
- [x] Project refactor
- [x] Status protocol
- [x] Implement error handling
- [x] Add config file
- [ ] Login protocol
  - [x] Offline
  - [ ] Online
- [ ] Join game protocol

## Resources
MC Protocol: https://wiki.vg/Protocol


## Misleading Client Side Bugs
- [Sometimes the status requests ignore the response packet and stay at "Pinging..."](https://bugs.mojang.com/browse/MC-125762)