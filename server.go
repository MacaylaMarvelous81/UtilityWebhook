package main
import (
	"github.com/sandertv/mcwss"
	"github.com/sandertc/mcwss/protocol/event"
	"github.com/sandertv/mcwss/mctype"
)
func main() {
	server := mcwss.NewServer(nil)
	server.OnConnection(func(player *mcwss.Player) {
		var world = player.World()
		player.SendMessage("Welcome! You have been connected to the external service. Type .help in chat for commands!")
		player.OnPlayerMessage(func(event *event.PlayerMessage) {
			if event.MessageType == "chat" {
				if event.Sender == player.Name() {
					if event.Message == ".help" {
						player.SendMessage("Command List:\n.help - Shows this help menu\n.love - Spawns heart particles on you\n.angry - Spawns angry particles on you")
					}
					if event.Message == ".love" {
						player.Position(func(position mctype.Position) {
							world.SpawnParticle("minecraft:heart_particle", position)
							player.SendMessage("Heart particles created!")
						})
					}
					if event.Message == ".angry" {
						player.Position(func(position mctype.Position) {
							world.SpawnParticle("minecraft:villager_angry", position)
							player.SendMessage("Angry particles created!")
						})
					}
				}
			}
		})
	})
	server.OnDisconnection(func(player *mcwss.Player) {})
	server.Run()
}
