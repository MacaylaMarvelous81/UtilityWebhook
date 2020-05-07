package main
import (
	"github.com/sandertv/mcwss"
)
import (
	"github.com/sandertv/mcwss/protocol/event"
)
func main() {
	server := mcwss.NewServer(nil)
	server.OnConnection(func(player *mcwss.Player) {
		player.SendMessage("Welcome! You have been connected to the external service. Type .help in chat for commands!")
		player.OnPlayerMessage(func(event *event.PlayerMessage) {
			if event.MessageType == "chat" {
				if event.Message == ".help" {
					player.SendMessage("Command List:\n.help - Shows this help menu")
				}
			}
		})
	})
	server.OnDisconnection(func(player *mcwss.Player) {})
	server.Run()
}