package chatopsbot

import (
	"guanyu.dev/chatopsbot/internal/chatopsbot/eventhandler"
)

// Register all event handlers to discordgo instance
func RegisterHandlers(b *ChatopsBot) {
	eventHandlers := [](interface{}){}

	// Ready Event
	eventHandlers = append(eventHandlers, eventhandler.Ready)
	// InteractionCreate Event
	eventHandlers = append(eventHandlers, eventhandler.SlashCommand(b.scm))

	// Iterate all handlers and register to discordgo instance
	for _, handler := range eventHandlers {
		b.GetSession().AddHandler(handler)
	}
}
