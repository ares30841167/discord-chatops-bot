package eventhandler

import (
	"github.com/bwmarrin/discordgo"
	"guanyu.dev/chatopsbot/internal/chatopsbot/slashcmd"
)

// This function will make an "InteractionCreate" event handler that
// can access all registered slash command handler
func SlashCommand(scm *slashcmd.SlashCommandManager) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// This function will be called when the bot receives
	// the "InteractionCreate" event from Discord.
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandHandlers := scm.GetSlashCommandHandlers()
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	}
}
